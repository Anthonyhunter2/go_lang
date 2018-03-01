package main

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
)

//Person is a test struct for now
type Person struct {
	//	ID   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	ID   bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Name string        `json:"Name" bson:"Name"`
	Hole string        `json:"Hole" bson:"Hole"`
}

func createNew() {
	newID := bson.NewObjectId()
	playerone := &Person{ID: newID, Name: "Anthony", Hole: "2"}
	err := moncol.Insert(playerone)
	if err != nil {
		fmt.Println(err)
	}
}

func findOneByID(idstring string) {
	result := Person{}
	err := moncol.Find(bson.M{"_id": bson.ObjectIdHex(idstring)}).One(&result)
	if err != nil {
		fmt.Println("Couldn't Find that ID string")
		return
	}
	fmt.Println(result)
}
func updateOneByID(idstring string) {
	result := &Person{Name: "Steve"}
	err := moncol.Update(bson.M{"_id": bson.ObjectIdHex(idstring)}, &result)
	if err != nil {
		fmt.Println("Couldn't Find that ID string")
		return
	}
	fmt.Println(result)
}
func main() {
	initdb()
	defer moncon.Close()
	updateOneByID("5a975a41c9a0d81b5896bfa2")
}
