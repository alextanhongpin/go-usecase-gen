package toc

import (
	"log"
	"os"
	"text/template"
)

const readme = `
# Table of Content
{{ range .Items }}
	{{ .Name }}
	{{ range .Tocs }}
		- [{{ .FileName }}](/docs/{{ .FileName }})
	{{ end }}
{{ end }}
`

type TocData struct {
	Name string
	Tocs []Toc
}
type MarkdownData struct {
	Items []TocData
}

func GenerateMarkdown(in, out string) {
	// Get all yaml files in directory.
	files, err := GetFilesInDir(in)
	if err != nil {
		log.Fatal(err)
	}

	// Setup target destination.
	rp := relativePath(out)
	f, err := os.Create(rp)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data MarkdownData
	for name, file := range files {
		data.Items = append(data.Items, TocData{
			Name: name,
			Tocs: file,
		})
	}
	t := template.Must(template.New("").Parse(readme))
	if err := t.Execute(f, data); err != nil {
		log.Fatal(err)
	}
}
