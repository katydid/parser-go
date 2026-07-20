//  Copyright 2025 Walter Schulze
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

package cast

import (
	"math"
	"testing"
)

func TestCastInt64(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := int64(123)
	bs := FromInt64(want, alloc)
	got := ToInt64(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastMaxInt64(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := int64(math.MaxInt64)
	bs := FromInt64(want, alloc)
	got := ToInt64(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastMinInt64(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := int64(math.MinInt64)
	bs := FromInt64(want, alloc)
	got := ToInt64(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastUint64(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := uint64(123)
	bs := FromUint64(want, alloc)
	got := ToUint64(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastMaxUint64(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := uint64(math.MaxUint64)
	bs := FromUint64(want, alloc)
	got := ToUint64(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastFloat64(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := float64(123)
	bs := FromFloat64(want, alloc)
	got := ToFloat64(bs)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCastMaxFloat64(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := float64(math.MaxFloat64)
	bs := FromFloat64(want, alloc)
	got := ToFloat64(bs)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCastSmallestNonzeroFloat64(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := float64(math.SmallestNonzeroFloat64)
	bs := FromFloat64(want, alloc)
	got := ToFloat64(bs)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCastString(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := "abc"
	bs := FromString(want, alloc)
	got := ToString(bs)
	if got != want {
		t.Fatalf("want %s got %s", want, got)
	}
}
