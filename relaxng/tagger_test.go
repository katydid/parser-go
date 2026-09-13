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

package relaxng_test

import (
	"strings"
	"testing"

	"katydid.org.za/go/parser-go/expect"
	. "katydid.org.za/go/parser-go/hedge"
	"katydid.org.za/go/parser-go/jsonschema"
	"katydid.org.za/go/parser-go/parse"
	"katydid.org.za/go/parser-go/relaxng"
)

func TestTags(t *testing.T) {
	x := NewRelaxNGableParser(RelaxNGableHedge)
	p := relaxng.NewTagger(x, relaxng.WithTags())

	expect.Hint(t, p, parse.EnterHint) // root

	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "elems")
	expect.Hint(t, p, parse.EnterHint) // root.elems

	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "elem_A")
	expect.Hint(t, p, parse.EnterHint) // elem_A

	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "attrs")
	expect.Hint(t, p, parse.EnterHint) // elem_A.attrs
	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "attr_a")
	expect.Hint(t, p, parse.ValueHint)
	expect.Tag(t, p, "z")
	expect.Hint(t, p, parse.LeaveHint) // elem_A.attrs

	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "elems")
	expect.Hint(t, p, parse.EnterHint) // elem_A.elems

	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "elem_B")
	expect.Hint(t, p, parse.EnterHint) // elem_B
	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "attrs")
	expect.Hint(t, p, parse.EnterHint) // elem_B.attrs
	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "attr_b")
	expect.Hint(t, p, parse.ValueHint)
	expect.Tag(t, p, "y")
	expect.Hint(t, p, parse.LeaveHint) // elem_B.attrs
	expect.Hint(t, p, parse.LeaveHint) // elem_B

	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "elem_C")
	expect.Hint(t, p, parse.EnterHint) // elem_C
	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "attrs")
	expect.Hint(t, p, parse.EnterHint) // elem_C.attrs
	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "attr_c")
	expect.Hint(t, p, parse.ValueHint)
	expect.Tag(t, p, "x")
	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "attr_d")
	expect.Hint(t, p, parse.ValueHint)
	expect.Tag(t, p, "w")
	expect.Hint(t, p, parse.LeaveHint) // elem_C.attrs
	expect.Hint(t, p, parse.LeaveHint) // elem_C

	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "elem_D")
	expect.Hint(t, p, parse.EnterHint) // elem_D
	expect.Hint(t, p, parse.LeaveHint) // elem_D

	expect.Hint(t, p, parse.LeaveHint) // elem_A.elems

	expect.Hint(t, p, parse.LeaveHint) // elem_A

	expect.Hint(t, p, parse.LeaveHint) // root.elems

	expect.Hint(t, p, parse.LeaveHint) // root
	expect.EOF(t, p)
}

// <A a="z"><B b="y"/><C c="x" d="w"></C><D/></A>
var RelaxNGableHedge = Hedge{
	Node{Label: NewStringToken("elem_A"), Children: Hedge{
		Node{Label: NewStringToken("attr_a"), Children: Hedge{Node{Label: NewStringToken("z"), Children: nil}}},
		Node{Label: NewStringToken("elem_B"), Children: Hedge{
			Node{Label: NewStringToken("attr_b"), Children: Hedge{Node{Label: NewStringToken("y"), Children: nil}}},
		}},
		Node{Label: NewStringToken("elem_C"), Children: Hedge{
			Node{Label: NewStringToken("attr_c"), Children: Hedge{Node{Label: NewStringToken("x"), Children: nil}}},
			Node{Label: NewStringToken("attr_d"), Children: Hedge{Node{Label: NewStringToken("w"), Children: nil}}},
		}},
		Node{Label: NewStringToken("elem_D"), Children: nil},
	}},
}

type parser struct {
	p    jsonschema.Parser
	prev string
}

func NewRelaxNGableParser(h Hedge) relaxng.RelaxNGableParser {
	return &parser{p: NewParser(h), prev: "elem_"}
}

func (p *parser) Reset() {
	p.p.Reset()
}

func (p *parser) Next() (parse.Hint, error) {
	hint, err := p.p.Next()
	if err == nil && hint == parse.FieldHint {
		kind, tok, err := p.p.Token()
		if err == nil && kind == parse.StringKind {
			p.prev = string(tok)
		}
	}
	return hint, err
}

func (p *parser) Skip() error {
	return p.p.Skip()
}

func (p *parser) Token() (parse.Kind, []byte, error) {
	return p.p.Token()
}

func (p *parser) RelaxNGType() relaxng.RelaxNGType {
	if strings.HasPrefix(p.prev, "elem_") {
		return relaxng.RelaxNGTypeElem
	}
	if strings.HasPrefix(p.prev, "attr_") {
		return relaxng.RelaxNGTypeAttr
	}
	return relaxng.RelaxNGTypeUnknown
}
