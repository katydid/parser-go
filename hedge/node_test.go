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
	"encoding/json"
	"math"
	"testing"
	"time"
)

func TestNodeJSON(t *testing.T) {
	nodes := []Node{
		// nodes without children
		NewUnknownNode(),
		NewNullNode(),
		NewFalseNode(),
		NewTrueNode(),
		NewBytesNode([]byte{0, 1, 2, 'a', 'b', 'c', 255}),
		NewBytesNode(nil),
		NewStringNode("abc"),
		NewStringNode(""),
		NewInt64Node(123),
		NewInt64Node(math.MaxInt64),
		NewInt64Node(math.MaxInt64),
		NewInt64Node(0),
		NewFloat64Node(123.456),
		NewFloat64Node(math.MaxFloat32),
		NewFloat64Node(math.MaxFloat64),
		NewFloat64Node(0),
		NewDecimalNode("123.456"),
		NewNanosecondsNode(time.Now().UnixNano()),
		NewDateTimeNode(time.Now()),
		NewTagNode("array"),

		// nodes with children
		NewUnknownNode(NewUnknownNode()),
		NewNullNode(NewNullNode()),
		NewFalseNode(NewFalseNode()),
		NewTrueNode(NewTrueNode()),
		NewBytesNode([]byte{0, 1, 2, 'a', 'b', 'c', 255}, NewBytesNode([]byte{0, 1, 2, 'a', 'b', 'c', 255}), NewBytesNode([]byte{0, 1, 2, 'a', 'b', 'c', 255})),
		NewBytesNode(nil, NewBytesNode(nil)),
		NewStringNode("abc", NewStringNode("abc")),
		NewStringNode("", NewStringNode("")),
		NewInt64Node(123, NewInt64Node(123)),
		NewInt64Node(math.MaxInt64, NewInt64Node(math.MaxInt64), NewInt64Node(math.MaxInt64), NewInt64Node(math.MaxInt64), NewInt64Node(math.MaxInt64)),
		NewInt64Node(math.MaxInt64, NewInt64Node(math.MaxInt64)),
		NewInt64Node(0, NewInt64Node(0)),
		NewFloat64Node(123.456, NewFloat64Node(123.456)),
		NewFloat64Node(math.MaxFloat32, NewFloat64Node(math.MaxFloat32)),
		NewFloat64Node(math.MaxFloat64, NewFloat64Node(math.MaxFloat64)),
		NewFloat64Node(0, NewFloat64Node(0)),
		NewDecimalNode("123.456", NewDecimalNode("123.456")),
		NewNanosecondsNode(time.Now().UnixNano(), NewNanosecondsNode(time.Now().UnixNano())),
		NewDateTimeNode(time.Now(), NewDateTimeNode(time.Now())),
		NewTagNode("array", NewTagNode("array")),
	}
	for _, want := range nodes {
		t.Run(want.String(), func(t *testing.T) {
			data, err := json.Marshal(want)
			if err != nil {
				t.Fatal(err)
			}
			got := Node{}
			if err := json.Unmarshal(data, &got); err != nil {
				t.Fatal(err)
			}
			if !got.Equal(want) {
				t.Fatalf("want %v got %v", want, got)
			}
		})
	}
}
