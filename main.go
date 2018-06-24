package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

const (
	outputFile = "output.pdf"
	delimiter  = "---"
)

func check(err interface{}) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func main() {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	check(err)

	mds := splitMarkdownFiles(os.Args[1])
	htmls := genHTML(mds)
	for _, html := range htmls {
		applyGFM(html)
		pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
		pdfg.AddPage(wkhtmltopdf.NewPage(html))
	}
	check(pdfg.Create())
	check(pdfg.WriteFile(outputFile))
	clean()
}

func applyGFM(html string) {
	cssStartTag := `<link rel="stylesheet" href="./assets/github-markdown.css">
<style>
	.markdown-body {
		box-sizing: border-box;
		min-width: 200px;
		max-width: 980px;
		margin: 0 auto;
		padding: 45px;
	}

	@media (max-width: 767px) {
		.markdown-body {
			padding: 15px;
		}
	}
</style>
<article class="markdown-body">`

	cssEndTag := `</article>`
	data, _ := ioutil.ReadFile(html)
	gfmmd := fmt.Sprintf("%s\n%s\n%s", cssStartTag, data, cssEndTag)
	writeFile(gfmmd, html)
}

// clean html and md file generated on the process
func clean() {
	files, err := filepath.Glob("page*.*")
	check(err)
	for _, f := range files {
		check(os.Remove(f))
	}
}

// genHTML generates html from given markdown file
func genHTML(markdownPathes []string) []string {
	var files []string
	for i, markdownPath := range markdownPathes {
		markdownFile, err := ioutil.ReadFile(markdownPath)
		check(err)

		filename := fmt.Sprintf("page%d.html", i)
		html := runBlackFriday(markdownFile)

		writeFile(html, filename)
		files = append(files, filename)
	}
	return files
}

// makeMarkdownFiles prepare markdown files splitted by '---'
func splitMarkdownFiles(markdownPath string) []string {
	file, err := os.Open(markdownPath)
	check(err)

	defer file.Close()

	var buffer bytes.Buffer
	var pages []string

	scanner := bufio.NewScanner(file)

	// split by delimiter
	for scanner.Scan() {
		if scanner.Text() == delimiter {
			pages = append(pages, buffer.String())
			buffer.Reset()
		} else {
			buffer.WriteString(fmt.Sprintf("%s\n", scanner.Text()))
		}
	}

	err = scanner.Err()
	check(err)

	// push last page to pages slice
	pages = append(pages, buffer.String())
	buffer.Reset()

	var files []string
	// save each page to files
	for i, page := range pages {
		filename := fmt.Sprintf("page%d.md", i)
		writeFile(page, filename)
		files = append(files, filename)
	}
	return files
}

// writeFile creates new file
func writeFile(content string, filename string) {
	f, err := os.Create(filename)
	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(content)
	check(err)
	w.Flush()
}

// runBlackFriday generates html from markdown filename
func runBlackFriday(mdfile []byte) string {
	htmlFlags := blackfriday.CommonHTMLFlags
	extFlags := blackfriday.CommonExtensions
	params := blackfriday.HTMLRendererParameters{Flags: htmlFlags, Title: "", CSS: ""}
	renderer := blackfriday.NewHTMLRenderer(params)
	html := blackfriday.Run(
		mdfile,
		blackfriday.WithExtensions(extFlags),
		blackfriday.WithRenderer(renderer),
	)
	return string(html)
}
