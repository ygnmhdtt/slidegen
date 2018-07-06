FROM ubuntu:18.04

# todo: refactor

RUN apt update
RUN apt install -y git golang wget libfontconfig1 libxrender1 language-pack-ja-base language-pack-ja

RUN wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.3/wkhtmltox-0.12.3_linux-generic-amd64.tar.xz && \
  tar vxf wkhtmltox-0.12.3_linux-generic-amd64.tar.xz && \
  cp wkhtmltox/bin/wk* /usr/local/bin/

RUN apt upgrade -y openssl

ENV GOPATH /go
RUN go get -u github.com/SebastiaanKlippert/go-wkhtmltopdf
RUN go get -u gopkg.in/russross/blackfriday.v2

ENV COLUMNS 200
ENV LINES 50

WORKDIR /go/src/slidegen
