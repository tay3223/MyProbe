package demo

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Hello
//@Tags 测试1
//@Summary 测试接口
//@Description 这是一个测试接口
//@Accept json
//@Produce json
//@Success 200 {string} Hello
//@Router /api/v1/hello [get]
func Hello(g *gin.Context) {
    g.JSON(http.StatusOK, "Hello World")
}
