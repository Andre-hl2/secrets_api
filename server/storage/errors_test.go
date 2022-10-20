package storage

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func compareErrMessage(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("Error message doesn't match expected. Want '%s', Got: '%s'", want, got)
	}
}

func TestNotFoundError(t *testing.T) {
	compareErrMessage(
		t,
		NotFound{Resource: "users", Id: uuid.Nil}.Error(),
		fmt.Sprintf("Resource 'users' not found with id: %s", uuid.Nil),
	)
}

func TestGetAllError(t *testing.T) {
	compareErrMessage(
		t,
		GetAllError{Resource: "users"}.Error(),
		"Failed to get all 'users'",
	)
}

func TestUpdateError(t *testing.T) {
	compareErrMessage(
		t,
		UpdateError{Resource: "users", Id: uuid.Nil}.Error(),
		fmt.Sprintf("Failed to update 'users' record with id '%s'", uuid.Nil),
	)
}

func TestCreationFailedError(t *testing.T) {
	compareErrMessage(
		t,
		CreationFailedError{Resource: "user", Err: "oops"}.Error(),
		"Failed to create 'user' with error: oops",
	)
}

func TestPostgresConnError(t *testing.T) {
	compareErrMessage(
		t,
		PostgresConnError{Err: "whoops"}.Error(),
		"Failed to connect to postgres: whoops",
	)

}
