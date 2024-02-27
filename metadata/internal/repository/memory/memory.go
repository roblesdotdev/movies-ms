package memory

import (
	"context"
	"sync"

	"github.com/roblesdotdev/movies-ms/metadata/internal/repository"
	model "github.com/roblesdotdev/movies-ms/metadata/pkg"
)

// Defines a memory movie metadata repository.
type Repository struct {
	data map[string]*model.Metadata
	sync.RWMutex
}

// Creates a new memory repository.
func New() *Repository {
	return &Repository{
		data: map[string]*model.Metadata{},
	}
}

// Get retrieves movie metadata for by movie id.
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()
	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrorNotFound
	}
	return m, nil
}

// Add movie metadata for a given movie id.
func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}
