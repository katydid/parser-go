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

//go:build !purego

package cast

import "testing"

// The deprecated unsafe version of FromInt64 never allocates.
func TestAllocCastInt64(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	f := func() {
		want := int64(1233456578)
		bs := FromInt64(want, alloc)
		got := ToInt64(bs)
		if got != want {
			t.Fatalf("want %d got %d", want, got)
		}
	}
	for i := 0; i < 10000; i++ {
		allocs := testing.AllocsPerRun(1, f)
		if allocs > 0 {
			t.Fatalf("Cast Allocs = %f", allocs)
		}
	}
}

// The deprecated unsafe version of FromUint64 never allocates.
func TestAllocCastUint64(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	f := func() {
		want := uint64(1233456578)
		bs := FromUint64(want, alloc)
		got := ToUint64(bs)
		if got != want {
			t.Fatalf("want %d got %d", want, got)
		}
	}
	for i := 0; i < 10000; i++ {
		allocs := testing.AllocsPerRun(1, f)
		if allocs > 0 {
			t.Fatalf("Cast Allocs = %f", allocs)
		}
	}
}

// The deprecated unsafe version of FromFloat64 never allocates.
func TestAllocCastFloat64(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	f := func() {
		want := float64(1233456578)
		bs := FromFloat64(want, alloc)
		got := ToFloat64(bs)
		if got != want {
			t.Fatalf("want %f got %f", want, got)
		}
	}
	for i := 0; i < 10000; i++ {
		allocs := testing.AllocsPerRun(1, f)
		if allocs > 0 {
			t.Fatalf("Cast Allocs = %f", allocs)
		}
	}
}
