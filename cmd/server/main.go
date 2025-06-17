package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiimlamxung/go-ecommerce-backend-api/internal/initialize"
)

func main() {
	gin.ForceConsoleColor()
	initialize.Run()
}
