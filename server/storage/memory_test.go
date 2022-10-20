package storage

import (
	"testing"

	"github.com/google/uuid"
)

func createUser(st MemoryStore, t *testing.T) User {
	user, err := st.CreateUser(NewUser{Name: "Test User"})
	if err != nil {
		t.Errorf("Failed to create user: %s", err.Error())
	}

	return user
}

func compareUsers(user1 User, user2 User, t *testing.T) {
	if user1.Name != user2.Name {
		t.Errorf("Users values doesn't match. Left: '%s', Right: '%s'", user1.Name, user2.Name)
	}

	if user1.Id != user2.Id {
		t.Errorf("Users values doesn't match. Left: '%s', Right: '%s'", user1.Id, user2.Id)
	}
}

func createSecret(st MemoryStore, t *testing.T) Secret {
	secret, err := st.CreateSecret(NewSecret{Secret: "shhhhh"})
	if err != nil {
		t.Errorf("Failed to create secret: %s", err.Error())
	}

	return secret
}

func compareSecrets(secret1 Secret, secret2 Secret, t *testing.T) {
	if secret1.Secret != secret2.Secret {
		t.Errorf("Secrets values doesn't match. Left: '%s', Right: '%s'", secret1.Secret, secret2.Secret)
	}
	if secret1.Id != secret2.Id {
		t.Errorf("Secrets values doesn't match. Left: '%s', Right: '%s'", secret1.Id, secret2.Id)
	}
	if secret1.GuesserId.String() != secret2.GuesserId.String() {
		t.Errorf("Secrets values doesn't match. Left: '%s', Right: '%s'", secret1.GuesserId, secret2.GuesserId)
	}
	if secret1.Guessed != secret2.Guessed {
		t.Errorf("Secrets values doesn't match. Left: '%t', Right: '%t'", secret1.Guessed, secret2.Guessed)
	}
}

func TestNewMemoryStore(t *testing.T) {
	st := NewMemoryStore()

	if st.users == nil {
		t.Error("Failed to correctly initialize 'users' map on 'NewmemoryStore'")
	}
	if st.secrets == nil {
		t.Error("Failed to correctly initialize 'secrets' map on 'NewmemoryStore'")
	}
}

func TestHealthCheck(t *testing.T) {
	st := NewMemoryStore()

	status := st.HealthCheck()

	if !status {
		t.Error("Failed to get good health check on Memory Store")
	}
}

func TestCreateUser(t *testing.T) {
	st := NewMemoryStore()

	oldCount := len(st.users)

	createUser(st, t)

	newCount := len(st.users)

	if newCount != (oldCount + 1) {
		t.Error("Users count incorrect after creating user")
	}
}

func TestGetUser(t *testing.T) {
	st := NewMemoryStore()

	createdUser := createUser(st, t)

	getUser, err := st.GetUser(createdUser.Id)
	if err != nil {
		t.Errorf("Failed to get user: %s", err.Error())
	}

	compareUsers(createdUser, getUser, t)

}

func TestFailToGetUser(t *testing.T) {
	st := NewMemoryStore()

	_, err := st.GetUser(uuid.New())
	if err == nil {
		t.Error("Failed to get err when getting invalid user")
	}
}

func TestGetAllUsers(t *testing.T) {
	st := NewMemoryStore()

	createdUser1 := createUser(st, t)
	createdUser2 := createUser(st, t)

	users, err := st.GetAllUsers()
	if err != nil {
		t.Errorf("Failed to get all users: %s", err.Error())
	}

	if len(users) != 2 {
		t.Errorf("Get all users lenght incorrect. Want 2, got %d", len(users))
	}

	compareUsers(createdUser1, users[0], t)
	compareUsers(createdUser2, users[1], t)
}

func TestCreateSecret(t *testing.T) {
	st := NewMemoryStore()

	oldCount := len(st.secrets)

	createSecret(st, t)

	newCount := len(st.secrets)

	if newCount != (oldCount + 1) {
		t.Error("Secrets count incorrect after creating secret")
	}
}

func TestGetSecret(t *testing.T) {
	st := NewMemoryStore()

	createdSecret := createSecret(st, t)

	getSecret, err := st.GetSecret(createdSecret.Id)
	if err != nil {
		t.Errorf("Failed to get secret: %s", err.Error())
	}

	compareSecrets(createdSecret, getSecret, t)

}

func TestFailToGetSecret(t *testing.T) {
	st := NewMemoryStore()

	_, err := st.GetSecret(uuid.New())
	if err == nil {
		t.Error("Failed to get err when getting invalid secret")
	}
}

func TestGetAllSecrets(t *testing.T) {
	st := NewMemoryStore()

	createdSecret1 := createSecret(st, t)
	createdSecret2 := createSecret(st, t)

	secrets, err := st.GetAllSecrets()
	if err != nil {
		t.Errorf("Failed to get all secrets: %s", err.Error())
	}

	if len(secrets) != 2 {
		t.Errorf("Get all secrets lenght incorrect. Want 2, got %d", len(secrets))
	}

	compareSecrets(createdSecret1, secrets[0], t)
	compareSecrets(createdSecret2, secrets[1], t)
}

func TestUpdateSecret(t *testing.T) {
	st := NewMemoryStore()

	createdSecret := createSecret(st, t)
	user := createUser(st, t)

	updatedSecret, err := st.UpdateSecret(createdSecret.Id, user.Id)
	if err != nil {
		t.Errorf("Failed to update secret: %s", err.Error())
	}

	if !updatedSecret.Guessed {
		t.Error("Failed to set 'guessed' correctly after updating secret")
	}

	if updatedSecret.GuesserId.String() != user.Id.String() {
		t.Errorf("GuesserId doesn't match expected value. Got '%s', Wanted: '%s'", updatedSecret.GuesserId.String(), user.Id.String())
	}
}

func TestFailToUpdateInvalidSecret(t *testing.T) {
	st := NewMemoryStore()

	_, err := st.UpdateSecret(uuid.New(), uuid.New())
	if err == nil {
		t.Error("Failed to get err when updating invalid secret")
	}
}
