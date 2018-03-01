package main

import (
	"fmt"

	"github.com/globalsign/mgo"
)

var moncon *mgo.Session
var mondb *mgo.Database
var moncol *mgo.Collection

func initdb() {
	dbhost := "172.17.0.2"
	dbname := "proGodb"
	col := "app"
	//dbhost := os.Getenv("DBHOST")
	//dbname := os.Getenv("DBNAME")
	//col := os.Getenv("COLNAME")
	constring := dbhost + "/" + dbname
	connect, err := mgo.Dial(constring)
	session := connect.Copy()
	if err != nil {
		fmt.Println("Couldnt create connection")
		return
	}
	moncon = session
	mondb = session.DB(dbname)
	moncol = mondb.C(col)

	//fmt.Print(s.Find(nil).All(&Something))
	//fmt.Println("Creating DB connection")

}
