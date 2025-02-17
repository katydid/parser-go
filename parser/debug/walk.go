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
	"io"

	"github.com/katydid/parser-go/parser"
)

// Walk walks through the whole parser in a top down manner.
func Walk(p parser.Interface) error {
	for {
		if err := p.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		if err := WalkValue(p); err != nil {
			return err
		}
		if !p.IsLeaf() {
			p.Down()
			err := Walk(p)
			if err != nil {
				return err
			}
			p.Up()
		}
	}
	return nil
}

// WalkValue tries to parse the current value and returns an error if no value is returned.
func WalkValue(p parser.Value) error {
	var err error
	_, err = p.Bool()
	if err == nil {
		return nil
	}
	_, err = p.Int()
	if err == nil {
		return nil
	}
	_, err = p.Uint()
	if err == nil {
		return nil
	}
	_, err = p.Double()
	if err == nil {
		return nil
	}
	_, err = p.String()
	if err == nil {
		return nil
	}
	_, err = p.Bytes()
	if err == nil {
		return nil
	}
	return parser.ErrNotValue
}

// RandomWalk does a random walk of the parser, given two probabilities.
//
//	next is passed to r.Intn and when a zero value is returned the Next method on the parser is called.
//	down is passed to r.Intn and when a non zero value is returned the Down method on the parser is called.
//
// RandomWalk is great for testing that the implemented parser can handle random skipping of parts of the tree.
func RandomWalk(p parser.Interface, r Rand, next, down int) error {
	for {
		if r.Intn(next) == 0 {
			break
		}
		if err := p.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		if err := WalkValue(p); err != nil {
			return err
		}
		if !p.IsLeaf() {
			if r.Intn(down) != 0 {
				p.Down()
				if err := RandomWalk(p, r, next, down); err != nil {
					return err
				}
				p.Up()
			}
		}
	}
	return nil
}
