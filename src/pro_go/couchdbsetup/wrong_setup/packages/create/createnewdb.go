package proGoCreateDB

import (
	"fmt"
	"time"

	couchdb "github.com/rhinoman/couchdb-go"
)

func ProGoCreate(dbname string, dbip string, admin string, passwd string) (string, error) {
	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection(dbip, 5984, timeout)
	if err != nil {
		fmt.Println(err)
	}
	auth := couchdb.BasicAuth{Username: admin, Password: passwd}
	conn.CreateDB(dbname, &auth)
	if err != nil {
		return "", err
	}
	return "DB Created", err

}
