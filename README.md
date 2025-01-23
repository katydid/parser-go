## parser-go

Parser interface for Go.

This includes tools for developing implementations of the parser interface, for example the `debug` package.

## Implementations

* [JSON](https://github.com/katydid/parser-go-json)
* [Protobufs](https://github.com/katydid/parser-go-proto)
* [XML](https://github.com/katydid/parser-go-xml)
* [YAML](https://github.com/katydid/parser-go-yaml)
* [Reflect](https://github.com/katydid/parser-go-reflect)

## Implementing your own parser

The katydid validator supports validating any serialization format that implements the following parser interface:

```go
type Interface interface {
    Next() error
    IsLeaf() bool
    Up()
    Down()
    Value
}

type Value interface {
    Double() (float64, error)
    Int() (int64, error)
    Uint() (uint64, error)
    Bool() (bool, error)
    String() (string, error)
    Bytes() ([]byte, error)
}
```

This interface allows for the implementation of an online pull based parser. 
That is a parser that lazily parses the input as the methods are called and only parses the input once, without backtracking. 
Exercising the parser can be done with the [debug.Walk](https://pkg.go.dev/github.com/katydid/parser-go/parser/debug#Walk) function. 
The Walk function also returns some debugging output, which should be useful in the development of your own parser.

Your parser should also be able to handle skipping of some of the input. 
This happens when the Walk function returns before encountering an `io.EOF`. 
The [debug.RandomWalk](https://pkg.go.dev/github.com/katydid/parser-go/parser/debug#RandomWalk) function is useful for testing this type of robustness in your parser.
