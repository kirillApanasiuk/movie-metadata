package memory

import (
	"context"
	"sync"

	"movie.com/internal/repository"
	"movie.com/model"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

func New() *Repository {
	return &Repository{
		data: make(map[string]*model.Metadata),
	}
}

// Get retrieves microservices_in_go metadata for by by microservices_in_go id.
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()

	metadata, ok := r.data[id]
	if !ok {
		return nil, repository.ErrorNotFound
	}

	return metadata, nil
}

// Put adds microservices_in_go metadata for a given microservices_in_go id.
func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}
