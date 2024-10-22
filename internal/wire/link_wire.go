//go:build wireinject

package wire

import (
	"github.com/google/wire"

	"github.com/a6slab/shorter_url/internal/handler"
	"github.com/a6slab/shorter_url/internal/repository"
)

func InitShorterRouterHandler() (*handler.ShorterHandler, error) {
	wire.Build(
		repository.NewShorterURLRepo,
		handler.NewShorterHandler,
	)
	return new(handler.ShorterHandler), nil
}