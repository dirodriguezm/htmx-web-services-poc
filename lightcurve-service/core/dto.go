package core

type Detection struct {
	Candid   int64   `json:"candid"`
	Oid      string  `json:"oid"`
	Mjd      float32 `json:"mjd"`
	Magpsf   float32 `json:"mag"`
	Sigmapsf float32 `json:"e_mag"`
	Fid      int8    `json:"fid"`
}

type NonDetection struct {
	Oid        string  `json:"oid"`
	Mjd        float32 `json:"mjd"`
	Diffmaglim float32 `json:"diffmaglim"`
	Fid        int     `json:"fid"`
}
