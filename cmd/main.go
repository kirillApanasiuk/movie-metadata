package main

import (
	"log"
	metadata2 "movie.com/metadata/internal/controller/metadata"
	httpHandler "movie.com/metadata/internal/handler"
	"movie.com/metadata/internal/repository/memory"
	"net/http"
)

func main() {
	log.Println("Starting the microservices_in_go metadata service")
	repo := memory.New()
	ctrl := metadata2.NewController(repo)
	h := httpHandler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
