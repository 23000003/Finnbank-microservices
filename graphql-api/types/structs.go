package types

import (
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StructGrpcServiceConnections struct {
	ProductServer       string
	BankCardServer      string
	AuthServer          string
	StatementServer     string
	TransactionServer   string
	NotificationServer  string
	OpenedAccountServer string
}

type StructServiceDatabasePools struct {
	BankCardDBPool      *pgxpool.Pool
	AccountDBPool       *pgxpool.Pool
	TransactionDBPool   *pgxpool.Pool
	NotificationDBPool  *pgxpool.Pool
	OpenedAccountDBPool *pgxpool.Pool
}

type StructWebSocketConnections struct {
	TransactionConn *websocket.Conn
	NotificationConn   *websocket.Conn
}