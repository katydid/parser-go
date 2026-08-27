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

package hedge

import (
	"fmt"
	"strings"
)

// Node is a type that represents a node in a tree.
// It has a label an children nodes.
type Node struct {
	Label    Token
	Children Hedge
}

// String returns a string representation of Node.
func (n Node) String() string {
	if len(n.Children) == 0 {
		return n.Label.String()
	}
	return n.Label.String() + ":" + n.Children.String()
}

// Hedge is a list of Node.
type Hedge []Node

// Equal returns whether two Hedges are the same.
func (n Node) Equal(m Node) bool {
	if !n.Label.Equal(m.Label) {
		return false
	}
	if !n.Children.Equal(m.Children) {
		return false
	}
	return true
}

func (n Node) VerboseEqual(m Node) error {
	if !n.Label.Equal(m.Label) {
		return fmt.Errorf("%v != %v", n.Label, m.Label)
	}
	if len(n.Children) != len(m.Children) {
		return fmt.Errorf("%v: has different number of children", n.Label)
	}
	for i := range n.Children {
		if err := n.Children[i].VerboseEqual(m.Children[i]); err != nil {
			return fmt.Errorf("%v.%v", n.Label, err)
		}
	}
	return nil
}

// String returns a string representation of Nodes.
func (h Hedge) String() string {
	ss := make([]string, len(h))
	for i := range h {
		ss[i] = h[i].String()
	}
	return "{" + strings.Join(ss, ",") + "}"
}

// Equal returns whether two Node lists are equal.
func (h Hedge) Equal(g Hedge) bool {
	if len(h) != len(g) {
		return false
	}
	for i := range h {
		if !h[i].Equal(g[i]) {
			return false
		}
	}
	return true
}

func (h Hedge) VerboseEqual(g Hedge) error {
	if len(h) != len(g) {
		return fmt.Errorf("different number of nodes")
	}
	for i := range h {
		if err := h[i].VerboseEqual(g[i]); err != nil {
			return fmt.Errorf("%d.%v", i, err)
		}
	}
	return nil
}

// Field is a helper function for creating a Node with a label and one child label.
// This is how a field with a value is typically represented.
func Field(name string, value string) Node {
	return Node{
		Label: NewStringToken(name),
		Children: Hedge{
			Node{
				Label: NewStringToken(value),
			},
		},
	}
}

// Nested is a helper function for creating a Node.
func Nested(name string, fs ...Node) Node {
	return Node{
		Label:    NewStringToken(name),
		Children: Hedge(fs),
	}
}
