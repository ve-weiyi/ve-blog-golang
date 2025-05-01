package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
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

	// 创建客户端
	client := l.NewClient(w, r)

	// 注册
	err := l.RegisterHandler(client)
	if err != nil {
		return err
	}

	// 关闭连接（前端关闭）
	client.WsConn.SetCloseHandler(func(code int, text string) error {
		l.Info("关闭连接", code, text)
		l.svcCtx.WebsocketManager.PushMessage(client.ClientId, text)

		return l.UnregisterHandler(client)
	})

	// 监听消息
	go client.ReadMessage(func(msg []byte) (err error) {
		return l.ReceiveHandler(client, msg)
	})
	return nil
}

func (l *WebsocketLogic) NewClient(w http.ResponseWriter, r *http.Request) *ws.Client {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic("http转换升级为websocket失败：")
	}

	ip := restx.GetRemoteIP(r)
	is, _ := ipx.GetIpSourceByBaidu(ip)

	client := &ws.Client{
		ClientId:    r.RemoteAddr,
		UserId:      cast.ToString(l.ctx.Value(restx.HeaderUid)),
		DeviceId:    cast.ToString(l.ctx.Value(restx.HeaderTerminal)),
		IpAddress:   ip,
		IpSource:    is,
		ConnectTime: time.Now().Unix(),
		ActiveTime:  time.Now().Unix(),
		WsConn:      wsConn,
	}

	l.Infof("新的客户端连接：%+v", client)
	return client
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

func (l *WebsocketLogic) RegisterHandler(client *ws.Client) (err error) {
	l.svcCtx.WebsocketManager.RegisterClient(client)

	l.onClientInfoEvent(client)

	return l.onOnlineCountEvent(client)
}

func (l *WebsocketLogic) UnregisterHandler(client *ws.Client) (err error) {
	l.svcCtx.WebsocketManager.UnRegisterClient(client)

	return l.onOnlineCountEvent(client)
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

func (l *WebsocketLogic) onClientInfoEvent(client *ws.Client) (err error) {
	info := types.ClientInfoResp{
		ClientId:  client.ClientId,
		UserId:    client.UserId,
		DeviceId:  client.DeviceId,
		Nickname:  "",
		IpAddress: client.IpAddress,
		IpSource:  client.IpSource,
	}

	reply := &types.ReplyMsg{
		Type:      constant.ClientInfo,
		Data:      jsonconv.AnyToJsonNE(info),
		Timestamp: time.Now().Unix(),
	}

	err = l.svcCtx.WebsocketManager.PushMessage(client.ClientId, reply)
	if err != nil {
		return err
	}
	return nil
}

func (l *WebsocketLogic) onOnlineCountEvent(client *ws.Client) (err error) {
	event := &types.OnlineCountResp{
		Msg:   fmt.Sprintf("当前在线人数：%d", l.svcCtx.WebsocketManager.GetOnlineCount()),
		Count: l.svcCtx.WebsocketManager.GetOnlineCount(),
	}

	reply := &types.ReplyMsg{
		Type:      constant.OnlineCount,
		Data:      jsonconv.AnyToJsonNE(event),
		Timestamp: time.Now().Unix(),
	}

	err = l.svcCtx.WebsocketManager.PushBroadcast(reply)
	if err != nil {
		return err
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

	err = l.svcCtx.WebsocketManager.PushMessage(client.ClientId, reply)
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

	err = l.svcCtx.WebsocketManager.PushMessage(client.ClientId, reply)
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
		UserId:     client.UserId,
		TerminalId: client.ClientId,
		IpAddress:  client.IpAddress,
		IpSource:   client.IpSource,
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

	err = l.svcCtx.WebsocketManager.PushBroadcast(reply)
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

	err = l.svcCtx.WebsocketManager.PushBroadcast(reply)
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
