package search

import (
	// Provides support for accessing operating system
	// functionality like reading and writing files.
	"os"

	"encoding/json"
)

const feedsDataFile = "data/data.json"

// Struct that represents feed data,
// and it is used for its processing.
type Feed struct {
	// Last part containing 'json:<field>' defines id tag with JSON annotation
	// which can be interpreted by Go's JSON encoder and decoder.
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// Function that reads and parse the feed data file.
// Returns slice of @Feed structs and corresponding error code.
func RetrieveFeeds() (feeds []*Feed, err error) {
	file, err := os.Open(feedsDataFile)
	if err != nil {
		return
	}

	// Defer to Golang is finally to Java.
	// Such a statement defers the execution of a function
	// until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately,
	// but the call is not executed until the surrounding function returns.
	defer file.Close()

	// Decode the file into a slice of pointers to @Feed values.
	err = json.NewDecoder(file).Decode(&feeds)

	return
}
