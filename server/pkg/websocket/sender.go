package websocket

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
)

type Client struct {
	id   string
	conn *websocket.Conn
	ch   chan []byte
}

func (c *Client) run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case msg, ok := <-c.ch:
			if !ok {
				return nil
			}
			if err := c.conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
				return err
			}
		}
	}
}

type Pusher struct {
	mutex   *sync.RWMutex
	clients map[model.UserID]*Client
}

const (
	writeWait = 20 * time.Millisecond
)

func NewPusher() *Pusher {
	return &Pusher{
		mutex:   &sync.RWMutex{},
		clients: make(map[model.UserID]*Client, 1024),
	}
}

// Send implements service.Pusher.
func (p *Pusher) Send(ctx context.Context, userID model.UserID, event []byte) error {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	client, ok := p.clients[userID]
	if !ok {
		slog.WarnContext(ctx, "client not found",
			slog.Attr{Key: "userID", Value: slog.StringValue(userID.String())},
			slog.Attr{Key: "event", Value: slog.AnyValue(event)},
		)
		return errors.New("client not found")
	}

	select {
	case client.ch <- event:
	case <-time.After(writeWait):
		return errors.New("websocket write timeout")
	}

	return nil
}

func (p *Pusher) SendAll(ctx context.Context, userIDs []model.UserID, event []byte) error {
	ch := make(chan error, len(userIDs))

	for _, userID := range userIDs {
		go func(userID model.UserID) {
			ch <- p.Send(ctx, userID, event)
		}(userID)
	}

	var err error
	for range len(userIDs) {
		if e := <-ch; e != nil {
			err = errors.Join(e)
		}
	}

	return err
}

func (p *Pusher) Register(ctx context.Context, userID model.UserID, conn *websocket.Conn) (func(ctx context.Context) error, func()) {
	id, err := uuid.NewRandom()
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate uuid", "error", err)
		return func(context.Context) error {
			return nil
		}, func() {}
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()
	client := &Client{
		id:   id.String(),
		conn: conn,
		ch:   make(chan []byte, 1000),
	}
	p.clients[userID] = client

	return client.run, func() {
		p.mutex.Lock()
		defer p.mutex.Unlock()

		delete(p.clients, userID)
		close(client.ch)
	}
}
