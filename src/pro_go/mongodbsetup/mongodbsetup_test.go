package main

import (
	"testing"

	"github.com/globalsign/mgo/bson"
)

var obID string

func TestMongoCon(t *testing.T) {
	Initdb()
	defer moncon.Close()
	if moncon == nil {
		t.Errorf("Couldn't make a connection")
	}
}

func TestCreateDoc(t *testing.T) {
	Initdb()
	defer moncon.Close()
	result := CreateNew("testgolfer")
	if !bson.IsObjectIdHex(result) {
		t.Errorf("Could not create new doc")
	}
	obID = result
}

func TestFindDoc(t *testing.T) {
	Initdb()
	defer moncon.Close()
	if doc, _ := FindOneByID(obID); doc == nil {
		t.Errorf("Something went wrong returning the doc")
	}
}

func TestUpdateDoc(t *testing.T) {
	Initdb()
	defer moncon.Close()
	UpdateNameByID(obID, "Testgolfer")
	returnval, _ := FindOneByID(obID)
	if returnval.Name != "Testgolfer" {
		t.Errorf("Something went wrong returning the doc")
	}
}

func TestUpdateHole(t *testing.T) {
	Initdb()
	defer moncon.Close()
	UpdateSingleHole(obID, "Hole 1", 5)
	returnval, _ := FindOneByID(obID)
	if returnval.Round.Hole1 != 5 {
		t.Errorf("Something went wrong updating a single hole doc")
	}
}

func TestCurHole(t *testing.T) {
	Initdb()
	defer moncon.Close()
	if check := CurrentHole(obID); check != "Hole 2" {
		t.Errorf("Something went wrong returning current hole")
	}
}

func TestNextHole(t *testing.T) {
	Initdb()
	defer moncon.Close()
	if check := NextHole(obID); check != "Hole 3" {
		t.Errorf("Something went wrong returning current hole")
	}
}
func TestPreviousHole(t *testing.T) {
	Initdb()
	defer moncon.Close()
	if check := PreviousHole(obID); check != "Hole 1" {
		t.Errorf("Something went wrong returning current hole")
	}
}
func TestDeleteDoc(t *testing.T) {
	Initdb()
	defer moncon.Close()
	DeleteByID(obID)
	if _, err := FindOneByID(obID); err == nil {
		t.Errorf("Something went wrong deleteing the doc")
	}
}
