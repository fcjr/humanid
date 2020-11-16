# humanid - Human readable unique(-ish) ids
[![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Go Report Card][report-card-img]][report-card] [![GolangCI][golangci-img]][golangci] [![Coverage Status][cov-img]][cov]

humanid provides human readable unique(-ish) ids for consumer facing use.

## Installation

`go get github.com/fcjr/humanid`

## Usage

```go
// default separator: '-'
id := humanid.New()
fmt.Println(id)

// custom separator
hid := humanid.NewProvider("_")
fmt.Println(hid.New())
```

[doc-img]: https://img.shields.io/static/v1?label=godoc&message=reference&color=blue
[doc]: https://pkg.go.dev/github.com/fcjr/humanid?tab=doc
[ci-img]: https://travis-ci.org/fcjr/humanid.svg?branch=master
[ci]: https://travis-ci.org/fcjr/humanid
[report-card-img]: https://goreportcard.com/badge/github.com/fcjr/humanid
[report-card]: https://goreportcard.com/report/github.com/fcjr/humanid
[golangci-img]: https://golangci.com/badges/github.com/fcjr/humanid.svg
[golangci]: https://golangci.com/r/github.com/fcjr/humanid
[cov-img]: https://codecov.io/gh/fcjr/humanid/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/fcjr/humanid