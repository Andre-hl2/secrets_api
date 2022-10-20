package controllers

import (
	"secret_api/storage"

	"github.com/google/uuid"
)

func CreateUser(st storage.Store, name string) (storage.User, error) {
	return st.CreateUser(storage.NewUser{Name: name})
}

func GetUser(st storage.Store, id uuid.UUID) (storage.User, error) {
	return st.GetUser(id)
}

func GetAllUsers(st storage.Store) ([]storage.User, error) {
	return st.GetAllUsers()
}
