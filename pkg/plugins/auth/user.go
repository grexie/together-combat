package auth

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Name interface {
	First() string
	SetFirst(first string)
	WithFirst(first string) Name

	Last() string
	SetLast(last string)
	WithLast(last string) Name
}

type User interface {
	ID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	WithID(id primitive.ObjectID) User

	Email() string
	SetEmail(email string)
	WithEmail(email string) User
	
	Password() string
	SetPassword(password string) error
	WithPassword(password string) (User, error)

	BirthDate() time.Time
	SetBirthDate(birthDate time.Time)
	WithBirthDate(birthDate time.Time) User

	AvatarURL() string
	SetAvatarURL(avatarUrl string)
	WithAvatarURL(avatarUrl string) User

	Name() Name
	SetName(name Name)
	WithName(name Name) User

	Enabled() bool
	SetEnabled(enabled bool)
	WithEnabled(enabled bool) User

	Created() time.Time
	SetCreated(created time.Time)
	WithCreated(created time.Time) User

	Updated() time.Time
	SetUpdated(updated time.Time)
	WithUpdated(updated time.Time) User

	KeyEncryptingKey() *[]byte
	KeyEncryptingKeySalt() *[]byte
	KeyEncryptingKeyRaw() *string
	SetKeyEncryptingKey(keyEncryptingKeySalt []byte, keyEncryptingKey []byte)
	UnsetKeyEncryptingKey()
	WithKeyEncryptingKey(keyEncryptingKeySalt []byte, keyEncryptingKey []byte) User
	WithoutKeyEncryptingKey() User

	App() App
	SetApp(app App)
	UnsetApp()
	WithApp(app App) User
	WithoutApp() User

	SendCampaign(ctx context.Context, campaign Campaign, data any) error
	SendCampaignAction(ctx context.Context, action CampaignAction, data any) error
}

type CreateUserOptions interface {
	Email() string
	SetEmail(email string)
	WithEmail(email string) CreateUserOptions

	Password() string
	SetPassword(password string)
	WithPassword(password string) CreateUserOptions

	AvatarURL() string
	SetAvatarURL(avatarUrl string)
	WithAvatarURL(avatarUrl string) CreateUserOptions

	Name() Name
	SetName(name Name)
	WithName(name Name) CreateUserOptions

	BirthDate() time.Time
	SetBirthDate(birthDate time.Time)
	WithBirthDate(birthDate time.Time) CreateUserOptions
}

type UpdateUserOptions interface {
	Email() *string
	SetEmail(email string)
	UnsetEmail()
	WithEmail(email string) UpdateUserOptions
	WithoutEmail() UpdateUserOptions

	Password() *string
	SetPassword(password string)
	UnsetPassword()
	WithPassword(password string) UpdateUserOptions
	WithoutPassword() UpdateUserOptions

	AvatarURL() *string
	SetAvatarURL(avatarUrl string)
	UnsetAvatarURL()
	WithAvatarURL(avatarUrl string) UpdateUserOptions
	WithoutAvatarURL() UpdateUserOptions

	Name() Name
	SetName(name Name)
	UnsetName()
	WithName(name Name) UpdateUserOptions
	WithoutName() UpdateUserOptions

	BirthDate() *time.Time
	SetBirthDate(birthDate time.Time)
	UnsetBirthDate()
	WithBirthDate(birthDate time.Time) UpdateUserOptions
	WithoutBirthDate() UpdateUserOptions

	Updated() *time.Time
	SetUpdated(updated time.Time)
	UnsetUpdated()
	WithUpdated(updated time.Time) UpdateUserOptions
	WithoutUpdated() UpdateUserOptions
}

type ListUsersResponse interface {
	Count() int64
	SetCount(count int64)
	WithCount(count int64) ListUsersResponse

	Page() []User
	SetPage(page []User)
	WithPage(page []User) ListUsersResponse
}

type ListUsersOptions interface {
	App() *primitive.ObjectID
	SetApp(app primitive.ObjectID)
	UnsetApp()
	WithApp(app primitive.ObjectID) ListUsersOptions
	WithoutApp() ListUsersOptions

	OrganisationalUsers() *bool
	SetOrganisationalUsers(organisationalUsers bool)
	UnsetOrganisationalUsers()
	WithOrganisationalUsers(organisationalUsers bool) ListUsersOptions
	WithoutOrganisationalUsers() ListUsersOptions

	Roles() *[]primitive.ObjectID
	SetRoles(roles []primitive.ObjectID)
	UnsetRoles()
	WithRoles(roles []primitive.ObjectID) ListUsersOptions
	WithoutRoles() ListUsersOptions

	Offset() *int64
	SetOffset(offset int64)
	UnsetOffset()
	WithOffset(offset int64) ListUsersOptions
	WithoutOffset() ListUsersOptions

	Count() *int64
	SetCount(count int64)
	UnsetCount()
	WithCount(count int64) ListUsersOptions
	WithoutCount() ListUsersOptions
}