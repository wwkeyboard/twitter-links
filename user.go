package twitter_links

import (
	"log"

	"github.com/boltdb/bolt"
)

type User struct {
	Token    string
	Verifier string
}

func GetUser() {
	db, err := bolt.Open("twitter-links.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
