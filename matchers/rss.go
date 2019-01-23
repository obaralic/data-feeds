package matchers

import (
	"data-feeds/search"
	. "data-feeds/search"
)

// rssMatcher implements the Matcher interface.
type rssMatcher struct{}

// init registers the matcher with the program.
func init() {
	var matcher rssMatcher

	search.Register(RSS_MATCHER, matcher)
}

func (rss rssMatcher) Search(feed *Feed, term string) ([]*Result, error) {
	return nil, nil
}
