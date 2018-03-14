package Cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "kltn"
	Session, err = cluster.CreateSession()

	if err != nil {
		panic(err)
	}

	fmt.Println("Cassandra init done")

}
