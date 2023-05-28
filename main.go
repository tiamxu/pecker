package main

import (
	"fmt"

	"github.com/tiamxu/pecker/routers"

	"github.com/gin-gonic/gin"

	"github.com/tiamxu/kit/log"
)

var (
	cfg *Config
)

func (c *Config) Initial() (err error) {
	return
}
func init() {
	loadConfig()
	//log level
	if err := cfg.InitConfig(); err != nil {
		log.Fatalf("配置文件错误,%v\n", err)
	}
}
func main() {

	InitServer()

}

func InitServer() {
	r := gin.Default()
	routers.InitRouter(r)
	err := r.Run(cfg.HttpSrv.Address)
	if err != nil {
		panic(fmt.Sprintf("start server err: %v", err))
	}
}
