package rating

import (
	"context"

	"github.com/roblesdotdev/movies-ms/rating/pkg/model"
)

type ratingRepository interface {
	Get(ctx context.Context, recordId model.RecordId, recordType model.RecordType) ([]model.Rating, error)
	Put(ctx context.Context, recordId model.RecordId, recordType model.RecordType, rating *model.Rating) error
}

// Controller defines a rating service controller.
type Controller struct {
	repo ratingRepository
}

// New create a rating service controller.
func New(repo ratingRepository) *Controller {
	return &Controller{repo}
}

// GetAggregatedRating returns the aggregated rating for a record
// or err if there no ratings for it.
func (c *Controller) GetAggregatedRating(ctx context.Context, recordId model.RecordId, recordType model.RecordType) (float64, error) {
	ratings, err := c.repo.Get(ctx, recordId, recordType)
	if err != nil {
		return 0, err
	}
	sum := float64(0)
	for _, r := range ratings {
		sum += float64(r.Value)
	}
	return sum / float64(len(ratings)), nil
}

func (c *Controller) PutRating(ctx context.Context, recordId model.RecordId, recordType model.RecordType, rating *model.Rating) error {
	return c.repo.Put(ctx, recordId, recordType, rating)
}
