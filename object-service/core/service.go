package core

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetObjectById(id string,
	handle_success func(result Object),
	handle_error func(err error),
	db *pgxpool.Pool,
) {
	log.Printf("Fetching object %v\n", id)
	conn, err := db.Acquire(context.Background())
	if err != nil {
		handle_error(err)
	}
	defer conn.Release()
	row := conn.QueryRow(context.Background(), "select  oid, corrected, stellar, ndet, meanra, meandec, firstmjd, lastmjd from object where oid=$1", id)
	var object Object
	err = row.Scan(&object.Oid, &object.Corrected, &object.Stellar, &object.Ndet, &object.Meanra, &object.Meandec, &object.Firstmjd, &object.Lastmjd)
	if err != nil {
		handle_error(err)
	} else {
		handle_success(object)
	}
}

func DegreeToHms(ra float64) string {
	sign := ""
	if ra < 0 {
		sign = "-"
	}
	ra = math.Abs(ra)
	raH := math.Floor(ra / 15)
	raM := math.Floor((ra/15 - raH) * 60)
	raS := (((ra/15 - raH) * 60) - raM) * 60
	return sign + strconv.FormatFloat(raH, 'g', 3, 64) + ":" + strconv.FormatFloat(raM, 'g', 3, 64) + ":" + strconv.FormatFloat(raS, 'g', 3, 64)
}

func DegreeToDms(dec float64) string {
	sign := "+"
	if dec < 0 {
		sign = "-"
	}
	dec = math.Abs(dec)
	decD := math.Floor(dec)
	decM := math.Abs(math.Floor((dec - decD) * 60))
	decS := math.Abs((((dec - decD) * 60) - decM) * 60)
	return sign + strconv.FormatInt(int64(decD), 10) + ":" + strconv.FormatInt(int64(decM), 10) + ":" + strconv.FormatFloat(decS, 'g', 3, 64)

}

func HmsToDegree(hms string) (float64, error) {
	// Split the input string into hours, minutes, and seconds
	parts := strings.Split(hms, ":")
	if len(parts) != 3 {
		return 0.0, fmt.Errorf("Invalid RA format")
	}

	// Parse hours, minutes, and seconds
	hours, err := parseHMS(parts[0])
	if err != nil {
		return 0.0, err
	}
	minutes, err := parseHMS(parts[1])
	if err != nil {
		return 0.0, err
	}
	seconds, err := parseHMS(parts[2])
	if err != nil {
		return 0.0, err
	}

	// Calculate the total degrees
	totalDegrees := hours*15.0 + minutes*0.25 + seconds*0.00416667

	return totalDegrees, nil
}

func parseHMS(s string) (float64, error) {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, err
	}
	if value < 0 || value >= 60 {
		return 0.0, fmt.Errorf("Invalid HMS value: %s", s)
	}
	return value, nil
}

func DmsToDegrees(dec string) (float64, error) {
	// Split the input string into degrees, minutes, and seconds
	dec = strings.Trim(dec, " ")
	parts := strings.Split(dec, ":")
	if len(parts) != 3 {
		return 0.0, fmt.Errorf("Invalid Dec format")
	}

	// Parse degrees, minutes, and seconds
	degrees, err := parseDMS(parts[0])
	if err != nil {
		return 0.0, err
	}
	minutes, err := parseDMS(parts[1])
	if err != nil {
		return 0.0, err
	}
	seconds, err := parseDMS(parts[2])
	if err != nil {
		return 0.0, err
	}

	// Calculate the total degrees
	totalDegrees := degrees + minutes/60.0 + seconds/3600.0

	return totalDegrees, nil
}

func parseDMS(s string) (float64, error) {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, err
	}
	if value < 0 || value >= 60 {
		return 0.0, fmt.Errorf("Invalid DMS value: %s", s)
	}
	return value, nil
}

func MjdToGreg(mjd float64) string {
	jd := mjd + 2400000.5
	var a float64

	if jd < 2299160.0 {
		a = jd
	} else {
		alpha := int((jd - 1867216.25) / 36524.25)
		a = jd + 1.0 + float64(alpha) - float64(alpha)/4.0
	}

	b := a + 1524.0
	c := int((b - 122.1) / 365.25)
	d := int(365.25 * float64(c))
	e := int((b - float64(d)) / 30.6001)

	day := int(b - float64(d) - float64(int(30.6001*float64(e))))
	var month time.Month

	if e < 14 {
		month = time.Month(e - 1)
	} else {
		month = time.Month(e - 13)
	}

	year := c - 4715

	if month > 2 {
		year--
	}

	// Calculate the fractional part of the day
	dayFraction := jd - float64(int(jd))

	// Calculate the number of seconds in a day
	secondsInDay := 86400.0

	// Calculate the UTC time by multiplying the fractional part by the number of seconds in a day
	utcTime := time.Duration(dayFraction * secondsInDay * float64(time.Second))

	// Create a time.Time object with the extracted year, month, day, and the calculated UTC time
	gregorianDateTime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Add(utcTime)

	return gregorianDateTime.Format("2006-01-02 15:04:05")
}

func GregToMjd(gregDate string) (float64, error) {
	// This does not work correctly. Should be replaced by the correct calculation
	t, err := time.Parse("2006-01-02 15:04:05", gregDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return 0, err
	}

	// Calculate the Modified Julian Date (MJD)
	year := float64(t.Year())
	month := float64(t.Month())
	day := float64(t.Day())
	hour := float64(t.Hour())
	minute := float64(t.Minute())
	second := float64(t.Second())

	if month <= 2 {
		year -= 1
		month += 12
	}

	a := year / 100.0
	b := 2 - a + (a / 4)

	jd := (365.25 * (year + 4716)) + (30.6001 * (month + 1)) + day + b - 1524.5
	jd += (hour + minute/60 + second/3600) / 24.0

	mjd := jd - 2400000.5
	return mjd, nil
}
