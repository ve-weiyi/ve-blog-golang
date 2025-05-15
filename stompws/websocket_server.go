/*
Package server contains a simple STOMP server implementation.
*/
package stompws

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-stomp/stomp/v3/frame"
	"github.com/go-stomp/stomp/v3/server"
	"github.com/gorilla/websocket"

	"github.com/go-stomp/stomp/v3"
	"github.com/go-stomp/stomp/v3/server/client"
	"github.com/go-stomp/stomp/v3/server/queue"
	"github.com/go-stomp/stomp/v3/server/topic"
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
				s.Log.Infof("stomp: SubscribeOp to %s", r.Sub.Destination())
				for _, h := range s.hm.Find(r.Sub.Destination()) {
					s.Log.Infof("stomp: OnTopicSubscribe to %s", r.Sub.Destination())
					h.OnTopicSubscribe(r.Sub)
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
					h.OnTopicUnsubscribe(r.Sub)
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
				find := s.tm.Find(destination)
				find.Enqueue(r.Frame)
				for _, h := range s.hm.Find(destination) {
					h.OnTopicPublish(r.Sub, r.Frame)
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

func (s *StompWebsocketServer) RegisterTopicHooks(topic string, hooks ...TopicHook) {
	s.hm.RegisterHooks(topic, hooks...)
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
		return nil
	}

	wsConn := &WebSocketConn{Conn: conn}
	_ = client.NewConn(s, wsConn, s.ch)
	return nil
}

type Config struct {
	Authenticator server.Authenticator // Authenticates login/passcodes. If nil no authentication is performed
	QueueStorage  server.QueueStorage  // Implementation of queue storage. If nil, in-memory queues are used.
	HeartBeatTime time.Duration        // Preferred value for heart-beat read/write timeout, if zero, then DefaultHeartBeat.
	Log           stomp.Logger
}

func (c *Config) HeartBeat() time.Duration {
	if c.HeartBeatTime == time.Duration(0) {
		return 10 * time.Second
	}
	return c.HeartBeatTime
}

func (c *Config) Authenticate(login, passcode string) bool {
	if c.Authenticator != nil {
		return c.Authenticator.Authenticate(login, passcode)
	}

	// no authentication defined
	return true
}

func (c *Config) Logger() stomp.Logger {
	return c.Log
}

const QueuePrefix = "/queue"

func isQueueDestination(dest string) bool {
	return strings.HasPrefix(dest, QueuePrefix)
}
