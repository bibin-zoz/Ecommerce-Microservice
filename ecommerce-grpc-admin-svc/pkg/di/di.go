package di

import (
	server "admin/service/pkg/api"
	"admin/service/pkg/api/service"
	"admin/service/pkg/config"
	"admin/service/pkg/db"
	"admin/service/pkg/repository"
	"admin/service/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminServiceServer := service.NewAdminServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}
