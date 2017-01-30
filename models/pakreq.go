package models

import (
	"log"
	"strings"
	"time"
)

// Pakreq stores a #pakreq
type Pakreq struct {
	PkgName   string
	PkgURL    string
	Author    string
	OpenTime  time.Time
	CloseTime time.Time
}

// IsPakreq returns true if the string is a valid pakreq
func IsPakreq(text string) bool {
	if strings.Contains(text, `#pakreq`) == false {
		return false
	}
	if len(text) <= 7 {
		return false
	}
	if text[7] != ' ' {
		return false
	}
	StrPakreqSlices := strings.Split(text, " ")
	if StrPakreqSlices[0] != `#pakreq` || len(StrPakreqSlices) <= 1 {
		return false
	}
	return true
}

// PakreqToPkgString returns a slice of string containing package names
func PakreqToPkgString(text string) []string {
	if IsPakreq(text) == false {
		log.Fatalf("String is not a valid pakreq: %s", text)
	}
	StrPakreqSlices := strings.Split(text, " ")
	return StrPakreqSlices[1:]
}
