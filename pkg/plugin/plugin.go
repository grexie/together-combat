package plugin

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TogetherPlugin interface {
	Initialize(ctx context.Context) error
	Context(ctx context.Context) (context.Context, error)
	TestContext(ctx context.Context) (context.Context, error)
}

type TogetherCommand func(ctx context.Context, args ...string) error

type TogetherServer interface {
	http.Handler
	Now() time.Time
	SetTime(time.Time)
	AdvanceTime(time.Duration)
	ResetTime()

	App() *fiber.App

	RequestContext(context.Context, *http.Request) (context.Context, error)

	LookupPlugin(name string, export string) (any, error)

	InitialURL() string
	SetInitialURL(initialUrl string)

	AnonymizeField(alias string) graphql.FieldResolveFn
	Anonymize(alias string, ctx context.Context, id *primitive.ObjectID) (*string, error)
	MustAnonymize(alias string, ctx context.Context, id primitive.ObjectID) string
	Deanonymize(alias string, ctx context.Context, id *string) (*primitive.ObjectID, error)
	Auth(options AuthOptions, fn ...graphql.FieldResolveFn) graphql.FieldResolveFn

	RegisterCommand(name string, command TogetherCommand)
	ExecuteCommand(ctx context.Context, name string, args ...string) error
	ShowHelp(executableName string) error

	DB() *mongo.Database
	Redis() *redis.Client

	Schema() *graphql.Schema

	Input(name string) *graphql.InputObject
	AddInput(*graphql.InputObject)
	Type(name string) *graphql.Object
	AddType(*graphql.Object)
	OnTypes([]string, func())

	Void() *graphql.Scalar
	JSON() *graphql.Scalar
	Byte() *graphql.Scalar

	UpdateType() *graphql.Enum

	RootQuery() *graphql.Object
	RootMutation() *graphql.Object
	RootSubscription() *graphql.Object
}

type TogetherContextType string
const (
	TogetherServerContext TogetherContextType = "@grexie/together:server-context"
	TogetherRequestContext TogetherContextType = "@grexie/together:request-context"
	TogetherWebSocketContext TogetherContextType = "@grexie/together:websocket-context"
)

func WithRequest(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, TogetherRequestContext, r)
}

func WithServer(ctx context.Context, s TogetherServer) context.Context {
	return context.WithValue(ctx, TogetherServerContext, s)
}

func WithWebSocket(ctx context.Context) context.Context {
	return context.WithValue(ctx, TogetherWebSocketContext, true)
}

func GetServer(ctx context.Context) TogetherServer {
	if server := ctx.Value(TogetherServerContext); server == nil {
		log.Fatal(fmt.Errorf("TogetherServer not found in context"))
		return nil
	} else if s, ok := server.(TogetherServer); !ok {
		log.Fatal(fmt.Errorf("TogetherServer not found in context, was another interface: %s", server))
		return nil
	} else {
		return s
	}
}

func GetRequest(ctx context.Context) *http.Request {
	if request := ctx.Value(TogetherRequestContext); request == nil {
		log.Fatal(fmt.Errorf("http.Request not found in context"))
		return nil
	} else if r, ok := request.(*http.Request); !ok {
		log.Fatal(fmt.Errorf("http.Request not found in context, was another interface: %s", request))
		return nil
	} else {
		return r
	}
}

func IsWebSocket(ctx context.Context) bool {
	if ws := ctx.Value(TogetherWebSocketContext); ws == nil {
		return false
	} else {
		return ws.(bool)
	}
}