//go:build wireinject
// +build wireinject

package depinject

import (
	"github.com/google/wire"

	http "stock-challenge-go/pkg/api"
	config "stock-challenge-go/pkg/config"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
