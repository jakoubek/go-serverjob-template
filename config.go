package main

import (
"encoding/json"
"fmt"
"log"
"os"
)

type Config struct {
	Database struct {
		Dsn string `json:"dsn"`
	} `json:"database"`
	Email   EmailConfig `json:"email"`
}

type EmailConfig struct {
	Host          string `json:"host"`
	Port          int    `json:"port"`
	Auth     bool   `json:"auth"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	SenderName    string `json:"sender_name"`
	SenderAddress string `json:"sender_address"`
}

//func (c Config) IsProd() bool {
//	return c.Env == "prod"
//}

func DefaultConfig() Config {
	return Config{
	}
}

func LoadConfig(configReq bool) Config {
	f, err := os.Open("config.json")
	if err != nil {
		if configReq {
			panic(err)
		}
		fmt.Println("Using the default config...")
		return DefaultConfig()
	}
	var c Config
	dec := json.NewDecoder(f)
	err = dec.Decode(&c)
	if err != nil {
		panic(err)
	}
	log.Printf("Successfully loaded config file %s", f.Name())
	return c
}
