package templatehandlers

import (
	"bytes"
	"html/template"

	"github.com/microcosm-cc/bluemonday"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type Markdown struct{}

func NewMarkdown() *Markdown {
	return &Markdown{}
}

func (th *Markdown) ShowPage(data []byte) []byte {
	unsafe := blackfriday.Run(data)

	return bluemonday.UGCPolicy().SanitizeBytes(unsafe)
}

const MarkdownEditPageTemplate = `
<h1>Editing {{.Title}}</h1>
<form action="{{.Path}}" method="POST">
<div><textarea name="body">{{printf "%s" .Content}}</textarea></div>
<div><input type="submit" value="Save"></div>
</form>
`

type MarkdownEditPageContent struct {
	Path    string
	Title   string
	Content []byte
}

func (th *Markdown) EditPage(path string, data []byte) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})

	t, err := template.New("github.com/adamluzsi/httpwiki-go/templatehandlers/markdown/edit").Parse(MarkdownEditPageTemplate)

	if err != nil {
		return nil, err
	}

	err = t.Execute(buffer, MarkdownEditPageContent{Path: path, Title: path, Content: data})

	return buffer.Bytes(), err
}
