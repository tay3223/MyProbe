package utils

import "github.com/gin-gonic/gin"

func InitRunningMode() {
    ginMode := "release"
    switch ginMode {
    case "debug":
        gin.SetMode(gin.DebugMode)
    case "release":
        gin.SetMode(gin.ReleaseMode)
    default:
        //默认为"release"
        gin.SetMode(gin.ReleaseMode)
    }
    Log().Infof("运行模式加载为%v...", ginMode)
}
