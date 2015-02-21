package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ChimeraCoder/anaconda"
	//"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"

	"github.com/wwkeyboard/twitter-links"
)

func SignIn(c *gin.Context) {
	twitter_links.SetKeys()

	// Need to generate an ID for the user first, and include it in the oauth callback
	redirect_url, _, err := anaconda.AuthorizationURL("http://www.wellwornkeyboard.com/signin/callback/123456")
	if err != nil {
		log.Print(err)
		c.String(200, "There was an error")
	} else {
		c.Redirect(302, redirect_url)
		return
	}
}

func ListLinks(c *gin.Context) {
	u := new(twitter_links.User)
	api := twitter_links.Api(*u)

	searchResult, err := api.GetHomeTimeline(nil)
	if err != nil {
		log.Print(err)
		return
	}

	var links []twitter_links.Link
	for _, tweet := range searchResult {
		for _, url := range tweet.Entities.Urls {
			l := twitter_links.Link{
				Url:    url.Expanded_url,
				Sender: tweet.User.Name,
				Text:   tweet.Text,
			}
			links = append(links, l)
		}
	}

	res, err := json.Marshal(links)
	if err != nil {
		c.String(500, err.Error())
	}

	c.String(200, string(res))
}

func SignInCallback(c *gin.Context) {
	query := c.Request.URL.Query()
	token := query.Get("oauth_token")
	verifier := query.Get("oauth_verifier")

	realCreds := twitter_links.GetCredsFromCallback(token, verifier)
	fmt.Print(realCreds)

	u := twitter_links.User{
		Username: "whatev",
		Token:    realCreds.Token,
		Verifier: verifier,
	}
	twitter_links.SaveUser(&u)
}

func main() {
	router := gin.Default()
	router.GET("/signin", SignIn)
	router.GET("/signin/callback/:user_id", SignInCallback)
	router.GET("/links", ListLinks)

	router.Run(":8080")
}
