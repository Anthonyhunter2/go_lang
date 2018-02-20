//Package proGoDeletedb is a package to dynamically delete a db
package proGoDeletedb

import (
	"time"

	"github.com/rhinoman/couchdb-go"
)

func DelDB(dbip string, admin string, passwd string, dbname string) (string, error) {
	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection(dbip, 5984, timeout)
	auth := couchdb.BasicAuth{Username: admin, Password: passwd}
	db := conn.DeleteDB(dbname, &auth)
	if err != nil {
		return "DB not deleted", err
	} else if db == nil {
		return "DB has been deleted", err
	}
	return "", err
}
