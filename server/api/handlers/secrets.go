package handlers

import (
	"fmt"
	"net/http"
	"secret_api/api/controllers"
	"secret_api/storage"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type GuessSecretRequest struct {
	Guess     string    `json:"guess"`
	GuesserId uuid.UUID `json:"guesser_id"`
}

type SecretResponse struct {
	Id        uuid.UUID `json:"id"`
	Guessed   bool      `json:"guessed"`
	GuesserId uuid.UUID `json:"guesser_id,omitempty"`
}

type SecretListResponse struct {
	Secrets []SecretResponse `json:"secrets"`
}

type GuessSecretResponse struct {
	Correct bool           `json:"correct"`
	Secret  SecretResponse `json:"secret"`
}

func secretToResponse(s storage.Secret) SecretResponse {
	return SecretResponse{
		Id:        s.Id,
		Guessed:   s.Guessed,
		GuesserId: s.GuesserId,
	}
}

func secretsToRespone(s []storage.Secret) SecretListResponse {
	secrets := []SecretResponse{}
	for _, secret := range s {
		secrets = append(secrets, secretToResponse(secret))
	}

	return SecretListResponse{Secrets: secrets}
}

func CreateSecretHandler(res http.ResponseWriter, req *http.Request) {
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

	// Decodes the new secret data
	var newSecret storage.NewSecret
	if err := DecodeJsonBody(res, req, &newSecret); err != nil {
		return
	}

	// Validates that the secret field was properly set
	if newSecret.Secret == "" {
		http.Error(res, "Field 'secret' can't be null or empty", http.StatusBadRequest)
		return
	}

	// Creates new secret in the store
	secret, err := controllers.CreateSecret(st, newSecret.Secret)
	if err != nil {
		http.Error(res, "Failed to create secret", http.StatusInternalServerError)
		return
	}

	// Encode the created secret response
	if err := EncodeJsonBody(res, secretToResponse(secret)); err != nil {
		return
	}
}

func GetAllSecretsHandler(res http.ResponseWriter, req *http.Request) {
	// Gets store from context
	st, err := GetStore(req)
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

	secrets, err := st.GetAllSecrets()
	if err != nil {
		http.Error(res, "Failed to get secrets", http.StatusInternalServerError)
		return
	}

	if err := EncodeJsonBody(res, secretsToRespone(secrets)); err != nil {
		return
	}
}

func GetSecretHandler(res http.ResponseWriter, req *http.Request) {
	// Gets store from context
	st, err := GetStore(req)
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

	secret_id_param := chi.URLParam(req, "secret_id")

	secret_id, err := uuid.Parse(secret_id_param)
	if err != nil {
		err_txt := fmt.Sprintf("Invalid secret_id param: %s", err.Error())
		http.Error(res, err_txt, http.StatusBadRequest)
		return
	}

	secret, err := controllers.GetSecret(st, secret_id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}

	if err := EncodeJsonBody(res, secret); err != nil {
		return
	}
}

func GuessSecretHandler(res http.ResponseWriter, req *http.Request) {
	// Gets store from context
	st, err := GetStore(req)
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Gets secret_id from path param
	secret_id_param := chi.URLParam(req, "secret_id")
	secret_id, err := uuid.Parse(secret_id_param)
	if err != nil {
		err_txt := fmt.Sprintf("Invalid secret_id param: %s", err.Error())
		http.Error(res, err_txt, http.StatusBadRequest)
		return
	}

	var guessRequest GuessSecretRequest
	if err := DecodeJsonBody(res, req, &guessRequest); err != nil {
		return
	}

	// Validates that the guess fields was properly set
	if guessRequest.Guess == "" {
		http.Error(res, "Field 'guess' can't be null or empty", http.StatusBadRequest)
		return
	}

	secret, correct, err := controllers.GuessSecret(st, secret_id, guessRequest.Guess, guessRequest.GuesserId)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	if err := EncodeJsonBody(res, GuessSecretResponse{
		Correct: correct,
		Secret:  secretToResponse(secret),
	}); err != nil {
		return
	}
}
