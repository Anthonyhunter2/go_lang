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
	result := createNew()
	if !bson.IsObjectIdHex(result) {
		t.Errorf("Could not create new doc")
	}
	obID = result
}

func TestFindDoc(t *testing.T) {
	initdb()
	defer moncon.Close()
	check, _ := findOneByID(obID)
	if check == nil {
		t.Errorf("Something went wrong returning the doc")
	}
}

func TestUpdateDoc(t *testing.T) {
	initdb()
	defer moncon.Close()
	check := updateOneByID(obID)
	if check != "Updated" {
		t.Errorf("Something went wrong returning the doc")
	}
}

func TestDeleteDoc(t *testing.T) {
	initdb()
	defer moncon.Close()
	check := deleteByID(obID)
	if check != "Deleted" {
		t.Errorf("Something went wrong returning the doc")
	}
}
