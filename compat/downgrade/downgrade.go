//  Copyright 2013 Walter Schulze
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

// downgrade package is used for compatibility.
// DowngradeParser downgrades a new parse.Parser implementation to an old parser.Interface implementation.
package downgrade

import (
	"io"

	"github.com/katydid/parser-go/cast"
	"github.com/katydid/parser-go/parse"
	"github.com/katydid/parser-go/parser"
	"github.com/katydid/parser-go/pool"
)

type ParserWithInit interface {
	parse.Parser
	Init([]byte)
}

type InterfaceWithInit interface {
	parser.Interface
	Init([]byte) error
}

type downgradeParser struct {
	action  action
	actions []action
	state   state
	stack   []state
	parser  ParserWithInit
	pool    pool.Pool
}

// DowngradeParser downgrades a new parse.Parser implementation to an old parser.Interface implementation with an Init method.
func DowngradeParser(parser parse.Parser) parser.Interface {
	p := pool.New()
	parserWithInit, ok := parser.(ParserWithInit)
	if !ok {
		parserWithInit = &noopInit{parser}
	}
	return &downgradeParser{
		stack:   make([]state, 0, 10),
		actions: make([]action, 0, 10),
		parser:  parserWithInit,
		pool:    p,
	}
}

type noopInit struct {
	parse.Parser
}

func (n *noopInit) Init([]byte) {}

// DowngradeParserWithInit downgrades a new parse.Parser implementation to an old parser.Interface implementation with an Init method.
func DowngradeParserWithInit(parser ParserWithInit) InterfaceWithInit {
	p := pool.New()
	return &downgradeParser{
		stack:   make([]state, 0, 10),
		actions: make([]action, 0, 10),
		parser:  parser,
		pool:    p,
	}
}

func (p *downgradeParser) Init(buf []byte) error {
	p.parser.Init(buf)
	p.action = nextAction
	p.actions = p.actions[:0]
	p.state = state{}
	p.stack = p.stack[:0]
	p.pool.FreeAll()
	return nil
}

func (p *downgradeParser) nextAtStartState(action action) error {
	switch action {
	case nextAction:
		parseHint, err := p.parser.Next()
		if err != nil {
			return err
		}
		switch parseHint {
		case parse.ObjectOpenHint:
			p.state.kind = inObjectAtKeyStateKind
			parseHintNext, err := p.parser.Next()
			if err != nil {
				return err
			}
			if parseHintNext == parse.ObjectCloseHint {
				return p.eof()
			}
			return nil
		case parse.ArrayOpenHint:
			p.state.kind = inArrayIndexStateKind
			parseHintNext, err := p.parser.Next()
			if err != nil {
				return err
			}
			if parseHintNext == parse.ArrayCloseHint {
				return p.eof()
			}
			p.state.arrayElemHint = parseHintNext
			return nil
		case parse.ValueHint, parse.KeyHint:
			p.state.kind = inLeafStateKind
			return nil
		}
		panic("unreachable")
	case downAction:
		return errNextShouldBeCalled
	case upAction:
		return errNextShouldBeCalled
	}
	panic("unreachable")
}

func (p *downgradeParser) nextInLeafState(action action) error {
	switch action {
	case nextAction:
		// We already parsed the leaf, so there is no next element.
		return p.eof()
	case downAction:
		// Cannot call Down when in leaf, since we are the bottom.
		return errDownLeaf
	case upAction:
		// We can go up, if we are an array element or value for a key in an object.
		if err := p.pop(); err != nil {
			return err
		}
		// If we were in an object, then move onto next key.
		// If we were in an array, them ove onto the next element.
		return p.next()
	}
	panic("unreachable")
}

func (p *downgradeParser) nextAtEOF(action action) error {
	switch action {
	case nextAction:
		// If Next is called too many times, just keep on return EOF
		return p.eof()
	case downAction:
		// We cannot go down if we are at the EOF
		return errDownEOF
	case upAction:
		// We can go up, if we are an array element or value for a key in an object.
		if err := p.pop(); err != nil {
			return err
		}
		// If we were in an object, then move onto next key.
		// If we were in an array, them ove onto the next element.
		return p.next()
	}
	panic("unreachable")
}

func (p *downgradeParser) nextInObjectAtKeyState(action action) error {
	// inObjectAtKeyStateKind represents that we have scanned a key
	switch action {
	case nextAction:
		// We want to skip over value and move onto next key.
		// We start by skipping over the value.
		if err := p.parser.Skip(); err != nil {
			return err
		}
		// Next we move onto the Next key.
		parseKind, err := p.parser.Next()
		if parseKind == parse.ObjectCloseHint {
			// If the Object has ended, we return eof
			return p.eof()
		}
		return err
	case downAction:
		// Set the state to be ready to parse to next key, when Up is called.
		p.state.kind = inObjectAtValueStateKind
		// We do not want to skip over the value, we want to continue into value.
		// We start by pushing the current state to the stack.
		if err := p.push(); err != nil {
			return err
		}
		// The state is reset to be the start state.
		// We can call this' Next method, instead of the parser's Next method.
		return p.next()
	case upAction:
		// We want to skip over value and the rest of the object and move onto the next key.
		// We start by skipping over the the value.
		if err := p.parser.Skip(); err != nil {
			return err
		}
		// Next we skip over the rest of the object.
		if err := p.parser.Skip(); err != nil {
			return err
		}
		// Now we pop the stack
		if err := p.pop(); err != nil {
			return err
		}
		// Finally we move onto the next key or element, if we were in an array.
		return p.next()
	}
	panic("unreachable")
}

func (p *downgradeParser) nextInObjectAtValueState(action action) error {
	// inObjectAtValueStateKind represents that we have scanned a value and Up was called.
	switch action {
	case nextAction:
		// Up was just called and we need to scan to the Next key.
		parseKind, err := p.parser.Next()
		if parseKind == parse.ObjectCloseHint {
			// If the Object has ended, we return eof
			return p.eof()
		}
		// Set the state to the next key
		p.state.kind = inObjectAtKeyStateKind
		return err
	case downAction:
		// We can't call Down right, while at the end of value.
		return errDown
	case upAction:
		// We want to skip over the rest of the object and move onto the next key.
		// We skip over the rest of the object.
		if err := p.parser.Skip(); err != nil {
			return err
		}
		// Now we pop the stack
		if err := p.pop(); err != nil {
			return err
		}
		// Finally we move onto the next key or element, if we were in an array.
		return p.next()
	}
	panic("unreachable")
}

func (p *downgradeParser) nextInArrayIndexState(action action) error {
	// inArrayIndexState represents that we have scanned an element, if it was null, bool, number or string and the first key of an object or .
	switch action {
	case nextAction:
		p.state.arrayIndex += 1
		switch p.state.arrayElemHint {
		case parse.ObjectOpenHint, parse.ArrayOpenHint:
			if err := p.parser.Skip(); err != nil {
				return err
			}
		case parse.ValueHint:
		case parse.KeyHint:
			panic("unreachable: unexpected key hint in array")
		default:
			panic("unreachable")
		}
		parseHint, err := p.parser.Next()
		if err != nil {
			return err
		}
		if parseHint == parse.ArrayCloseHint {
			return p.eof()
		}
		p.state.arrayElemHint = parseHint
		return nil
	case downAction:
		// We are at an array element that we are representing as an index.
		// We do not need parse another thing, simply update the state.
		p.state.kind = inArrayAfterIndexStateKind
		if err := p.push(); err != nil {
			return err
		}
		switch p.state.arrayElemHint {
		case parse.ObjectOpenHint:
			p.state.kind = inObjectAtKeyStateKind
			parseHintNext, err := p.parser.Next()
			if err != nil {
				return err
			}
			if parseHintNext == parse.ObjectCloseHint {
				return p.eof()
			}
			return nil
		case parse.ArrayOpenHint:
			p.state.kind = inArrayIndexStateKind
			parseHintNext, err := p.parser.Next()
			if err != nil {
				return err
			}
			if parseHintNext == parse.ArrayCloseHint {
				return p.eof()
			}
			p.state.arrayElemHint = parseHintNext
			return nil
		case parse.ValueHint:
			p.state.kind = inLeafStateKind
			return nil
		case parse.KeyHint:
			panic("unreachable: unexpected key hint in array")
		}
		panic("unreachable")
	case upAction:
		switch p.state.arrayElemHint {
		case parse.ObjectOpenHint, parse.ArrayOpenHint:
			// skip the element
			if err := p.parser.Skip(); err != nil {
				return err
			}
		case parse.ValueHint:
		case parse.KeyHint:
			panic("unreachable: unexpected key hint in array")
		default:
			panic("unreachable")
		}
		// Skip the rest of the array
		if err := p.parser.Skip(); err != nil {
			return err
		}
		if err := p.pop(); err != nil {
			return err
		}
		return p.next()
	}
	panic("unreachable")
}

func (p *downgradeParser) nextInArrayAfterIndexState(action action) error {
	// This is after Up was called on an element.
	switch action {
	case nextAction:
		p.state.arrayIndex += 1
		parseHint, err := p.parser.Next()
		if err != nil {
			return err
		}
		if parseHint == parse.ArrayCloseHint {
			return p.eof()
		}
		p.state.kind = inArrayIndexStateKind
		p.state.arrayElemHint = parseHint
		return nil
	case downAction:
		return errDown
	case upAction:
		// Skip the rest of the array
		if err := p.parser.Skip(); err != nil {
			return err
		}
		if err := p.pop(); err != nil {
			return err
		}
		return p.next()
	}
	panic("unreachable")
}

func (p *downgradeParser) eof() error {
	if len(p.stack) == 0 {
		// if we are at the top of stack, then check that there is no more input left.
		_, err := p.parser.Next()
		if err == nil {
			return errExpectedEOF
		}
		if err != io.EOF {
			return err
		}
	}
	// When EOF is returned also set the state to an EOF state.
	// This state allows us to call Up.
	p.state.kind = atEOFStateKind
	return io.EOF
}

func (p *downgradeParser) next() error {
	action := p.action
	// do not forget to reset action
	p.action = nextAction
	switch p.state.kind {
	case atStartStateKind:
		return p.nextAtStartState(action)
	case inLeafStateKind:
		return p.nextInLeafState(action)
	case inArrayIndexStateKind:
		return p.nextInArrayIndexState(action)
	case inArrayAfterIndexStateKind:
		return p.nextInArrayAfterIndexState(action)
	case inObjectAtKeyStateKind:
		return p.nextInObjectAtKeyState(action)
	case inObjectAtValueStateKind:
		return p.nextInObjectAtValueState(action)
	case atEOFStateKind:
		return p.nextAtEOF(action)
	}
	panic("unreachable")
}

func (p *downgradeParser) nexts() error {
	lastAction := p.action
	for i := 0; i < len(p.actions); i++ {
		p.action = p.actions[i]
		if err := p.next(); err != nil {
			// ignore EOF if we still have more actions to perform.
			if err != io.EOF || i == len(p.actions) {
				return err
			}
		}
	}
	p.actions = p.actions[:0]
	p.action = lastAction
	return p.next()
}

func (p *downgradeParser) Next() error {
	return p.nexts()
}

func (p *downgradeParser) Down() {
	p.pushAction(downAction)
}

func (p *downgradeParser) Up() {
	p.pushAction(upAction)
}

func (p *downgradeParser) pushAction(newAction action) {
	if p.action == nextAction {
		p.action = newAction
		return
	}
	// when Up is called straight after Down, we simply call next.
	if p.action == downAction && newAction == upAction {
		p.action = nextAction
		return
	}
	p.actions = append(p.actions, p.action)
	p.action = newAction
}

func (p *downgradeParser) push() error {
	// Append the current state to the stack.
	p.stack = append(p.stack, p.state)
	p.state.kind = atStartStateKind
	p.state.arrayIndex = 0
	return nil
}

func (p *downgradeParser) pop() error {
	if len(p.stack) == 0 {
		return errPop
	}
	top := len(p.stack) - 1
	// Set the current state to the state on top of the stack.
	p.state = p.stack[top]
	// Remove the state on the top the stack from the stack,
	// but do it in a way that keeps the capacity,
	// so we can reuse it the next time Down is called.
	p.stack = p.stack[:top]
	return nil
}

func (p *downgradeParser) IsLeaf() bool {
	return p.state.kind == inLeafStateKind
}

func (p *downgradeParser) Bool() (bool, error) {
	tokenKind, _, err := p.parser.Token()
	if err != nil {
		return false, err
	}
	if tokenKind == parse.FalseKind {
		return false, nil
	}
	if tokenKind == parse.TrueKind {
		return true, nil
	}
	return false, errNotBool
}

func (p *downgradeParser) Int() (int64, error) {
	if p.state.kind == inArrayIndexStateKind {
		return p.state.arrayIndex, nil
	}
	tokenKind, bs, err := p.parser.Token()
	if err != nil {
		return 0, err
	}
	if tokenKind != parse.Int64Kind {
		return 0, errNotInt
	}
	return cast.ToInt64(bs), nil
}

func (p *downgradeParser) Uint() (uint64, error) {
	if p.state.kind == inArrayIndexStateKind {
		return 0, errNotUint
	}
	i, err := p.Int()
	if err != nil {
		return 0, err
	}
	if i >= 0 {
		return uint64(i), nil
	}
	return 0, errNotUint
}

func (p *downgradeParser) Double() (float64, error) {
	tokenKind, bs, err := p.parser.Token()
	if err != nil {
		return 0, err
	}
	if tokenKind != parse.Float64Kind {
		return 0, errNotFloat
	}
	return cast.ToFloat64(bs), nil
}

func (p *downgradeParser) String() (string, error) {
	tokenKind, bs, err := p.parser.Token()
	if err != nil {
		return "", err
	}
	if tokenKind != parse.StringKind && tokenKind != parse.DecimalKind {
		return "", errNotString
	}
	return cast.ToString(bs), nil
}

var nullBytes = []byte{'n', 'u', 'l', 'l'}

func (p *downgradeParser) Bytes() ([]byte, error) {
	tokenKind, bs, err := p.parser.Token()
	if err != nil {
		return nil, err
	}
	if tokenKind == parse.NullKind {
		return nullBytes, nil
	}
	return bs, nil
}
