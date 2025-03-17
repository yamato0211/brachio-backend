package memdb

import (
	"context"
	"fmt"
	"sync"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type GameStateRepository struct {
	locks sync.Map
	db    sync.Map
}

func NewGameStateRepository() repository.GameStateRepository {
	return &GameStateRepository{
		locks: sync.Map{},
		db:    sync.Map{},
	}
}

func (r *GameStateRepository) Find(ctx context.Context, id model.RoomID) (*model.GameState, error) {
	if !r.isLocked(ctx) {
		r.lock(id)
		defer r.unlock(id)
	}

	data, ok := r.db.Load(id)
	if !ok {
		return nil, model.ErrRoomNotFound
	}

	state, ok := data.(*model.GameState)
	if !ok {
		return nil, model.ErrRoomNotFound
	}

	return state, nil
}

func (r *GameStateRepository) Save(ctx context.Context, state *model.GameState) error {
	if !r.isLocked(ctx) {
		r.lock(state.RoomID)
		defer r.unlock(state.RoomID)
	}

	r.db.Store(state.RoomID, state)

	return nil
}

func (r *GameStateRepository) Delete(ctx context.Context, id model.RoomID) error {
	if !r.isLocked(ctx) {
		r.lock(id)
		defer r.unlock(id)
	}

	r.db.Delete(id)

	return nil
}

func (r *GameStateRepository) Transaction(ctx context.Context, roomID model.RoomID, fn func(ctx context.Context) error) error {
	r.lock(roomID)
	defer r.unlock(roomID)
	ctx = r.setLock(ctx)

	return fn(ctx)
}

type lockKey struct{}

func (r *GameStateRepository) lock(roomID model.RoomID) bool {
	l := &sync.RWMutex{}
	l.Lock()
	if lock, loaded := r.locks.LoadOrStore(roomID, l); loaded {
		lock, ok := lock.(*sync.RWMutex)
		if !ok {
			fmt.Println("lock is not *sync.RWMutex")
			return false
		}
		lock.Lock()
		return true
	}
	return true
}

func (r *GameStateRepository) unlock(roomID model.RoomID) {
	lock, _ := r.locks.Load(roomID)
	safeLock, ok := lock.(*sync.RWMutex)
	if !ok {
		return
	}
	safeLock.Unlock()
}

func (r *GameStateRepository) setLock(ctx context.Context) context.Context {
	return context.WithValue(ctx, lockKey{}, struct{}{})
}

func (r *GameStateRepository) isLocked(ctx context.Context) bool {
	return ctx.Value(lockKey{}) != nil
}
