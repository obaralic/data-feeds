package search

import (
	"log"
)

// Struct that represents search result.
type Result struct {
	Field   string
	Content string
}

// Interface that defines the behavior required by types
// that want to implement a new search type.
type Matcher interface {
	Search(feed *Feed, term string) ([]*Result, error)
}

// Function that performs matching of a given term through a given @Feed.
// It is executed as a go rutine thus passing @Result
// channel as a mean of synchronization.
func Match(matcher Matcher, feed *Feed, term string, resultChannel chan<- *Result) {
	searchResults, error := matcher.Search(feed, term)
	if error != nil {
		log.Println(error)
		return
	}

	// Write the results to the channel.
	for _, result := range searchResults {
		resultChannel <- result
	}
}

// Function that writes results to the console window
// as they are received by the individual goroutines.
func Display(resultsChannel chan *Result) {

	// Channel is blocked while is being writen to it.
	// Iterate through results when they become available.
	// Once the channel gets closed loop will be terminated.
	for result := range resultsChannel {
		log.Printf("%s:\n%s:\n\n", result.Field, result.Content)
	}
}
