FROM ubuntu:18.04

# install tools and japanese fonts
RUN apt update && \
    apt install -y git golang wget libfontconfig1 libxrender1 language-pack-ja-base language-pack-ja fonts-takao-gothic fonts-takao-mincho fontconfig && \
    echo "ja_JP.UTF-8 UTF-8" >> /etc/locale.gen && \
    locale-gen && \
    bash -c "fc-cache -f -v"

# install wkhtmltopdf binary
RUN wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.3/wkhtmltox-0.12.3_linux-generic-amd64.tar.xz && \
  tar vxf wkhtmltox-0.12.3_linux-generic-amd64.tar.xz && \
  cp wkhtmltox/bin/wk* /usr/local/bin/

ENV GOPATH /go
ENV COLUMNS 200
ENV LINES 50
ENV LANG=ja_JP.UTF-8

WORKDIR /go/src/slidegen
