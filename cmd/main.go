package main

import (
	controller "github.com/kirillApanasiuk/movie-metadata/internal/controller/metadata"
	httpHandler "github.com/kirillApanasiuk/movie-metadata/internal/handler"
	"github.com/kirillApanasiuk/movie-metadata/internal/repository/memory"
	"log"
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
