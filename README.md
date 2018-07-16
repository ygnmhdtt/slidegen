# slidegen - generate presentation slide from markdown file

## About

slidegen is the CLI tool to generate presentation slide from markdown file.  
Markdown style is always GFM([GitHub Flavored Markdown](https://github.github.com/gfm/)).

## Prerequisite

* Golang
* [wkhtmltopdf](https://wkhtmltopdf.org/downloads.html)

### Running in docker container

If you use docker, running slidegen in container is the easiest way.

### Markdown file format and generated pdf

I prepared [sample markdown]() and [generated pdf]() for example.

## Installation and Usage

* Both ways, the name of generated pdf is `output.pdf`

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

### Linux

If you use Linux, use docker for installation.

```sh
$ go get github.com/ygnmhdtt/slidegen
$ cd $GOPATH/src/github.com/ygnmhdtt/slidegen
$ make build
```

#### Usage

```
$ make gen F=your/markdown/file.md
```

## LISENCE

MIT
