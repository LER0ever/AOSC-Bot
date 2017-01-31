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
	PkgDesc   string
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
	if strings.Contains(text, "\n") {
		text = strings.Split(text, "\n")[0]
	}
	StrPakreqSlices := strings.Split(text, " ")
	if StrPakreqSlices[0] != `#pakreq` || len(StrPakreqSlices) != 4 || StrPakreqSlices[2] != `@` || IsURL(StrPakreqSlices[3]) == false {
		return false
	}
	return true
}

// PakreqToPkgString returns a slice of string containing package names
func PakreqToArray(text string) []string {
	if IsPakreq(text) == false {
		log.Fatalf("String is not a valid pakreq: %s", text)
	}
	StrPakreqSlices := strings.Split(text, " ")
	return StrPakreqSlices[1:]
}

// IsURL use a very efficient way to judge if it is a URL ;)
func IsURL(text string) bool {
	if text[0:4] != "http" {
		return false
	}
	return true
}

// PakreqExists returns if the pakreq already exists in the database
func PakreqExists(pakreq Pakreq) bool {
	for _, pkg := range CurConfig.CurPakreqs {
		if pkg.PkgName == pakreq.PkgName {
			return true
		}
	}
	return false
}
