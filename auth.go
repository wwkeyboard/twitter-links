package twitter_links

import (
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
)

func AppSecret() string {
	return os.Getenv("CONSUMER_SECRET")
}

func GetCredsFromCallback(token string, verifier string) (realCreds *oauth.Credentials) {
	tmpCreds := &oauth.Credentials{Token: token, Secret: AppSecret()}

	realCreds, _, err := anaconda.GetCredentials(tmpCreds, verifier)
	if err != nil {
		log.Print(err)
		return
	}

	return
}

func SetKeys() {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(AppSecret())
}

func Api(u User) (api *anaconda.TwitterApi) {
	SetKeys()

	oauth_token := os.Getenv("OAUTH_TOKEN")
	oauth_verifier := os.Getenv("OAUTH_VERIFIER")
	api = anaconda.NewTwitterApi(oauth_token, oauth_verifier)

	return
}
