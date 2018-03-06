package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
)

//CreateNew enters a new document into the db and returns the objectID, this will act as a blank scorecard
func CreateNew(golfer string) string {
	newID := bson.NewObjectId()
	currentdate := time.Now().Format("2006-01-02")
	playerone := &Person{ID: newID, Date: currentdate, Name: golfer, CurHole: "Hole 1"}
	if err := moncol.Insert(playerone); err != nil {
		log.Fatalf("Couldn't Create new round for %v", golfer)
	}
	return newID.Hex()
}

//FindOneByID returns back the full document from the ID based on the objectID
func FindOneByID(idstring string) (*Person, error) {
	result := &Person{}
	err := moncol.Find(bson.M{"_id": bson.ObjectIdHex(idstring)}).One(&result)
	if err != nil {
		//fmt.Println("Couldn't Find that ID string")
		//log.Fatal(err)
		return result, err
	}
	return result, err
}

//UpdateNameByID right now only changes the name of the golfer on the scorecard, but later could be used to
//update any of the top level information in the doc, ie date, currenthole, or meta data
func UpdateNameByID(idstring string, golfer string) {
	result := bson.M{"$set": bson.M{"Name": golfer}}
	if err := moncol.Update(bson.M{"_id": bson.ObjectIdHex(idstring)}, &result); err != nil {
		log.Fatal(err)
	}

}

//UpdateSingleHole will update the nested hole passed to it and also update the current hole to the updated hole +1
func UpdateSingleHole(idstring string, feild string, score int) {
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
		log.Fatalf("%v is not a vaild hole number", feild)
	}
	//Instead of doing 1 read and 1 write to the db and using our next hole func, we are just going to reuse some of the code
	//to figure out what your next hole will be
	newHoleNum := holeNum + 1
	if newHoleNum == 19 {
		newHoleNum = 1
	}

	//converting our number back to a string so it can be put into our doc
	newHoleStr := "Hole " + strconv.Itoa(newHoleNum)
	findDoc := bson.M{"_id": bson.ObjectIdHex(idstring)}
	updateDoc := bson.M{"$set": bson.M{"Round." + feild: score, "Current Hole": newHoleStr}}
	err := moncol.Update(findDoc, updateDoc)
	if err != nil {
		log.Fatalf("%v not a vaild id", idstring)
	}

}

// CurrentHole Returns the current hole defined in the doc
func CurrentHolef(idstring string) string {
	result := &Person{}
	err := moncol.Find(bson.M{"_id": bson.ObjectIdHex(idstring)}).One(&result)
	if err != nil {
		log.Fatalf("%v not a vaild id", idstring)
	}
	return result.CurHole
}

//NextHole Returns the next hole, if your on Hole 18 it will return Hole 1
func NextHole(idstring string) string {
	result := &Person{}
	err := moncol.Find(bson.M{"_id": bson.ObjectIdHex(idstring)}).One(&result)
	if err != nil {
		log.Fatalf("%v not a vaild id", idstring)
	}
	holenumstr := result.CurHole[5:]
	holenum, _ := strconv.Atoi(holenumstr)
	newHoleNum := holenum + 1
	if newHoleNum == 19 {
		newHoleNum = 1
	}
	newHoleStr := "Hole " + strconv.Itoa(newHoleNum)
	findDoc := bson.M{"_id": bson.ObjectIdHex(idstring)}
	updateDoc := bson.M{"$set": bson.M{"Current Hole": newHoleStr}}
	err2 := moncol.Update(findDoc, updateDoc)
	if err2 != nil {
		log.Fatalf("%v not a vaild id", idstring)
	}
	return newHoleStr
}

//PreviousHole returns the previous hole, if your on Hole 1 it will return Hole 18
func PreviousHole(idstring string) string {
	result := &Person{}
	err := moncol.Find(bson.M{"_id": bson.ObjectIdHex(idstring)}).One(&result)
	if err != nil {
		log.Fatalf("%v not a vaild id", idstring)
	}
	holenumstr := result.CurHole[5:]
	holenum, _ := strconv.Atoi(holenumstr)
	newHoleNum := holenum - 1
	if newHoleNum == 0 {
		newHoleNum = 18
	}
	newHoleStr := "Hole " + strconv.Itoa(newHoleNum)
	findDoc := bson.M{"_id": bson.ObjectIdHex(idstring)}
	updateDoc := bson.M{"$set": bson.M{"Current Hole": newHoleStr}}
	err2 := moncol.Update(findDoc, updateDoc)
	if err2 != nil {
		log.Fatalf("%v not a vaild id", idstring)
	}
	return newHoleStr
}

//DeleteByID deletes the document by the ObjectID
func DeleteByID(idstring string) {
	err := moncol.Remove(bson.M{"_id": bson.ObjectIdHex(idstring)})
	if err != nil {
		log.Fatalf("Couldn't Find that ID string to delete: %v", idstring)
	}
}

//func main() {
//	Initdb()
//	defer moncon.Close()
// for value := 0; value < 2000; value++ {
// 	val := strconv.Itoa(value)
// 	newval := "Entry" + val
// 	newID := createNew(newval)
// 	fmt.Println(findOneByID(newID))
// }
// fmt.Println(createNew("mazzaa"))
// fmt.Println(findOneByID("5a9de0e6c9a0d8525e0cdcfd"))
// updateSingleHole("5a9de0e6c9a0d8525e0cdcfd", "hole2", 7)
// updateNameByID("5a9de0e6c9a0d8525e0cdcfd", "Changed")
// fmt.Println(currentHole("5a9de0e6c9a0d8525e0cdcfd"))
//fmt.Println(nextHole("5a9de0e6c9a0d8525e0cdcfd"))
//fmt.Println(previousHole("5a9de0e6c9a0d8525e0cdcfd"))
//fmt.Println(deleteByID("5a98afa0c9a0d86be221e816")

//}
