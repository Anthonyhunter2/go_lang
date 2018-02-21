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

//We're going to pass in the static variables from environemnt variables passed in the docker file
//dbAddr := os.Getenv("DBSERVER")
//dbAdmin := os.Getenv("DBADMIN")
//dbPasswd := os.Getenv("DBPASSWORD")
//dbName := os.Getenv("DBNAME")

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
	card := ScoreCard{User: username, CHole: 1}
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
	// This grabs the current hole value and sets it to the hole that was just updated
	curHoleUpdated, _ := strconv.ParseInt(hole[len(hole)-1:], 10, 64)
	reflect.ValueOf(&card.CHole).Elem().SetInt(curHoleUpdated)

	_, err2 := db.Save(&card, key, revnumber)
	if err2 != nil {
		w.WriteHeader(401)
		w.Write([]byte("Something went wrong saving new scorecard"))
	}
	json.NewEncoder(w).Encode(card)
}

//CurrentHole will show the current hole and score
func CurrentHole(w http.ResponseWriter, r *http.Request) {
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
	curHoleNum := strconv.Itoa(card.CHole)
	curHoleString := "Hole" + curHoleNum
	curHoleVal := reflect.ValueOf(&card.Round).Elem().FieldByName(curHoleString).Int()
	type jresponse struct {
		Hole  string
		Score int64
	}
	json.NewEncoder(w).Encode(jresponse{Hole: curHoleNum, Score: curHoleVal})
}

//PrevHole will show the current hole and score
func PrevHole(w http.ResponseWriter, r *http.Request) {
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
	if card.CHole-1 == 0 {
		card.CHole = 19
	}
	prevHoleNum := strconv.Itoa(card.CHole - 1)
	prevHoleString := "Hole" + prevHoleNum
	prevHoleVal := reflect.ValueOf(&card.Round).Elem().FieldByName(prevHoleString).Int()
	type jresponse struct {
		Hole  string
		Score int64
	}
	json.NewEncoder(w).Encode(jresponse{Hole: prevHoleNum, Score: prevHoleVal})
}

//NextHole will show the current hole and score
func NextHole(w http.ResponseWriter, r *http.Request) {
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
	if card.CHole+1 == 19 {
		card.CHole = 0
	}
	NextHoleNum := strconv.Itoa(card.CHole + 1)
	NextHoleString := "Hole" + NextHoleNum
	NextHoleVal := reflect.ValueOf(&card.Round).Elem().FieldByName(NextHoleString).Int()
	type jresponse struct {
		Hole  string
		Score int64
	}
	json.NewEncoder(w).Encode(jresponse{Hole: NextHoleNum, Score: NextHoleVal})
}

//CurrentRound will return just the current round
func CurrentRound(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	card := ScoreCard{}
	v := url.Values{}
	_, err := db.Read(key, &card, &v)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Make sure your passing all arguments"))
		return
	}
	json.NewEncoder(w).Encode(&card.Round)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getscore/{id}", GetScore).Methods("GET")
	router.HandleFunc("/currentround/{id}", CurrentRound).Methods("GET")
	router.HandleFunc("/currenthole/{id}", CurrentHole).Methods("GET")
	router.HandleFunc("/prevhole/{id}", PrevHole).Methods("GET")
	router.HandleFunc("/nexthole/{id}", NextHole).Methods("GET")
	router.HandleFunc("/newround/{user}", NewRound).Methods("POST")
	router.HandleFunc("/round/update&{id}&{hole}={num}", UpdateRound).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
