package transmit

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

type Session interface {
	ID() string
	Close()
}

var sid int64 = 0

func NewSessionID(prex string) string {
	return fmt.Sprintf("%s_%d_%d", prex, atomic.AddInt64(&sid, 1), time.Now().Unix())
}

type MessageType = int

const (
	MessageType_Method MessageType = 1
	MessageType_Async  MessageType = 2
)

type SessionStat = int32

const (
	SessionConnected    SessionStat = 1
	SessionDisconnected SessionStat = 2
)

type sessionKey struct{}

func SessionFromContext(ctx context.Context) (Session, bool) {
	s, ok := ctx.Value(sessionKey{}).(Session)
	return s, ok
}

func NewSessionContext(ctx context.Context, s Session) context.Context {
	return context.WithValue(ctx, sessionKey{}, s)
}
