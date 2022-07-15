package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"html/template"
	"io/fs"
	"ipProbe/routers"
	"ipProbe/utils"
	"net/http"
)

//go:embed assets/static
var Static embed.FS

//go:embed assets/templates
var Temp embed.FS

//go:embed assets/templates/favicon.png
var Icon string

func initStatic(g *gin.Engine) {

	//加载静态文件和模板文件（方式一）
	//g.Static("/static", "./static")           //指定/static目录对应的实际地址是./static
	//g.LoadHTMLGlob("./templates/*.html")      //加载指定目录下的html，这样才能在别的地方中使用

	//加载静态文件和模板文件（方式二）
	fe, _ := fs.Sub(Static, "assets/static")
	g.StaticFS("/static", http.FS(fe))

	templ := template.Must(template.New("").ParseFS(Temp, "assets/templates/*.html"))
	g.SetHTMLTemplate(templ)

	//加载favicon图标
	//g.StaticFile("/favicon.ico", "./favicon.png")
	g.StaticFile("favicon.ico", Icon)
}

func init() {
	//初始化Swagger全局配置
	utils.InitSwagger()
	//初始化gin的运行模式
	utils.InitRunningMode()
}

func main() {
	r := gin.Default()

	//初始化静态资源
	initStatic(r)

	//加载路由表
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
