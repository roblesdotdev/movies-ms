package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/roblesdotdev/movies-ms/pkg/discovery"
	"github.com/roblesdotdev/movies-ms/pkg/discovery/consul"
	"github.com/roblesdotdev/movies-ms/rating/internal/controller/rating"
	httphandler "github.com/roblesdotdev/movies-ms/rating/internal/handler/http"
	"github.com/roblesdotdev/movies-ms/rating/internal/repository/memory"
)

const serviceName = "rating"

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.Parse()
	log.Printf("starting rating service on port %d", port)
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
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
