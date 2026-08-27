//  Copyright 2026 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package log

import (
	"testing"

	"katydid.org.za/go/parser-go/expect"
	. "katydid.org.za/go/parser-go/hedge"
	"katydid.org.za/go/parser-go/parse"
)

func TestParseFloat64(t *testing.T) {
	n := Node{
		Label: NewStringToken("num"),
		Children: []Node{
			{Label: NewFloat64Token(3.14)},
		},
	}
	var p parse.Parser = NewParser([]Node{n})
	p = WrapParser(p)
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "num")

	expect.Hint(t, p, parse.ValueHint)
	// There used to be a bug where this expect.Float(t, p, 3.14) failed,
	// because the value was not properly copied by log.WrapParser.
	expect.Float(t, p, 3.14)

	expect.Hint(t, p, parse.LeaveHint)
	expect.EOF(t, p)
}

func TestParseInt64(t *testing.T) {
	n := Node{
		Label: NewInt64Token(1),
		Children: []Node{
			{
				Label: NewInt64Token(0),
				Children: []Node{
					{Label: NewInt64Token(123)},
				},
			},
		},
	}
	var p parse.Parser = NewParser([]Node{n})
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.Int(t, p, 1)
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.Int(t, p, 0)
	expect.Hint(t, p, parse.ValueHint)
	expect.Int(t, p, 123)

	expect.Hint(t, p, parse.LeaveHint)

	expect.Hint(t, p, parse.LeaveHint)
	expect.EOF(t, p)
}
