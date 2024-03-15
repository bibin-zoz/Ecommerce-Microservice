package di

import (
	server "cart/service/pkg/api"
	"cart/service/pkg/api/service"
	"cart/service/pkg/client"
	"cart/service/pkg/config"
	"cart/service/pkg/db"
	"cart/service/pkg/repository"
	"cart/service/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	cartRepository := repository.NewCartRepository(gormDB)
	productClient := client.NewProductClient(&cfg)
	adminUseCase := usecase.NewCartUseCase(cartRepository, productClient)

	adminServiceServer := service.NewCartServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
