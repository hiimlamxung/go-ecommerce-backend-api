package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hiimlamxung/go-ecommerce-backend-api/internal/controller"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	// group route
	{
		{
			api := router.Group("/api")
			{
				v1 := api.Group("/v1")
				{
					v1.GET("/ping/:name", controller.NewPongController().Pong)

					testBind := v1.Group("/test-bind-form-data")
					{
						testBind.GET("/get-data-b", controller.GetDataB)
						testBind.GET("/get-data-c", controller.GetDataC)
						testBind.GET("/get-data-d", controller.GetDataD)
					}

					user := v1.Group("/user")
					{
						user.GET("/:id", controller.NewUserController().GetUserById)
						user.POST("/create", controller.NewUserController().CreateUser)
					}
				}
			}
		}
	}
	return router
}
