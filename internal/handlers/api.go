package handler

import (
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	middleware "github.com/houcine7/my1s-goapi/internal/middlewares"
)


func Handler(router *chi.Mux){
	router.Use(chiMiddleware.StripSlashes) // this will remove any trailing slashes from the URL

	router.Route("/api/v1/account", func(router chi.Router){
		router.Use(middleware.Authorization)
		router.Get("/coins",GetCoinBalance)
	})
}