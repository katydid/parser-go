// Copyright 2025 Walter Schulze
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parse

import (
	"math/rand"
	"testing"
)

func BenchmarkCompareFloat(b *testing.B) {
	f1 := float64(1)
	f2 := float64(22398234)
	f3 := float64(309908980)
	f1a := float64(1)
	f4 := rand.Float64()
	for b.Loop() {
		for i := 0; i < 250; i++ {
			if f1 == f2 {
				b.Fatal("want not equal")
			}
			if f1 == f3 {
				b.Fatal("want not equal")
			}
			if f3 == f2 {
				b.Fatal("want not equal")
			}
			if f1 != f1a {
				b.Fatal("want equal")
			}
			if f4 == f2 {
				b.Fatal("want not equal")
			}
		}
	}
}

func BenchmarkCompareInt(b *testing.B) {
	f1 := int64(1)
	f2 := int64(28234980230984)
	f3 := int64(32309980234)
	f1a := int64(1)
	f4 := int64(rand.Int())
	for b.Loop() {
		for i := 0; i < 250; i++ {
			if f1 == f2 {
				b.Fatal("want not equal")
			}
			if f1 == f3 {
				b.Fatal("want not equal")
			}
			if f3 == f2 {
				b.Fatal("want not equal")
			}
			if f1 != f1a {
				b.Fatal("want equal")
			}
			if f4 == f2 {
				b.Fatal("want not equal")
			}
		}
	}
}
