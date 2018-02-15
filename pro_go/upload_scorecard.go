package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Holes This holds just the struct for our golf holes
type Holes struct {
	Hole1  int `json:"Hole 1"`
	Score1 int `json:"Score 1"`
	Hole2  int `json:"Hole 2"`
	Score2 int `json:"Score 2"`
	// Hole3  int `json:"Hole 3"`
	// Score3  int `json:"Score 3"`
	// Hole4   int `json:"Hole 4"`
	// Score4  int `json:"Score 4"`
	// Hole5   int `json:"Hole 5"`
	// Score5  int `json:"Score 5"`
	// Hole6   int `json:"Hole 6"`
	// Score6  int `json:"Score 6"`
	// Hole7   int `json:"Hole 7"`
	// Score7  int `json:"Score 7"`
	// Hole8   int `json:"Hole 8"`
	// Score8  int `json:"Score 8"`
	// Hole9   int `json:"Hole 9"`
	// Score9  int `json:"Score 9"`
	// Hole10  int `json:"Hole 10"`
	// Score10 int `json:"Score 10"`
	// Hole11  int `json:"Hole 11"`
	// Score11 int `json:"Score 11"`
	// Hole12  int `json:"Hole 12"`
	// Score12 int `json:"Score 12"`
	// Hole13  int `json:"Hole 13"`
	// Score13 int `json:"Score 13"`
	// Hole14  int `json:"Hole 14"`
	// Score14 int `json:"Score 14"`
	// Hole15  int `json:"Hole 15"`
	// Score15 int `json:"Score 15"`
	// Hole16  int `json:"Hole 16"`
	// Score16 int `json:"Score 16"`
	// Hole17  int `json:"Hole 17"`
	// Score17 int `json:"Score 17"`
	// Hole18  int `json:"Hole 18"`
	// Score18 int `json:"Score 18"`
}

//ScoreCard Holds our complete data set, including name & date
type ScoreCard struct {
	//	Date  time.Time
	User  string `json:"User"`
	Round Holes
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	// timeout := time.Duration(500 * time.Millisecond)
	// conn, err := couchdb.NewConnection("172.17.0.2", 5984, timeout)
	// auth := couchdb.BasicAuth{Username: "golfer", Password: "Easy123!"}
	// db := conn.SelectDB("project_under_par", &auth)
	dat, err := ioutil.ReadFile("/home/anthony/go/pro_go/holedoc.json")
	check(err)
	docToSend := ScoreCard{}
	err2 := json.Unmarshal(dat, &docToSend)
	fmt.Println(docToSend.Round.Score1)
	check(err2)
	//	passedIn := os.Args[1:]
	// if len(passedIn) == 0 {
	// 	newID := uuid.New().String()
	// 	simble, err := db.Save(docToSend, newID, "")
	// 	fmt.Println(simble, newID)
	// 	check(err)
	// } else {

	// }
}
