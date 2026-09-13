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
	"time"
)

// Node is a type that represents a node in a tree.
// It has a Label and Children list of Nodes (Hedge).
type Node struct {
	Label    Token `json:"Label"`
	Children Hedge `json:"Children,omitempty"`
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

// @deprecated Field is a helper function for creating a Node with a label and one child label.
// This is how a field with a value is typically represented.
func Field(name string, value string) Node {
	return NewStringNode(name, NewStringNode(value))
}

// @deprecated Nested is a helper function for creating a Node.
func Nested(name string, fs ...Node) Node {
	return NewStringNode(name, fs...)
}

func NewUnknownNode(children ...Node) Node {
	return Node{
		Label:    NewUnknownToken(),
		Children: children,
	}
}

func NewNullNode(children ...Node) Node {
	return Node{
		Label:    NewNullToken(),
		Children: children,
	}
}

func NewFalseNode(children ...Node) Node {
	return Node{
		Label:    NewFalseToken(),
		Children: children,
	}
}

func NewTrueNode(children ...Node) Node {
	return Node{
		Label:    NewTrueToken(),
		Children: children,
	}
}

func NewBytesNode(b []byte, children ...Node) Node {
	return Node{
		Label:    NewBytesToken(b),
		Children: children,
	}
}

func NewStringNode(s string, children ...Node) Node {
	return Node{
		Label:    NewStringToken(s),
		Children: children,
	}
}

func NewInt64Node(i int64, children ...Node) Node {
	return Node{
		Label:    NewInt64Token(i),
		Children: children,
	}
}

func NewFloat64Node(f float64, children ...Node) Node {
	return Node{
		Label:    NewFloat64Token(f),
		Children: children,
	}
}

func NewDecimalNode(d string, children ...Node) Node {
	return Node{
		Label:    NewDecimalToken(d),
		Children: children,
	}
}

func NewNanosecondsNode(n int64, children ...Node) Node {
	return Node{
		Label:    NewNanosecondsToken(n),
		Children: children,
	}
}

func NewDateTimeNode(v time.Time, children ...Node) Node {
	return Node{
		Label:    NewDateTimeToken(v),
		Children: children,
	}
}

func NewTagNode(v string, children ...Node) Node {
	return Node{
		Label:    NewTagToken(v),
		Children: children,
	}
}
