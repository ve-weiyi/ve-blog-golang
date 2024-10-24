package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/ws"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/ipx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/chatrpc"
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
	// 接收消息
	receive := func(msg []byte) (tx []byte, err error) {
		logx.Info(string(msg))

		var req types.ChatSocketMsg
		err = json.Unmarshal(msg, &req)
		if err != nil {
			return nil, fmt.Errorf("msg content must be :%v", jsonconv.AnyToJsonNE(req))
		}

		if req.Content == "" {
			return nil, fmt.Errorf("content is empty")
		}

		token := cast.ToString(r.Context().Value(constant.HeaderAuthorization))
		uid := cast.ToString(r.Context().Value(constant.HeaderUid))
		device := cast.ToString(r.Context().Value(constant.HeaderTerminal))
		// 如果有uid,则需要校验用户是否登录
		if token != "" || uid != "" {
			_, err = l.svcCtx.TokenHolder.VerifyToken(r.Context(), token, cast.ToString(uid))
			if err != nil {
				return nil, err
			}

			_, err := l.svcCtx.AccountRpc.GetUserInfo(r.Context(), &accountrpc.UserIdReq{UserId: cast.ToInt64(uid)})
			if err != nil {
				return nil, err
			}
		}
		host, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
		if err != nil {
			return nil, err
		}

		ip, err := ipx.GetIpInfoByBaidu(host)
		if err != nil {
			return nil, err
		}

		chat := &chatrpc.ChatMessageNewReq{
			Id:          0,
			UserId:      cast.ToString(uid),
			DeviceId:    device,
			ChatId:      "",
			ReplyMsgId:  "",
			ReplyUsers:  "",
			IpAddress:   ip.Location,
			IpSource:    ip.Origip,
			ChatContent: req.Content,
			Type:        req.Type,
		}

		out, err := l.svcCtx.ChatRpc.AddChatMessage(r.Context(), chat)
		if err != nil {
			return nil, err
		}

		var resp types.ChatSocketMsgResp
		resp = types.ChatSocketMsgResp{
			Type:    out.Type,
			Content: out.ChatContent,
			Time:    out.CreatedAt,
		}

		return []byte(jsonconv.AnyToJsonNE(resp)), nil
	}

	ws.HandleWebSocket(w, r, receive)

	return nil
}

const (
	ChatTypeWebsocket = 1
)
