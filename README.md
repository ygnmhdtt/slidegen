# slidegen - generate presentation slide from markdown file

![](https://github.com/ygnmhdtt/slidegen/blob/master/samples/demo.gif)

## About

slidegen is the CLI tool to generate presentation slide from markdown file.  
Markdown style is always GFM([GitHub Flavored Markdown](https://github.github.com/gfm/)).

## Prerequisite

* Golang
* [wkhtmltopdf](https://wkhtmltopdf.org/downloads.html)

### Markdown file format and generated pdf

I prepared [sample markdown and pdf](https://github.com/ygnmhdtt/slidegen/tree/master/samples) for example.
Markdown file must be contain `---` for delimiter of pdf pages. ([like this](https://raw.githubusercontent.com/ygnmhdtt/slidegen/master/samples/awscli-on-container.md))

## Installation and Usage

* MacOSX is supported.

### Mac OSX

* install wkhtmltopdf(ref: [How To Install and Run wkhtmltopdf on Mac OsX](https://stackoverflow.com/questions/10375168/how-to-install-and-run-wkhtmltopdf-on-mac-osx-10-7-3-for-use-in-a-php-applicatio))

* get binary

```sh
$ go get github.com/ygnmhdtt/slidegen
```

#### Usage

```sh
$ slidegen your/markdown/file.md
```

## Development

Pull Requests are always welcome!

## LISENCE

MIT
