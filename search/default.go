package search

type defaultMatcher struct{}

// Init registers the default matcher with the program.
// Executed as a part of package loading process.
func init() {
	var matcher defaultMatcher
	Register(DEFAULT_MATCHER, matcher)
}

// Search implements the behavior for the default matcher.
func (matcher defaultMatcher) Search(feed *Feed, term string) ([]*Result, error) {
	return nil, nil
}
