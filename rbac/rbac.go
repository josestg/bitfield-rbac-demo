package rbac

import (
	"github.com/josestg/bitfield"
)

const (
	N               = 3
	maxBitsPerField = 52 // only 52 bit are secure for JavaScript (Frontend)
)

// Permission is a bitfield representing a permission.
type Permission = bitfield.BitField

// These permissions must be in first field.
// Splitting the permission to 3 groups of 52 bits is useful to demonstrate the behavior of the bitfield.
const (
	SeeUsers Permission = iota
	AddUsers
	DelUsers
	// add more here
)

// These permissions must be in second field.
const (
	SeeRoles = 52 + iota
	AddRoles
	DelRoles
	// add more here
)

// These permissions must be in third field.
const (
	AddEmails = 104 + iota
	PutEmails
	SeeEmails
	// add more here
)

// Role represents a role with a set of permissions.
type Role [N]Permission

// NewRole creates a new role with the given permissions.
func NewRole(permissions ...Permission) Role {
	var r Role
	for _, p := range permissions {
		r.AddPermission(p)
	}
	return r
}

// HasPermission returns true if the role has the given permission.
func (r *Role) HasPermission(p Permission) bool {
	f, b, ok := r.addr(p)
	if !ok {
		return false
	}
	return r[f].IsSet(b)
}

// AddPermission adds the given permission to the role.
func (r *Role) AddPermission(p Permission) {
	f, b, ok := r.addr(p)
	if ok {
		r[f] = r[f].SetBit(b)
	}
}

func (r *Role) addr(p Permission) (uint8, uint8, bool) {
	field := uint8(p / maxBitsPerField)
	if field >= N {
		return 0, 0, false
	}
	bit := uint8(p % maxBitsPerField)
	return field, bit, true
}
