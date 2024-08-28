package auth

import (
	"strings"
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

type PermissionsSet map[string]bool

func FromPermissionsSlice(permissions []Permission) PermissionsSet {
	set := PermissionsSet{}
	for _, p := range permissions {
		set.Add(p)
	}
	return set
}

func (s PermissionsSet) String() string {
	permissions := []string{}
	for k := range s {
		permissions = append(permissions, k)
	}
	return strings.Join(permissions, ", ")
}

func (s PermissionsSet) ToSlice() []Permission {
	out := []Permission{}
	for k := range s {
		out = append(out, k)
	}
	return out
}

func (s PermissionsSet) Has(p Permission) bool {
	if _, ok := s[p]; ok {
		return true
	} else {
		return false
	}
}

func (s PermissionsSet) Add(p Permission) {
	s[p] = true
}

func (s PermissionsSet) Remove(p Permission) {
	delete(s, p)
}

func (s PermissionsSet) IsOwner() bool {
	return s.Has(TogetherPermissionOwner)
}

func (s PermissionsSet) IsAdmin() bool {
	return s.Has(TogetherPermissionAdmin) || s.IsOwner()
}

func (s PermissionsSet) IsDeveloper() bool {
	return s.Has(TogetherPermissionDeveloper) || s.IsAdmin()
}

func (s PermissionsSet) IsManager() bool {
	return s.Has(TogetherPermissionManager) || s.IsAdmin()
}

func (s PermissionsSet) IsSupport() bool {
	return s.Has(TogetherPermissionSupport) || s.IsManager()
}

func (s PermissionsSet) IsOrganizationalUser() bool {
	return s.IsSupport() || s.IsDeveloper()
}

func (s PermissionsSet) IsUser() bool {
	return s.Has(TogetherPermissionUser)
}

func (s PermissionsSet) IsConnected() bool {
	return s.IsUser() || s.IsOrganizationalUser()
}

func (s PermissionsSet) IsLogin() bool {
	return s.Has(TogetherPermissionLogin)
}