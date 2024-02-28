package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/roblesdotdev/movies-ms/movie/internal/controller/movie"
	metadatagateway "github.com/roblesdotdev/movies-ms/movie/internal/gateway/metadata/http"
	ratinggateway "github.com/roblesdotdev/movies-ms/movie/internal/gateway/rating/http"
	httphandler "github.com/roblesdotdev/movies-ms/movie/internal/handler/http"
	"github.com/roblesdotdev/movies-ms/pkg/discovery"
	"github.com/roblesdotdev/movies-ms/pkg/discovery/consul"
)

const serviceName = "movie"

func main() {
	var port int
	flag.IntVar(&port, "port", 8082, "API handler port")
	flag.Parse()
	log.Printf("starting the movie service on port %d", port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceId := discovery.GenerateInstanceId(serviceName)
	if err := registry.Register(ctx, instanceId, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.ReportHealthyState(instanceId, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceId, serviceName)
	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
