package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// Configuration 项目配置
type Configuration struct {
	AccessToken string `json:"access_token"`
	AllowOrigin string `json:"allow_origin"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Proxy       string `json:"proxy"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 给配置赋默认值
		config = &Configuration{
			AccessToken: "",
			Host:        "0.0.0.0",
			Port:        "8080",
			AllowOrigin: "*",
			Proxy:       "",
		}

		// 判断配置文件是否存在，存在直接JSON读取
		_, err := os.Stat("config.json")
		if err == nil {
			f, err := os.Open("config.json")
			if err != nil {
				log.Fatalf("open config err: %v", err)
				return
			}
			defer f.Close()
			encoder := json.NewDecoder(f)
			err = encoder.Decode(config)
			if err != nil {
				log.Fatalf("decode config err: %v", err)
				return
			}
		}
		// 有环境变量使用环境变量
		AccessToken := os.Getenv("ACCESS_TOKEN")
		if AccessToken != "" {
			config.AccessToken = AccessToken
		}

		AllowOrigin := os.Getenv("ALLOW_ORIGIN")
		if AllowOrigin != "" {
			config.AllowOrigin = AllowOrigin
		}

		Proxy := os.Getenv("PROXY")
		if Proxy != "" {
			config.Proxy = Proxy
		}

		Host := os.Getenv("HOST")
		if Host != "" {
			config.Host = Host
		}

		Port := os.Getenv("PORT")
		if Port != "" {
			config.Port = Port
		}
	})

	return config
}
