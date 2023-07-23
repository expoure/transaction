package routes

import (
	"github.com/expoure/pismo/account/internal/adapter/input/controller"
	"github.com/gin-gonic/gin"
)

func InitAccountRoutes(
	r *gin.RouterGroup,
	accountController controller.AccountControllerInterface,
) {

	r.GET("/:id", accountController.FindAccountByID)
	r.POST("", accountController.CreateAccount)
}
