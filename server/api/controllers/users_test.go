package controllers

import (
	"secret_api/storage"
	"testing"
)

func createUser(st storage.MemoryStore, t *testing.T) storage.User {
	user, err := CreateUser(&st, "Test User")
	if err != nil {
		t.Errorf("Failed to create user: %s", err.Error())
	}

	return user
}

func compareUsers(user1 storage.User, user2 storage.User, t *testing.T) {
	if user1.Name != user2.Name {
		t.Errorf("Users values doesn't match. Left: '%s', Right: '%s'", user1.Name, user2.Name)
	}

	if user1.Id != user2.Id {
		t.Errorf("Users values doesn't match. Left: '%s', Right: '%s'", user1.Id, user2.Id)
	}
}

func TestCreateUser(t *testing.T) {
	st := storage.NewMemoryStore()

	name := "Test User"

	user := createUser(st, t)

	if user.Name != name {
		t.Errorf("Created user name doesn't match. Want: '%s', Got: '%s'", name, user.Name)
	}
}

func TestGetUser(t *testing.T) {
	st := storage.NewMemoryStore()

	createdUser := createUser(st, t)

	gotUser, err := GetUser(&st, createdUser.Id)
	if err != nil {
		t.Errorf("Failed to get user: %s", err.Error())
	}

	compareUsers(createdUser, gotUser, t)
}

func TestGetAllUsers(t *testing.T) {
	st := storage.NewMemoryStore()

	createUser(st, t)
	createUser(st, t)

	users, err := GetAllUsers(&st)
	if err != nil {
		t.Errorf("Failed to get all users: %s", err.Error())
	}

	if len(users) != 2 {
		t.Errorf("Len of users doesn't match expected value. Want '2', Got '%d'", len(users))
	}
}
