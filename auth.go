package twitter_links

import (
	"github.com/ChimeraCoder/anaconda"
	"os"
)

func Api() (api *anaconda.TwitterApi) {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))

	oauth_token := os.Getenv("OAUTH_TOKEN")
	oauth_verifier := os.Getenv("OAUTH_VERIFIER")
	api = anaconda.NewTwitterApi(oauth_token, oauth_verifier)

	return
}
