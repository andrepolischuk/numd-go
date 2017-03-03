# numd [![Build Status][travis-image]][travis-url]

> Склонение слов после числительных

[Docs](https://godoc.org/github.com/andrepolischuk/numd-go)

* 1 рубль
* 2 рубля
* 5 рублей

## Установка

```go
import "github.com/andrepolischuk/numd-go"
```

## API

### numd.Decline(num, nominative, genitiveSingular, genitivePlural)

Получаем слово в нужном склонении

```go
price := numd.Decline(24, "рубль", "рубля", "рублей") // 24 рубля
length := numd.Decline(51, "метр", "метра", "метров") // 51 метр
```

## License

MIT

[travis-url]: https://travis-ci.org/andrepolischuk/numd-go
[travis-image]: https://travis-ci.org/andrepolischuk/numd-go.svg?branch=master
