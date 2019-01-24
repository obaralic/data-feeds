package matchers

import (
	"data-feeds/search"
	. "data-feeds/search"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"encoding/xml"
)

type (
	// Item struct defines fields associated
	// with the item tag in RSS document.
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// Image struct defines fields associated
	// with the image tag in RSS document.
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// Channel struct defines fields associated
	// with the image tag in RSS document.
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// Document struct defines fields associated
	// with the image tag in RSS document.
	document struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// rssMatcher implements the Matcher interface.
type rssMatcher struct{}

// init registers the matcher with the program.
func init() {
	var matcher rssMatcher
	search.Register(RSS_MATCHER, matcher)
}

// Function that performs HTTP GET request for the RSS feed and decodes it.
func (matcher rssMatcher) httpGetRss(feed *Feed) (*document, error) {
	if feed.URI == "" {
		return nil, errors.New("No RSS feed URI provided")
	}

	// Get RSS feed from the web
	httpResponce, error := http.Get(feed.URI)
	if error != nil {
		return nil, error
	}

	// Close the HTTP responce once the function ends.
	defer httpResponce.Body.Close()

	if httpResponce.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Responce Error %d\n", httpResponce.StatusCode)
	}

	// DEcode RSS feed into specified structures
	var doc document
	error = xml.NewDecoder(httpResponce.Body).Decode(&doc)

	return &doc, error
}

// Applies Search function to the rssMatcher type
// so that it's compliant with the Matcher interface.
func (matcher rssMatcher) Search(feed *Feed, term string) ([]*Result, error) {

	log.Printf("Searching [%s] in feed...\nType: %s\nName: %s\nURI: %s\n\n",
		term, feed.Type, feed.Name, feed.URI)

	doc, error := matcher.httpGetRss(feed)
	if error != nil {
		return nil, error
	}

	var results []*Result
	for _, channelItem := range doc.Channel.Item {
		matched, error := regexp.MatchString(term, channelItem.Title)
		if error != nil {
			return nil, error
		}

		// Save the result if matched
		if matched {
			results = append(results, &Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		matched, error = regexp.MatchString(term, channelItem.Description)
		if error != nil {
			return nil, error
		}

		if matched {
			results = append(results, &Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}
