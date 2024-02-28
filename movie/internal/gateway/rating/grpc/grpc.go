package grpc

import (
	"context"

	"github.com/roblesdotdev/movies-ms/gen"
	"github.com/roblesdotdev/movies-ms/internal/grpcuitil"
	"github.com/roblesdotdev/movies-ms/pkg/discovery"
	"github.com/roblesdotdev/movies-ms/rating/pkg/model"
)

type Gateway struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry}
}

func (g *Gateway) GetAggregatedRating(ctx context.Context, recordId model.RecordId, recordType model.RecordType) (float64, error) {
	conn, err := grpcuitil.ServiceConnection(ctx, "rating", g.registry)
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	client := gen.NewRatingServiceClient(conn)
	resp, err := client.GetAggregatedRating(ctx, &gen.GetAggregatedRatingRequest{RecordId: string(recordId), RecordType: string(recordType)})
	if err != nil {
		return 0, err
	}
	return resp.RatingValue, nil
}
