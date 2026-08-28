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

func TestCastInt64Ptr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := int64(123)
	bs := FromInt64Ptr(&want, alloc)
	var got int64
	ToInt64Ptr(bs, &got)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastMaxInt64Ptr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := int64(math.MaxInt64)
	bs := FromInt64Ptr(&want, alloc)
	var got int64
	ToInt64Ptr(bs, &got)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastMinInt64Ptr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := int64(math.MinInt64)
	bs := FromInt64Ptr(&want, alloc)
	var got int64
	ToInt64Ptr(bs, &got)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastInt32Ptr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := int32(123)
	bs := FromInt32Ptr(&want, alloc)
	var got int32
	ToInt32Ptr(bs, &got)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastMaxInt32Ptr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := int32(math.MaxInt32)
	bs := FromInt32Ptr(&want, alloc)
	var got int32
	ToInt32Ptr(bs, &got)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastMinInt32Ptr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := int32(math.MinInt32)
	bs := FromInt32Ptr(&want, alloc)
	var got int32
	ToInt32Ptr(bs, &got)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastUint64Ptr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := uint64(123)
	bs := FromUint64Ptr(&want, alloc)
	var got uint64
	ToUint64Ptr(bs, &got)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastMaxUint64Ptr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := uint64(math.MaxUint64)
	bs := FromUint64Ptr(&want, alloc)
	var got uint64
	ToUint64Ptr(bs, &got)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastUint32Ptr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := uint32(123)
	bs := FromUint32Ptr(&want, alloc)
	var got uint32
	ToUint32Ptr(bs, &got)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastMaxUint32Ptr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := uint32(math.MaxUint32)
	bs := FromUint32Ptr(&want, alloc)
	var got uint32
	ToUint32Ptr(bs, &got)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCastFloat64BitsPtr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := float64(123)
	u := math.Float64bits(want)
	bs := FromFloat64BitsPtr(&u, alloc)
	var gotu uint64
	ToFloat64BitsPtr(bs, &gotu)
	got := math.Float64frombits(gotu)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCastMaxFloat64BitsPtr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := float64(math.MaxFloat64)
	u := math.Float64bits(want)
	bs := FromFloat64BitsPtr(&u, alloc)
	var gotu uint64
	ToFloat64BitsPtr(bs, &gotu)
	got := math.Float64frombits(gotu)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCastSmallestNonzeroFloat64BitsPtr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := float64(math.SmallestNonzeroFloat64)
	u := math.Float64bits(want)
	bs := FromFloat64BitsPtr(&u, alloc)
	var gotu uint64
	ToFloat64BitsPtr(bs, &gotu)
	got := math.Float64frombits(gotu)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCastFloat32BitsPtr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := float32(123)
	u := math.Float32bits(want)
	bs := FromFloat32BitsPtr(&u, alloc)
	var gotu uint32
	ToFloat32BitsPtr(bs, &gotu)
	got := math.Float32frombits(gotu)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCastMaxFloat32BitsPtr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := float32(math.MaxFloat32)
	u := math.Float32bits(want)
	bs := FromFloat32BitsPtr(&u, alloc)
	var gotu uint32
	ToFloat32BitsPtr(bs, &gotu)
	got := math.Float32frombits(gotu)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCastSmallestNonzeroFloat32BitsPtr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := float32(math.SmallestNonzeroFloat32)
	u := math.Float32bits(want)
	bs := FromFloat32BitsPtr(&u, alloc)
	var gotu uint32
	ToFloat32BitsPtr(bs, &gotu)
	got := math.Float32frombits(gotu)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCastStringPtr(t *testing.T) {
	alloc := func(size int) []byte { return make([]byte, size) }
	want := "abc"
	bs := FromStringPtr(&want, alloc)
	var got string
	ToStringPtr(bs, &got)
	if got != want {
		t.Fatalf("want %s got %s", want, got)
	}
}
