package middleware

import (
	"context"
	"net/http"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var httpCtxKey = &contextKey{"httpReqRes"}

type contextKey struct {
	name string
}

type httpReqRes struct {
	Req *http.Request
	Res http.ResponseWriter
}

// Middleware decodes the share session cookie and packs the session into context
func GetReqRes() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// reqRes := &httpReqRes{r, w}
			// httpReqRes{r, w}.

			// put it in context
			ctx := context.WithValue(r.Context(), httpCtxKey, &httpReqRes{r, w})

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return
		})
	}
}

// ForContext finds the req/res from the context
func GetReqResCtx(ctx context.Context) *httpReqRes {
	raw, _ := ctx.Value(httpCtxKey).(*httpReqRes)
	return raw
}
