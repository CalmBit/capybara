package util

import (
	"time"
	"strconv"
	"fmt"
)

func Ago(date time.Time) string {
	duration := time.Now().UTC().Sub(date.UTC())
	fmt.Printf("Now: %s -  Before: %s\n", time.Now().UTC(), date.UTC())
	if duration.Hours() > 48 {
		return date.Format("Mon, 02 Jan 2006")
	} else if duration.Hours() >= 24 {
		return "1d"
	} else if duration.Hours() >= 1 {
		return strconv.Itoa(int(duration.Hours()))+"h"
	} else if duration.Minutes() >= 1 {
		return strconv.Itoa(int(duration.Minutes()))+"m"
	} else if duration.Seconds() >= 1 {
		return strconv.Itoa(int(duration.Seconds()))+"s"
	} else {
		return "now"
	}
}
