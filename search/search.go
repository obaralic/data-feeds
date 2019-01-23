package search

import (
	"log"
	"sync"
)

const (
	DEFAULT_MATCHER = "default"
	RSS_MATCHER     = "rss"
)

// Map of registered matchers that are using for the search.
var matchers = make(map[string]Matcher)

// Function that registers matchers that will be used throughout the program.
func Register(feedType string, matcher Matcher) {
	if _, registered := matchers[feedType]; registered {
		log.Fatalln("Matcher", feedType, "is already registered!")
	}

	matchers[feedType] = matcher
	log.Println("Matcher", feedType, "succesfully registered")
}

// Run performs the search logic.
func Run(term string) {

	// Get the list of available feeds to search through.
	feeds, error := RetrieveFeeds()
	if error != nil {
		log.Fatal(error)
	}

	// Create unbound channel to receive messages from the matcher.
	resultChannel := make(chan *Result)

	// Sort of a countdown latch principle where the WaitGroup is a semaphore.
	var latch sync.WaitGroup
	latch.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	for _, feed := range feeds {
		matcher := getMatcher(*feed)

		// Launch the goroutine as a cloasure to perform the search.
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, term, resultChannel)
			latch.Done()
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all the work is done.
	go func() {
		// Wait for the all match routines
		latch.Wait()

		// Close the channel to signal to the matcher Display function
		// that we can exit the program.
		close(resultChannel)
	}()

	Display(resultChannel)
}

func getMatcher(feed Feed) (matcher Matcher) {
	matcher, contained := matchers[feed.Type]
	if !contained {
		matcher = matchers[DEFAULT_MATCHER]
	}
	return
}
