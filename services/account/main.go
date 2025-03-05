package main

import (
	"context"
	"finnbank/services/account/db"
	"finnbank/services/account/middleware"
	"finnbank/services/account/server"
	"finnbank/services/account/service"
	"finnbank/services/common/grpc/account"
	"finnbank/services/common/utils"
	"sync"
)

/*Transfer http configuration to http.go*/
// 	router := gin.New()
// serviceAPI := router.Group("/api/account") // base path
// handlers.AccountRouter(serviceAPI, accountService)

//	if err := router.Run("localhost:8082"); err != nil {
//		logger.Fatal("Failed to start server: %v", err)
//	}

// I'll transfer allat later, too lazy for it atm
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx := context.Background()
	logger, err := utils.NewLogger()
	if err != nil {
		panic(err)
	}

	database, err := db.InitDb(ctx)
	if err != nil {
		logger.Fatal("Error connecting to Db: %v", err)
	}
	defer database.Close(ctx)
	auth := &middleware.AuthService{}
	accountService := service.AccountService{
		DB:                                database,
		Logger:                            logger,
		Auth:                              auth,
		UnimplementedAccountServiceServer: account.UnimplementedAccountServiceServer{},
	}
	go func() {
		if err := server.GrpcServer(accountService, logger); err != nil {
			logger.Fatal("Failed to start gRPC server")
			return
		}
	}()
	logger.Info("Starting the server...")
	logger.Info("Server running on localhost:8082")
	wg.Wait()
}
