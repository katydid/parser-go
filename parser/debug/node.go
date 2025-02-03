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

import "strings"

// Node is a type that represents a node in a tree.
// It has a label an children nodes.
type Node struct {
	Label    string
	Children Nodes
}

// String returns a string representation of Node.
func (n Node) String() string {
	if len(n.Children) == 0 {
		return n.Label
	}
	return n.Label + ":" + n.Children.String()
}

// Equal returns whether two Nodes are the same.
func (n Node) Equal(m Node) bool {
	if n.Label != m.Label {
		return false
	}
	if !n.Children.Equal(m.Children) {
		return false
	}
	return true
}

// Nodes is a list of Node.
type Nodes []Node

// String returns a string representation of Nodes.
func (n Nodes) String() string {
	ss := make([]string, len(n))
	for i := range n {
		ss[i] = n[i].String()
	}
	return "{" + strings.Join(ss, ",") + "}"
}

// Equal returns whether two Node lists are equal.
func (n Nodes) Equal(m Nodes) bool {
	if len(n) != len(m) {
		return false
	}
	for i := range n {
		if !n[i].Equal(m[i]) {
			return false
		}
	}
	return true
}

// Field is a helper function for creating a Node with a label and one child label.
// This is how a field with a value is typically represented.
func Field(name string, value string) Node {
	return Node{
		Label: name,
		Children: Nodes{
			Node{
				Label: value,
			},
		},
	}
}

// Nested is a helper function for creating a Node.
func Nested(name string, fs ...Node) Node {
	return Node{
		Label:    name,
		Children: Nodes(fs),
	}
}
