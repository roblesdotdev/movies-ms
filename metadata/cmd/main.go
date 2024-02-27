package main

import (
	"log"
	"net/http"

	"github.com/roblesdotdev/movies-ms/metadata/internal/controller/metadata"
	httphandler "github.com/roblesdotdev/movies-ms/metadata/internal/handler/http"
	"github.com/roblesdotdev/movies-ms/metadata/internal/repository/memory"
)

func main() {
	log.Println("starting the movie metadata service")
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)

	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
