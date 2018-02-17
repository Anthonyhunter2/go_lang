package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	scoreCard "pro_go/packages/card"
	"time"

	couchdb "github.com/rhinoman/couchdb-go"
)

// queryb)trying to grab the unique id were tring to query
type idReturn struct {
	ID string `json:"_id"`
}

func getJSON(jsonFile string) idReturn {
	dat, err := ioutil.ReadFile(jsonFile)
	incomingID := idReturn{}
	err2 := json.Unmarshal(dat, &incomingID)
	if err2 != nil {
		fmt.Println(err2)
	} else if err != nil {
		fmt.Println(err)
	}
	return incomingID

}

func main() {
	timeout := time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("172.17.0.2", 5984, timeout)
	auth := couchdb.BasicAuth{Username: "golfer", Password: "Easy123!"}
	db := conn.SelectDB("project_under_par", &auth)
	if err != nil {
		fmt.Println(err)
	}
	IDData := getJSON("/home/anthony/go/src/pro_go/query.json")
	currentScoreCard := scoreCard.ScoreCard{}
	v := url.Values{}
	revNumber, _ := db.Read(IDData.ID, &currentScoreCard, &v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(revNumber, currentScoreCard)

}
