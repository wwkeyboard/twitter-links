package main

import (
	"log"

	"encoding/json"

	"github.com/ChimeraCoder/anaconda"
	//"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"

	"github.com/wwkeyboard/twitter-links"
)

type Link struct {
	Url    string
	Sender string
	Text   string
}

func SignIn(c *gin.Context) {
	twitter_links.SetKeys()

	redirect_url, _, err := anaconda.AuthorizationURL("http://www.wellwornkeyboard.com/signin/callback")
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

	var links []Link
	for _, tweet := range searchResult {
		for _, url := range tweet.Entities.Urls {
			l := Link{
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
	/*	oauth_token := r.URL.Query().Get("oauth_token")
		oauth_verifier := r.URL.Query().Get("oauth_verifier")
		fmt.Fprintf(w, "oauth_token    %s\n", oauth_token)
		fmt.Fprintf(w, "oauth_verifier %s\n", oauth_verifier)
	*/
}

func main() {
	router := gin.Default()
	router.GET("/signin", SignIn)
	router.GET("/signin/callback", SignInCallback)
	router.GET("/links", ListLinks)

	router.Run(":8080")
}
