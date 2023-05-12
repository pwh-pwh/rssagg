package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/pwh-pwh/rssagg/handlers"
	"github.com/pwh-pwh/rssagg/middlewares"
	"github.com/pwh-pwh/rssagg/resp"
	"net/http"
)

func Router(router chi.Router) {
	//Create sub-router for the /v1 namespace and mount it to the main router.
	router.Route("/v1", func(r chi.Router) {
		r.Get("/readiness", func(writer http.ResponseWriter, request *http.Request) {
			resp.RespondWithJSON(writer, 200, struct {
				Status string `json:"status"`
			}{
				"ok",
			})
		})
		r.Get("/err", func(writer http.ResponseWriter, request *http.Request) {
			resp.RespondWithError(writer, 500, "Internal Server Error")
		})
		r.Route("/users", func(r chi.Router) {
			r.Post("/", handlers.CreateUserHandler)
			r.Get("/", middlewares.AuthMiddleware(handlers.GetUserHandler))
		})
		r.Route("/feeds", func(r chi.Router) {
			r.Post("/", middlewares.AuthMiddleware(handlers.CreateFeedsHandler))
			r.Get("/", handlers.GetFeedsHandler)
		})
		r.Route("/feed_follows", func(r chi.Router) {
			r.Post("/", middlewares.AuthMiddleware(handlers.CreateFeedFollowHandler))
			r.Delete("/{feedFollowID}", middlewares.AuthMiddleware(handlers.DelFeedFollowHandler))
			r.Get("/", middlewares.AuthMiddleware(handlers.GetAllFF4UserHandler))
		})
		r.Get("/posts", middlewares.AuthMiddleware(handlers.GetPostsLimit))
	})
}
