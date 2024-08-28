package auth

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token interface {
	ID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	WithID(id primitive.ObjectID) Token

	Type() TokenType
	SetType(tokenType TokenType)
	WithType(tokenType TokenType) Token

	App() primitive.ObjectID
	SetApp(app primitive.ObjectID)
	WithApp(app primitive.ObjectID) Token

	User() *primitive.ObjectID
	SetUser(user primitive.ObjectID)
	UnsetUser()
	WithUser(user primitive.ObjectID) Token
	WithoutUser() Token

	Session() *primitive.ObjectID
	SetSession(session primitive.ObjectID)
	UnsetSession()
	WithSession(session primitive.ObjectID) Token
	WithoutSession() Token

	Roles() *[]primitive.ObjectID
	SetRoles(roles []primitive.ObjectID)
	UnsetRoles()
	WithRoles(roles []primitive.ObjectID) Token
	WithoutRoles() Token

	TokenHash() string
	SetTokenHash(tokenHash string)
	WithTokenHash(tokenHash string) Token

	Expires() *time.Time
	SetExpires(expires time.Time)
	UnsetExpires()
	WithExpires(expires time.Time) Token
	WithoutExpires() Token
	Expire(ctx context.Context, expires *time.Duration) (TokenWithJWT, error) 

	Created() time.Time
	SetCreated(created time.Time)
	WithCreated(created time.Time) Token

	Updated() time.Time
	SetUpdated(updated time.Time)
	WithUpdated(updated time.Time) Token
}

type TokenVerifyResult interface {
	App() App
	SetApp(app App)
	WithApp(app App) TokenVerifyResult

	User() User
	SetUser(user User)
	WithUser(user User) TokenVerifyResult

	Session() Session
	SetSession(session Session)
	WithSession(session Session) TokenVerifyResult

	Token() Token
	SetToken(token Token)
	WithToken(token Token) TokenVerifyResult
}

type TokenWithJWT interface {
	Token

	SetToken(token Token)
	WithToken(token Token) TokenWithJWT

	JWT() string
	SetJWT(jwt string)
	WithJWT(jwt string) TokenWithJWT
}

type AuthToken interface {
	AuthToken() string
	SetAuthToken(authToken string)
	WithAuthToken(authToken string) AuthToken

	RedirectURL() string
	SetRedirectURL(redirectUrl string)
	WithRedirectURL(redirectUrl string) AuthToken 
}

type ExchangeToken interface {
	RedirectURL() string
	SetRedirectURL(redirectUrl string)
	WithRedirectURL(redirectUrl string) ExchangeToken

	Tokens() AccessTokenExchange
	SetTokens(tokens AccessTokenExchange)
	WithTokens(tokens AccessTokenExchange) ExchangeToken
}

type AccessToken interface {
	AccessToken() string
	SetAccessToken(accessToken string)
	WithAccessToken(accessToken string) AccessToken

	RefreshToken() string
	SetRefreshToken(refreshToken string)
	WithRefreshToken(refreshToken string) AccessToken
}

type AccessTokenExchange interface {
	AccessToken() string
	SetAccessToken(accessToken string)
	WithAccessToken(accessToken string) AccessTokenExchange

	RefreshToken() string
	SetRefreshToken(refreshToken string)
	WithRefreshToken(refreshToken string) AccessTokenExchange

	KeyManagementToken() string
	SetKeyManagementToken(keyManagementToken string)
	WithKeyManagementToken(keyManagementToken string) AccessTokenExchange
}

type LoginToken interface {
	Token() string
	SetToken(token string)
	WithToken(token string) LoginToken
}

type AppToken interface {
	Token() string
	SetToken(token string)
	WithToken(token string) AppToken
}