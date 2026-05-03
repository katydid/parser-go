## parser-go

[Parser](https://github.com/katydid/parser) interface for Go.

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
import "github.com/katydid/parser-go/parse"

func Walk(p parse.Parser) error {
	for {
		_, err := p.Next()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if _, _, err := p.Token(); err != nil {
			return err
		}
	}
	return nil
}
```

We should always first call the `Next` method and check the error (for io.EOF), to see if we have reached the end of the list of fields. 
Then we can retrieve a field value or field name via the `Token` method.

## Implementing your own parser

The katydid validator supports validating any serialization format that implements the following parser interface:

```go
type Parser interface {
	// Next returns the Hint of the token or an error.
	Next() (Hint, error)

	// Skip allows the user to skip over uninteresting parts of the parse tree.
	// Based on the Hint skip has different intuitive behaviours.
	// If the Hint was:
	// * '{': the whole Map is skipped.
	// * 'k': the key's value is skipped.
	// * '[': the whole List is skipped.
	// * 'v': the rest of the Map or List is skipped.
	// * ']': same as calling Next and ignoring the Hint.
	// * '}': same as calling Next and ignoring the Hint.
	Skip() error

	// Tokenize parses the current token.
	Token() (Kind, []byte, error)
}
```

This interface allows for the implementation of an online pull based parser.
That is a parser that lazily parses the input as the methods are called and only parses the input once, without backtracking. 
Exercising the parser can be done with the [debug.Parse](https://pkg.go.dev/github.com/katydid/parser-go/parse/debug#Parse) function. 
The `Walk` function also returns some debugging output, which should be useful in the development of your own parser.

Your parser should also be able to handle skipping of some of the input, via the `Skip` method. 
The [debug.RandomParse](https://pkg.go.dev/github.com/katydid/parser-go/parse/debug#RandomParse) function is useful for testing this type of robustness in your parser.

See the [Parser Documentation](https://github.com/katydid/parser) for more details.