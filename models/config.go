package models

// Config struct, stores in file
type Config struct {
	TgBotKey      string
	GhToken       string
	Version       string
	RepoURL       string
	CurPakreqs    []Pakreq
	CurUpdates    []Update
	ClosedPakreqs []Pakreq
	ClosedUpdates []Update
}

// CurConfig Current Config
var CurConfig Config

func init() {
	ReadFromFile()
}
