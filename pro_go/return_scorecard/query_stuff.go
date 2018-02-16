package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	scoreCard "github.com/pro_go/cards"
	couchdb "github.com/rhinoman/couchdb-go"
)

// queryb)trying to grab the unique id were tring to query
type idReturn struct {
	ID string `json:"_id"`
}

func main() {
	timeout := time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("172.17.0.2", 5984, timeout)
	auth := couchdb.BasicAuth{Username: "golfer", Password: "Easy123!"}
	db := conn.SelectDB("project_under_par", &auth)
	dat, err := ioutil.ReadFile("/home/anthony/go/pro_go/query.json")
	if err != nil {
		fmt.Println(err)
	}
	incomingID := idReturn{}
	err2 := json.Unmarshal(dat, &incomingID)
	if err2 != nil {
		fmt.Println(err2)
	}
	currentScoreCard := scoreCard.ScoreCard{}
	v := url.Values{}
	revNumber, _ := db.Read(incomingID.ID, &currentScoreCard, &v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(revNumber, currentScoreCard)

}
