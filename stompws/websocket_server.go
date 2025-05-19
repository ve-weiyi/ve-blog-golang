/*
Package server contains a simple STOMP server implementation.
*/
package stompws

import (
	"net/http"
	"strings"

	"github.com/go-stomp/stomp/v3/frame"
	"github.com/gorilla/websocket"

	"github.com/ve-weiyi/ve-blog-golang/stompws/server/client"
	"github.com/ve-weiyi/ve-blog-golang/stompws/server/queue"
	"github.com/ve-weiyi/ve-blog-golang/stompws/server/topic"
)

// A Server defines parameters for running a STOMP server.
type StompWebsocketServer struct {
	Config

	ch   chan client.Request
	tm   *topic.Manager
	qm   *queue.Manager
	hm   *HookManager
	stop bool // has stop been requested
}

func NewWebsocketServer(c Config) *StompWebsocketServer {
	return &StompWebsocketServer{
		Config: c,
		ch:     make(chan client.Request, 128),
		tm:     topic.NewManager(),
		qm:     queue.NewManager(queue.NewMemoryQueueStorage()),
		hm:     NewHookManager(),
		stop:   false,
	}
}

func (s *StompWebsocketServer) Run() {
	for {
		r := <-s.ch
		switch r.Op {
		case client.SubscribeOp:
			s.Log.Infof("stomp: subscribe to %s", r.Sub.Destination())
			if isQueueDestination(r.Sub.Destination()) {
				q := s.qm.Find(r.Sub.Destination())
				// todo error handling
				q.Subscribe(r.Sub)
			} else {
				t := s.tm.Find(r.Sub.Destination())
				t.Subscribe(r.Sub)
				for _, h := range s.hm.Find(r.Sub.Destination()) {
					h.OnTopicSubscribe(t, r.Sub)
				}
			}

		case client.UnsubscribeOp:
			s.Log.Infof("stomp: unsubscribe from %s", r.Sub.Destination())
			if isQueueDestination(r.Sub.Destination()) {
				q := s.qm.Find(r.Sub.Destination())
				// todo error handling
				q.Unsubscribe(r.Sub)
			} else {
				t := s.tm.Find(r.Sub.Destination())
				t.Unsubscribe(r.Sub)
				for _, h := range s.hm.Find(r.Sub.Destination()) {
					h.OnTopicUnsubscribe(t, r.Sub)
				}
			}

		case client.EnqueueOp:
			s.Log.Infof("stomp: enqueue to %s", r.Frame.Header.Get(frame.Destination))
			destination, ok := r.Frame.Header.Contains(frame.Destination)
			if !ok {
				// should not happen, already checked in lower layer
				panic("missing destination")
			}

			if isQueueDestination(destination) {
				q := s.qm.Find(destination)
				q.Enqueue(r.Frame)
			} else {
				t := s.tm.Find(destination)
				t.Enqueue(r.Frame)
				for _, h := range s.hm.Find(destination) {
					h.OnTopicPublish(t, r.Frame)
				}
			}

		case client.RequeueOp:
			s.Log.Infof("stomp: requeue to %s", r.Frame.Header.Get(frame.Destination))
			destination, ok := r.Frame.Header.Contains(frame.Destination)
			if !ok {
				// should not happen, already checked in lower layer
				panic("missing destination")
			}

			// only requeue to queues, should never happen for topics
			if isQueueDestination(destination) {
				q := s.qm.Find(destination)
				q.Requeue(r.Frame)
			}
		}
	}
}

func (s *StompWebsocketServer) HandleWebSocket(w http.ResponseWriter, r *http.Request) error {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{"stomp", "mqtt", "v12.stomp", "v11.stomp", "v10.stomp"},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	wsConn := NewWebSocketConn(conn)
	_ = client.NewConn(s, wsConn, s.ch)
	return nil
}

func (s *StompWebsocketServer) RegisterTopicHooks(topic string, hooks ...TopicHook) {
	s.hm.RegisterHooks(topic, hooks...)
}

const QueuePrefix = "/queue"

func isQueueDestination(dest string) bool {
	return strings.HasPrefix(dest, QueuePrefix)
}
