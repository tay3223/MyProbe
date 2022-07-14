package utils

import "ipProbe/docs"

func InitSwagger() {
    docs.SwaggerInfo.BasePath = ""
    Log().Info("swagger模块初始化完成...")
}
