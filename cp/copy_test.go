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

package cp

import (
	"math"
	"testing"
)

func TestCopyInt64(t *testing.T) {
	want := int64(123)
	bs := FromInt64(want)
	got := ToInt64(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCopyMaxInt64(t *testing.T) {
	want := int64(math.MaxInt64)
	bs := FromInt64(want)
	got := ToInt64(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCopyMinInt64(t *testing.T) {
	want := int64(math.MinInt64)
	bs := FromInt64(want)
	got := ToInt64(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCopyInt32(t *testing.T) {
	want := int32(123)
	bs := FromInt32(want)
	got := ToInt32(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCopyMaxInt32(t *testing.T) {
	want := int32(math.MaxInt32)
	bs := FromInt32(want)
	got := ToInt32(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCopyMinInt32(t *testing.T) {
	want := int32(math.MinInt32)
	bs := FromInt32(want)
	got := ToInt32(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCopyUint64(t *testing.T) {
	want := uint64(123)
	bs := FromUint64(want)
	got := ToUint64(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCopyMaxUint64(t *testing.T) {
	want := uint64(math.MaxUint64)
	bs := FromUint64(want)
	got := ToUint64(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCopyUint32(t *testing.T) {
	want := uint32(123)
	bs := FromUint32(want)
	got := ToUint32(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCopyMaxUint32(t *testing.T) {
	want := uint32(math.MaxUint32)
	bs := FromUint32(want)
	got := ToUint32(bs)
	if got != want {
		t.Fatalf("want %d got %d", want, got)
	}
}

func TestCopyFloat64(t *testing.T) {
	want := float64(123)
	bs := FromFloat64(want)
	got := ToFloat64(bs)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCopyMaxFloat64(t *testing.T) {
	want := float64(math.MaxFloat64)
	bs := FromFloat64(want)
	got := ToFloat64(bs)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCopySmallestNonzeroFloat64(t *testing.T) {
	want := float64(math.SmallestNonzeroFloat64)
	bs := FromFloat64(want)
	got := ToFloat64(bs)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCopyFloat32(t *testing.T) {
	want := float32(123)
	bs := FromFloat32(want)
	got := ToFloat32(bs)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCopyMaxFloat32(t *testing.T) {
	want := float32(math.MaxFloat32)
	bs := FromFloat32(want)
	got := ToFloat32(bs)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCopySmallestNonzeroFloat32(t *testing.T) {
	want := float32(math.SmallestNonzeroFloat32)
	bs := FromFloat32(want)
	got := ToFloat32(bs)
	if got != want {
		t.Fatalf("want %f got %f", want, got)
	}
}

func TestCopyString(t *testing.T) {
	want := "abc"
	bs := FromString(want)
	got := ToString(bs)
	if got != want {
		t.Fatalf("want %s got %s", want, got)
	}
}
