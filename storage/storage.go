package storage

import (
	"context"
	"errors"
)

var ErrEmpty = errors.New("No notifications found")

type Storage interface {
	Push(ctx context.Context, clientID int, notification entity.Notification) error
}
