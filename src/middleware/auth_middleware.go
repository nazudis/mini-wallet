package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/nazudis/mini-wallet/src/helper"
)

type CtxKey string

var CustomerXidCtxKey CtxKey = "CustomerXidCtx"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := helper.PlugResponse(w)

		auth := r.Header.Get("Authorization")
		if auth == "" {
			data := helper.MapJSON{"error": "empty authorization"}
			res.SetHttpStatusCode(http.StatusUnauthorized).ReplyFail(data)
			return
		}

		auths := strings.Split(auth, " ")

		if auths[0] != "Token" {
			data := helper.MapJSON{"error": "invalid auth schema on index 0"}
			res.SetHttpStatusCode(http.StatusUnauthorized).ReplyFail(data)
			return
		}

		if len(auths) < 2 {
			data := helper.MapJSON{"error": "token not provide"}
			res.SetHttpStatusCode(http.StatusUnauthorized).ReplyFail(data)
			return
		}

		token := auths[1]
		customerXid, err := helper.DecodeToken(token)
		if err != nil {
			data := helper.MapJSON{"error": "invalid token"}
			res.SetHttpStatusCode(http.StatusUnauthorized).ReplyFail(data)
			return
		}

		ctx := context.WithValue(r.Context(), CustomerXidCtxKey, customerXid)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
