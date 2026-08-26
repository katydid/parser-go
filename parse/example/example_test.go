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

package example

import (
	"testing"

	"katydid.org.za/go/parser-go/expect"
	"katydid.org.za/go/parser-go/hedge"
	"katydid.org.za/go/parser-go/parse"
)

func TestLog(t *testing.T) {
	p := hedge.NewParser(Output)
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "A")
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "1")

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "B")
	expect.Hint(t, p, parse.EnterHint)
	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "0")
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "b2")
	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "1")
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "b3")
	expect.Hint(t, p, parse.LeaveHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "C")
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "A")
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "2")

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "D")
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "3")

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "E")
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "0")
	expect.Hint(t, p, parse.EnterHint)
	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "A")
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "0")
	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "B")
	expect.Hint(t, p, parse.EnterHint)
	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "0")
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "b4")
	expect.Hint(t, p, parse.LeaveHint) // B
	expect.Hint(t, p, parse.LeaveHint) // 0

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "1")
	expect.Hint(t, p, parse.EnterHint)
	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "A")
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "0")
	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "B")
	expect.Hint(t, p, parse.EnterHint)
	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "0")
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "b5")
	expect.Hint(t, p, parse.LeaveHint) // B
	expect.Hint(t, p, parse.LeaveHint) // 1

	expect.Hint(t, p, parse.LeaveHint) // E

	expect.Hint(t, p, parse.LeaveHint) // C

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "D")
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "4")

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "F")
	expect.Hint(t, p, parse.EnterHint)
	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "0")
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "5")
	expect.Hint(t, p, parse.LeaveHint)

	expect.Hint(t, p, parse.LeaveHint)
	expect.EOF(t, p)
}
