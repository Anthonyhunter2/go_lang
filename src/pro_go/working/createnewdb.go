package main

import (
	"fmt"
	"time"

	couchdb "github.com/rhinoman/couchdb-go"
)

func main() {
	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("172.17.0.2", 5984, timeout)
	if err != nil {
		fmt.Println(err)
	}
	auth := couchdb.BasicAuth{Username: "golfer", Password: "Easy123!"}
	conn.CreateDB("project_under_par", &auth)
	if err != nil {
		fmt.Println("Created New db")
	} else {
		fmt.Println(err)
	}
}
