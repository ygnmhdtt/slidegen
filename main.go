package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"gopkg.in/russross/blackfriday.v2"
)

func panicOnErr(err interface{}) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func main() {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	panicOnErr(err)

	mds := makeMarkdownFiles()
	for _, md := range mds {
		html := genHTML(md)
		pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))
	}

	panicOnErr(pdfg.Create())
	panicOnErr(pdfg.WriteFile("output.pdf"))

}

// genHTML generates html from given markdown file
func genHTML(markdownPath string) string {
	markdownFile, err := ioutil.ReadFile(markdownPath)
	panicOnErr(err)

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
	return string(html)
}

// makeMarkdownFiles prepare markdown files
// this function splits md file by '---'
// file format like this:
// # Page 1
// This is page 1.
// ---
// # Page 2
// This is page 2.
func makeMarkdownFiles() []string {
	files := []string{"apple", "orange", "lemon"}
	return files
}
