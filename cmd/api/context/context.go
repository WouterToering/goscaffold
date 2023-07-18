package context

import (
	"net/http"

	"context"
)

func Get(r *http.Request, key any) any {
	return r.Context().Value(key)
}

func Set(r *http.Request, key, val any) *http.Request {
	if val == nil {
		return r
	}

	return r.WithContext(context.WithValue(r.Context(), key, val))
}
