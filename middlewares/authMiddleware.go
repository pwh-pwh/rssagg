package middlewares

import (
	"context"
	"github.com/pwh-pwh/rssagg/config"
	"github.com/pwh-pwh/rssagg/internal/auth"
	"github.com/pwh-pwh/rssagg/internal/database"
	"github.com/pwh-pwh/rssagg/resp"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func AuthMiddleware(handler authedHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		apiKey, err := auth.GetAPIKey(request.Header)
		if err != nil {
			resp.RespondWithError(writer, 400, "can not find api_key")
			return
		}
		userDb, err := config.Config.DB.GetUserByApiKey(context.Background(), apiKey)
		if err != nil {
			resp.RespondWithError(writer, 400, "can not find api_key")
			return
		}
		handler(writer, request, userDb)
	}
}
