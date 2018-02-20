package main

import (
	"fmt"
	"pro_go/wrong_setup/packages/del"
)

func main() {
	gopher, _ := proGoDeletedb.DelDB("172.17.0.2", "golfer", "Easy123!", "project_under_par")
	fmt.Println(gopher)
}
