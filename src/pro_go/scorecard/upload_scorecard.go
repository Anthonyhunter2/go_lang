package makeScores

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func UploadScores() {
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
