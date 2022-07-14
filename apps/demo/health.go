package demo

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Health(g *gin.Context) {
    g.JSON(http.StatusOK, "Health monitoring is ok")
}
