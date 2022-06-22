package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	gh "golang.org/x/oauth2/github"
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
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.User, c.Passwd, c.Host, c.Port, c.Name,
	)
}

func createOAuthConfig(ghClientId, ghClientSecret string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     ghClientId,
		ClientSecret: ghClientSecret,
		Endpoint:     gh.Endpoint,
	}
}

func getAuthUrl(config *oauth2.Config) string {
	return config.AuthCodeURL("state")
}

func getTokenFromAuthCode(ctx context.Context, config *oauth2.Config, authCode string) (*oauth2.Token, error) {
	token, err := config.Exchange(ctx, authCode)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func main() {
	// ctx := context.Background()

	// appConfig := AppConfig{}
	// appConfig.loadFromPath("./config.json")

	// oauthConfig := createOAuthConfig(appConfig.GHClientId, appConfig.GHClientSecret)
	// fmt.Println(getAuthUrl(oauthConfig))

	// var authCode string
	// fmt.Scanln(&authCode)

	// fmt.Println(getTokenFromAuthCode(ctx, oauthConfig, authCode))

	// fmt.Println(config)
	// fmt.Println(config.DBInfo.createDBString())

	fmt.Println("Hello World")
}
