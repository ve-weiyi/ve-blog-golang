package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/global/constant"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/messagerpc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/ws"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/ipx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
)

type WebsocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	header *restx.RestHeader
}

// Websocket消息
func NewWebsocketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebsocketLogic {
	return &WebsocketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebsocketLogic) Websocket(w http.ResponseWriter, r *http.Request) error {
	// 校验token
	//err := l.checkSign()
	//if err != nil {
	//	return err
	//}

	uid := cast.ToString(l.ctx.Value(restx.HeaderUid))
	tid := cast.ToString(l.ctx.Value(restx.HeaderTerminal))
	ip := restx.GetRemoteIP(r)
	is, _ := ipx.GetIpSourceByBaidu(ip)

	client, err := ws.Upgrade(w, r, map[string]string{
		restx.HeaderUid:      uid,
		restx.HeaderTerminal: tid,
		"IP":                 ip,
		"IS":                 is,
	})

	if err != nil {
		l.Errorf("websocket连接失败：%+v", err)
		return err
	}

	client.SetHandleConnect(func() error {
		l.Infof("websocket连接成功：%+v", client.ClientID)
		event := &types.OnlineCountResp{
			Msg:   fmt.Sprintf("当前在线人数：%d", l.svcCtx.Hub.Metrics.Connections),
			Count: int(l.svcCtx.Hub.Metrics.Connections),
		}

		reply := &types.ReplyMsg{
			Type:      constant.OnlineCount,
			Data:      jsonconv.AnyToJsonNE(event),
			Timestamp: time.Now().Unix(),
		}

		// 发送在线人数
		l.svcCtx.Hub.Broadcast([]byte(jsonconv.AnyToJsonNE(reply)))

		info := types.ClientInfoResp{
			ClientId:  r.RemoteAddr,
			UserId:    client.Claims[restx.HeaderUid],
			DeviceId:  client.Claims[restx.HeaderTerminal],
			Nickname:  "",
			IpAddress: client.Claims["IP"],
			IpSource:  client.Claims["IS"],
		}

		reply = &types.ReplyMsg{
			Type:      constant.ClientInfo,
			Data:      jsonconv.AnyToJsonNE(info),
			Timestamp: time.Now().Unix(),
		}

		// 发送用户信息
		return l.svcCtx.Hub.Broadcast([]byte(jsonconv.AnyToJsonNE(reply)))
	})

	client.SetHandleDisconnect(func() error {
		l.Infof("websocket连接断开：%+v", client.ClientID)
		event := &types.OnlineCountResp{
			Msg:   fmt.Sprintf("当前在线人数：%d", l.svcCtx.Hub.Metrics.Connections),
			Count: int(l.svcCtx.Hub.Metrics.Connections),
		}

		reply := &types.ReplyMsg{
			Type:      constant.OnlineCount,
			Data:      jsonconv.AnyToJsonNE(event),
			Timestamp: time.Now().Unix(),
		}

		return l.svcCtx.Hub.Broadcast([]byte(jsonconv.AnyToJsonNE(reply)))
	})

	client.SetHandleMessage(func(message []byte) error {
		return l.ReceiveHandler(client, message)
	})

	return client.Connect(l.svcCtx.Hub)
}

func (l *WebsocketLogic) checkSign() (err error) {
	// 校验是否登录
	token := cast.ToString(l.ctx.Value(restx.HeaderToken))
	uid := cast.ToString(l.ctx.Value(restx.HeaderUid))
	// 如果有uid,则需要校验用户是否登录
	if token != "" || uid != "" {
		err := l.svcCtx.TokenHolder.VerifyToken(l.ctx, token, cast.ToString(uid))
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *WebsocketLogic) ReceiveHandler(client *ws.Client, msg []byte) (err error) {
	var receiveMsg types.ReceiveMsg
	err = json.Unmarshal(msg, &receiveMsg)
	if err != nil {
		return
	}

	switch receiveMsg.Type {
	case constant.OnlineCount:
		l.Info("在线人数")

	case constant.HistoryRecord:
		l.Info("历史记录")
		err = l.onHistoryRecordEvent(client)
		if err != nil {
			return err
		}

	case constant.HeartBeat:
		l.Info("心跳检测")
		err = l.onHeartBeatEvent(client, receiveMsg)
		if err != nil {
			return err
		}

	case constant.SendMessage:
		l.Info("发送消息")
		err = l.onSendMessageEvent(client, receiveMsg)
		if err != nil {
			return err
		}

	case constant.RecallMessage:
		l.Info("撤回消息")
		err = l.onRecallMessageEvent(client, receiveMsg)
		if err != nil {
			return err
		}

	default:

	}

	return nil
}

func (l *WebsocketLogic) onHistoryRecordEvent(client *ws.Client) (err error) {

	in := &messagerpc.FindChatListReq{
		After:  0,
		Before: time.Now().Unix(),
	}

	out, err := l.svcCtx.MessageRpc.FindChatList(context.Background(), in)
	if err != nil {
		return err
	}

	list := make([]*types.ChatRecordResp, 0)
	for _, v := range out.List {
		m := ConvertChatTypes(v)

		list = append(list, m)
	}

	reply := &types.ReplyMsg{
		Type:      constant.HistoryRecord,
		Data:      jsonconv.AnyToJsonNE(list),
		Timestamp: time.Now().Unix(),
	}

	err = l.svcCtx.Hub.Broadcast([]byte(jsonconv.AnyToJsonNE(reply)))
	if err != nil {
		return err
	}
	return nil
}

func (l *WebsocketLogic) onHeartBeatEvent(client *ws.Client, msg types.ReceiveMsg) (err error) {
	reply := &types.ReplyMsg{
		Type:      constant.HeartBeat,
		Data:      fmt.Sprintf("pong"),
		Timestamp: time.Now().Unix(),
	}

	err = l.svcCtx.Hub.Broadcast([]byte(jsonconv.AnyToJsonNE(reply)))
	if err != nil {
		return err
	}
	return nil
}

func (l *WebsocketLogic) onSendMessageEvent(client *ws.Client, msg types.ReceiveMsg) (err error) {
	var chat types.SendMessageReq
	err = json.Unmarshal([]byte(msg.Data), &chat)
	if err != nil {
		return err
	}

	in := &messagerpc.ChatNewReq{
		UserId:     client.Claims[restx.HeaderUid],
		TerminalId: client.Claims[restx.HeaderTerminal],
		IpAddress:  client.Claims["IP"],
		IpSource:   client.Claims["IS"],
		Type:       chat.Type,
		Content:    chat.Content,
	}

	out, err := l.svcCtx.MessageRpc.AddChat(context.Background(), in)
	if err != nil {
		return err
	}

	reply := &types.ReplyMsg{
		Type:      constant.SendMessage,
		Data:      jsonconv.AnyToJsonNE(ConvertChatTypes(out)),
		Timestamp: time.Now().Unix(),
	}

	err = l.svcCtx.Hub.Broadcast([]byte(jsonconv.AnyToJsonNE(reply)))
	if err != nil {
		return err
	}

	return nil
}

func (l *WebsocketLogic) onRecallMessageEvent(client *ws.Client, msg types.ReceiveMsg) (err error) {
	var req types.RecallMessageReq
	err = json.Unmarshal([]byte(msg.Data), &req)
	if err != nil {
		return err
	}

	in := &messagerpc.IdsReq{
		Ids: []int64{req.Id},
	}

	_, err = l.svcCtx.MessageRpc.DeletesChat(context.Background(), in)
	if err != nil {
		return err
	}

	event := &types.RecallMessageResp{
		Id: req.Id,
	}

	reply := &types.ReplyMsg{
		Type:      constant.RecallMessage,
		Data:      jsonconv.AnyToJsonNE(event),
		Timestamp: time.Now().Unix(),
	}

	err = l.svcCtx.Hub.Broadcast([]byte(jsonconv.AnyToJsonNE(reply)))
	if err != nil {
		return err
	}

	return nil
}

func ConvertChatTypes(in *messagerpc.ChatDetails) *types.ChatRecordResp {
	return &types.ChatRecordResp{
		Id:          in.Id,
		UserId:      in.UserId,
		DeviceId:    in.TerminalId,
		Nickname:    "",
		Avatar:      "",
		IpAddress:   in.IpAddress,
		IpSource:    in.IpSource,
		Type:        in.Type,
		ChatContent: in.Content,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
}
