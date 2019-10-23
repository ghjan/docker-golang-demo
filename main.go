package main

import (
	"docker-golang-demo/config"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-xweb/log"
	"net/http"
)

var (
	tomlFile = flag.String("config", "test.toml", "config file")
)

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "docker test",
	})
}

func main() {
	flag.Parse()
	// 解析配置文件
	tomlConfig, err := config.UnmarshalConfig(*tomlFile)
	if err != nil {
		log.Errorf("UnmarshalConfig: err:%v\n", err)
		return
	}
	address := tomlConfig.GetListenAddr()
	if tomlConfig.GetGinMode() == "release" {
		fmt.Println("to set gin mode to release mode")
		gin.SetMode(gin.ReleaseMode)
		fmt.Printf("Listening and serving HTTP on %s\n", address)
	}
	router := gin.New()
	router.GET("/", indexHandler)
	router.Run(address)
}
