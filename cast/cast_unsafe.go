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
	"reflect"
	"unsafe"
)

func ToInt64Ptr(bs []byte, v *int64) {
	*v = *(*int64)(unsafe.Pointer(&bs[0]))
}

// FromInt64Ptr is very unsafe, you have make sure to keep the int64 in a value that won't be freed, until you are done using this slice.
func FromInt64Ptr(i *int64, _alloc func(size int) []byte) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  8,
		Cap:  8,
		Data: uintptr(unsafe.Pointer(i)),
	}))
}

func ToInt32Ptr(bs []byte, v *int32) {
	*v = *(*int32)(unsafe.Pointer(&bs[0]))
}

// FromInt32Ptr is very unsafe, you have make sure to keep the int32 in a value that won't be freed, until you are done using this slice.
func FromInt32Ptr(i *int32, _alloc func(size int) []byte) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  4,
		Cap:  4,
		Data: uintptr(unsafe.Pointer(i)),
	}))
}

func ToUint64Ptr(bs []byte, v *uint64) {
	*v = *(*uint64)(unsafe.Pointer(&bs[0]))
}

// FromUint64Ptr is very unsafe, you have make sure to keep the uint64 in a value that won't be freed, until you are done using this slice.
func FromUint64Ptr(i *uint64, _alloc func(size int) []byte) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  8,
		Cap:  8,
		Data: uintptr(unsafe.Pointer(i)),
	}))
}

func ToUint32Ptr(bs []byte, v *uint32) {
	*v = *(*uint32)(unsafe.Pointer(&bs[0]))
}

// FromUint32Ptr is very unsafe, you have make sure to keep the uint32 in a value that won't be freed, until you are done using this slice.
func FromUint32Ptr(i *uint32, _alloc func(size int) []byte) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  4,
		Cap:  4,
		Data: uintptr(unsafe.Pointer(i)),
	}))
}

func ToFloat64BitsPtr(bs []byte, v *uint64) {
	*v = *(*uint64)(unsafe.Pointer(&bs[0]))
}

// FromFloat64BitsPtr is very unsafe, you have make sure to keep the float64 in a value that won't be freed, until you are done using this slice.
func FromFloat64BitsPtr(u *uint64, _alloc func(size int) []byte) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  8,
		Cap:  8,
		Data: uintptr(unsafe.Pointer(u)),
	}))
}

func ToFloat32BitsPtr(bs []byte, v *uint32) {
	*v = *(*uint32)(unsafe.Pointer(&bs[0]))
}

// FromFloat32BitsPtr is very unsafe, you have make sure to keep the float32 in a value that won't be freed, until you are done using this slice.
func FromFloat32BitsPtr(v *uint32, _alloc func(size int) []byte) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  4,
		Cap:  4,
		Data: uintptr(unsafe.Pointer(v)),
	}))
}

// ToString uses unsafe to cast a byte slice to a string without copying or allocating memory.
func ToStringPtr(buf []byte, v *string) {
	*v = unsafe.String(unsafe.SliceData(buf), len(buf))
}

// FromStringPtr uses unsafe to cast from a string to a slice of byte.
func FromStringPtr(s *string, _alloc func(size int) []byte) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		int
	}{*s, len(*s)}))
}

func unsafeBetBytesClassic(s *string) []byte {
	return *(*[]byte)(unsafe.Pointer(s))
}

// https://stackoverflow.com/a/69231355
func unsafeGetBytesYenForYang(s *string) []byte {
	const MaxInt32 = 1<<31 - 1
	return (*[MaxInt32]byte)(unsafe.Pointer((*reflect.StringHeader)(
		unsafe.Pointer(s)).Data))[: len(*s)&MaxInt32 : len(*s)&MaxInt32]
}

// https://stackoverflow.com/questions/59209493/how-to-use-unsafe-get-a-byte-slice-from-a-string-without-memory-copy/69231355#comment130999637_69231355
func unsafeGetBytesNunoCruces(s *string) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		int
	}{*s, len(*s)}))
}
