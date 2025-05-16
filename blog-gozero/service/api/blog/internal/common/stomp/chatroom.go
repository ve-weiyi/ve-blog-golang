package stomp

import (
	"context"
	"time"

	"github.com/go-stomp/stomp/v3/frame"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/global/constant"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/messagerpc"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/ipx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/stompws/server/client"
	"github.com/ve-weiyi/ve-blog-golang/stompws/server/topic"
)

type ChatRoomTopicHook struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	count  int64
}

func NewChatRoomTopicHook(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRoomTopicHook {
	return &ChatRoomTopicHook{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatRoomTopicHook) OnTopicSubscribe(t *topic.Topic, sub *client.Subscription) error {

	l.Logger.Infof("Topic subscribed :%s", sub.Destination())
	l.count++

	// 发送在线人数
	f := frame.New(frame.MESSAGE)
	f.Header.Set(frame.Subscription, sub.Id())
	f.Header.Set(frame.Destination, sub.Destination())
	f.Body = []byte(jsonconv.AnyToJsonNE(types.MessageEvent{
		Type: constant.OnlineCount,
		Data: jsonconv.AnyToJsonNE(types.OnlineEvent{
			Count:    l.count,
			IsOnline: true,
		}),
	}))
	// 广播
	t.Enqueue(f)

	ip := sub.NetConn().RemoteAddr().String()
	is, _ := ipx.GetIpSourceByBaidu(ip)

	// 发送客户端信息
	info := frame.New(frame.MESSAGE)
	info.Header.Set(frame.Subscription, sub.Id())
	info.Header.Set(frame.Destination, sub.Destination())
	info.Body = []byte(jsonconv.AnyToJsonNE(types.MessageEvent{
		Type: constant.ClientInfo,
		Data: jsonconv.AnyToJsonNE(types.ClientInfoEvent{
			IpAddress: ip,
			IpSource:  is,
		}),
	}))
	// 私发
	sub.SendTopicFrame(info)

	out, err := l.svcCtx.MessageRpc.FindChatList(context.Background(), &messagerpc.FindChatListReq{
		After:   time.Now().Add(-365 * 24 * time.Hour).UnixMilli(),
		Before:  time.Now().UnixMilli(),
		Limit:   10,
		UserId:  "",
		Type:    "",
		Content: "",
		Status:  0,
	})
	if err != nil {
		return err
	}

	list := make([]*types.ChatMessageEvent, 0)
	for _, msg := range out.List {
		list = append(list, &types.ChatMessageEvent{
			Id:         msg.Id,
			UserId:     msg.UserId,
			TerminalId: msg.TerminalId,
			Nickname:   msg.Nickname,
			Avatar:     msg.Avatar,
			IpAddress:  msg.IpAddress,
			IpSource:   msg.IpSource,
			Type:       msg.Type,
			Content:    msg.Content,
			Status:     msg.Status,
			CreatedAt:  msg.CreatedAt,
			UpdatedAt:  msg.UpdatedAt,
		})
	}

	// 发送历史消息
	history := frame.New(frame.MESSAGE)
	history.Header.Set(frame.Subscription, sub.Id())
	history.Header.Set(frame.Destination, sub.Destination())
	history.Body = []byte(jsonconv.AnyToJsonNE(types.MessageEvent{
		Type: constant.HistoryRecord,
		Data: jsonconv.AnyToJsonNE(types.HistoryMessageEvent{
			List: list,
		}),
	}))
	sub.SendTopicFrame(history)
	return nil
}

func (l *ChatRoomTopicHook) OnTopicUnsubscribe(t *topic.Topic, sub *client.Subscription) error {
	l.Logger.Infof("Topic unsubscribed :%s", sub.Destination())
	l.count--
	f := frame.New(frame.MESSAGE)
	f.Header.Set(frame.Subscription, sub.Id())
	f.Header.Set(frame.Destination, sub.Destination())
	f.Body = []byte(jsonconv.AnyToJsonNE(types.MessageEvent{
		Type: constant.OnlineCount,
		Data: jsonconv.AnyToJsonNE(types.OnlineEvent{
			Count:    l.count,
			IsOnline: false,
		}),
	}))
	// 广播
	t.Enqueue(f)
	return nil
}

func (l *ChatRoomTopicHook) OnTopicPublish(t *topic.Topic, f *frame.Frame) error {
	l.Logger.Infof("Topic published msg:%s", f.Command)
	var event types.MessageEvent
	jsonconv.JsonToAny(string(f.Body), &event)
	switch event.Type {
	case constant.SendMessage:
		var msg types.ChatMessageEvent
		jsonconv.JsonToAny(event.Data, &msg)
		_, err := l.svcCtx.MessageRpc.AddChat(context.Background(), &messagerpc.AddChatReq{
			UserId:     msg.UserId,
			TerminalId: msg.TerminalId,
			IpAddress:  msg.IpAddress,
			IpSource:   msg.IpSource,
			Nickname:   msg.Nickname,
			Avatar:     msg.Avatar,
			Type:       msg.Type,
			Content:    msg.Content,
			Status:     msg.Status,
			CreatedAt:  msg.CreatedAt,
			UpdatedAt:  msg.UpdatedAt,
		})
		if err != nil {
			return err
		}
	case constant.RecallMessage:
		var recall types.RecallMessageEvent
		jsonconv.JsonToAny(event.Data, &recall)
		_, err := l.svcCtx.MessageRpc.UpdateChatStatus(context.Background(), &messagerpc.UpdateChatStatusReq{
			Id:     recall.Id,
			Status: constant.ChatStatusRecall,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
