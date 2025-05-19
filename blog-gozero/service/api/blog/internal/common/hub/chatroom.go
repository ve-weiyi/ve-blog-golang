package hub

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/global/constant"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/messagerpc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/ws"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
)

type ChatRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRoomLogic {
	return &ChatRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatRoomLogic) Websocket(w http.ResponseWriter, r *http.Request) error {
	client, err := ws.Upgrade(w, r)

	if err != nil {
		l.Errorf("websocket连接失败：%+v", err)
		return err
	}

	client.SetHandleConnect(func() error {
		return l.HandleConnect(client)
	})

	client.SetHandleDisconnect(func() error {
		return l.HandleDisconnect(client)
	})

	client.SetHandleMessage(func(message []byte) error {
		return l.HandleMessage(client, message)
	})

	return client.Connect(l.svcCtx.Hub)
}

func (l *ChatRoomLogic) HandleConnect(client *ws.Client) error {
	l.Infof("websocket连接成功：%+v", client.ClientID)
	event := &types.OnlineEvent{
		Msg:      fmt.Sprintf("当前在线人数：%d", l.svcCtx.Hub.Metrics.Connections),
		Count:    l.svcCtx.Hub.Metrics.Connections,
		IsOnline: true,
	}

	reply := &types.MessageEvent{
		Type:      constant.OnlineCount,
		Data:      jsonconv.AnyToJsonNE(event),
		Timestamp: time.Now().UnixMilli(),
	}

	// 发送在线人数
	return l.svcCtx.Hub.Broadcast([]byte(jsonconv.AnyToJsonNE(reply)))
}

func (l *ChatRoomLogic) HandleDisconnect(client *ws.Client) error {
	l.Infof("websocket连接断开：%+v", client.ClientID)
	event := &types.OnlineEvent{
		Msg:      fmt.Sprintf("当前在线人数：%d", l.svcCtx.Hub.Metrics.Connections),
		Count:    l.svcCtx.Hub.Metrics.Connections,
		IsOnline: false,
	}

	reply := &types.MessageEvent{
		Type:      constant.OnlineCount,
		Data:      jsonconv.AnyToJsonNE(event),
		Timestamp: time.Now().UnixMilli(),
	}

	return l.svcCtx.Hub.Broadcast([]byte(jsonconv.AnyToJsonNE(reply)))
}

func (l *ChatRoomLogic) HandleMessage(client *ws.Client, msg []byte) (err error) {
	var receiveMsg *types.MessageEvent
	err = json.Unmarshal(msg, &receiveMsg)
	if err != nil {
		return
	}

	switch receiveMsg.Type {
	case constant.HistoryRecord:
		l.Info("历史记录")
		err = l.onHistoryRecordEvent(client)
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

	case constant.HeartBeat:
		l.Info("心跳检测")
		err = l.onHeartBeatEvent(client)
		if err != nil {
			return err
		}
	default:

	}

	return nil
}

func (l *ChatRoomLogic) onHeartBeatEvent(client *ws.Client) (err error) {
	reply := &types.MessageEvent{
		Type:      constant.HeartBeat,
		Data:      fmt.Sprintf("pong"),
		Timestamp: time.Now().UnixMilli(),
	}

	err = client.Send([]byte(jsonconv.AnyToJsonNE(reply)))
	if err != nil {
		return err
	}
	return nil
}

func (l *ChatRoomLogic) onHistoryRecordEvent(client *ws.Client) (err error) {

	in := &messagerpc.FindChatListReq{
		After:  0,
		Before: time.Now().UnixMilli(),
	}

	out, err := l.svcCtx.MessageRpc.FindChatList(context.Background(), in)
	if err != nil {
		return err
	}

	list := make([]*types.ChatMessageEvent, 0)
	for _, v := range out.List {
		m := ConvertChatTypes(v)

		list = append(list, m)
	}

	reply := &types.MessageEvent{
		Type: constant.HistoryRecord,
		Data: jsonconv.AnyToJsonNE(types.HistoryMessageEvent{
			List: list,
		}),
		Timestamp: time.Now().UnixMilli(),
	}

	err = client.Send([]byte(jsonconv.AnyToJsonNE(reply)))
	if err != nil {
		return err
	}
	return nil
}

func (l *ChatRoomLogic) onSendMessageEvent(client *ws.Client, msg *types.MessageEvent) (err error) {
	var chat types.ChatMessageEvent
	err = json.Unmarshal([]byte(msg.Data), &chat)
	if err != nil {
		return err
	}

	in := &messagerpc.AddChatReq{
		UserId:     chat.UserId,
		TerminalId: chat.TerminalId,
		IpAddress:  chat.IpAddress,
		IpSource:   chat.IpSource,
		Nickname:   chat.Nickname,
		Avatar:     chat.Avatar,
		Type:       chat.Type,
		Content:    chat.Content,
		Status:     chat.Status,
		CreatedAt:  chat.CreatedAt,
		UpdatedAt:  chat.UpdatedAt,
	}

	out, err := l.svcCtx.MessageRpc.AddChat(context.Background(), in)
	if err != nil {
		return err
	}

	reply := &types.MessageEvent{
		Type:      constant.SendMessage,
		Data:      jsonconv.AnyToJsonNE(ConvertChatTypes(out)),
		Timestamp: time.Now().UnixMilli(),
	}

	err = l.svcCtx.Hub.Broadcast([]byte(jsonconv.AnyToJsonNE(reply)))
	if err != nil {
		return err
	}

	return nil
}

func (l *ChatRoomLogic) onRecallMessageEvent(client *ws.Client, msg *types.MessageEvent) (err error) {
	var req types.RecallMessageEvent
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

	event := &types.RecallMessageEvent{
		Id: req.Id,
	}

	reply := &types.MessageEvent{
		Type:      constant.RecallMessage,
		Data:      jsonconv.AnyToJsonNE(event),
		Timestamp: time.Now().UnixMilli(),
	}

	err = l.svcCtx.Hub.Broadcast([]byte(jsonconv.AnyToJsonNE(reply)))
	if err != nil {
		return err
	}

	return nil
}

func ConvertChatTypes(in *messagerpc.ChatDetails) *types.ChatMessageEvent {
	out := &types.ChatMessageEvent{
		Id:         in.Id,
		UserId:     in.UserId,
		TerminalId: in.TerminalId,
		Nickname:   in.Nickname,
		Avatar:     in.Avatar,
		IpAddress:  in.IpAddress,
		IpSource:   in.IpSource,
		Type:       in.Type,
		Content:    in.Content,
		Status:     in.Status,
		CreatedAt:  in.CreatedAt,
		UpdatedAt:  in.UpdatedAt,
	}

	return out
}
