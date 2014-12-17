package main

import (
	"log"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/ChimeraCoder/anaconda"
)

type Link struct {
	Url string
	Sender string
	Text string
}

func SignIn(c *gin.Context) { //w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	anaconda.SetConsumerKey("c5LBQ5awsg1XwVyCb5oOhh84W")
	anaconda.SetConsumerSecret("rqebVQvAwR3Os5jYDUToM5YlzeFFzCkR6dg9SO86JSD4Hzzc2K")

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
	anaconda.SetConsumerKey("c5LBQ5awsg1XwVyCb5oOhh84W")
	anaconda.SetConsumerSecret("rqebVQvAwR3Os5jYDUToM5YlzeFFzCkR6dg9SO86JSD4Hzzc2K")

	oauth_token := "22762491-Z0o9Ag9J8zUZjwVFQ0cgRRNpulfz9EnijVNzNPH9t"
	oauth_verifier := "3IjBz7P0FZrEKVqo5moZq06QPdz97NxeGyB3nnlhg5q77"
	api := anaconda.NewTwitterApi(oauth_token, oauth_verifier)

	searchResult, err := api.GetHomeTimeline(nil)
	if err != nil {
		log.Print(err)
		return
	}

	var links []Link
	for _ , tweet := range searchResult {
		for _ , url := range tweet.Entities.Urls {
			l := Link{
				Url: url.Expanded_url,
				Sender: tweet.User.Name,
				Text: tweet.Text,
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

func SignInCallback(c *gin.Context){
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
