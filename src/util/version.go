package util

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	// Name ...
	Name = "urbs-console"
	// Version ...
	Version = "unknown"
	// BuildTime ...
	BuildTime = "unknown"
	// GitSHA1 ...
	GitSHA1 = "unknown"
)

// PrintVersion ...
func PrintVersion() {
	fmt.Println(string(GetVersion()))
	os.Exit(0)
}

// GetVersion ...
func GetVersion() []byte {
	v := map[string]string{
		"name":      Name,
		"version":   Version,
		"buildTime": BuildTime,
		"gitSHA1":   GitSHA1,
	}
	d, _ := json.MarshalIndent(v, "", "    ")
	return d
}
