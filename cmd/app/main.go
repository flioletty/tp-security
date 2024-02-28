package main

import (
	"context"
	"log"
	"net/http"
	"proxy/pkg/handler"
	"proxy/pkg/repository"
	"proxy/pkg/service"
	"time"
)

func main() {
	ctx := context.Background()

	dbConn, err := repository.NewMongoClient(ctx, repository.MongoConfig{
		ConnectionString: "mongodb://root:rootpassword@mongodb:27017",
		DatabaseName:     "security_hw",
	})
	if err != nil {
		log.Fatal("error when connecting to the mongodb")
	}

	repos := repository.NewRepository(dbConn)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := http.Server{
		Addr:         ":8000",
		Handler:      handlers.InitRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Starting proxy server on :8000")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting web-app server: ", err)
	}
}
