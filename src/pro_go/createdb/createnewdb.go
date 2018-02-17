package main

import (
	"pro_go/packages/create"
	"pro_go/scorecard"
)

func main() {
	action, _ := proGoCreateDB.ProGoCreate("project_under_par", "172.17.0.2", "golfer", "Easy123!")
	if action == "DB Created" {
		//	newRecords := makeScores.UploadScores()
		for i := 0; i < 10; i++ {
			makeScores.UploadScores()
		}
	}
}
