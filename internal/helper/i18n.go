package helper

import (
	"fmt"
)

const (
	Success           = "Success"
	ErrFailedReadBody = "Failed to read body"
	ErrInternalError  = "Encountered an error"

	// Auth
	ErrFailHash = "Failed to hash password"

	// User
	ErrExistingUsername = "Existing username"
	ErrExistingEmail    = "Existing email"
	ErrFailCreateUser   = "Failed to create user"
	ErrFailDelUser      = "failed to delete user"
	ErrFailGetUsers     = "Failed to retrieve users"
)

func NotFound(entity string) string {
	return fmt.Sprintf("%s not found", entity)
}
