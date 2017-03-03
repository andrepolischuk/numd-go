# numd [![Build Status][travis-image]][travis-url]

> Pluralize a word

[Docs](https://godoc.org/github.com/andrepolischuk/numd-go)

* 1 dollar, 5 dollars
* 1 рубль, 2 рубля, 5 рублей

## Install

```sh
go get github.com/andrepolischuk/numd-go
```

## Usage

```go
package main

import "github.com/andrepolischuk/numd-go"

func main() {
  priceInRub := numd.Decline(24, "рубль", "рубля", "рублей") // 24 рубля
  priceInUsd := numd.Decline(51, "dollar", "dollars") // 51 dollars
}
```

## License

MIT

[travis-url]: https://travis-ci.org/andrepolischuk/numd-go
[travis-image]: https://travis-ci.org/andrepolischuk/numd-go.svg?branch=master
