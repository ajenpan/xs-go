package transmit

import "context"

type OnMethodMessage func(context.Context, interface{}) (interface{}, error)
type OnAsyncMessage func(context.Context, interface{}) error
type OnSessionStat func(context.Context, SessionStat)
