package handlers

import (
	"context"
	"net/http"
	"secret_api/storage"
)

func AddStore(next http.Handler, st storage.Store) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), "store", st)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}

func GetStore(req *http.Request) (storage.Store, error) {
	if st := req.Context().Value("store"); st != nil {
		return st.(storage.Store), nil
	}
	return &storage.MemoryStore{}, MissingStore{}
}
