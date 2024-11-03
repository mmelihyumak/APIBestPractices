package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"goapibestpractices/handlers"
	"goapibestpractices/middlewares"
	httpSwagger "github.com/swaggo/http-swagger"
    _ "goapibestpractices/docs"
)

func routes() http.Handler{
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middlewares.LoggerMiddleware)
	//mux.Use(middlewares.JwtAuthMiddleware)

	mux.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	mux.Post("/api/brand", handlers.BrandHandler)
	mux.Post("/api/auth", handlers.AuthHandler)

	return mux
}
