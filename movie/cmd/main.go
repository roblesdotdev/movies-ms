package main

import (
	"log"
	"net/http"

	"github.com/roblesdotdev/movies-ms/movie/internal/controller/movie"
	metadatagateway "github.com/roblesdotdev/movies-ms/movie/internal/gateway/metadata/http"
	ratinggateway "github.com/roblesdotdev/movies-ms/movie/internal/gateway/rating/http"
	httphandler "github.com/roblesdotdev/movies-ms/movie/internal/handler/http"
)

func main() {
	log.Println("starting the movie service")
	metadataGateway := metadatagateway.New("http://127.0.0.1:8080")
	ratingGateway := ratinggateway.New("http://127.0.0.1:8081")
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
