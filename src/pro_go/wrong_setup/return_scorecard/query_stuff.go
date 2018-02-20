package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"pro_go/packages/card"
	"reflect"
	"time"

	couchdb "github.com/rhinoman/couchdb-go"
)

func getJSON(jsonFile string) (string, int, int64) {
	dat, err := ioutil.ReadFile(jsonFile)
	incomingID := scoreCard.IDReturn{}
	err2 := json.Unmarshal(dat, &incomingID)
	if err2 != nil {
		fmt.Println(err2)
	} else if err != nil {
		fmt.Println(err)
	}
	parsing := reflect.ValueOf(incomingID)
	for i := 1; i < parsing.NumField(); i++ {
		if parsing.Field(i).Int() != 0 {
			return incomingID.ID, i, parsing.Field(i).Int()
		}
	}
	return incomingID.ID, 0, 0
}

func main() {
	timeout := time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("172.17.0.2", 5984, timeout)
	auth := couchdb.BasicAuth{Username: "golfer", Password: "Easy123!"}
	db := conn.SelectDB("project_under_par", &auth)
	if err != nil {
		fmt.Println(err)
	}
	IDNum, HoleNum, NewVal := getJSON("/home/anthony/go/src/pro_go/query.json")
	currentScoreCard := scoreCard.ScoreCard{}
	v := url.Values{}
	revNumber, _ := db.Read(IDNum, &currentScoreCard, &v)
	if err != nil {
		fmt.Println(err)
	}
	switch {
	case HoleNum == 1:
		currentScoreCard.Round.Hole1 = NewVal
	case HoleNum == 2:
		currentScoreCard.Round.Hole2 = NewVal
	case HoleNum == 3:
		currentScoreCard.Round.Hole3 = NewVal
	case HoleNum == 4:
		currentScoreCard.Round.Hole4 = NewVal
	case HoleNum == 5:
		currentScoreCard.Round.Hole5 = NewVal
	case HoleNum == 6:
		currentScoreCard.Round.Hole6 = NewVal
	case HoleNum == 7:
		currentScoreCard.Round.Hole7 = NewVal
	case HoleNum == 8:
		currentScoreCard.Round.Hole8 = NewVal
	case HoleNum == 9:
		currentScoreCard.Round.Hole9 = NewVal
	case HoleNum == 10:
		currentScoreCard.Round.Hole10 = NewVal
	case HoleNum == 11:
		currentScoreCard.Round.Hole11 = NewVal
	case HoleNum == 12:
		currentScoreCard.Round.Hole12 = NewVal
	case HoleNum == 13:
		currentScoreCard.Round.Hole13 = NewVal
	case HoleNum == 14:
		currentScoreCard.Round.Hole14 = NewVal
	case HoleNum == 15:
		currentScoreCard.Round.Hole15 = NewVal
	case HoleNum == 16:
		currentScoreCard.Round.Hole16 = NewVal
	case HoleNum == 17:
		currentScoreCard.Round.Hole17 = NewVal
	case HoleNum == 18:
		currentScoreCard.Round.Hole18 = NewVal
	}
	fmt.Println(revNumber, currentScoreCard)
	simble, err3 := db.Save(currentScoreCard, IDNum, revNumber)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(simble)
}
