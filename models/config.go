package models

// Config struct, stores in file
type Config struct {
	CurPakreqs    []Pakreq
	CurUpdates    []Update
	ClosedPakreqs []Pakreq
	ClosedUpdates []Update
}

// CurConfig Current Config
var CurConfig Config
