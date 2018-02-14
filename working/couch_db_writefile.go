package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rhinoman/couchdb-go"
)

type TestDocument struct {
	Date time.Time
	User string
	//Round map[string]interface{
	Hole1 map[string]int
	Hole2 map[string]int
}

func main() {
	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("172.17.0.2", 5984, timeout)
	auth := couchdb.BasicAuth{Username: "golfer", Password: "Easy123!"}
	db := conn.SelectDB("project_under_par", &auth)
	theDoc := TestDocument{Date: time.Now(), User: "Anthony",
		Hole1: map[string]int{"Score": 56},
		Hole2: map[string]int{"Score": 110}}

	newID := uuid.New().String()
	simble, err := db.Save(theDoc, newID, "")
	fmt.Println(simble)
	if err != nil {
		fmt.Println(err)
	}
}

// "Hole 1": map[string]int{
// 	"Score": 5},
// "Hole 2": map[string]int{
// 	"Score": 3},
// "Hole 3": map[string]int{
// 	"Score": 3},
// "Hole 4": map[string]int{
// 	"Score": 4},
// "Hole 5": map[string]int{
// 	"Score": 76},
// "Hole 6": map[string]int{
// 	"Score": 33},
// "Hole 7": map[string]int{
// 	"Score": 31},
// "Hole 8": map[string]int{
// 	"Score": 3},
// "Hole 9": map[string]int{
