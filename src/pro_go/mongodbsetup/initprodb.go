package main

import (
	"fmt"

	"github.com/globalsign/mgo"
)

var moncon *mgo.Session
var mondb *mgo.Database
var moncol *mgo.Collection

func initdb() {
	var err error
	dbhost := "172.17.0.2"
	dbname := "proGodb"
	col := "app"
	//dbhost := os.Getenv("DBHOST")
	//dbname := os.Getenv("DBNAME")
	//col := os.Getenv("COLNAME")
	moncon, err = mgo.Dial(dbhost + "/" + dbname)
	if err != nil {
		fmt.Println("Couldnt create connection")
		return
	}
	mondb = moncon.DB(dbname)
	moncol = mondb.C(col)

}
