package handlers

import (
	"fmt"
	"net/http"

	"secret_api/api/controllers"
	"secret_api/storage"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type GetUsersResponse struct {
	Users []storage.User `json:"users"`
}

func CreateUserHandler(res http.ResponseWriter, req *http.Request) {
	// Check Content-Type
	if !ValidateContentType(res, req) {
		return
	}

	// Gets store from context
	st, err := GetStore(req)
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Decodes the new user data
	var newUser storage.NewUser
	if err := DecodeJsonBody(res, req, &newUser); err != nil {
		return
	}

	// Validates that the name field was properly set
	if newUser.Name == "" {
		http.Error(res, "Field 'name' can't be null or empty", http.StatusBadRequest)
		return
	}

	// Creates new user in the store
	user, err := controllers.CreateUser(st, newUser.Name)
	if err != nil {
		http.Error(res, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Encode the created user response
	if err := EncodeJsonBody(res, user); err != nil {
		return
	}
}

func GetUsersHandler(res http.ResponseWriter, req *http.Request) {
	// Gets store from context
	st, err := GetStore(req)
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

	users, err := controllers.GetAllUsers(st)
	if err != nil {
		http.Error(res, "Failed to get users", http.StatusInternalServerError)
		return
	}

	if err := EncodeJsonBody(res, GetUsersResponse{Users: users}); err != nil {
		return
	}
}

func GetUserHandler(res http.ResponseWriter, req *http.Request) {
	// Gets store from context
	st, err := GetStore(req)
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

	user_id_param := chi.URLParam(req, "user_id")

	user_id, err := uuid.Parse(user_id_param)
	if err != nil {
		err_txt := fmt.Sprintf("Invalid user_id param: %s", err.Error())
		http.Error(res, err_txt, http.StatusBadRequest)
		return
	}

	user, err := controllers.GetUser(st, user_id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}

	if err := EncodeJsonBody(res, user); err != nil {
		return
	}
}
