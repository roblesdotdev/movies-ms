package metadata

import (
	"context"

	"github.com/roblesdotdev/movies-ms/metadata/pkg/model"
)

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Controller defines a metadata service.
type Controller struct {
	repo metadataRepository
}

// New creates a metadata service controller.
func New(repo metadataRepository) *Controller {
	return &Controller{repo}
}

// Get retuns movie metadata by id.
func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
