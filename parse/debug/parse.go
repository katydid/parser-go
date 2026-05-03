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

package debug

import (
	"fmt"
	"io"

	"github.com/katydid/parser-go/parse"
)

// Parse parses through the whole parser in a top down manner and records the values into a Nodes structute.
func Parse(p parse.Parser) (Nodes, error) {
	nodes := make(Nodes, 0)
	for {
		hint, err := p.Next()
		if err != nil {
			if err == io.EOF {
				return nodes, nil
			}
			return nil, err
		}
		switch hint {
		case parse.EnterHint:
			children, err := Parse(p)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, children...)
		case parse.FieldHint:
			name, err := parse.GetValue(p)
			if err != nil {
				return nil, err
			}
			childHint, err := p.Next()
			if err != nil {
				return nil, err
			}
			switch childHint {
			case parse.ValueHint:
				val, err := parse.GetValue(p)
				if err != nil {
					return nil, err
				}
				nodes = append(nodes, Node{Label: fmt.Sprintf("%v", name), Children: []Node{{Label: fmt.Sprintf("%v", val)}}})
			case parse.EnterHint:
				children, err := Parse(p)
				if err != nil {
					return nil, err
				}
				nodes = append(nodes, Node{Label: fmt.Sprintf("%v", name), Children: children})
			}
		case parse.ValueHint:
			val, err := parse.GetValue(p)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, Node{Label: fmt.Sprintf("%v", val), Children: nil})
		case parse.LeaveHint:
			return nodes, nil
		}
	}
}

// RandomParse does a random parse of the parser, given two probabilities.
//
//	next is passed to r.Intn and when a non zero value is returned the Next method on the parser is called.
//	skip is passed to r.Intn and when a zero value is returned the Skip method on the parser is called.
//
// RandomParse is great for testing that the implemented parser can handle random skipping of parts of the tree.
func RandomParse(p parse.Parser, r Rand, next, skip int) (Nodes, error) {
	nodes := make(Nodes, 0)
	for {
		if r.Intn(next) == 0 {
			return nodes, nil
		}
		hint, err := p.Next()
		if err != nil {
			if err == io.EOF {
				return nodes, nil
			}
			return nil, err
		}
		switch hint {
		case parse.EnterHint:
			children, err := Parse(p)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, children...)
		case parse.FieldHint:
			name, err := parse.GetValue(p)
			if err != nil {
				return nil, err
			}
			if r.Intn(skip) == 0 {
				p.Skip()
			} else {
				childHint, err := p.Next()
				if err != nil {
					return nil, err
				}
				switch childHint {
				case parse.ValueHint:
					val, err := parse.GetValue(p)
					if err != nil {
						return nil, err
					}
					nodes = append(nodes, Node{Label: fmt.Sprintf("%v", name), Children: []Node{{Label: fmt.Sprintf("%v", val)}}})
				case parse.EnterHint:
					children, err := Parse(p)
					if err != nil {
						return nil, err
					}
					nodes = append(nodes, Node{Label: fmt.Sprintf("%v", name), Children: children})
				}
			}
		case parse.ValueHint:
			val, err := parse.GetValue(p)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, Node{Label: fmt.Sprintf("%v", val), Children: nil})
		case parse.LeaveHint:
			return nodes, nil
		}
	}
}
