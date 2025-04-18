package handlers

import (
	"context"
	"finnbank/common/grpc/products"
	"finnbank/common/utils"
	"finnbank/graphql-api/db"
	"finnbank/graphql-api/graphql_config/resolvers"
	"finnbank/graphql-api/grpc_client"
	"finnbank/graphql-api/db"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"finnbank/graphql-api/types"
	sv "finnbank/graphql-api/services"
)

type StructGraphQLHandler struct {
	l *utils.Logger
	r *resolvers.StructGraphQLResolvers
	db *types.StructServiceDatabasePools
}

func NewGraphQLServicesHandler(l *utils.Logger, r *resolvers.StructGraphQLResolvers, db *types.StructServiceDatabasePools) *StructGraphQLHandler {
	return &StructGraphQLHandler{
		l: l,
		r: r,
		db: db,
	}
}

func (g *StructGraphQLHandler) ProductServicesHandler(connAddress string) *handler.Handler {

	grpcConnection := grpc_client.NewGRPCClient(connAddress)
	prodsServer := products.NewProductServiceClient(grpcConnection)

	productSchema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    g.r.GetProductQueryType(prodsServer),
			Mutation: g.r.GetProductMutationType(prodsServer),
		},
	)

	if err != nil {
		g.l.Fatal("Failed to configure Product Handler Schema %v", err)
	}

	productHandler := handler.New(&handler.Config{
		Schema:   &productSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	return productHandler
}

// HIGHLIGHT
// Instead of adding gRPC here, initialize DB connector add pass it to the Query and Mutations

func (g *StructGraphQLHandler) AccountServicesHandler() *handler.Handler {

	ACCService := sv.NewAccountService(g.db.AccountDBPool, g.l)

	accountSchema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    g.r.GetAccountQueryType(ACCService),
			Mutation: g.r.GetAccountMutationType(ACCService),
		},
	)
	if err != nil {
		g.l.Fatal("Failed to configure Account Handler Schema: %v", err)
	}

	accountHandler := handler.New(&handler.Config{
		Schema:   &accountSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	return accountHandler
}

// >>>>>>>>>>>>>============= ADD HERE ========== <<<<<<<<<<<<<<

func (g *StructGraphQLHandler) BankCardServicesHandler(connAddress string) *handler.Handler {

	// Follow AccountServiceHandler | OpenedAccountServicesHandler

	// grpcConnection := grpc_client.NewGRPCClient(connAddress)
	// proto server

	return nil
}

func (g *StructGraphQLHandler) NotificationServicesHandler(connAddress string) *handler.Handler {

	// Follow AccountServiceHandler | OpenedAccountServicesHandler

	// grpcConnection := grpc_client.NewGRPCClient(connAddress)
	// proto server

	return nil
}

func (g *StructGraphQLHandler) StatementServicesHandler(connAddress string) *handler.Handler {

	// Follow AccountServiceHandler | OpenedAccountServicesHandler

	// grpcConnection := grpc_client.NewGRPCClient(connAddress)
	// proto server

	return nil
}

func (g *StructGraphQLHandler) TransactionServicesHandler(connAddress string) *handler.Handler {

	// Follow AccountServiceHandler | OpenedAccountServicesHandler

	// grpcConnection := grpc_client.NewGRPCClient(connAddress)
	// proto server

	return nil
}

func (g *StructGraphQLHandler) OpenedAccountServicesHandler(connAddress string) *handler.Handler {
	
	// grpcConnection := grpc_client.NewGRPCClient(connAddress)
	// proto server

	OAService := sv.NewOpenedAccountService(g.db.OpenedAccountDBPool, g.l)

	openedAccountSchema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    g.r.GetOpenedAccountQueryType(OAService),
			Mutation: g.r.GetOpenedAccountMutationType(OAService),
		},
	)
	if err != nil {
		g.l.Fatal("Failed to configure Opened Account Handler Schema: %v", err)
	}

	openedAccountHandler := handler.New(&handler.Config{
		Schema:   &openedAccountSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	return openedAccountHandler

}