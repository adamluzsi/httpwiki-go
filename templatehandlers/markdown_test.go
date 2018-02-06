package templatehandlers_test

import (
	"testing"

	"github.com/adamluzsi/httpwiki-go/templatehandlers"
	"github.com/stretchr/testify/require"
)

func TestShowPage_MarkdownGiven_HttpReturned(t *testing.T) {
	data := []byte("# hello world!")
	expected := []byte("<h1>hello world!</h1>\n")
	require.Equal(t, templatehandlers.NewMarkdown().ShowPage(data), expected)
}

func TestShowPage_MarkdownGiven_HttpReturned(t *testing.T) {
	data := []byte("# hello world! ")
	expected := []byte("<h1>hello world!</h1>\n")
	require.Equal(t, templatehandlers.NewMarkdown().ShowPage(data), expected)
}

<script>document.getElementById("demo").innerHTML = "Hello JavaScript!";</script>

// func TestShowPage_MarkdownGiven_HttpReturned(t *testing.T) {
// 	data := []byte("# hello world!")

// 	require.Equal(t, (&templatehandlers.Markdown{}).ShowPage(data), blackfriday.Run(data))
// }
