package grpc

import (
	"context"

	"github.com/roblesdotdev/movies-ms/gen"
	"github.com/roblesdotdev/movies-ms/internal/grpcuitil"
	"github.com/roblesdotdev/movies-ms/metadata/pkg/model"
	"github.com/roblesdotdev/movies-ms/pkg/discovery"
)

type Gateway struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry}
}

func (g *Gateway) Get(ctx context.Context, id string) (*model.Metadata, error) {
	conn, err := grpcuitil.ServiceConnection(ctx, "metadata", g.registry)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := gen.NewMetadataServiceClient(conn)
	resp, err := client.GetMetadata(ctx, &gen.GetMetadataRequest{MovieId: id})
	if err != nil {
		return nil, err
	}
	return model.MetadataFromProto(resp.Metadata), nil
}
