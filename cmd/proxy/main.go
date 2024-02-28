package main

import (
	"context"
	"log"
	"net/http"

	"proxy/pkg/proxy"
	"proxy/pkg/repository"
	"proxy/pkg/service"
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
	proxy := proxy.NewProxy(services)
	
	srv := http.Server{
		Addr:    ":8080",
		Handler: proxy.InitRoutes(),
	}

	log.Println("Starting proxy server on :8080")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting proxy server: ", err)
	}
}
