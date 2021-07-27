package utils

import (
	"bytes"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// custom markdown parser
var markdown = goldmark.New(
	goldmark.WithRendererOptions(
		html.WithUnsafe(),
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithExtensions(
		meta.Meta,
	),
)

// v1
func BlogPostParser(filename string) (string, map[string]interface{}) {
	var buf bytes.Buffer
	context := parser.NewContext()

	file, err := os.ReadFile("./blog/" + filename + ".markdown")
	if err != nil {
		panic(err)
	}

	if err := markdown.Convert(file, &buf, parser.WithContext(context)); err != nil {
		panic(err)
	}

	metaData := meta.Get(context)

	return buf.String(), metaData
}

// v1
func BlogIndexParser() []map[string]interface{} {
	var blogSlice []map[string]interface{}

	var buf bytes.Buffer
	context := parser.NewContext()

	files, err := os.ReadDir("./blog")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		file, err := os.ReadFile("./blog/" + f.Name())
		if err != nil {
			panic(err)
		}

		if err := markdown.Convert(file, &buf, parser.WithContext(context)); err != nil {
			panic(err)
		}
		metaData := meta.Get(context)
		metaData["fileName"] = strings.Split(f.Name(), ".markdown")[0]
		blogSlice = append(blogSlice, metaData)
	}

	sort.Slice(blogSlice, func(i, j int) bool {
		format := "January 2, 2006"
		t1, err := time.Parse(format, blogSlice[i]["date"].(string))
		if err != nil {
			panic(err)
		}
		t2, err := time.Parse(format, blogSlice[j]["date"].(string))
		if err != nil {
			panic(err)
		}
		return t1.Unix() > t2.Unix()
	})
	return blogSlice
}
