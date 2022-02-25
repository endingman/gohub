package main

import (
	"flag"
	"fmt"
	"gohub/bootstrap"
	"gohub/pkg/config"

	"github.com/gin-gonic/gin"

	btsConfig "gohub/config"
)

/**
利用 Go 的 init() 方法来注册 config 目录下的配置信息。下面补充一下 init() 的知识。

Go 里面有两个特殊的函数：

main 包中的 main 函数，它是所有 Go 可执行程序的入口函数；
包级别的 init 函数。
init 函数是一个无参无返回值的函数：
*/
func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Logger
	bootstrap.SetupLogger()

	// new 一个 Gin Engine 实例
	router := gin.New()

	// 初始化数据库
	bootstrap.SetupDB()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
