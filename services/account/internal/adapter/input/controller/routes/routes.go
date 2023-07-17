package routes

import (
	"github.com/expoure/pismo/account/internal/adapter/input/controller"
	// "github.com/expoure/pismo/account/internal/adapter/input/controller/middlewares"
	"github.com/gin-gonic/gin"
)

func InitAccountRoutes(
	r *gin.RouterGroup,
	accountController controller.AccountControllerInterface,
) {

	r.GET("/:id", accountController.FindAccountByID)
	// r.GET("/ByDocumentNumber/:documentNumber", middlewares.VerifyTokenMiddleware, accountController.FindAccountByDocumentNumber)
	r.POST("", accountController.CreateAccount)
}
