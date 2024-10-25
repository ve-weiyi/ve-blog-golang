package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/ipx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/messagerpc"
)

type WebSocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// WebSocket消息
func NewWebSocketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebSocketLogic {
	return &WebSocketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebSocketLogic) WebSocket(w http.ResponseWriter, r *http.Request) error {
	l.svcCtx.WebsocketManager.RegisterWebSocket(w, r, r.RemoteAddr)

	// 接收消息
	receive := func(msg []byte) (tx []byte, err error) {
		logx.Info(string(msg))

		token := cast.ToString(r.Context().Value(constant.HeaderAuthorization))
		uid := cast.ToString(r.Context().Value(constant.HeaderUid))
		// 如果有uid,则需要校验用户是否登录
		if token != "" || uid != "" {
			_, err = l.svcCtx.TokenHolder.VerifyToken(r.Context(), token, cast.ToString(uid))
			if err != nil {
				return nil, err
			}
		}

		var req types.ChatSocketMsg
		err = json.Unmarshal(msg, &req)
		if err != nil {
			return nil, fmt.Errorf("msg content must be :%v", jsonconv.AnyToJsonNE(req))
		}

		if req.Data == "" {
			return nil, fmt.Errorf("msg data is empty")
		}

		var resp types.ChatSocketMsgResp
		switch req.Type {
		case HEART_BEAT:
			// 心跳
			_, err = l.onHeartbeat(w, r, uid)
			if err != nil {
				return nil, err
			}

		case ONLINE_COUNT:
			// 在线人数
			count, err := l.onOnline(w, r)
			if err != nil {
				return nil, err
			}

			return []byte(jsonconv.AnyToJsonNE(count)), nil

		case HISTORY_RECORD:
			// 历史记录

		case SEND_MESSAGE:
			// 发送消息

		case RECALL_MESSAGE:
		// 撤回消息
		default:

		}

		return []byte(jsonconv.AnyToJsonNE(resp)), nil
	}

	l.svcCtx.WebsocketManager.OnReceiveMsg(r.RemoteAddr, receive)
	return nil
}

func (l *WebSocketLogic) onHeartbeat(w http.ResponseWriter, r *http.Request, uid string) (data any, err error) {
	key := KeyPrefix + uid
	// 更新用户的心跳时间，设置过期时间为10分钟
	err = l.svcCtx.Redis.SetexCtx(l.ctx, key, time.Now().String(), 10*60)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (l *WebSocketLogic) onOnline(w http.ResponseWriter, r *http.Request) (data any, err error) {
	// 获取当前在线用户的数量
	keys, err := l.svcCtx.Redis.KeysCtx(l.ctx, KeyPrefix)
	if err != nil {
		return nil, err
	}

	return len(keys), nil
}

func (l *WebSocketLogic) onHistoryRecord(w http.ResponseWriter, r *http.Request) (list []*messagerpc.ChatMessageDetails, err error) {
	// 获取历史记录
	return nil, nil
}

func (l *WebSocketLogic) onSendMessage(w http.ResponseWriter, r *http.Request, req types.ChatSocketMsg) (data any, err error) {
	uid := cast.ToString(r.Context().Value(constant.HeaderUid))
	device := cast.ToString(r.Context().Value(constant.HeaderTerminal))

	// 发送消息
	host, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		return nil, err
	}

	ip, err := ipx.GetIpInfoByBaidu(host)
	if err != nil {
		return nil, err
	}

	chat := &messagerpc.ChatMessageNewReq{
		Id:          0,
		UserId:      uid,
		DeviceId:    device,
		TopicId:     "",
		ReplyMsgId:  "",
		ReplyUserId: "",
		IpAddress:   ip.Location,
		IpSource:    ip.Origip,
		ChatContent: req.Data,
		Type:        "",
	}

	out, err := l.svcCtx.MessageRpc.AddChatMessage(r.Context(), chat)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (l *WebSocketLogic) onRecallMessage(w http.ResponseWriter, r *http.Request, uid string) (resp types.ChatSocketMsgResp, err error) {
	// 撤回消息
	return resp, nil
}

const (
	ChatTypeWebsocket = 1

	KeyPrefix = "chat:online:"
)

const (
	ONLINE_COUNT   = 1 // 在线人数
	HISTORY_RECORD = 2 // 历史记录
	SEND_MESSAGE   = 3 // 发送消息
	RECALL_MESSAGE = 4 // 撤回消息
	HEART_BEAT     = 5 // 心跳
)
