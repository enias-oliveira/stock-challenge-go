// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package depinject

import (
	"stock-challenge-go/pkg/client"
	"stock-challenge-go/pkg/config"
	"stock-challenge-go/pkg/db"
	"stock-challenge-go/pkg/http"
	"stock-challenge-go/pkg/http/handler"
	"stock-challenge-go/pkg/repository"
	"stock-challenge-go/pkg/service"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	accountRepository := repository.NewAccountRepository(gormDB)
	accountService := services.NewAccountService(accountRepository)
	accountHandler := handler.NewAccountHandler(accountService)
	stooqClient := client.NewStooqClient(cfg)
	stockRepository := repository.NewStockRepository(stooqClient)
	historyRepository := repository.NewHistoryRepository(gormDB)
	historyService := services.NewHistoryService(historyRepository)
	stockService := services.NewStockService(stockRepository, historyService)
	stockHandler := handler.NewStockHandler(stockService)
	serverHTTP := http.NewServerHTTP(accountHandler, stockHandler)
	return serverHTTP, nil
}
