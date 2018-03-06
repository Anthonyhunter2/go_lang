package main

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var moncon *mgo.Session
var mondb *mgo.Database
var moncol *mgo.Collection

//Initdb initalizes the db connection
func Initdb() {
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

//Holes holds just our hole information
type Holes struct {
	Hole1  int `json:"Hole 1" bson:"Hole 1"`
	Hole2  int `json:"Hole 2" bson:"Hole 2"`
	Hole3  int `json:"Hole 3" bson:"Hole 3"`
	Hole4  int `json:"Hole 4" bson:"Hole 4"`
	Hole5  int `json:"Hole 5" bson:"Hole 5"`
	Hole6  int `json:"Hole 6" bson:"Hole 6"`
	Hole7  int `json:"Hole 7" bson:"Hole 7"`
	Hole8  int `json:"Hole 8" bson:"Hole 8"`
	Hole9  int `json:"Hole 9" bson:"Hole 9"`
	Hole10 int `json:"Hole 10" bson:"Hole 10"`
	Hole11 int `json:"Hole 11" bson:"Hole 11"`
	Hole12 int `json:"Hole 12" bson:"Hole 12"`
	Hole13 int `json:"Hole 13" bson:"Hole 13"`
	Hole14 int `json:"Hole 14" bson:"Hole 14"`
	Hole15 int `json:"Hole 15" bson:"Hole 15"`
	Hole16 int `json:"Hole 16" bson:"Hole 16"`
	Hole17 int `json:"Hole 17" bson:"Hole 17"`
	Hole18 int `json:"Hole 18" bson:"Hole 18"`
}

//Person is a test struct for now
type Person struct {
	ID      bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Date    string        `json:"Date" bson:"Date"`
	Name    string        `json:"Name" bson:"Name"`
	CurHole string        `json:"Current Hole" bson:"Current Hole"`
	Round   Holes         `json:"Round" bson:"Round"`
}
