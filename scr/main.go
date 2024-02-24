package ssg

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v2"
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
func getMdPosts() []Post {
	return nil
}

func generatePage() {}

func parseFrontMatter() {}
