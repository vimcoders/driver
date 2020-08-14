package driver

import "time"

type Session interface {
	GetUserID() int64
	GetExpire() time.Time
	OnMessage(c Context)
	SendMessage(v interface{})
	Close() error
}

type SessionMgr interface {
	Get(userID int64) Session
	Del(userID int64)
	Add(session Session)
}
