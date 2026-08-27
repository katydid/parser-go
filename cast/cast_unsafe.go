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

//go:build !purego

package cast

import (
	"math"
	"reflect"
	"unsafe"
)

func ToInt64(bs []byte) int64 {
	return *(*int64)(unsafe.Pointer(&bs[0]))
}

// FromInt64 is very unsafe, you have make sure to keep the int64 in a value that won't be freed, until you are doing using this slice.
func FromInt64(i int64, _alloc func(size int) []byte) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  8,
		Cap:  8,
		Data: uintptr(unsafe.Pointer(&i)),
	}))
}

// FromInt64Ptr is very unsafe, you have make sure to keep the int64 in a value that won't be freed, until you are doing using this slice.
func FromInt64Ptr(i *int64, _alloc func(size int) []byte) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  8,
		Cap:  8,
		Data: uintptr(unsafe.Pointer(i)),
	}))
}

func ToInt32(bs []byte) int32 {
	return *(*int32)(unsafe.Pointer(&bs[0]))
}

// FromInt32 is very unsafe, you have make sure to keep the int32 in a value that won't be freed, until you are doing using this slice.
func FromInt32(i int32, _alloc func(size int) []byte) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  4,
		Cap:  4,
		Data: uintptr(unsafe.Pointer(&i)),
	}))
}

func ToUint64(bs []byte) uint64 {
	return *(*uint64)(unsafe.Pointer(&bs[0]))
}

// FromUint64 is very unsafe, you have make sure to keep the uint64 in a value that won't be freed, until you are doing using this slice.
func FromUint64(i uint64, _alloc func(size int) []byte) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  8,
		Cap:  8,
		Data: uintptr(unsafe.Pointer(&i)),
	}))
}

func ToUint32(bs []byte) uint32 {
	return *(*uint32)(unsafe.Pointer(&bs[0]))
}

// FromUint32 is very unsafe, you have make sure to keep the uint32 in a value that won't be freed, until you are doing using this slice.
func FromUint32(i uint32, _alloc func(size int) []byte) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  4,
		Cap:  4,
		Data: uintptr(unsafe.Pointer(&i)),
	}))
}

func ToFloat64(bs []byte) float64 {
	u := *(*uint64)(unsafe.Pointer(&bs[0]))
	return math.Float64frombits(u)
}

// FromFloat64 is very unsafe, you have make sure to keep the float64 in a value that won't be freed, until you are doing using this slice.
func FromFloat64(f float64, _alloc func(size int) []byte) []byte {
	u := math.Float64bits(f)
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  8,
		Cap:  8,
		Data: uintptr(unsafe.Pointer(&u)),
	}))
}

func FromFloat64BitsPtr(u *uint64, _alloc func(size int) []byte) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  8,
		Cap:  8,
		Data: uintptr(unsafe.Pointer(u)),
	}))
}

func ToFloat32(bs []byte) float32 {
	u := *(*uint32)(unsafe.Pointer(&bs[0]))
	return math.Float32frombits(u)
}

// FromFloat32 is very unsafe, you have make sure to keep the float32 in a value that won't be freed, until you are doing using this slice.
func FromFloat32(f float32, _alloc func(size int) []byte) []byte {
	u := math.Float32bits(f)
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  4,
		Cap:  4,
		Data: uintptr(unsafe.Pointer(&u)),
	}))
}

// ToString uses unsafe to cast a byte slice to a string without copying or allocating memory.
func ToString(buf []byte) string {
	return unsafe.String(unsafe.SliceData(buf), len(buf))
}

// FromString uses unsafe to cast from a string to a slice of byte.
func FromString(s string, _alloc func(size int) []byte) []byte {
	return unsafeBetBytesClassic(s)
}

func unsafeBetBytesClassic(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// https://stackoverflow.com/a/69231355
func unsafeGetBytesYenForYang(s string) []byte {
	const MaxInt32 = 1<<31 - 1
	return (*[MaxInt32]byte)(unsafe.Pointer((*reflect.StringHeader)(
		unsafe.Pointer(&s)).Data))[: len(s)&MaxInt32 : len(s)&MaxInt32]
}

// https://stackoverflow.com/questions/59209493/how-to-use-unsafe-get-a-byte-slice-from-a-string-without-memory-copy/69231355#comment130999637_69231355
func unsafeGetBytesNunoCruces(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		int
	}{s, len(s)}))
}
