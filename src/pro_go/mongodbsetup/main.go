package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
)

//Holes holds just our hole information
type Holes struct {
	Hole1  int `json:"Hole 1" bson:"Hole 1"`
	Hole2  int `json:"Hole 2" bson:"Hole 2"`
	Hole3  int `json:"Hole 3" bson:"Hole 3"`
	Hole4  int `json:"Hole 4" bson:"Hole 4"`
	Hole5  int `json:"Hole 5" bson:"Hole 5"`
	Hole6  int `json:"Hole 6" bson:"Hole 6"`
	Hole7  int `json:"Hole 7" bson:"Hole 7"`
	Hole8  int `json:"Hole 8" bson:"Hole 8"`
	Hole9  int `json:"Hole 9" bson:"Hole 9"`
	Hole10 int `json:"Hole 10" bson:"Hole 10"`
	Hole11 int `json:"Hole 11" bson:"Hole 11"`
	Hole12 int `json:"Hole 12" bson:"Hole 12"`
	Hole13 int `json:"Hole 13" bson:"Hole 13"`
	Hole14 int `json:"Hole 14" bson:"Hole 14"`
	Hole15 int `json:"Hole 15" bson:"Hole 15"`
	Hole16 int `json:"Hole 16" bson:"Hole 16"`
	Hole17 int `json:"Hole 17" bson:"Hole 17"`
	Hole18 int `json:"Hole 18" bson:"Hole 18"`
}

//Person is a test struct for now
type Person struct {
	ID    bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Date  string        `json:"Date" bson:"Date"`
	Name  string        `json:"Name" bson:"Name"`
	Round Holes         `json:"Round" bson:"Round"`
}

func createNew(golfer string) string {
	newID := bson.NewObjectId()
	currentdate := time.Now().Format("2006-01-02")
	playerone := &Person{ID: newID, Date: currentdate, Name: golfer}
	err := moncol.Insert(playerone)
	if err != nil {
		fmt.Println(err)
	}
	return newID.Hex()
}

func findOneByID(idstring string) (*Person, error) {
	result := &Person{}
	err := moncol.Find(bson.M{"_id": bson.ObjectIdHex(idstring)}).One(&result)
	if err != nil {
		fmt.Println("Couldn't Find that ID string")
		return result, err
	}
	return result, err
}
func updateOneByID(idstring string, golfer string) string {
	result := &Person{}
	err := moncol.Update(bson.M{"_id": bson.ObjectIdHex(idstring)}, &result)
	if err != nil {
		return "Couldn't Find that ID string"
	}
	return "Updated"
}
func deleteByID(idstring string) string {
	err := moncol.Remove(bson.M{"_id": bson.ObjectIdHex(idstring)})
	if err != nil {
		return "Couldn't Find that ID string"
	}
	return "Deleted"
}

func updateSingleHole(idstring string, feild string, score int) string {
	// tests to make sure the feild to update is given in the correct systax
	if !strings.Contains(feild, " ") {
		feild = feild[:4] + " " + feild[4:]
	}
	//couldn't find an easy way to test to make sure the first char of feild was uppercase
	//so were just setting it to uppercase here
	feild = strings.ToUpper(string(feild[0])) + feild[1:]
	findDoc := bson.M{"_id": bson.ObjectIdHex(idstring)}
	updateDoc := bson.M{"$set": bson.M{"Round." + feild: score}}
	err := moncol.Update(findDoc, updateDoc)
	if err != nil {
		return "Could not update record"
	}
	return "Updated"
}
func main() {
	initdb()
	defer moncon.Close()
	//fmt.Println(createNew("Anthony"))
	fmt.Println(updateSingleHole("5a9aceeec9a0d85451c38031", "hole12", 3))
	// fmt.Println(findOneByID("5a9aceeec9a0d85451c38031"))
	//fmt.Println(deleteByID("5a98afa0c9a0d86be221e816")
	//fmt.Println(updateSingleHole("5a98afa0c9a0d86be221e816", "Hole 1", 3))
}
