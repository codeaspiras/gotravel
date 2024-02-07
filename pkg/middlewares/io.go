package middlewares

import (
	"gotravel/pkg/altcontext"
	"gotravel/pkg/stdio"
	"net/http"
)

func WithIO(io stdio.IO) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			next.ServeHTTP(w, r.WithContext(altcontext.CtxWithIO(ctx, io)))
		})
	}
}
