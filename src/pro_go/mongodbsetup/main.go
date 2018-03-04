package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
)

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
	//Here we're making sure the hole number passed is betweeen 1 and 18
	holeNum, _ := strconv.Atoi(string(feild[5:]))
	if 1 <= holeNum && holeNum >= 18 {
		return "Thats not a vaild hole number"
	}
	findDoc := bson.M{"_id": bson.ObjectIdHex(idstring)}
	updateDoc := bson.M{"$set": bson.M{"Round." + feild: score, "Current Hole": feild}}
	err := moncol.Update(findDoc, updateDoc)
	if err != nil {
		return "Could not update record"
	}
	return "Updated"
}

func currentHole(idstring string) string {
	result := &Person{}
	err := moncol.Find(bson.M{"_id": bson.ObjectIdHex(idstring)}).One(&result)
	if err != nil {
		return "Couldn't Find That ID String"
	}
	return result.CurHole
}
func nextHole(idstring string) string {
	result := &Person{}
	err := moncol.Find(bson.M{"_id": bson.ObjectIdHex(idstring)}).One(&result)
	if err != nil {
		return "Couldn't Find That ID String"
	}
	holenumstr := result.CurHole[5:]
	holenum, _ := strconv.Atoi(holenumstr)
	newHoleNum := holenum + 1
	if newHoleNum == 19 {
		newHoleNum = 1
	}
	newHoleStr := strconv.Itoa(newHoleNum)
	return "Hole " + newHoleStr
}
func previousHole(idstring string) string {
	result := &Person{}
	err := moncol.Find(bson.M{"_id": bson.ObjectIdHex(idstring)}).One(&result)
	if err != nil {
		return "Couldn't Find That ID String"
	}
	holenumstr := result.CurHole[5:]
	holenum, _ := strconv.Atoi(holenumstr)
	newHoleNum := holenum - 1
	if newHoleNum == 0 {
		newHoleNum = 18
	}
	newHoleStr := strconv.Itoa(newHoleNum)
	return "Hole " + newHoleStr
}
func main() {
	initdb()
	defer moncon.Close()
	//fmt.Println(createNew("Anthony"))
	// fmt.Println(updateSingleHole("5a9aceeec9a0d85451c38031", "hole1", 7))
	// fmt.Println(findOneByID("5a9aceeec9a0d85451c38031"))
	//fmt.Println(deleteByID("5a98afa0c9a0d86be221e816")
	// fmt.Println(currentHole("5a9aceeec9a0d85451c38031"))
	// fmt.Println(nextHole("5a9aceeec9a0d85451c38031"))
	fmt.Println(previousHole("5a9aceeec9a0d85451c38031"))
}
