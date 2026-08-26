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
	"fmt"
	"io"

	"katydid.org.za/go/parser-go/cast"
	"katydid.org.za/go/parser-go/parse"
)

type state byte

const unknownState = state(0)
const openedState = state(1)
const valueState = state(2)
const pairState = state(3)
const fieldState = state(4)
const eofState = state(5)

type parser struct {
	current state
	k       string
	v       string
	h       Hedge
	stack   []Hedge
	alloc   func(size int) []byte
}

func NewParser(h Hedge) parse.Parser {
	return &parser{
		current: unknownState,
		h:       h,
		alloc: func(size int) []byte {
			return make([]byte, size)
		},
	}
}

func (p *parser) pop() {
	if len(p.stack) == 0 {
		p.current = eofState
		return
	}
	p.current = openedState
	p.h = p.stack[len(p.stack)-1]
	p.stack = p.stack[:len(p.stack)-1]
}

func (p *parser) nextNode(current Node, nexts Hedge) parse.Hint {
	if len(current.Children) == 0 {
		p.current = valueState
		p.v = current.Label
		p.h = nexts
		return parse.ValueHint
	}
	if len(current.Children) == 1 {
		if len(current.Children[0].Children) == 0 {
			p.current = pairState
			p.k = current.Label
			p.v = current.Children[0].Label
			p.h = nexts
			return parse.FieldHint
		}
	}
	p.stack = append(p.stack, nexts)
	p.current = fieldState
	p.k = current.Label
	p.h = current.Children
	return parse.FieldHint
}

func (p *parser) Next() (parse.Hint, error) {
	switch p.current {
	case unknownState:
		p.current = openedState
		return parse.EnterHint, nil
	case openedState:
		if len(p.h) == 0 {
			p.pop()
			return parse.LeaveHint, nil
		}
		return p.nextNode(p.h[0], p.h[1:]), nil
	case valueState:
		if len(p.h) == 0 {
			p.pop()
			return parse.LeaveHint, nil
		}
		return p.nextNode(p.h[0], p.h[1:]), nil
	case pairState:
		p.current = valueState
		return parse.ValueHint, nil
	case fieldState:
		p.current = openedState
		return parse.EnterHint, nil
	case eofState:
		return parse.UnknownHint, io.EOF
	}
	panic("unreachable")
}

func (p *parser) Skip() error {
	switch p.current {
	case unknownState:
		p.pop()
		return nil
	case openedState:
		p.pop()
		return nil
	case valueState:
		p.pop()
		return nil
	case pairState:
		p.current = openedState
		return nil
	case fieldState:
		p.pop()
		return nil
	case eofState:
		return nil
	}
	panic("unreachable")
}

func (p *parser) Token() (parse.Kind, []byte, error) {
	switch p.current {
	case unknownState:
		return parse.UnknownKind, nil, fmt.Errorf("unknown Token")
	case openedState:
		return parse.UnknownKind, nil, fmt.Errorf("unknown Token")
	case valueState:
		return parse.StringKind, cast.FromString(p.v, p.alloc), nil
	case pairState:
		return parse.StringKind, cast.FromString(p.k, p.alloc), nil
	case fieldState:
		return parse.StringKind, cast.FromString(p.k, p.alloc), nil
	case eofState:
		return parse.UnknownKind, nil, fmt.Errorf("unknown Token")
	}
	panic("unreachable")
}
