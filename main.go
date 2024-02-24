package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwaklog/ssg/src"
)

func main() {
	// get all posts from ./posts from os
	files, err := os.ReadDir("./posts/")
	if err != nil {
		log.Fatal(err)
	}

	// list files in ./post
	fmt.Printf("%q", files)
}
