package handlers

import (
	"github.com/gin-gonic/gin"
)

func Usuarios(c *gin.Context) {

	c.JSON(200, "ok")
}
