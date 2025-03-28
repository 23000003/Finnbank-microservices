package routers

import (
	"finnbank/common/utils"
	"finnbank/api-gateway/middleware"
	"finnbank/api-gateway/types"
	"github.com/gin-gonic/gin"
)

type StructGatewayRouter struct {
	s *types.ApiGatewayServices
	r *gin.RouterGroup
	l *utils.Logger
}

func NewGatewayRouter(l *utils.Logger, r *gin.RouterGroup, s *types.ApiGatewayServices) *StructGatewayRouter {
	return &StructGatewayRouter{
		r: r,
		l: l,
		s: s,
	}
}

func (gr *StructGatewayRouter) ConfigureGatewayRouter() {

	// ======================== TEST ========================
	product := gr.r.Group("/product")
	{
		product.GET("", gr.s.ProductService.GetAllProduct)
		product.GET("/:id", gr.s.ProductService.GetByIdProduct)
		product.POST("", gr.s.ProductService.CreateProduct)
		product.PATCH("/:id", gr.s.ProductService.UpdateProduct)
		product.DELETE("/:id", gr.s.ProductService.DeleteProduct)
	}

	// ======================== Services ========================

	account := gr.r.Group("/account")
	{
		account.POST("/login", gr.s.AccountService.LoginUser)
		account.POST("/signup", gr.s.AccountService.SignupUser)
	}

	statement := gr.r.Group("/statement")
	statement.Use(middleware.AuthMiddleware())
	{
		statement.POST("/generate-statement", gr.s.StatementService.GenerateStatement)
	}

	bankcard := gr.r.Group("/bankcard")
	bankcard.Use(middleware.AuthMiddleware())
	{
		bankcard.GET("/get-user-bankcard", gr.s.BankcardService.GetUserBankcard)
		bankcard.POST("/generate-bankcard", gr.s.BankcardService.GenerateBankcardForUser)
		bankcard.PATCH("/renew-bankcard/:id", gr.s.BankcardService.RenewBankcardForUser)
	}

	transaction := gr.r.Group("/transaction")
	transaction.Use(middleware.AuthMiddleware())
	{
		transaction.GET("/get-all", gr.s.TransactionService.GetAllTransaction)
		transaction.GET("/:id", gr.s.TransactionService.GetTransaction)
		transaction.POST("/generate-transaction", gr.s.TransactionService.GenerateTransaction)
	}

	notification := gr.r.Group("/notification")
	notification.Use(middleware.AuthMiddleware())
	{
		notification.GET("/get-all", gr.s.NotificationService.GetUserNotifications)
		notification.POST("/generate-notif", gr.s.NotificationService.GenerateNotification)
		notification.PATCH("/mark-as-read/:id", gr.s.NotificationService.MarkAsReadNotification)
		notification.DELETE("/clear-notif/:id", gr.s.NotificationService.DeleteNotification)
	}
}
