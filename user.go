package twitter_links

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/boltdb/bolt"
)

type User struct {
	Username string
	Token    string
	Verifier string
}

func openDB() (db *bolt.DB) {
	db, err := bolt.Open("twitter-links.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func SaveUser(u *User) (err error) {
	db := openDB()

	// first encode the user to a gob
	var serializedUser bytes.Buffer
	enc := gob.NewEncoder(&serializedUser)

	// Encode (send) some values.
	err = enc.Encode(u)
	if err != nil {
		log.Print("encode error:", err)
		return
	}

	// then write that gob to the DB
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return err
		}
		return b.Put([]byte(u.Username), serializedUser.Bytes())
	})
	return
}

func GetUser() (user *User, err error) {
	db := openDB()

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("posts"))
		if err != nil {
			return err
		}
		return b.Put([]byte("2015-01-01"), []byte("My New Year post"))
	})

	defer db.Close()
	return
}
