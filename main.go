package main

import (
	"context"
	"io"
	"log"
	"net/url"

	"golang.org/x/oauth2/clientcredentials"
)

const (
	apiURL         = "https://api.twitter.com/oauth/access_token"
	tokenURL       = "https://api.twitter.com/oauth2/token"
	consumerKey    = "IQKbtAYlXLripLGPWd0HUA"
	consumerSecret = "GgDYlkSvaPxGxC4X8liwpUoqKwwr3lCADbz8A7ADU"
)

type UserInfo struct {
	OauthToken       string
	OauthTokenSecret string
	ScreenName       string
	UserID           string
}

func Login() *UserInfo {
	userName := "9ee7k5y"
	passWord := "Xamakush420"
	config := &clientcredentials.Config{
		ClientID:     consumerKey,
		ClientSecret: consumerSecret,
		TokenURL:     tokenURL,
	}
	client := config.Client(context.Background())
	params := url.Values{}
	params["x_auth_mode"] = []string{"client_auth"}
	params["x_auth_username"] = []string{userName}
	params["x_auth_password"] = []string{passWord}
	resp, err := client.PostForm(apiURL, params)
	log.Printf("Status : %v\n", resp.StatusCode)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	content := string(b)
	m, _ := url.ParseQuery(content)
	oauthToken := m["oauth_token"][0]
	oauthSecret := m["oauth_token_secret"][0]
	userID := m["user_id"][0]
	screenName := m["screen_name"][0]
	userinfo := &UserInfo{
		OauthToken:       oauthToken,
		OauthTokenSecret: oauthSecret,
		UserID:           userID,
		ScreenName:       screenName,
	}
	return userinfo

}

func main() {
	info := Login()
	log.Printf("SCREEN NAME : %s\nUSER ID : %s\nOAUTH TOKEN : %s\nOAUTH TOKEN SECRET : %s\n",
		info.ScreenName,
		info.UserID,
		info.OauthToken,
		info.OauthTokenSecret)

}
