package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Port string `json:"port"`
	MySQL struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"mysql"`
	Redis struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Password string `json:"password"`
	} `json:"redis"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &config, nil
}
