package main

import (
	"testing"

	"github.com/globalsign/mgo/bson"
)

var obID string

func TestMongoCon(t *testing.T) {
	initdb()
	defer moncon.Close()
	if moncon == nil {
		t.Errorf("Couldn't make a connection")
	}
}

func TestCreateDoc(t *testing.T) {
	initdb()
	defer moncon.Close()
	result := createNew("testgolfer")
	if !bson.IsObjectIdHex(result) {
		t.Errorf("Could not create new doc")
	}
	obID = result
}

func TestFindDoc(t *testing.T) {
	initdb()
	defer moncon.Close()
	_, err := findOneByID(obID)
	if err != nil {
		t.Errorf("Something went wrong returning the doc")
	}
}

func TestUpdateDoc(t *testing.T) {
	initdb()
	defer moncon.Close()
	check := updateOneByID(obID, "Testgolfer")
	if check != "Updated" {
		t.Errorf("Something went wrong returning the doc")
	}
}
func TestUpdateHole(t *testing.T) {
	initdb()
	defer moncon.Close()
	check := updateSingleHole(obID, "Hole 1", 5)
	if check != "Updated" {
		t.Errorf("Something went wrong updating a single hole doc")
	}
}

func TestDeleteDoc(t *testing.T) {
	initdb()
	defer moncon.Close()
	check := deleteByID(obID)
	if check != "Deleted" {
		t.Errorf("Something went wrong deleteing the doc")
	}
}
