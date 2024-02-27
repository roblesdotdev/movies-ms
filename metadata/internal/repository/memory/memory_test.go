package memory

import (
	"context"
	"testing"

	model "github.com/roblesdotdev/movies-ms/metadata/pkg"
)

func TestMemoryRepository(t *testing.T) {
	repo := New()

	md := &model.Metadata{
		Title:       "test movie",
		Description: "test description",
	}

	id := "1"
	err := repo.Put(context.TODO(), id, md)
	if err != nil {
		t.Errorf("error adding metadata %v", err)
	}

	res, err := repo.Get(context.TODO(), id)
	if err != nil {
		t.Errorf("error retrieving metadata %v", err)
	}
	if res == nil {
		t.Error("retrieved metadata is nil")
	}

	if res.Title != md.Title {
		t.Error("retrieved metadata not match added metadata.")
	}
}
