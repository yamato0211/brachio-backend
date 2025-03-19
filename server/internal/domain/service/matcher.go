package service

import (
	"context"
	"sync"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
)

type MatcherCallbackFunc func(model.RoomID)

type Matcher interface {
	Apply(ctx context.Context, password string, callback MatcherCallbackFunc) error
}

type MatcherService struct {
	lock sync.RWMutex
	mp   map[string]MatcherCallbackFunc
}

func NewMatcherService() Matcher {
	return &MatcherService{
		lock: sync.RWMutex{},
		mp:   make(map[string]MatcherCallbackFunc),
	}
}

func (s *MatcherService) Apply(ctx context.Context, password string, callback MatcherCallbackFunc) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if f, ok := s.mp[password]; ok {
		delete(s.mp, password)

		roomID := model.NewRoomID()

		f(roomID)
		callback(roomID)

		return nil
	}

	s.mp[password] = callback

	return nil
}
