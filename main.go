package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "ipProbe/routers"
    "ipProbe/utils"
)

func init() {
    //初始化Swagger全局配置
    utils.InitSwagger()
    //初始化gin的运行模式
    utils.InitRunningMode()
}

func main() {
    r := gin.Default()

    routers.Path(r)

    //swagger默认路由
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    //服务启动地址（default：127.0.0.1:8080）
    runUrl := viper.GetString("server.host") + ":" + viper.GetString("server.port")
    utils.Log().Info("服务启动成功：http://" + runUrl)
    fmt.Println("服务启动成功：http://" + runUrl)
    if err := r.Run(runUrl); err != nil {
        return
    }
}
