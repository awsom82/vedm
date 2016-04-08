# vedm

`vedm` is a vedomosti.ru news client for the terminal.  There's not much to this program code-wise, but this was a another happy surprise from GO as the XML parsing worked the first time.

Quick install: `go get github.com/awsom82/vedm`

## Features

* Grabs latest Vedomosti.ru headlines from the [RSS Feed](http://www.vedomosti.ru/rss/news) and displays them in the terminal
* Shows news text in terminal
* You get unlimited access to read latest news for free

![Vedomosti Terminal](screenshot.png?raw=true "Vedomosti Terminal Screenshot")

## Dependencies

* Working Go environment
* utf-8 terminal with colors

## Setup

* Install the `vedm` binary into your $GOPATH with `go get github.com/awsom82/vedm`
* Invoked `vedm` (assuming go/bin is in your path)
* If you prefer an executable named `vedm` or just want to play and reinstall the code, make sure `$GOBIN` is set, `cd $GOPATH/src/github.com/awsom82/vedm` and type `go install vedm.go`.

## Why?
Project made for golang learning and made easy reading not bad news paper for free to everyone.

Enjoy!