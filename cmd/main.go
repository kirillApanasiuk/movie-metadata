package main

import (
	"log"
	controller "movie.com/internal/controller/metadata"
	httpHandler "movie.com/internal/handler"
	"movie.com/internal/repository/memory"
	"net/http"
)

func main() {
	log.Println("Starting the microservices_in_go metadata service")
	repo := memory.New()
	ctrl := controller.NewController(repo)
	h := httpHandler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
