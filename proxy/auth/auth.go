package auth

import (
	"context"
	"fmt"

	"github.com/gliderlabs/ssh"
)

type Permissions []Permission
type Permission string

func addPermissions(ctx ssh.Context, p ...Permission) {
	for _, perm := range p {
		addPermission(ctx, perm)
	}
}

func addPermission(ctx ssh.Context, p Permission) {
	if ctx.Value("permissions") == nil {
		ctx.SetValue("permissions", Permissions{})
	}

	ctx.SetValue("permissions", append(ctx.Value("permissions").(Permissions), p))
}

func getPermissions(ctx ssh.Context) Permissions {
	if ctx.Value("permissions") == nil {
		return nil
	} else {
		return ctx.Value("permissions").(Permissions)
	}
}

func hasPermission(ctx ssh.Context, p Permission) bool {
	if getPermissions(ctx) != nil {
		for _, perm := range getPermissions(ctx) {
			if perm == p {
				return true
			}
		}
	}
	return false
}

func HasPermission(ctx context.Context, p Permission) bool {
	if getPermissions(ctx.(ssh.Context)) != nil {
		for _, perm := range getPermissions(ctx.(ssh.Context)) {
			if perm == p {
				return true
			}
		}
	}
	return false
}

func VerifyRequiredPermissions(s ssh.Session, perms ...Permission) {
	missingPerms := Permissions{}

	for _, perm := range perms {
		if !HasPermission(s.Context(), perm) {
			missingPerms = append(missingPerms, perm)
		}
	}

	if len(missingPerms) != 0 {
		s.Write([]byte("You don't have the required permissions " + fmt.Sprintf("%v", missingPerms) + "\n"))
		s.Close()
	}
}
