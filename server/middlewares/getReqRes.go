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

// Middleware to pass the req/res objects to the context
// so we can use them in the sessions parameters
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

//  finds the req/res from the context
func GetReqResCtx(ctx context.Context) *httpReqRes {
	raw, _ := ctx.Value(httpCtxKey).(*httpReqRes)
	return raw
}
