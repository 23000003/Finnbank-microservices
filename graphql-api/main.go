package main

import (
	"context"
	"finnbank/common/utils"
	"finnbank/graphql-api/db"
	"finnbank/graphql-api/graphql_config"
	"finnbank/graphql-api/graphql_config/handlers"
	"finnbank/graphql-api/graphql_config/resolvers"
	q "finnbank/graphql-api/queue"
	"finnbank/graphql-api/types"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/cors"
)

func CorsMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowCredentials: true,
	})
}

func main() {
	logger, err := utils.NewLogger()
	if err != nil {
		panic(err)
	}
	logger.Info("Starting the application...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbServicesPool := db.InitializeServiceDatabases(logger)
	q := q.NewQueue(logger, ctx)

	wsConn := initializeWebsockets(logger, ctx)

	server := initializeGraphQL(logger, dbServicesPool, q, wsConn)

	q.StartAutoDequeue(dbServicesPool.OpenedAccountDBPool, dbServicesPool.TransactionDBPool, dbServicesPool.AccountDBPool, wsConn.TransactionConn)
	startAndShutdownServer(server, logger, dbServicesPool, wsConn, cancel)
}

func initializeGraphQL(logger *utils.Logger, dbPool *types.StructServiceDatabasePools, q *q.Queue, wsConn *types.StructWebSocketConnections) *http.Server {
	resolvers := resolvers.NewGraphQLResolvers(logger)
	handlers := handlers.NewGraphQLServicesHandler(logger, resolvers, dbPool)
	graphql := graphql_config.NewGraphQL(logger, handlers)
	graphql.ConfigureGraphQLHandlers(q, wsConn)

	http.HandleFunc("/graphql/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("GraphQL API is OK."))
	})

	return &http.Server{
		Addr:    ":8083",
		Handler: CorsMiddleware().Handler(http.DefaultServeMux),
	}
}

func startAndShutdownServer(server *http.Server, logger *utils.Logger, dbServicesPool *types.StructServiceDatabasePools, wsConn *types.StructWebSocketConnections, cancel context.CancelFunc) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("Server running on http://localhost:8083")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server: %v", err)
		}
	}()

	<-done
	logger.Info("Server is shutting down...")

	// Cancel context to stop goroutines
	cancel()

	// Give goroutines time to stop
	time.Sleep(1 * time.Second)

	// Close WebSocket connections
	if wsConn.TransactionConn != nil {
		wsConn.TransactionConn.Close()
	}
	if wsConn.NotificationConn != nil {
		wsConn.NotificationConn.Close()
	}

	ctx, cancelTimeout := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelTimeout()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed: %v", err)
	}
	logger.Info("Server exited properly")

	db.CleanupDatabase(dbServicesPool, logger)
}
