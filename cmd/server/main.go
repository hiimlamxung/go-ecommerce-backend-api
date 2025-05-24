package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiimlamxung/go-ecommerce-backend-api/internal/routers"
)

func main() {
	gin.ForceConsoleColor()
	router := routers.NewRouter()
	router.Run(":8010") // listen and serve on 0.0.0.0:8080
}
