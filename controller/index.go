package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

/**
 * @description 初始化controller
 */
func IndexApi(c *gin.Context)  {
	println("init controller")
	c.String(http.StatusOK,"default")
}