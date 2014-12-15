package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/julienschmidt/httprouter"
	"github.com/ChimeraCoder/anaconda"
)

func SignIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	anaconda.SetConsumerKey("c5LBQ5awsg1XwVyCb5oOhh84W")
//	anaconda.SetConsumerSecret("rqebVQvAwR3Os5jYDUToM5YlzeFFzCkR6dg9SO86JSD4Hzzc2K")

	redirect_url, _, err := anaconda.AuthorizationURL("http://www.wellwornkeyboard.com/signin/callback")
	if err != nil {
		log.Print(err)
		fmt.Fprint(w, "There was an error")
	} else {
		log.Print("sending to callback")
		fmt.Fprint(w, redirect_url)

		// This doesn't work, figure out why
		http.Redirect(w,r, redirect_url, 302)
		return

	}
}

func ListLinks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	for _ , tweet := range searchResult {
		fmt.Fprintf(w, "%s: %s\n", tweet.User.Name, tweet.Text)
		for _ , url := range tweet.Entities.Urls {
			fmt.Fprintf(w, "<a href='%s'>%s</a>", url.Expanded_url, url.Expanded_url)
		}
	}
}

func SignInCallback(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	oauth_token := r.URL.Query().Get("oauth_token")
	oauth_verifier := r.URL.Query().Get("oauth_verifier")
	fmt.Fprintf(w, "oauth_token    %s\n", oauth_token)
	fmt.Fprintf(w, "oauth_verifier %s\n", oauth_verifier)
}

func main() {
	router := httprouter.New()
	router.GET("/signin", SignIn)
	router.GET("/signin/callback", SignInCallback)
	router.GET("/links", ListLinks)

	log.Fatal(http.ListenAndServe(":8080", router))
}
