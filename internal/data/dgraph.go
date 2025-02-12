package data

import (
	"log"

	"github.com/dgraph-io/dgo/v240"
	"github.com/dgraph-io/dgo/v240/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TODO: login for security
func Init() (*dgo.Dgraph, func()) {
	conn, err := grpc.NewClient("localhost:9080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(api.NewDgraphClient(conn)), func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
