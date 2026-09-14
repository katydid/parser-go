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

package hedge

import (
	"encoding/json"
	"testing"

	"katydid.org.za/go/parser-go/expect"
	"katydid.org.za/go/parser-go/parse"
)

func TestJSONUnmarshal(t *testing.T) {
	str := `[{"Label":{"Kind":"string","Value":"A"},"Children":[{"Label":{"Kind":"string","Value":"B"}}]}]`
	h := Hedge{}
	if err := json.Unmarshal([]byte(str), &h); err != nil {
		t.Fatal(err)
	}
	t.Log(h.String())
	p := NewParser(h)
	expect.Hint(t, p, parse.EnterHint)
	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "A")
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "B")
	expect.Hint(t, p, parse.LeaveHint)
	expect.EOF(t, p)
}
