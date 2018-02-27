package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/rhinoman/couchdb-go"
)

//This will return our scorecard
// type ScoreCard struct {
// 	date time.Time
// 	User string
// 	//Round map[string]map[string]string
// 	Round map[string]interface {
// 	}
// }

func main() {
	var m interface{}
	//returnDoc := ScoreCard{}
	v := url.Values{}
	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("172.17.0.2", 5984, timeout)
	if err != nil {
		fmt.Println(err)
	}
	auth := couchdb.BasicAuth{Username: "golfer", Password: "Easy123!"}
	db := conn.SelectDB("project_under_par", &auth)
	db.Read("1263434af76ea9903d9e6807f400a8c5", &m, &v)
	//	document, _ := db.Read("1263434af76ea9903d9e6807f400a8c5", &m, &v)
	//fmt.Println(m)
	b := []byte(`m`)
	error5 := json.Unmarshal(b, m)
	if error5 != nil {
		fmt.Println(m)
	}
	for k, vc := range m.(map[string]interface{}) {
		jabby := fmt.Sprintf("%T", vc)
		fmt.Println(jabby)
		switch vv := vc.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case map[string]interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				for n, o := range u.(map[string]interface{}) {
					fmt.Println(i, n, o)
				}
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
