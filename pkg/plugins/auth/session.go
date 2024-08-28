package auth

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session interface {
	ID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	WithID(id primitive.ObjectID) Session

	App() primitive.ObjectID
	SetApp(app primitive.ObjectID)
	WithApp(app primitive.ObjectID) Session

	User() primitive.ObjectID
	SetUser(user primitive.ObjectID)
	WithUser(user primitive.ObjectID) Session

	InitialRemoteAddress() string
	SetInitialRemoteAddress(initialRemoteAddress string)
	WithInitialRemoteAddress(initialRemoteAddress string) Session

	RemoteAddress() string
	SetRemoteAddress(remoteAddress string)
	WithRemoteAddress(remoteAddress string) Session
	Location() (Location, error)

	UserAgent() *string
	SetUserAgent(userAgent string)
	UnsetUserAgent()
	WithUserAgent(userAgent string) Session
	WithoutUserAgent() Session
	UserAgentParsed() UserAgent

	LastActive() time.Time
	SetLastActive(lastActive time.Time)
	WithLastActive(lastActive time.Time) Session

	Expires() time.Time
	SetExpires(expires time.Time)
	WithExpires(expires time.Time) Session
	Expire(ctx context.Context, expires time.Duration) error

	Created() time.Time
	SetCreated(created time.Time)
	WithCreated(created time.Time) Session

	Updated() time.Time
	SetUpdated(updated time.Time)
	WithUpdated(updated time.Time) Session
}

type ListSessionsOptions interface {
	App() *primitive.ObjectID
	SetApp(app primitive.ObjectID)
	UnsetApp()
	WithApp(app primitive.ObjectID) ListSessionsOptions
	WithoutApp() ListSessionsOptions

	User() *primitive.ObjectID
	SetUser(user primitive.ObjectID)
	UnsetUser()
	WithUser(user primitive.ObjectID) ListSessionsOptions
	WithoutUser() ListSessionsOptions

	Offset() *int64
	SetOffset(offset int64)
	UnsetOffset()
	WithOffset(offset int64) ListSessionsOptions
	WithoutOffset() ListSessionsOptions

	Count() *int64
	SetCount(count int64)
	UnsetCount()
	WithCount(count int64) ListSessionsOptions
	WithoutCount() ListSessionsOptions
}

type ListSessionsResponse interface {
	Count() int64
	SetCount(count int64)
	WithCount(count int64) ListSessionsResponse

	Page() []Session
	SetPage(page []Session)
	WithPage(page []Session) ListSessionsResponse
}