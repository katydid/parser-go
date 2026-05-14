// Copyright 2025 Walter Schulze
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parse

import (
	"errors"
	"fmt"

	"github.com/katydid/parser-go/cast"
)

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
	Token
}

type ParserWithInit interface {
	Parser
	Init([]byte)
}

type Token interface {
	Token() (Kind, []byte, error)
}

// GetValue returns the current value.
func GetValue(p Token) (any, error) {
	kind, val, err := p.Token()
	if err != nil {
		return nil, err
	}
	switch kind {
	case NullKind:
		return nil, nil
	case FalseKind:
		return false, nil
	case TrueKind:
		return true, nil
	case BytesKind:
		return val, nil
	case StringKind:
		return cast.ToString(val), nil
	case Int64Kind:
		return cast.ToInt64(val), nil
	case Float64Kind:
		return cast.ToFloat64(val), nil
	case DecimalKind:
		return cast.ToString(val), nil
	case NanosecondsKind:
		return cast.ToInt64(val), nil
	case DateTimeKind:
		return cast.ToString(val), nil
	case TagKind:
		return cast.ToString(val), nil
	case UnknownKind:
		return nil, errUnknownKind
	default:
		panic("unreachable")
	}
}

var errUnknownKind = errors.New("unknown kind")

// Sprint returns a value printed as a string.
func Sprint(value Token) string {
	v, err := GetValue(value)
	if err != nil {
		return fmt.Sprintf("error:<%v>", err)
	}
	return fmt.Sprintf("%#v", v)
}
