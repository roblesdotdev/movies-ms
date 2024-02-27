package main

import (
	"log"
	"net/http"

	"github.com/roblesdotdev/movies-ms/rating/internal/controller/rating"
	httphandler "github.com/roblesdotdev/movies-ms/rating/internal/handler/http"
	"github.com/roblesdotdev/movies-ms/rating/internal/repository/memory"
)

func main() {
	log.Println("starting rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
