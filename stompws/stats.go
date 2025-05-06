package stompws

import (
	"sync"
)

type StatsCollector struct {
	connections      int64
	subscriptions    int64
	messagesSent     int64
	messagesReceived int64
	messagesDropped  int64
	errors           int64
	mu               sync.Mutex
}

func NewStatsCollector() *StatsCollector {
	return &StatsCollector{}
}

func (s *StatsCollector) IncrementConnections() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.connections++
}

func (s *StatsCollector) DecrementConnections() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.connections--
}

func (s *StatsCollector) IncrementSubscriptions() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.subscriptions++
}

func (s *StatsCollector) DecrementSubscriptions() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.subscriptions--
}

func (s *StatsCollector) IncrementMessagesSent() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messagesSent++
}

func (s *StatsCollector) IncrementMessagesReceived() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messagesReceived++
}

func (s *StatsCollector) IncrementMessagesDropped() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messagesDropped++
}

func (s *StatsCollector) IncrementErrors() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.errors++
}
