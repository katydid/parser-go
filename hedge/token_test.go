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

func TestTokenJSON(t *testing.T) {
	tokens := []Token{
		NewUnknownToken(),
		NewNullToken(),
		NewFalseToken(),
		NewTrueToken(),
		NewBytesToken([]byte{0, 1, 2, 'a', 'b', 'c', 255}),
		NewBytesToken(nil),
		NewStringToken("abc"),
		NewStringToken(""),
		NewInt64Token(123),
		NewInt64Token(math.MaxInt64),
		NewInt64Token(math.MaxInt64),
		NewInt64Token(0),
		NewFloat64Token(123.456),
		NewFloat64Token(math.MaxFloat32),
		NewFloat64Token(math.MaxFloat64),
		NewFloat64Token(0),
		NewDecimalToken("123.456"),
		NewNanosecondsToken(time.Now().UnixNano()),
		NewDateTimeToken(time.Now()),
		NewTagToken("array"),
	}
	for _, want := range tokens {
		t.Run(want.String(), func(t *testing.T) {
			data, err := json.Marshal(want)
			if err != nil {
				t.Fatal(err)
			}
			got := Token{}
			if err := json.Unmarshal(data, &got); err != nil {
				t.Fatal(err)
			}
			if !got.Equal(want) {
				t.Fatalf("want %v got %v", want, got)
			}
		})
	}
}
