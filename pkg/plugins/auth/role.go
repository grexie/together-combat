package auth

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role interface {
	ID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	WithID(id primitive.ObjectID) Role

	App() primitive.ObjectID
	SetApp(app primitive.ObjectID)
	WithApp(app primitive.ObjectID) Role

	Name() string
	SetName(name string)
	WithName(name string) Role

	Permissions() []Permission
	SetPermissions(permissions []Permission)
	WithPermissions(permissions []Permission) Role
	HasTogetherPermission() bool
	PermissionsSet() PermissionsSet
	HasPermission(permission Permission) bool

	IsUserRole() bool
	SetUserRole(userRole bool)
	WithUserRole(userRole bool) Role

	Enabled() bool
	SetEnabled(enabled bool)
	WithEnabled(enabled bool) Role

	Created() time.Time
	SetCreated(created time.Time)
	WithCreated(created time.Time) Role

	Updated() time.Time
	SetUpdated(updated time.Time)
	WithUpdated(updated time.Time) Role

	AddUser(ctx context.Context, user primitive.ObjectID) (bool, error)
	RemoveUser(ctx context.Context, user primitive.ObjectID) (bool, error)
}

type CreateRoleOptions interface {
	Name() string
	SetName(name string)
	WithName(name string) CreateRoleOptions

	Permissions() []Permission
	SetPermissions(permissions []Permission)
	WithPermissions(permissions []Permission) CreateRoleOptions
	HasTogetherPermission() bool
  HasPermission(permission Permission) bool
  HasInvalidPermission() bool
  HasDuplicatePermissions() bool

	IsUserRole() bool
	SetUserRole(userRole bool)
	WithUserRole(userRole bool) CreateRoleOptions
}

type UpdateRoleOptions interface {
	Name() *string
	SetName(name string)
	UnsetName()
	WithName(name string) UpdateRoleOptions
	WithoutName() UpdateRoleOptions

	Permissions() *[]Permission
	SetPermissions(permissions []Permission)
	UnsetPermissions()
	WithPermissions(permissions []Permission) UpdateRoleOptions
	WithoutPermissions() UpdateRoleOptions
	IsAddingTogetherPermission(actOnRole Role) bool
	IsRemovingTogetherPermission(actOnRole Role) bool
	HasPermission(permission Permission) bool
	HasInvalidPermission() bool
	HasDuplicatePermissions() bool

	IsUserRole() *bool
	SetUserRole(userRole bool)
	UnsetUserRole()
	WithUserRole(userRole bool) UpdateRoleOptions
	WithoutUserRole() UpdateRoleOptions

	Updated() *time.Time
	SetUpdated(updated time.Time)
	UnsetUpdated()
	WithUpdated(updated time.Time) UpdateRoleOptions
	WithoutUpdated() UpdateRoleOptions
}

type ListRolesResponse interface {
	Count() int64
	SetCount(count int64)
	WithCount(count int64) ListRolesResponse

	Page() []Role
	SetPage(page []Role)
	WithPage(page []Role) ListRolesResponse

	HasPermission(permission Permission) bool
	Has(id primitive.ObjectID) bool
}

type ListRolesOptions interface {
	App() *primitive.ObjectID
	SetApp(app primitive.ObjectID)
	UnsetApp()
	WithApp(app primitive.ObjectID) ListRolesOptions
	WithoutApp() ListRolesOptions

	User() *primitive.ObjectID
	SetUser(user primitive.ObjectID)
	UnsetUser()
	WithUser(user primitive.ObjectID) ListRolesOptions
	WithoutUser() ListRolesOptions

	Roles() *[]primitive.ObjectID
	SetRoles(roles []primitive.ObjectID)
	UnsetRoles()
	WithRoles(roles []primitive.ObjectID) ListRolesOptions
	WithoutRoles() ListRolesOptions

	Offset() *int64
	SetOffset(offset int64)
	UnsetOffset()
	WithOffset(offset int64) ListRolesOptions
	WithoutOffset() ListRolesOptions

	Count() *int64
	SetCount(count int64)
	UnsetCount()
	WithCount(count int64) ListRolesOptions
	WithoutCount() ListRolesOptions
}