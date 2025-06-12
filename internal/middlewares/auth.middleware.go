package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiimlamxung/go-ecommerce-backend-api/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		fmt.Printf("token: %s\n Type: %T\n", token, token)
		if token != "valid-token" {
			response.ResponseWithInternalCode(c, http.StatusUnauthorized, response.ErrInvalidToken, nil)
			c.Abort() // c.Abort() được dùng để dừng việc thực thi các middleware tiếp theo và handler cho request hiện tại.
			return
		}
		c.Next()
	}
}
