package main

import (
	_ "github.com/lib/pq"
	"goapibestpractices/services"
	"goapibestpractices/database"
	"goapibestpractices/repository"
	"goapibestpractices/cache"
	"log"
	"net/http"
)

var portNumber string = ":8080"

func main() {
	var err error

	db, err := database.InitDB("postgres://localhost:5432/bookings?sslmode=disable")
    if err != nil {
        log.Fatalf("Veritabanı bağlantısı başarısız oldu: %s", err)
    }
    defer db.Close()

	redisCache := cache.NewRedisCache("localhost:6379", "", 0)

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo, redisCache)

	// HTTP sunucusunu başlat
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(userService), // routes fonksiyonuna userService'i parametre olarak geçin
	}

	log.Println("Starting server on :8080...")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
