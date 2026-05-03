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

	"github.com/katydid/parser-go/parse"
)

// Walk walks through the whole parser in a top down manner.
func Walk(p parse.Parser) error {
	for {
		_, err := p.Next()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if _, _, err := p.Token(); err != nil {
			return err
		}
	}
	return nil
}

// RandomWalk does a random walk of the parser, given two probabilities.
//
//	The next parameter is passed to r.Intn and when a zero value is returned the Next method on the parser is called.
//	The skip parameter is passed to r.Intn and when a non zero value is returned the Skip method on the parser is called.
//
// RandomWalk is great for testing that the implemented parser can handle random skipping of parts of the tree.
func RandomWalk(p parse.Parser, r Rand, next, skip int) error {
	for {
		if r.Intn(next) == 0 {
			continue
		}
		_, err := p.Next()
		if err != nil && err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if _, _, err := p.Token(); err != nil {
			return err
		}
		if r.Intn(skip) == 0 {
			continue
		}
		err = p.Skip()
		if err != nil && err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}
