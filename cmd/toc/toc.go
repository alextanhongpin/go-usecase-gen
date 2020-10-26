package main

import "github.com/alextanhongpin/usecase/pkg/toc"

func main() {
	toc.GenerateMarkdown("./src", "./dst/README.md")
}
