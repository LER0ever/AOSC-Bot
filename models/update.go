package models

import (
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
