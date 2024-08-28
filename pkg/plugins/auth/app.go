package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type App interface {
	ID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	WithID(id primitive.ObjectID) App

	Name() string
	SetName(name string)
	WithName(name string) App

	Website() string
	SetWebsite(website string)
	WithWebsite(website string) App

	Description() string
	SetDescription(description string)
	WithDescription(description string) App

	LogoURL() string
	SetLogoURL(logoUrl string)
	WithLogoURL(logoUrl string) App

	AvatarURL() string
	SetAvatarURL(avatarUrl string)
	WithAvatarURL(avatarUrl string) App

	ReferrerURLs() []string
	SetReferrerURLs(referrerUrls []string)
	WithReferrerURLs(referrerUrls []string) App

	RedirectURL() string
	SetRedirectURL(redirectUrl string)
	WithRedirectURL(redirectUrl string) App

	Public() bool
	SetPublic(public bool)
	WithPublic(public bool) App

	Anonymization() AppAnonymization
	SetAnonymization(anonymization AppAnonymization)
	WithAnonymization(anonymization AppAnonymization) App

	Enabled() bool
	SetEnabled(enabled bool)
	WithEnabled(enabled bool) App

	Locked() bool
	SetLocked(locked bool)
	WithLocked(locked bool) App

	Created() time.Time
	SetCreated(created time.Time)
	WithCreated(created time.Time) App

	Updated() time.Time
	SetUpdated(updated time.Time)
	WithUpdated(updated time.Time) App

	IsTogether() bool
	SetTogether(together bool)
	WithTogether(together bool) App

	User() User
	SetUser(user User)
	WithUser(user User) App
}

type CreateAppOptions interface {
	Name() string
	SetName(name string)
	WithName(name string) CreateAppOptions

	Website() string
	SetWebsite(website string)
	WithWebsite(website string) CreateAppOptions

	Description() string
	SetDescription(description string)
	WithDescription(description string) CreateAppOptions

	LogoURL() string
	SetLogoURL(logoUrl string)
	WithLogoURL(logoUrl string) CreateAppOptions

	AvatarURL() string
	SetAvatarURL(avatarUrl string)
	WithAvatarURL(avatarUrl string) CreateAppOptions

	ReferrerURLs() []string
	SetReferrerURLs(referrerUrls []string)
	WithReferrerURLs(referrerUrls []string) CreateAppOptions

	RedirectURL() string
	SetRedirectURL(redirectUrl string)
	WithRedirectURL(redirectUrl string) CreateAppOptions

	Public() *bool
	SetPublic(public bool)
	UnsetPublic()
	WithPublic(public bool) CreateAppOptions
	WithoutPublic() CreateAppOptions

	Enabled() *bool
	SetEnabled(enabled bool)
	UnsetEnabled()
	WithEnabled(enabled bool) CreateAppOptions
	WithoutEnabled() CreateAppOptions
}

type UpdateAppOptions interface {
	Name() *string
	SetName(name string)
	UnsetName()
	WithName(name string) UpdateAppOptions
	WithoutName() UpdateAppOptions

	Website() *string
	SetWebsite(website string)
	UnsetWebsite()
	WithWebsite(website string) UpdateAppOptions
	WithoutWebsite() UpdateAppOptions

	Description() *string
	SetDescription(description string)
	UnsetDescription()
	WithDescription(description string) UpdateAppOptions
	WithoutDescription() UpdateAppOptions

	LogoURL() *string
	SetLogoURL(logoUrl string)
	UnsetLogoURL()
	WithLogoURL(logoUrl string) UpdateAppOptions
	WithoutLogoURL() UpdateAppOptions

	AvatarURL() *string
	SetAvatarURL(avatarUrl string)
	UnsetAvatarURL()
	WithAvatarURL(avatarUrl string) UpdateAppOptions
	WithoutAvatarURL() UpdateAppOptions

	ReferrerURLs() *[]string
	SetReferrerURLs(referrerUrls []string)
	UnsetReferrerURLs()
	WithReferrerURLs(referrerUrls []string) UpdateAppOptions
	WithoutReferrerURLs() UpdateAppOptions

	RedirectURL() *string
	SetRedirectURL(redirectUrl string)
	UnsetRedirectURL()
	WithRedirectURL(redirectUrl string) UpdateAppOptions
	WithoutRedirectURL() UpdateAppOptions

	Public() *bool
	SetPublic(public bool)
	UnsetPublic()
	WithPublic(public bool) UpdateAppOptions
	WithoutPublic() UpdateAppOptions

	Enabled() *bool
	SetEnabled(enabled bool)
	UnsetEnabled()
	WithEnabled(enabled bool) UpdateAppOptions
	WithoutEnabled() UpdateAppOptions

	Updated() *time.Time
	SetUpdated(updated time.Time)
	UnsetUpdated()
	WithUpdated(updated time.Time) UpdateAppOptions
	WithoutUpdated() UpdateAppOptions
}

type ListAppsResponse interface {
	Count() int64
	SetCount(count int64)
	WithCount(count int64) ListAppsResponse

	Page() []App
	SetPage(page []App)
	WithPage(page []App) ListAppsResponse
}

type ListAppsOptions interface {
	Public() *bool
	SetPublic(public bool)
	UnsetPublic()
	WithPublic(public bool) ListAppsOptions
	WithoutPublic() ListAppsOptions

	User() *primitive.ObjectID
	SetUser(user primitive.ObjectID)
	UnsetUser()
	WithUser(user primitive.ObjectID) ListAppsOptions
	WithoutUser() ListAppsOptions

	Offset() *int64
	SetOffset(offset int64)
	UnsetOffset()
	WithOffset(offset int64) ListAppsOptions
	WithoutOffset() ListAppsOptions

	Count() *int64
	SetCount(count int64)
	UnsetCount()
	WithCount(count int64) ListAppsOptions
	WithoutCount() ListAppsOptions
}