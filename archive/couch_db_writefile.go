package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rhinoman/couchdb-go"
)

type TestDocument struct {
	Title string
	Note  string
}

func main() {
	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("172.17.0.2", 5984, timeout)
	auth := couchdb.BasicAuth{Username: "golfer", Password: "Easy123!"}
	db := conn.SelectDB("project_under_par", &auth)
	theDoc := TestDocument{Title: "The first doc uploaded", Note: "I guess scores or could go here"}
	idNew := uuid.New().String()
	simble, err := db.Save(theDoc, idNew, "")
	fmt.Println(simble)
	if err != nil {
		fmt.Println(err)
	}
}
