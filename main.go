package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
	// "github.com/yuin/goldmark"
	// "github.com/yuin/goldmark/extension"
	// "github.com/yuin/goldmark/parser"
	// "github.com/yuin/goldmark/renderer/html"
	// "gopkg.in/yaml.v2"
)

type FrontMatters struct {
	Title  string `yaml:"title"`
	Author string `yaml:"author"`
	Date   string `yaml:"date"`
}

type Post struct {
	frontMatters FrontMatters
	Body         string
}

// get posts from ./posts
func getMdPosts() ([]Post, error) {
	// get all posts from ./posts from os
	files, err := os.ReadDir(postDir)

	// fetch only md files from the list
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".md" {
			continue
		}

        inpMdFile := filepath.Join(postDir, file.Name())
        outFile := filepath.Join(outputDir, file.Name())

        fmt.Printf("inp : %q\tout: %q\n", inpMdFile, outFile)

        mdFileContent, err := os.ReadFile(inpMdFile)
        if err != nil {
            log.Fatal(err)
        }

        postFrontMatter := parseFrontMatter(mdFileContent)

        fmt.Printf("frontmatter: %q\n", postFrontMatter)
	}

	if err != nil {
		log.Fatal(err)
	}


	// list files in ./post

	return []Post{}, nil
}

func generateHtmlFils() {}

func parseFrontMatter(fileContentString []byte) FrontMatters {

    frontmatter := &FrontMatters{}

    if err := yaml.Unmarshal(fileContentString, frontmatter); err != nil {
        log.Fatal(err)
    }

    // fmt.Printf("frontmatter: %q\n", frontmatter)
    

	return *frontmatter 
}

func serveSite() {}

const (
	postDir   = "posts"
	outputDir = "output"
)

func main() {
	_, err := getMdPosts()
	if err != nil {
		log.Fatal(err)
	}
}
