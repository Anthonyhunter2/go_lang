package main

import (
	"fmt"
	"time"

	"github.com/rhinoman/couchdb-go"
)

type readdoc struct {
	_id   string
	_rev  string
	User  string
	Hole  int
	Score int
}

func main() {
	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("192.168.1.52", 5984, timeout)
	auth := couchdb.BasicAuth{Username: "golfer", Password: "Easy123!"}
	db := conn.SelectDB("project_under_par", &auth)
	query := readdoc{
		_id:   "",
		_rev:  "",
		User:  "horan116",
		Hole:  0,
		Score: 0,
	}
	v := couchdb.FindQueryParams{}

	//newDoc := readdoc{}
	//	rev, err := db.Read("9fad9897-d871-4cd7-ae95-96ad3c473aed", &newDoc, &v)
	e := db.Find(&query, &v)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(e)
	}
	fmt.Println(query)

}
