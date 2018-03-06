package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

//NewRound will create a new record for the user passed in
func NewRound(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// token ;+ params["token"]
	username := params["user"]
	//Need to add the validate function to check token
	//Validate(token)
	hexcode := CreateNew(username)
	if !bson.IsObjectIdHex(hexcode) {
		w.WriteHeader(401)
		w.Write([]byte("That ID was not found"))
		return
	}
	json.NewEncoder(w).Encode(hexcode)
}

//CurrentRound will create a new record for the user passed in
func GetScore(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// token ;+ params["token"]
	idstring := params["id"]
	//Need to add the validate function to check token
	//Validate(token)
	round, err := FindOneByID(idstring)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("That ID was not found"))
		return
	}
	json.NewEncoder(w).Encode(round.Round)
}

func NexHole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// token ;+ params["token"]
	idstring := params["id"]
	//Need to add the validate function to check token
	//Validate(token)
	nexthole := NextHole(idstring)
	if nexthole[:4] != "Hole" {
		w.WriteHeader(401)
		w.Write([]byte("That ID was not found"))
		return
	}
	json.NewEncoder(w).Encode(nexthole)
}
func PrevHole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// token ;+ params["token"]
	idstring := params["id"]
	//Need to add the validate function to check token
	//Validate(token)
	prevhole := PreviousHole(idstring)
	if prevhole[:4] != "Hole" {
		w.WriteHeader(401)
		w.Write([]byte("That ID was not found"))
		return
	}
	json.NewEncoder(w).Encode(prevhole)
}
func CurrentHole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// token ;+ params["token"]
	idstring := params["id"]
	//Need to add the validate function to check token
	//Validate(token)
	curhole := CurrentHolef(idstring)
	if curhole[:4] != "Hole" {
		w.WriteHeader(401)
		w.Write([]byte("That ID was not found"))
		return
	}
	json.NewEncoder(w).Encode(curhole)
}
func main() {
	Initdb()
	defer moncon.Close()
	//	muxport := os.Getenv("MUX_PORT")
	muxport := ":8080"
	//mux needs the proceeding ":" we are checking for it here and adding it if not specified
	if string(muxport[0]) != ":" {
		muxport = ":" + muxport
		//also if no port is specified the default port will be 8080
	} else if muxport == "" {
		muxport = ":8080"
	}
	router := mux.NewRouter()
	router.HandleFunc("/newround/{user}", NewRound).Methods("POST")
	router.HandleFunc("/currentround/{id}", GetScore).Methods("GET")
	router.HandleFunc("/nexthole/{id}", NexHole).Methods("PUT")
	router.HandleFunc("/prevhole/{id}", PrevHole).Methods("PUT")
	router.HandleFunc("/currenthole/{id}", CurrentHole).Methods("GET")
	//	router.HandleFunc("/getscore/{id}", GetScore).Methods("GET")
	// router.HandleFunc("/currentround/{id}", CurrentRound).Methods("GET")
	// router.HandleFunc("/currenthole/{id}", CurrentHole).Methods("GET")
	// router.HandleFunc("/prevhole/{id}", PrevHole).Methods("GET")
	// router.HandleFunc("/nexthole/{id}", NextHole).Methods("GET")
	// router.HandleFunc("/round/update&{id}&{hole}={num}", UpdateRound).Methods("POST")

	// This will the router functions when the validate function is incorpirated
	// router.HandleFunc("/newround/{token}?{user}", NewRound).Methods("POST")
	// router.HandleFunc("/currentround/{token}?{id}", GetScore).Methods("GET")
	// router.HandleFunc("/nexthole/{token}?{id}", NexHole).Methods("GET")
	// router.HandleFunc("/prevhole/{token}?{id}", PrevHole).Methods("GET")
	// router.HandleFunc("/currenthole/{token}?{id}", CurrentHole).Methods("GET")
	log.Fatal(http.ListenAndServe(muxport, router))
}
