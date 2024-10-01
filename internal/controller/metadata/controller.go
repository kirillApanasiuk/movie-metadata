package metadata

import (
	"context"
	"errors"

	"movie.com/internal/repository"
	"movie.com/model"
)

var ErrNotFound = errors.New("not found")

type metadataGeter interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

type Controller struct {
	repo metadataGeter
}

func NewController(repo metadataGeter) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) GetMetadata(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrorNotFound) {
		return nil, ErrNotFound
	}
	return res, nil

}
