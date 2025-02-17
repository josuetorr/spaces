package data

import (
	"fmt"
	"log"
	"os"

	kivik "github.com/go-kivik/kivik/v4"
)

func Init() (*kivik.Client, func()) {
	couchdbUrl := fmt.Sprintf("http://localhost:%s", os.Getenv("COUCHDB_PORT"))
	client, err := kivik.New("couch", couchdbUrl)
	if err != nil {
		log.Fatal("Could not init couchDB client: ", err)
	}

	return client, func() {
		if err := client.Close(); err != nil {
			log.Fatal("Error closing couchDB client: ", err)
		}
	}
}
