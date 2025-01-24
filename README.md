## parser-go

Parser interface for Go.

This includes tools for developing implementations of the parser interface, for example the `debug` package.

## Implementations

* [JSON](https://github.com/katydid/parser-go-json)
* [Protobufs](https://github.com/katydid/parser-go-proto)
* [XML](https://github.com/katydid/parser-go-xml)
* [YAML](https://github.com/katydid/parser-go-yaml)
* [Reflect](https://github.com/katydid/parser-go-reflect)

## Using the parser

If you want to use a parser for you own use case, here is a simple walk function:

```go
import "github.com/katydid/parser-go/parser"

func walk(p parser.Interface) {
    for {
        if err := p.Next(); err != nil {
            if err == io.EOF {
                break
            } else {
                panic(err)
            }
        }
        value := print(p)
        if !p.IsLeaf() {
            p.Down()
            walk(p)
            p.Up()
        }
    }
    return
}
```

We should always first call the Next method and check the error, to see if we have reached the end of the list of fields. 
Then we can retrieve a field value or field name, which can be of type int, double, uint, string, bytes or a bool. 
Lastly we can check whether the current position is a terminating leaf in the tree or if we can go down the tree and finally come back up. 
Please see the godoc for the [parser.Interface](https://pkg.go.dev/github.com/katydid/parser-go/parser#Interface) and [parser.Value](https://pkg.go.dev/github.com/katydid/parser-go/parser#Value).

## Implementing your own parser

The katydid validator supports validating any serialization format that implements the following parser interface:

```go
type Interface interface {
    // Next skips to the next field or item in an array
    Next() error
    // IsLeaf is true if this is a value and you cannot traverse down anymore
    IsLeaf() bool
    // Down traverses down into a field value, which could be another message or array. Next must always be called after Down.
    Down()
    // Up traverses up out of a field value and back to the field's next sibling. Next must always be called after Up.
    Up()
    // Value is a collection of possible values that the field might have.
    Value
}

type Value interface {
    // String returns the string value if it is a string type or an error if it is not a string.
    String() (string, error)
    // Double returns the float64 value if it is a double type or an error if it is not a double.
    Double() (float64, error)
    // Int returns the int64 value if it is an integer type or an error if it is not an integer.
    Int() (int64, error)
    // Uint returns the uint64 value if it is an unsinged integer type or an error if it is not an unsinged integer.
    Uint() (uint64, error)
    // Bool returns the bool value if it is a boolean type or an error if it is not a boolean.
    Bool() (bool, error)
    // Bytes returns a byte slice value if it is a bytes type or an error if it is not bytes.
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
