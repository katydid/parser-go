// Copyright 2026 Walter Schulze
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

package relaxng

import (
	"io"

	"katydid.org.za/go/parser-go/parse"
)

type Parser interface {
	parse.Parser
	Reset()
}

type RelaxNGableParser interface {
	parse.Parser
	RelaxNGable
	Reset()
}

type tagger struct {
	p     RelaxNGableParser
	tag   bool
	alloc func(size int) []byte

	state state
	stack []state
}

var elemsTagToken = []byte("elems")
var attrsTagToken = []byte("attrs")

// NewTagger can tag elements and attributes.
// The following xml: `<A a="b" c="d"><B>C/B><D>E</D></A>`
// is parsed as: `{"elems": {"A": {"attrs": {"a": "b", "c": "d"}, "elems": {"B": "C", "D": "E"}}}}`.
// The kind returned from the Token method for
// "elems" and "attrs" will be parse.TagKind.
func NewTagger(p RelaxNGableParser, opts ...Option) Parser {
	t := &tagger{
		p:   p,
		tag: false,
		alloc: func(size int) []byte {
			return make([]byte, size)
		},
		state: state{},
		stack: make([]state, 0, 10),
	}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func (t *tagger) Reset() {
	// Reset the state.
	t.state = state{}
	// Shrink the stack's length, but keep its capacity,
	// so we can reuse it on the next parse.
	t.stack = t.stack[:0]
	// Reset the parser too.
	t.p.Reset()
}

func (t *tagger) nextStart(h parse.Hint) (parse.Hint, error) {
	if t.tag {
		switch h {
		case parse.EnterHint:
			t.down(enterState)
			return parse.EnterHint, nil
		case parse.LeaveHint:
			if err := t.up(); err != nil {
				return parse.UnknownHint, err
			}
			return parse.LeaveHint, nil
		case parse.FieldHint:
			switch t.p.RelaxNGType() {
			case RelaxNGTypeUnknown:
				t.state.kind = unknownFieldState
				return h, nil
			case RelaxNGTypeElem:
				t.state.kind = elemStartState
				return parse.FieldHint, nil
			case RelaxNGTypeAttr:
				t.state.kind = attrStartState
				return parse.FieldHint, nil
			default:
				panic("unreachable")
			}
		case parse.ValueHint:
			t.state.kind = valueState
			return h, nil
		default:
			panic("unreachable")
		}
	} else {
		switch h {
		case parse.EnterHint:
			t.down(startState)
			return parse.EnterHint, nil
		case parse.LeaveHint:
			if err := t.up(); err != nil {
				return parse.UnknownHint, err
			}
			return parse.LeaveHint, nil
		case parse.FieldHint:
			return h, nil
		case parse.ValueHint:
			return h, nil
		default:
			panic("unreachable")
		}
	}
}

func (t *tagger) Next() (parse.Hint, error) {
	switch t.state.kind {
	case startState:
		h, err := t.p.Next()
		if err != nil {
			return parse.UnknownHint, err
		}
		return t.nextStart(h)
	case enterState:

	case unknownFieldState:

	case elemStartState:
		t.down(elemStartedState)
		return parse.EnterHint, nil
	case elemStartedState:
		t.state.kind = elemFirstState
		return parse.FieldHint, nil
	case elemFirstState:

	case attrStartState:
		t.state.kind = attrLeaveState
		t.down(attrStartedState)
		return parse.EnterHint, nil
	case attrStartedState:
		t.state.kind = attrFieldState
		return parse.FieldHint, nil
	case attrFieldState:
		h, err := t.p.Next()
		if err != nil {
			return parse.UnknownHint, err
		}
		if h != parse.ValueHint {
			return parse.UnknownHint, errExpectedAttrValue
		}
		t.state.kind = attrValueState
		return parse.ValueHint, nil
	case attrValueState:
		h, err := t.p.Next()
		if err != nil {
			return parse.UnknownHint, err
		}
		switch h {
		case parse.LeaveHint:
			if err := t.up(); err != nil {
				return parse.UnknownHint, err
			}
			return parse.LeaveHint, nil
		case parse.FieldHint:
			switch t.p.RelaxNGType() {
			case RelaxNGTypeUnknown:
				t.state.kind = unknownFieldState
				return h, nil
			case RelaxNGTypeElem:
				t.state.kind = elemStartState
				return parse.FieldHint, nil
			case RelaxNGTypeAttr:
				t.state.kind = attrStartState
				return parse.FieldHint, nil
			default:
				panic("unreachable")
			}
		}
	case attrLeaveState:
		if err := t.up(); err != nil {
			return parse.UnknownHint, err
		}
		return parse.LeaveHint, nil
	case valueState:

	case endState:
		return parse.UnknownHint, io.EOF
	}
	return parse.UnknownHint, nil
}

func (t *tagger) Skip() error {
	return nil
}

func (t *tagger) Token() (parse.Kind, []byte, error) {
	switch t.state.kind {
	case elemStartState:
		return parse.TagKind, elemsTagToken, nil
	case attrStartState:
		return parse.TagKind, attrsTagToken, nil
	}
	return t.p.Token()
}

func (t *tagger) down(stateKind stateKind) {
	// Append the current state to the stack.
	t.stack = append(t.stack, t.state)
	// Create a new state.
	t.state.kind = stateKind
}

func (t *tagger) up() error {
	if len(t.stack) == 0 {
		return errUnexpectedLeave
	}
	top := len(t.stack) - 1
	// Set the current state to the state on top of the stack.
	t.state = t.stack[top]
	// Remove the state on the top the stack from the stack,
	// but do it in a way that keeps the capacity,
	// so we can reuse it the next time Down is called.
	t.stack = t.stack[:top]
	if len(t.stack) == 0 {
		t.state.kind = endState
	}
	return nil
}
