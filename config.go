package main

import (
	"fmt"

	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"
	"github.com/tiamxu/kit/log"
	"github.com/tiamxu/kit/sql"
)

const configPath = "config/config.yaml"

// yaml文件内容映射到结构体
type Config struct {
	ENV      string      `yaml:"env"`
	LogLevel string      `yaml:"log_level"`
	HttpSrv  HttpSrv     `yaml:"http"`
	DB       *sql.Config `yaml:"db"`
}

type HttpSrv struct {
	Address   string `yaml:"address"`
	KeepAlive bool   `yaml:"keep_alive"`
}

// set log level
func (c *Config) InitConfig() (err error) {
	defer func() {
		if err == nil {
			log.Printf("config initialed, env: %s", cfg.ENV)
		}
	}()
	err = sql.Connect(c.DB)
	if err != nil {
		fmt.Println("Init Mysql failed..")
		return err
	}
	fmt.Println("mysql conn succ")
	if level, err := logrus.ParseLevel(c.LogLevel); err != nil {
		return err
	} else {
		log.DefaultLogger().SetLevel(level)
	}

	return nil
}

// 读取配置文件

func loadConfig() {
	cfg = new(Config)
	multiconfig.MustLoadWithPath(configPath, cfg)
}

// func loadConfig() {
// 	filename := configPath
// 	data, err := ioutil.ReadFile(filename)
// 	if err != nil {

// 		return
// 	}
// 	// fmt.Printf("yaml文件内容:\n%v\n", string(data))
// 	cfg = new(Config)

// 	yaml.Unmarshal(data, cfg)
// 	// fmt.Printf("%+v\n", cfg)
// }
