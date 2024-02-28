package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/roblesdotdev/movies-ms/metadata/internal/controller/metadata"
	httphandler "github.com/roblesdotdev/movies-ms/metadata/internal/handler/http"
	"github.com/roblesdotdev/movies-ms/metadata/internal/repository/memory"
	"github.com/roblesdotdev/movies-ms/pkg/discovery"
	"github.com/roblesdotdev/movies-ms/pkg/discovery/consul"
)

const serviceName = "metadata"

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "API handler port")
	flag.Parse()
	log.Printf("starting the movie metadata service on port %d", port)
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
				log.Println("Failed to report healty state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceId, serviceName)
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)

	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
