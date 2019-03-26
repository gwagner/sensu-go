package routers

import (
	"context"
	"net/http"

	corev2 "github.com/sensu/sensu-go/api/core/v2"
)

// ListStoreFunc ...
type ListStoreFunc func(context.Context) ([]corev2.Resource, string, error)

type listerFunc func(ListStoreFunc) http.HandlerFunc

// Lister ...
var Lister listerFunc

func init() {
	Lister = lister
}

func lister(list ListStoreFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		results, continueToken, err := list(r.Context())
		if err != nil {
			WriteError(w, err)
			return
		}

		if continueToken != "" {
			w.Header().Set(corev2.PaginationContinueHeader, continueToken)
		}

		RespondWith(w, results)
	}
}

func listerHandler(list ListStoreFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Lister(list).ServeHTTP(w, r)
	}
}
