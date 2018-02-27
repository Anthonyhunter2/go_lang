//Package scoreCard is our base scorecard struct
package scoreCard

type Holes struct {
	Hole1  int64 `json:"Hole 1"`
	Hole2  int64 `json:"Hole 2"`
	Hole3  int64 `json:"Hole 3"`
	Hole4  int64 `json:"Hole 4"`
	Hole5  int64 `json:"Hole 5"`
	Hole6  int64 `json:"Hole 6"`
	Hole7  int64 `json:"Hole 7"`
	Hole8  int64 `json:"Hole 8"`
	Hole9  int64 `json:"Hole 9"`
	Hole10 int64 `json:"Hole 10"`
	Hole11 int64 `json:"Hole 11"`
	Hole12 int64 `json:"Hole 12"`
	Hole13 int64 `json:"Hole 13"`
	Hole14 int64 `json:"Hole 14"`
	Hole15 int64 `json:"Hole 15"`
	Hole16 int64 `json:"Hole 16"`
	Hole17 int64 `json:"Hole 17"`
	Hole18 int64 `json:"Hole 18"`
}

//ScoreCard Holds our complete data set, including name & date
type ScoreCard struct {
	//	Date  time.Time
	User  string `json:"User"`
	Round Holes
}

type IDReturn struct {
	ID     string `json:"_id"`
	Hole1  int64  `json:"Hole 1,omitempty"`
	Hole2  int64  `json:"Hole 2,omitempty"`
	Hole3  int64  `json:"Hole 3,omitempty"`
	Hole4  int64  `json:"Hole 4,omitempty"`
	Hole5  int64  `json:"Hole 5,omitempty"`
	Hole6  int64  `json:"Hole 6,omitempty"`
	Hole7  int64  `json:"Hole 7,omitempty"`
	Hole8  int64  `json:"Hole 8,omitempty"`
	Hole9  int64  `json:"Hole 9,omitempty"`
	Hole10 int64  `json:"Hole 10,omitempty"`
	Hole11 int64  `json:"Hole 11,omitempty"`
	Hole12 int64  `json:"Hole 12,omitempty"`
	Hole13 int64  `json:"Hole 13,omitempty"`
	Hole14 int64  `json:"Hole 14,omitempty"`
	Hole15 int64  `json:"Hole 15,omitempty"`
	Hole16 int64  `json:"Hole 16,omitempty"`
	Hole17 int64  `json:"Hole 17,omitempty"`
	Hole18 int64  `json:"Hole 18,omitempty"`
}
