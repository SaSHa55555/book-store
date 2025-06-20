package main

import (
	"context"
	"log"

	"github.com/SaSHa55555/book-store/config"
	"github.com/SaSHa55555/book-store/internal/store/handlers"
	"github.com/SaSHa55555/book-store/internal/store/repository"
	"github.com/SaSHa55555/book-store/internal/store/services"
	"github.com/SaSHa55555/book-store/pkg/db"
)

func main() {
	cfg := config.LoadConfig()

	pool, err := db.NewPool(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repo := repository.NewRepository(pool)
	service := services.NewStoreService(repo, cfg)
	handler := handlers.NewStoreHandler(service, cfg)

	router := handler.InitRoutes()
	log.Println("Starting server")
	log.Fatal(router.Run("0.0.0.0:" + cfg.StoreConf.Port))
}
