package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type AppConfig struct {
	BotToken       string   `json:"token"`
	GHClientId     string   `json:"client_id"`
	GHClientSecret string   `json:"client_secret"`
	DBInfo         DBConfig `json:"db"`
}

type DBConfig struct {
	Name   string `json:"name"`
	User   string `json:"user"`
	Passwd string `json:"passwd"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
}

func (c *AppConfig) loadFromPath(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&c); err != nil {
		return err
	}

	return nil
}

func (c DBConfig) createDBString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Passwd, c.Host, c.Port, c.Name)
}

func main() {
	// config := AppConfig{}
	// config.loadFromPath("./config.json")
	// fmt.Println(config)
	// fmt.Println(config.DBInfo.createDBString())
	fmt.Println("Hello World")
}
