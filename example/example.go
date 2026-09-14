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

package example

import (
	. "katydid.org.za/go/parser-go/hedge"
)

type Example struct {
	A int64
	B []string   `json:"B,omitempty"`
	C *Example   `json:"C,omitempty"`
	D *int32     `json:"D,omitempty"`
	E []*Example `json:"E,omitempty"`
	F []uint32   `json:"F,omitempty"`
	G *float64   `json:"G,omitempty"`
}

// Input is a sample instance of the Example struct.
var Input = &Example{
	A: int64(1),
	B: []string{"b2", "b3"},
	C: &Example{
		A: int64(2),
		D: ptr(int32(3)),
		E: []*Example{
			{
				B: []string{"b4"},
			},
			{
				B: []string{"b5"},
			},
		},
	},
	D: ptr(int32(4)),
	F: []uint32{5},
}

func ptr[A any](a A) *A {
	return &a
}

// Output is a sample instance of Hedge that repesents the Input variable after it has been parsed by Walk.
var Output = Hedge{
	Field(`A`, `1`),
	Nested(`B`,
		Field(`0`, `b2`),
		Field(`1`, `b3`),
	),
	Nested(`C`,
		Field(`A`, `2`),
		Field(`D`, `3`),
		Nested(`E`,
			Nested(`0`,
				Field(`A`, `0`),
				Nested(`B`,
					Field(`0`, `b4`),
				),
			),
			Nested(`1`,
				Field(`A`, `0`),
				Nested(`B`,
					Field(`0`, `b5`),
				),
			),
		),
	),
	Field(`D`, `4`),
	Nested(`F`,
		Field(`0`, `5`),
	),
}
