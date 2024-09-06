package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/timest/env"
)

type config struct {
	AppMode string `default:"debug"`
	Port    string
	Mysql   struct {
		Address string
		Port    string
		User    string
		Pass    string
		DbName  string
	}
	Huawei struct {
		User    string
		IAMUser string
		IAMPass string
	}
}

var (
	Cfg     = &config{}
	Version = "0.0.1"
)

func Init() {
	// 读取config
	_ = godotenv.Load()
	err := env.Fill(Cfg)
	if err != nil {
		panic(err)
	}
	log.Print("load env success")
}
