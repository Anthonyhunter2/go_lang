package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml <- HTTP status codes

// These are for testing without doing a lookup to the database. Not intended for production.
var userName string
var passwd string

// randToken creates a random token that will be used for auth. This currently is a 16 byte byte slice that is returned
// as a lower case base64 encoded string.
func randToken() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		// Should never get never.
		panic(err)
	}
	return fmt.Sprintf("%x", b)
}

// CreateUser function is used on user creation. It should check to see if the user exists, hash the password and then store this password.
// It will return a success or failure based on possible conditions.
// This function accepts an HTTP POST requests. Headers should include the following:
// id = "Some username"
// pw = "Some Password"
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var err error

	// Debug log entries.
	log.Println(r.FormValue("id"))
	log.Println(r.FormValue("pw"))

	// This needs to be changed to do a lookup against the database. We will use globals for testing.
	if userName != r.FormValue("id") {
		userName = r.FormValue("id")
	} else {
		// Return a 400 "Bad Request" error code in the event that the username already exists.
		http.Error(w, "Invalid Username", 400)
		return
	}

	// Take the value of the form POST "pw" and bcrypt hash this password.
	pw := r.FormValue("pw")
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Unable to hash password", 400)
		log.Fatalln("Unable to Hash Password")
		log.Println(err)
		return
	}

	// If we have gotten here, we should be now doing an insert into the database with the password provided,
	// and the new username, info.

	passwd = string(hash)
	http.Error(w, "Success", 200)

}

// Login is used to verify the password provided and generate a token. This token will be stored in Redis
// So that it can be validated against during future calls. It will return a 200 with the token on sucessful login
// or a 401 on invalid login.
func Login(w http.ResponseWriter, r *http.Request) {

	// Logic should be placed here to do a lookup in the database for the user and return a hash.
	// Once the hash in returned, we do our comparison.

	if err := bcrypt.CompareHashAndPassword([]byte(passwd), []byte(r.FormValue("pw"))); err != nil {
		http.Error(w, "Invalid", 401)
		return
	}

	token := randToken()

	// Need to add logic here to check to see if the token already exists, if it does then create a new one.
	// We also was to store this token with the username and some type of identifier to ensure this token
	// is being used by the browser it was intended.
	log.Println("Generating token for " + r.FormValue("id") + ": " + token)
	RedisSet(token, r.FormValue("id"))
	http.Error(w, token, 200)
}

// Validate is used to determine if a token that is provided is currently valid. Validate will determine if
// this token is currently stored in Redis and return the id of the user it belongs to if it is valid.
func Validate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	log.Println("Checking the following token: " + params["token"])

	// Submit a request to Redis for valid token. If the token is not found, type err will be returned.
	resp, err := RedisGet(params["token"])
	if err != nil {
		http.Error(w, "Unknown", 401)
	}

	http.Error(w, resp, 200)
}

func main() {
	RedisConnect()
	router := mux.NewRouter()
	router.HandleFunc("/createuser", CreateUser).Methods("POST")
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/validate/{token}", Validate).Methods("GET")

	log.Fatalln(http.ListenAndServe(":9000", router))
}
