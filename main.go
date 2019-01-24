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
	fmt.Println("Enter phrase to search for and press ENTER:\n")

	var term string
	fmt.Scanln(&term)
	search.Run(term)
}
