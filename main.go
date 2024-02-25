package main

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
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
	fileName     string
	FrontMatters FrontMatters
	Body         template.HTML
}

// get posts from ./posts
func getMdPosts() ([]Post, error) {
	// get all posts from ./posts from os
	files, err := os.ReadDir(postDir)
	var posts []Post

	// fetch only md files from the list
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".md" {
			continue
		}

		inpMdFile := filepath.Join(postDir, file.Name())

		mdFileContent, err := os.ReadFile(inpMdFile)
		if err != nil {
			log.Fatal(err)
		}

		postFrontMatter := parseFrontMatter(mdFileContent)
		fileContentString := strings.Split(string(mdFileContent), "---")[2]
		parsedMarkdownContent := markdownParser([]byte(fileContentString))

		posts = append(posts, Post{
			fileName:     file.Name(),
			FrontMatters: postFrontMatter,
			Body:         template.HTML(parsedMarkdownContent.String()),
		})
	}

	if err != nil {
		log.Fatal(err)
	}

	// list files in ./post
	// for _, post := range posts {
	// 	fmt.Printf("post: %q\n", post)
	// }

	return posts, nil
}

func generateHTML(posts []Post) {
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.Mkdir(outputDir, 0755); err != nil {
			log.Fatal(err)
		}
	}

	for _, post := range posts {

		postTemplate, err := template.ParseFiles("./template/post.html")
		if err != nil {
			log.Fatal(err)
		}

		var buffer bytes.Buffer
		err = postTemplate.ExecuteTemplate(&buffer, "post", post)
		if err != nil {
			log.Fatal(err)
		}

		outputFile := strings.Join([]string{outputDir, strings.TrimSuffix(post.fileName, ".md") + ".html"}, "/")

		if err := os.WriteFile(outputFile, buffer.Bytes(), 0644); err != nil {
			log.Fatal(err)
		}

	}
}

func parseFrontMatter(fileContentString []byte) FrontMatters {
	frontmatter := &FrontMatters{}

	if err := yaml.Unmarshal(fileContentString, frontmatter); err != nil {
		log.Fatal(err)
	}

	return *frontmatter
}

func markdownParser(markdownContent []byte) bytes.Buffer {
	var buffer bytes.Buffer
	if err := goldmark.Convert(markdownContent, &buffer); err != nil {
		log.Fatal(err)
	}
	return buffer
}

func serveSite() {}

const (
	postDir      = "posts"
	outputDir    = "rendered"
	templatePath = "./template/"
)

func main() {
	posts, err := getMdPosts()
	if err != nil {
		log.Fatal(err)
	}

	generateHTML(posts)

	///for _, post := range posts {
	//	fmt.Printf("post: %q\n", post)
	//}
}
