package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"goapibestpractices/handlers"
	"goapibestpractices/services"
	"goapibestpractices/middlewares"
	httpSwagger "github.com/swaggo/http-swagger"
    _ "goapibestpractices/docs"
)

func routes(userService *services.UserService) http.Handler {
    mux := chi.NewRouter()
    mux.Use(middleware.Recoverer)
    mux.Use(middlewares.LoggerMiddleware)

    mux.Get("/swagger", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

    userHandler := handlers.NewUserHandler(userService)

    mux.Post("/user", userHandler.CreateUser)      // Create işlemi için POST kullanmak
    mux.Get("/user", userHandler.GetUserByID)         // GET ile kullanıcıyı alma
    mux.Put("/user", userHandler.UpdateUser)       // Güncelleme işlemi için PUT kullanmak
    mux.Delete("/user", userHandler.DeleteUser)    // Silme işlemi için DELETE kullanmak

    return mux
}
