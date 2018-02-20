package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/gorilla/mux"

	couchdb "github.com/rhinoman/couchdb-go"
)

// Here we've made our varables global so we can take full advantage of the couchdb lib
var timeout = time.Duration(500 * time.Millisecond)
var conn, err = couchdb.NewConnection("172.17.0.2", 5984, timeout)
var auth = couchdb.BasicAuth{Username: "golfer", Password: "Easy123!"}
var db = conn.SelectDB("project_under_par", &auth)

//Holes just keeps the contents of our full 18 hole course scores
type Holes struct {
	Hole1  int64 `json:"Hole 1, omitempty"`
	Hole2  int64 `json:"Hole 2, omitempty"`
	Hole3  int64 `json:"Hole 3, omitempty"`
	Hole4  int64 `json:"Hole 4, omitempty"`
	Hole5  int64 `json:"Hole 5, omitempty"`
	Hole6  int64 `json:"Hole 6, omitempty"`
	Hole7  int64 `json:"Hole 7, omitempty"`
	Hole8  int64 `json:"Hole 8, omitempty"`
	Hole9  int64 `json:"Hole 9, omitempty"`
	Hole10 int64 `json:"Hole 10, omitempty"`
	Hole11 int64 `json:"Hole 11, omitempty"`
	Hole12 int64 `json:"Hole 12, omitempty"`
	Hole13 int64 `json:"Hole 13, omitempty"`
	Hole14 int64 `json:"Hole 14, omitempty"`
	Hole15 int64 `json:"Hole 15, omitempty"`
	Hole16 int64 `json:"Hole 16, omitempty"`
	Hole17 int64 `json:"Hole 17, omitempty"`
	Hole18 int64 `json:"Hole 18, omitempty"`
}

//ScoreCard Holds our complete data set, including name & date
type ScoreCard struct {
	User  string `json:"User"`
	CHole int    `json:"Current Hole"`
	Round Holes
}

//GetScore returns the current score based on id passed to the /getscore endpoint
func GetScore(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	card := ScoreCard{}
	v := url.Values{}
	_, err := db.Read(ID, &card, &v)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("That ID was not found"))
		return
	}
	json.NewEncoder(w).Encode(card)
}

//NewRound will create a new record for the user passed in
func NewRound(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["user"]
	randomuuid := uuid.New().String()
	card := ScoreCard{User: username}
	_, err := db.Save(&card, randomuuid, "")
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("That ID was not found"))
		return
	}
	json.NewEncoder(w).Encode(randomuuid)
}

//UpdateRound will create a new record for the user passed in
func UpdateRound(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	hole := params["hole"]
	newVal, _ := strconv.ParseInt(params["num"], 10, 64)
	card := ScoreCard{}
	v := url.Values{}
	revnumber, err := db.Read(key, &card, &v)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Make sure your passing all arguments"))
		return
	}
	//This updates the current hole# value based on the "Hole" input in the incoming request
	reflect.ValueOf(&card.Round).Elem().FieldByName(hole).SetInt(newVal)
	_, err2 := db.Save(&card, key, revnumber)
	if err2 != nil {
		w.WriteHeader(401)
		w.Write([]byte("Something went wrong saving new scorecard"))
	}
	json.NewEncoder(w).Encode(card)
}

//CurrentHole will show the current hole
// func CurrentHole(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	ID := params["id"]
// 	card := ScoreCard{User: username}
// 	_, err := db.Save(&card, randomuuid, "")
// 	if err != nil {
// 		w.WriteHeader(401)
// 		w.Write([]byte("That ID was not found"))
// 		return
// 	}
// 	json.NewEncoder(w).Encode(randomuuid)
// }
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getscore/{id}", GetScore).Methods("GET")
	//	router.HandleFunc("/currenthole/{id}", CurrentHole).Methods("GET")
	router.HandleFunc("/newround/{user}", NewRound).Methods("POST")
	router.HandleFunc("/round/update&{id}&{hole}={num}", UpdateRound).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
