package connect_nearby

import (
    "github.com/rhinoman/couchdb-go"
)



func connect_nearby {
    var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("172.17.0.2", 5984, timeout)
	auth := couchdb.BasicAuth{Username: "golfer", Password: "Easy123!"}
	db := conn.SelectDB("project_under_par", &auth)
}