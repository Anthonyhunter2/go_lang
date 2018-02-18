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

// idReturn returns the passed in json formatted id number
// type idReturn struct {
// 	ID   string `json:"_id"`
// 	Hole int    `json:",omitempty"`
// }

func getJSON(jsonFile string) (scoreCard.IDReturn, int, int64) {
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
			return incomingID, i, parsing.Field(i).Int()
		}
	}
	return incomingID, 0, 0
}
func main() {
	timeout := time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("172.17.0.2", 5984, timeout)
	auth := couchdb.BasicAuth{Username: "golfer", Password: "Easy123!"}
	db := conn.SelectDB("project_under_par", &auth)
	if err != nil {
		fmt.Println(err)
	}
	IDData, RefNum, ActVal := getJSON("/home/anthony/go/src/pro_go/query.json")
	HoleNumber := fmt.Sprintf("%s%d", "Hole", RefNum)
	currentScoreCard := scoreCard.ScoreCard{}
	v := url.Values{}
	revNumber, _ := db.Read(IDData.ID, &currentScoreCard, &v)
	if err != nil {
		fmt.Println(err)
	}
	val := reflect.ValueOf(currentScoreCard.Round)
	for k := 0; k < val.NumField(); k++ {
		CurrentHole := val.Type().Field(k).Name
		if CurrentHole == HoleNumber {
			switch {
			case RefNum == 4:
				currentScoreCard.Round.Hole4 = ActVal
			}

		}
	}

	fmt.Println(revNumber, currentScoreCard)
	simble, err3 := db.Save(currentScoreCard, IDData.ID, revNumber)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(simble)
}
