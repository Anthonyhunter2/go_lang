package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"pro_go/packages/card"
	"time"

	"github.com/google/uuid"
	couchdb "github.com/rhinoman/couchdb-go"
)

// // Holes This holds just the struct for our golf holes
// type Holes struct {
// 	Hole1  int `json:"Hole 1"`
// 	Hole2  int `json:"Hole 2"`
// 	Hole3  int `json:"Hole 3"`
// 	Hole4  int `json:"Hole 4"`
// 	Hole5  int `json:"Hole 5"`
// 	Hole6  int `json:"Hole 6"`
// 	Hole7  int `json:"Hole 7"`
// 	Hole8  int `json:"Hole 8"`
// 	Hole9  int `json:"Hole 9"`
// 	Hole10 int `json:"Hole 10"`
// 	Hole11 int `json:"Hole 11"`
// 	Hole12 int `json:"Hole 12"`
// 	Hole13 int `json:"Hole 13"`
// 	Hole14 int `json:"Hole 14"`
// 	Hole15 int `json:"Hole 15"`
// 	Hole16 int `json:"Hole 16"`
// 	Hole17 int `json:"Hole 17"`
// 	Hole18 int `json:"Hole 18"`
// }

// //ScoreCard Holds our complete data set, including name & date
// type ScoreCard struct {
// 	//	Date  time.Time
// 	User  string `json:"User"`
// 	Round Holes
// }

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	timeout := time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("172.17.0.2", 5984, timeout)
	auth := couchdb.BasicAuth{Username: "golfer", Password: "Easy123!"}
	db := conn.SelectDB("project_under_par", &auth)
	dat, err := ioutil.ReadFile("/home/anthony/go/src/pro_go/holedoc.json")
	check(err)

	docToSend := scoreCard.ScoreCard{}
	err2 := json.Unmarshal(dat, &docToSend)
	check(err2)
	passedIn := os.Args[1:]
	if len(passedIn) == 0 {
		newID := uuid.New().String()
		simble, err := db.Save(docToSend, newID, "")
		fmt.Println(simble, newID)
		check(err)
	} else {

	}
}
