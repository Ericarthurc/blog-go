package utils

import (
	"bytes"
	"os"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// v1
func ParserBlog() (string, map[string]interface{}) {
	markdown := goldmark.New(
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
	var buf bytes.Buffer
	context := parser.NewContext()

	file, err := os.ReadFile("./blog/post.md")
	if err != nil {
		panic(err)
	}

	if err := markdown.Convert(file, &buf, parser.WithContext(context)); err != nil {
		panic(err)
	}

	metaData := meta.Get(context)

	return buf.String(), metaData
}
