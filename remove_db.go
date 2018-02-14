package main

import (
	"fmt"
	"time"

	"github.com/rhinoman/couchdb-go"
)

func main() {
	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("172.17.0.2", 5984, timeout)
	auth := couchdb.BasicAuth{Username: "golfer", Password: "Easy123!"}
	db := conn.DeleteDB("project_under_par", &auth)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)

}
