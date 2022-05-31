package main

import (
	"ByteDance/conf"
	"flag"

	"ByteDance/config"
	"ByteDance/model"
	"ByteDance/router/douyin"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.Init()
	var env string
	flag.StringVar(&env, "env", "","自己配置 如--env=test")
	//如果是 --env=test的话那么就会读取.env.test文件
	flag.Parse()
	conf.InitConfig(env)


	douyin.InitRouter(r)
	model.Init()
	model.Init_Oss()
	port := conf.Get("app.port")

	r.Run(":" + port)

}
