package routers

import (
    "github.com/gin-gonic/gin"
    "ipProbe/apps/demo"
    "ipProbe/apps/ip"
    "ipProbe/routers/middleware"
    "net/http"
)

func Path(g *gin.Engine) {

    //g.Use(middleware.Cors())
    g.Use(middleware.CrosHandler())

    //api分组
    api := g.Group("/api")
    {
        v1 := api.Group("v1")
        {
            v1.GET("/hello", demo.Hello)
            v1.POST("/probe-ip", ip.ProbeIp)
        }
    }

    //路由重定向（打开根路由的时候直接进入到swagger页面）
    g.GET("/", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
    })

    //路由重定向
    g.GET("/docs", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
    })

    //健康监测接口（项目中的基础接口，应该始终保持可用）
    g.GET("/health", demo.Health)
    g.GET("/ok", demo.Health)
    g.GET("/ok.html", demo.Health)

    //加载favicon图标
    g.StaticFile("/favicon.ico", "./favicon.png")
}
