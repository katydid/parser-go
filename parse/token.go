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
	"fmt"

	"katydid.org.za/go/parser-go/cast"
)

// Sprint returns a value printed as a string.
func Sprint(value Token) string {
	v, err := GetValue(value)
	if err != nil {
		return fmt.Sprintf("error:<%v>", err)
	}
	return fmt.Sprintf("%#v", v)
}

// GetValue returns the current value, without copying if possible.
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
