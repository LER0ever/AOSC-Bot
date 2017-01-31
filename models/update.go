package models

import (
	"strings"
	"time"
)

// Update stores packages that need to be updated
type Update struct {
	PkgName   string
	Author    string
	ToVersion string
	OpenTime  time.Time
	CloseTime time.Time
}

// IsUpdate returns if it's a update
func IsUpdate(text string) bool {
	if strings.Contains(text, `#update`) == false {
		return false
	}
	if len(text) <= 7 {
		return false
	}
	if text[7] != ' ' {
		return false
	}
	if strings.Contains(text, "\n") {
		return false
	}

	StrUpdateSlices := strings.Split(text, " ")

	switch len(StrUpdateSlices) {
	case 2:
		if StrUpdateSlices[0] != `#update` {
			return false
		}
	case 4:
		if StrUpdateSlices[0] != `#update` || StrUpdateSlices[2] != "to" {
			return false
		}
	default:
		return false
	}

	return true
}
