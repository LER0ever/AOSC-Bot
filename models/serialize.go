package models

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// ExportToFile exports all current pakreqs and udpates to a config.toml
func ExportToFile() {
	f, err := os.OpenFile("config.toml", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error Open file: %s", err.Error())
	}
	if err = toml.NewEncoder(f).Encode(CurConfig); err != nil {
		log.Fatal("Error Marshaling data to config.toml: %s", err.Error())
	}
	f.Close()
}

// ReadFromFile reads config from config.toml
func ReadFromFile() {
	_, err := toml.DecodeFile("config.toml", &CurConfig)
	if err != nil {
		log.Fatalf("Error Unmarshaling data from config.toml: %s", err.Error())
	}
}
