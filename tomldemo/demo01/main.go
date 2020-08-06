package main
//@link https://github.com/BurntSushi/toml
//@link https://blog.csdn.net/Gusand/article/details/106094535

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"time"
)

type Config struct {
	Title    string
	App      app
	DB       mysql `toml:"mysql"`
	Redis    map[string]redis
	Releases releases
	Company  Company
}

type app struct {
	Author  string
	Org     string `toml:"organization"`
	Mark    string
	Release time.Time
}

type mysql struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type redis struct {
	Host string
	Port int
}

type releases struct {
	Release []string
	Tags    [][]interface{}
}

type Company struct {
	Name   string
	Detail detail
}

type detail struct {
	Type string
	Addr string
	ICP  string
}

func main() {
	var config Config
	if _, err := toml.DecodeFile("example.toml", &config); err != nil {
		panic(err)
	}

	fmt.Printf("全局信息: %+v\n\n", config.Title)

	fmt.Printf("App信息：%+v\n\n", config.App)

	fmt.Printf("Mysql配置：%+v\n\n", config.DB)

	fmt.Printf("版本信息：%+v\n\n", config.Releases)

	fmt.Printf("Redis主从：%+v\n\n", config.Redis)

	fmt.Printf("企业信息：%+v\n\n", config.Company)
}