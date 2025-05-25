package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct {
}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	fmt.Println("My Handler Pong")
	name := c.Param("name")
	uid := c.DefaultQuery("uid", "123456")
	// spew.Dump([]string{"user1", "user2", "user3"})
	c.JSON(http.StatusOK, gin.H{
		"name":  name,
		"uid":   uid,
		"users": []string{"user1", "user2", "user3"},
	})
}
