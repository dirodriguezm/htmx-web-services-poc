package core

type Object struct {
	Oid       string  `json:"oid"`
	Corrected bool    `json:"corrected"`
	Stellar   bool    `json:"stellar"`
	Ndet      int     `json:"ndet"`
	Meanra    float32 `json:"meanra"`
	Meandec   float32 `json:"meandec"`
	Firstmjd  float32 `json:"firstmjd"`
	Lastmjd   float32 `json:"lastmjd"`
}
