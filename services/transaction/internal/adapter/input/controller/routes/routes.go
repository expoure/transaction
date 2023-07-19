package routes

import (
	"github.com/expoure/pismo/transaction/internal/adapter/input/controller"
	// "github.com/expoure/pismo/transaction/internal/adapter/input/controller/middlewares"
	"github.com/gin-gonic/gin"
)

func InitTransactionRoutes(
	r *gin.RouterGroup,
	transactionController controller.TransactionControllerInterface,
) {

	r.GET("", transactionController.ListTransaction)
	// r.GET("/ByDocumentNumber/:documentNumber", middlewares.VerifyTokenMiddleware, accountController.FindAccountByDocumentNumber)
	r.POST("", transactionController.CreateTransaction)
}
