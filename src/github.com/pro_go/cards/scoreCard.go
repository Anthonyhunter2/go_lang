//Package scoreCard is our base scorecard struct
package scoreCard

type Holes struct {
	Hole1  int `json:"Hole 1"`
	Hole2  int `json:"Hole 2"`
	Hole3  int `json:"Hole 3"`
	Hole4  int `json:"Hole 4"`
	Hole5  int `json:"Hole 5"`
	Hole6  int `json:"Hole 6"`
	Hole7  int `json:"Hole 7"`
	Hole8  int `json:"Hole 8"`
	Hole9  int `json:"Hole 9"`
	Hole10 int `json:"Hole 10"`
	Hole11 int `json:"Hole 11"`
	Hole12 int `json:"Hole 12"`
	Hole13 int `json:"Hole 13"`
	Hole14 int `json:"Hole 14"`
	Hole15 int `json:"Hole 15"`
	Hole16 int `json:"Hole 16"`
	Hole17 int `json:"Hole 17"`
	Hole18 int `json:"Hole 18"`
}

//ScoreCard Holds our complete data set, including name & date
type ScoreCard struct {
	//	Date  time.Time
	User  string `json:"User"`
	Round Holes
}
