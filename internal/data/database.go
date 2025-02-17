package data

import (
	"fmt"
	"log"
	"os"

	kivik "github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb" // The CouchDB driver
)

func Init(db string, opts ...kivik.Option) (*kivik.DB, func()) {
	couchdbUrl := fmt.Sprintf("http://localhost:%s", os.Getenv("COUCHDB_PORT"))
	client, err := kivik.New("couch", couchdbUrl)
	if err != nil {
		log.Fatal("Could not init couchDB client: ", err)
	}

	return client.DB(db, opts...), func() {
		if err := client.Close(); err != nil {
			log.Fatal("Error closing couchDB client: ", err)
		}
	}
}
