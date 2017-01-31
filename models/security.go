package models

import (
	"strings"
	"time"
)

// Security stores packages that need to be updated
type Security struct {
	PkgName   string
	Author    string
	ToVersion string
	OpenTime  time.Time
	CloseTime time.Time
}

// IsSecurity returns if it's a update
func IsSecurity(text string) bool {
	if strings.Contains(text, `#security`) == false {
		return false
	}
	if len(text) <= 9 {
		return false
	}
	if text[9] != ' ' {
		return false
	}
	if strings.Contains(text, "\n") {
		return false
	}

	StrSecuritySlices := strings.Split(text, " ")

	switch len(StrSecuritySlices) {
	case 2:
		if StrSecuritySlices[0] != `#security` {
			return false
		}
	case 4:
		if StrSecuritySlices[0] != `#security` || StrSecuritySlices[2] != "to" {
			return false
		}
	default:
		return false
	}

	return true
}
