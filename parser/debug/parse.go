//  Copyright 2015 Walter Schulze
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

	"github.com/katydid/parser-go/parser"
)

// Parse parses through the whole parser in a top down manner and records the values into a Nodes structute.
func Parse(p parser.Interface) (Nodes, error) {
	nodes := make(Nodes, 0)
	for {
		if err := p.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		value, err := parser.GetValue(p)
		if err != nil {
			return nil, err
		}
		if p.IsLeaf() {
			nodes = append(nodes, Node{fmt.Sprintf("%v", value), nil})
		} else {
			name := fmt.Sprintf("%v", value)
			p.Down()
			v, err := Parse(p)
			if err != nil {
				return nil, err
			}
			p.Up()
			nodes = append(nodes, Node{name, v})
		}
	}
	return nodes, nil
}

// RandomParse does a random parse of the parser, given two probabilities.
//
//	next is passed to r.Intn and when a zero value is returned the Next method on the parser is called.
//	down is passed to r.Intn and when a non zero value is returned the Down method on the parser is called.
//
// RandomParse is great for testing that the implemented parser can handle random skipping of parts of the tree.
func RandomParse(p parser.Interface, r Rand, next, down int) (Nodes, error) {
	nodes := make(Nodes, 0)
	for {
		if err := p.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		if r.Intn(next) == 0 {
			break
		}
		value, err := parser.GetValue(p)
		if err != nil {
			return nil, err
		}
		if p.IsLeaf() {
			nodes = append(nodes, Node{fmt.Sprintf("%#v", value), nil})
		} else {
			name := fmt.Sprintf("%#v", value)
			var v Nodes
			if r.Intn(down) != 0 {
				p.Down()
				v, err = RandomParse(p, r, next, down)
				if err != nil {
					return nil, err
				}
				p.Up()
			}
			nodes = append(nodes, Node{name, v})
		}
	}
	return nodes, nil
}
