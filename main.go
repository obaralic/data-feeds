package main

import (
	"fmt"
	"log"
	"os"

	_ "data-feeds/matchers"
	"data-feeds/search"
)

// Init function is always called before main.
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	fmt.Println("Data Feeds - Entry point ...")
	search.Run("Search")
}
