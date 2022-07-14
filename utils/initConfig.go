package utils

import (
    "fmt"
    "github.com/spf13/viper"
)

func init() {
    //默认配置文件路径
    defaultConfigPath := "config/default.yaml"

    //加载的配置文件（写法二）
    viper.SetConfigFile(defaultConfigPath)

    //读取配置文件...（读取上面指定的配置文件）
    if err := viper.ReadInConfig(); err != nil {
        fmt.Println("读取配置文件失败：", err)
        return
    } else {
        fmt.Println(defaultConfigPath + "配置文件读取成功...")
    }

    //为一部分Config添加默认值
    setDefaultConfig()
}

//设置默认配置项
func setDefaultConfig() {
    /*
       如果 config/default.yaml 中没有设置下面这些内容，
       则下文中的默认值将会生效；
       配置文件中的优先级高于此处设定的优先级
    */

    //全局相关
    viper.SetDefault("global.env", "dev")              //当前运行环境
    viper.SetDefault("global.databaseType", "sqlite3") //当前连接数据库类型
    viper.SetDefault("global.ginMode", "release")      //当前gin的运行模式；（选项：debug、release）

    //运行相关
    viper.SetDefault("server.host", "127.0.0.1") //启动时监听的地址
    viper.SetDefault("server.port", "8080")      //启动时监听的端口

    //日志相关
    viper.SetDefault("log.level", "info")              //日志记录级别；（选项：debug、info）
    viper.SetDefault("log.format", "json")             //日志记录格式；（选项：json、text）
    viper.SetDefault("log.path", "./logs/default.log") //日志输出位置
}
