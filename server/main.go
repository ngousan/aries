package main

import (
	"aoBlog/config/app"
	"aoBlog/config/setting"
	_ "aoBlog/docs"
	"aoBlog/log"
)

// @title Gin Swagger
// @version 1.0
// @description Aries 开源博客项目 API 接口文档

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8088
func main() {
	engine := app.InitApp() // 初始化

	err := engine.Run(":" + setting.Config.Server.Port) // 运行
	if err != nil {
		log.Logger.Sugar().Panic("项目启动失败: ", err.Error())
	}
}
