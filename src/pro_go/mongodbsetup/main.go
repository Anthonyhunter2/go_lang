package main

import (
	"fmt"

	"github.com/globalsign/mgo"
)

func main() {
	connect, err := mgo.Dial("172.17.0.2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Creating DB connection")
	fmt.Println(connect)
}
