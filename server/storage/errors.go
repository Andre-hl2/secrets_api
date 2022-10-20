package storage

import (
	"fmt"

	"github.com/google/uuid"
)

type NotFound struct {
	Id       uuid.UUID
	Resource string
}

func (nf NotFound) Error() string {
	return fmt.Sprintf("Resource '%s' not found with id: %s", nf.Resource, nf.Id)
}

type CreationFailedError struct {
	Resource string
	Err      string
}

type GetAllError struct {
	Resource string
}

func (ga GetAllError) Error() string {
	return fmt.Sprintf("Failed to get all '%s'", ga.Resource)
}

type UpdateError struct {
	Id       uuid.UUID
	Resource string
}

func (ue UpdateError) Error() string {
	return fmt.Sprintf("Failed to update '%s' record with id '%s'", ue.Resource, ue.Id)
}

func (cf CreationFailedError) Error() string {
	return fmt.Sprintf("Failed to create '%s' with error: %s", cf.Resource, cf.Err)
}

type PostgresConnError struct {
	Err string
}

func (pc PostgresConnError) Error() string {
	return fmt.Sprintf("Failed to connect to postgres: %s", pc.Err)
}
