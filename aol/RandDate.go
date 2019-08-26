package aol

import (
	"math/rand"
	"strconv"
	"time"
)

type RandDate struct {
	Month string
	Day   string
	Year  string
}

func NewRandDate() RandDate {

	rand.Seed(time.Now().UnixNano())

	months := []string{"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}
	yearMin := 1990
	yearMax := 2000

	randDate := RandDate{
		Month: months[rand.Intn(len(months))],
		Day:   strconv.Itoa(rand.Intn(27) + 1),
		Year:  strconv.Itoa(rand.Intn(yearMax-yearMin) + yearMin),
	}

	return randDate
}
