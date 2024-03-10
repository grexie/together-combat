package plugin

type TogetherRole = string

const (
	TogetherRoleOwner TogetherRole = "owner"
	TogetherRoleAdmin TogetherRole = "admin"
	TogetherRoleManager TogetherRole = "manager"
	TogetherRoleDeveloper TogetherRole = "developer"
	TogetherRoleSupport TogetherRole = "support"
	TogetherRoleUser TogetherRole = "user"
	TogetherRoleOrg TogetherRole = "org"
	TogetherRoleConnected TogetherRole = "connected"
	TogetherRoleLogin TogetherRole = "login"
)

type TokenType = string

const (
	TokenTypeApp TokenType = "app"
	TokenTypeLogin TokenType = "login"
	TokenTypeChallenge TokenType = "challenge"
	TokenTypeAuth TokenType = "auth"
	TokenTypeRefresh TokenType = "refresh"
	TokenTypeAccess TokenType = "access"
)

type Permission = string

const (
	TogetherPermissionOwner Permission = "together:owner"
	TogetherPermissionAdmin Permission = "together:admin"
	TogetherPermissionDeveloper Permission = "together:developer"
	TogetherPermissionManager Permission = "together:manager"
	TogetherPermissionSupport Permission = "together:support"
	TogetherPermissionUser Permission = "together:user"
	TogetherPermissionLogin Permission = "together:login"
)

type AuthOptions struct {
	Together *bool
	Roles []TogetherRole
	Permissions []Permission
	TokenTypes []TokenType
}

func NewAuthOptions() AuthOptions {
	return AuthOptions{
		Together: nil,
		Roles: []TogetherRole{},
		Permissions: []Permission{},
		TokenTypes: []TokenType{TokenTypeApp, TokenTypeAccess},
	}
}

func (o AuthOptions) WithTogether() AuthOptions {
	_true := true
	o.Together = &_true
	return o
}

func (o AuthOptions) WithoutTogether() AuthOptions {
	_false := false
	o.Together = &_false
	return o
}

func (o AuthOptions) WithRoles(r ...TogetherRole) AuthOptions {
	o.Roles = r
	return o
}

func (o AuthOptions) WithPermissions(p ...Permission) AuthOptions {
	o.Permissions = p
	return o
}

func (o AuthOptions) WithTokenTypes(t ...TokenType) AuthOptions {
	o.TokenTypes = t
	return o
}