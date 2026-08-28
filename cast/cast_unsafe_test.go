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

import (
	"math"
	"testing"
)

func TestAllocCastInt64Ptr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	f := func() {
		want := int64(1233456578)
		bs := FromInt64Ptr(&want, alloc)
		var got int64
		ToInt64Ptr(bs, &got)
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

func TestAllocCastInt32Ptr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	f := func() {
		want := int32(1233456578)
		bs := FromInt32Ptr(&want, alloc)
		var got int32
		ToInt32Ptr(bs, &got)
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

func TestAllocCastUint64(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	f := func() {
		want := uint64(1233456578)
		bs := FromUint64Ptr(&want, alloc)
		var got uint64
		ToUint64Ptr(bs, &got)
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

// The deprecated unsafe version of FromUint32 never allocates.
func TestAllocCastUint32(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	f := func() {
		want := uint32(1233456578)
		bs := FromUint32Ptr(&want, alloc)
		var got uint32
		ToUint32Ptr(bs, &got)
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

func TestAllocCastFloat64BitsPtr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	f := func() {
		want := float64(1233456578)
		u := math.Float64bits(want)
		bs := FromFloat64BitsPtr(&u, alloc)
		var gotu uint64
		ToFloat64BitsPtr(bs, &gotu)
		got := math.Float64frombits(gotu)
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

func TestAllocCastFloat32Ptr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	f := func() {
		want := float32(1233456578)
		u := math.Float32bits(want)
		bs := FromFloat32BitsPtr(&u, alloc)
		var gotu uint32
		ToFloat32BitsPtr(bs, &gotu)
		got := math.Float32frombits(gotu)
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
