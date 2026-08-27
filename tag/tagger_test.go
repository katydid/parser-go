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

package tag_test

import (
	"strings"
	"testing"

	"katydid.org.za/go/parser-go/expect"
	. "katydid.org.za/go/parser-go/hedge"
	"katydid.org.za/go/parser-go/jsonschema"
	"katydid.org.za/go/parser-go/parse"
	"katydid.org.za/go/parser-go/tag"
)

func TestTaggerWithTagsAndIndexes(t *testing.T) {
	j := NewJSONSchemaAbleParser(JSONSchemaAbleHedge)
	p := tag.NewTagger(j, tag.WithTags(), tag.WithIndexes())
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "object")
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "A")
	expect.Hint(t, p, parse.ValueHint)
	expect.Int(t, p, 1)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "arrayB")
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "array")
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.Int(t, p, 0)
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "b2")
	expect.Hint(t, p, parse.FieldHint)
	expect.Int(t, p, 1)
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "b3")

	expect.Hint(t, p, parse.LeaveHint)
	expect.Hint(t, p, parse.LeaveHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "objectC")
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.Tag(t, p, "object")
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "D")
	expect.Hint(t, p, parse.ValueHint)
	expect.Int(t, p, 0)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "E")
	expect.Hint(t, p, parse.ValueHint)
	expect.Int(t, p, 1)

	expect.Hint(t, p, parse.LeaveHint)
	expect.Hint(t, p, parse.LeaveHint)

	expect.Hint(t, p, parse.LeaveHint)
	expect.Hint(t, p, parse.LeaveHint)
	expect.EOF(t, p)
}

func TestTaggerWithIndexes(t *testing.T) {
	j := NewJSONSchemaAbleParser(JSONSchemaAbleHedge)
	p := tag.NewTagger(j, tag.WithIndexes())
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "A")
	expect.Hint(t, p, parse.ValueHint)
	expect.Int(t, p, 1)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "arrayB")
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.Int(t, p, 0)
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "b2")
	expect.Hint(t, p, parse.FieldHint)
	expect.Int(t, p, 1)
	expect.Hint(t, p, parse.ValueHint)
	expect.String(t, p, "b3")

	expect.Hint(t, p, parse.LeaveHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "objectC")
	expect.Hint(t, p, parse.EnterHint)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "D")
	expect.Hint(t, p, parse.ValueHint)
	expect.Int(t, p, 0)

	expect.Hint(t, p, parse.FieldHint)
	expect.String(t, p, "E")
	expect.Hint(t, p, parse.ValueHint)
	expect.Int(t, p, 1)

	expect.Hint(t, p, parse.LeaveHint)

	expect.Hint(t, p, parse.LeaveHint)
	expect.EOF(t, p)
}

var JSONSchemaAbleHedge = Hedge{
	Node{Label: NewStringToken("A"), Children: Hedge{Node{Label: NewInt64Token(1)}}},
	Node{Label: NewStringToken("arrayB"), Children: Hedge{
		Node{Label: NewStringToken("b2")},
		Node{Label: NewStringToken("b3")},
	}},
	Node{Label: NewStringToken("objectC"), Children: Hedge{
		Node{Label: NewStringToken("D"), Children: Hedge{Node{Label: NewInt64Token(0)}}},
		Node{Label: NewStringToken("E"), Children: Hedge{Node{Label: NewInt64Token(1)}}},
	}},
}

type parser struct {
	p    tag.Parser
	prev string
}

func NewJSONSchemaAbleParser(h Hedge) tag.JSONSchemaAbleParser {
	return &parser{p: NewParser(h), prev: "object"}
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

func (p *parser) JSONSchemaType() jsonschema.JSONSchemaType {
	if strings.HasPrefix(p.prev, "array") {
		return jsonschema.JSONSchemaTypeArray
	}
	if strings.HasPrefix(p.prev, "object") {
		return jsonschema.JSONSchemaTypeObject
	}
	return jsonschema.JSONSchemaTypeUnknown
}
