package core

type Detection struct {
	Candid string  `json:"candid"`
	Oid    string  `json:"oid"`
	Mjd    float32 `json:"mjd"`
	Mag    float32 `json:"mag"`
	E_mag  float32 `json:"e_mag"`
}

type NonDetection struct {
	Oid        string  `json:"oid"`
	Mjd        float32 `json:"mjd"`
	Diffmaglim float32 `json:"diffmaglim"`
	Fid        int     `json:"fid"`
}
