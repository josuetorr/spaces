package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	kivik "github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb" // The CouchDB driver
)

func Init(dbName string, opts ...kivik.Option) (*kivik.DB, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := os.Getenv("COUCHDB_USER")
	passwd := os.Getenv("COUCHDB_PASSWORD")
	port := os.Getenv("COUCHDB_PORT")
	couchdbUrl := fmt.Sprintf("http://%s:%s@localhost:%s", user, passwd, port)
	client, err := kivik.New("couch", couchdbUrl)
	if err != nil {
		log.Fatal("Could not init couchDB client: ", err)
	}

	exists, err := client.DBExists(ctx, dbName, opts...)
	if !exists {
		if err := client.CreateDB(ctx, dbName, opts...); err != nil {
			log.Fatalf("Could not create '%s' db. Exiting with error: %v", dbName, err)
		}
	}

	return client.DB(dbName, opts...), func() {
		if err := client.Close(); err != nil {
			log.Fatal("Error closing couchDB client: ", err)
		}
	}
}
