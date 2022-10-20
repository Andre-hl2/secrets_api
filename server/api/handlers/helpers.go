package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/gddo/httputil/header"
)

func DefaultHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Not Found")
}

func ValidateContentType(res http.ResponseWriter, req *http.Request) bool {
	if req.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(req.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(res, msg, http.StatusUnsupportedMediaType)
			return false
		}
	}

	return true
}

func DecodeJsonBody(res http.ResponseWriter, req *http.Request, dst interface{}) error {
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&dst); err != nil {
		err_txt := fmt.Sprintf("Failed to parse request json: %s", err.Error())
		http.Error(res, err_txt, http.StatusBadRequest)
		return err
	}

	return nil
}

func EncodeJsonBody(res http.ResponseWriter, data interface{}) error {
	encodedData, err := json.Marshal(data)
	if err != nil {
		err_txt := fmt.Sprintf("Failed to encode response: %s", err.Error())
		http.Error(res, err_txt, http.StatusInternalServerError)
		return err
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(encodedData)
	return nil
}
