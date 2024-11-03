package main

import (
	_ "github.com/lib/pq"
	"goapibestpractices/models"
    "goapibestpractices/cache"
    "goapibestpractices/services"
	"log"
	"net/http"
	"fmt"
)

var portNumber string = ":8080"

func main() {
	var err error

	models.DB, err = models.ConnectDb()
	if err != nil {
		log.Fatalf("Veritabanı bağlantısı başarısız oldu %s", err)
	}

	redisClient := cache.GetRedisClient()
    redisCache := cache.NewRedisCache(redisClient)

	userService := services.NewUserService(redisCache)

    // Kullanıcı verisini cache üzerinden getirme
    userData, err := userService.GetUser("userID123")
    if err != nil {
        panic(err)
    }
    fmt.Println("User Data:", userData)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(),
	}

	log.Println("Starting server on :8080...")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
