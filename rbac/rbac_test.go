package rbac_test

import (
	"github.com/josestg/bitfield-rbac-demo/rbac"
	"math"
	"testing"
)

func TestNewRole(t *testing.T) {
	role := rbac.NewRole(
		rbac.SeeUsers, rbac.AddUsers, rbac.DelUsers,
		rbac.SeeRoles, rbac.AddRoles, rbac.DelRoles,
		rbac.AddEmails, rbac.PutEmails, rbac.SeeEmails,
	)

	expected := rbac.Role{7, 7, 7}
	if role != expected {
		t.Errorf("Expected %v, got %v", expected, role)
	}
}

func TestHasPermission(t *testing.T) {
	permissions := []rbac.Permission{
		rbac.SeeUsers, rbac.AddUsers, rbac.DelUsers,
		rbac.SeeRoles, rbac.AddRoles, rbac.DelRoles,
		rbac.AddEmails, rbac.PutEmails, rbac.SeeEmails,
	}

	role := rbac.NewRole(permissions[3:]...)

	for _, p := range permissions[:3] {
		if role.HasPermission(p) {
			t.Errorf("Expected role to not have %v", p)
		}
	}

	for _, p := range permissions[3:] {
		if !role.HasPermission(p) {
			t.Errorf("Expected role to have %v", p)
		}
	}

	// bit index out of range
	if role.HasPermission(math.MaxUint64) {
		t.Error("Expected role to not have permission")
	}
}
