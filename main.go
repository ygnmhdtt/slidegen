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
	"gopkg.in/russross/blackfriday.v2"
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
		pdfg.AddPage(wkhtmltopdf.NewPage(html))
	}
	check(pdfg.Create())
	check(pdfg.WriteFile("output.pdf"))
}

// genHTML generates html from given markdown file
func genHTML(markdownPathes []string) []string {
	var files []string

	for i, markdownPath := range markdownPathes {
		markdownFile, err := ioutil.ReadFile(markdownPath)
		check(err)

		//HTMLFlags and Renderer
		htmlFlags := blackfriday.CommonHTMLFlags //UseXHTML | Smartypants | SmartypantsFractions | SmartypantsDashes | SmartypantsLatexDashes
		// htmlFlags |= blackfriday.FootnoteReturnLinks     //Generate a link at the end of a footnote to return to the source
		// htmlFlags |= blackfriday.SmartypantsAngledQuotes //Enable angled double quotes (with Smartypants) for double quotes rendering
		// htmlFlags |= blackfriday.SmartypantsQuotesNBSP   //Enable French guillemets 損 (with Smartypants)
		renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{Flags: htmlFlags, Title: "", CSS: ""})

		//Extensions
		extFlags := blackfriday.CommonExtensions //NoIntraEmphasis | Tables | FencedCode | Autolink | Strikethrough | SpaceHeadings | HeadingIDs | BackslashLineBreak | DefinitionLists
		// extFlags |= blackfriday.Footnotes        //Pandoc-style footnotes
		// extFlags |= blackfriday.HeadingIDs       //specify heading IDs  with {#id}
		// extFlags |= blackfriday.Titleblock       //Titleblock ala pandoc
		// extFlags |= blackfriday.DefinitionLists  //Render definition lists

		html := blackfriday.Run(markdownFile, blackfriday.WithExtensions(extFlags), blackfriday.WithRenderer(renderer))

		filename := fmt.Sprintf("page%d.html", i)
		f, err := os.Create(filename)
		check(err)

		defer f.Close()

		w := bufio.NewWriter(f)
		_, err = w.WriteString(string(html))
		check(err)
		w.Flush()

		files = append(files, filename)
	}
	return files
}

// makeMarkdownFiles prepare markdown files
// this function splits md file by '---'
// file format like this:
//
// # Page 1
// This is page 1.
// ---
// # Page 2
// This is page 2.
//
func splitMarkdownFiles(markdownPath string) []string {
	file, err := os.Open(markdownPath)
	check(err)

	defer file.Close()

	var buffer bytes.Buffer
	var pages []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "---" {
			pages = append(pages, buffer.String())
			buffer.Reset()
		} else {
			buffer.WriteString(scanner.Text())
			buffer.WriteString("\n")
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
		f, err := os.Create(filename)
		check(err)

		defer f.Close()

		w := bufio.NewWriter(f)
		_, err = w.WriteString(page)
		check(err)
		w.Flush()

		files = append(files, filename)
	}
	return files
}
