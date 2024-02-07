package middlewares

import (
	"gotravel/pkg/altcontext"
	"gotravel/pkg/validator"
	"net/http"
)

func WithValidator(v validator.Validator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			next.ServeHTTP(w, r.WithContext(altcontext.CtxWithValidator(ctx, v)))
		})
	}
}
