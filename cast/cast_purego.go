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

//go:build purego

package cast

import (
	"encoding/binary"
)

func ToInt64Ptr(bs []byte, v *int64) {
	*v = int64(binary.LittleEndian.Uint64(bs))
}

func FromInt64Ptr(i *int64, alloc func(size int) []byte) []byte {
	bs := alloc(8)
	binary.LittleEndian.PutUint64(bs, uint64(*i))
	return bs
}

func ToInt32Ptr(bs []byte, v *int32) {
	*v = int32(binary.LittleEndian.Uint32(bs))
}

func FromInt32Ptr(i *int32, alloc func(size int) []byte) []byte {
	bs := alloc(4)
	binary.LittleEndian.PutUint32(bs, uint32(*i))
	return bs
}

func ToUint64Ptr(bs []byte, v *uint64) {
	*v = binary.LittleEndian.Uint64(bs)
}

func FromUint64Ptr(i *uint64, alloc func(size int) []byte) []byte {
	bs := alloc(8)
	binary.LittleEndian.PutUint64(bs, uint64(*i))
	return bs
}

func ToUint32Ptr(bs []byte, v *uint32) {
	*v = binary.LittleEndian.Uint32(bs)
}

func FromUint32Ptr(i *uint32, alloc func(size int) []byte) []byte {
	bs := alloc(4)
	binary.LittleEndian.PutUint32(bs, uint32(*i))
	return bs
}

func ToFloat64BitsPtr(bs []byte, v *uint64) {
	*v = binary.LittleEndian.Uint64(bs)
}

func FromFloat64BitsPtr(u *uint64, alloc func(size int) []byte) []byte {
	bs := alloc(8)
	binary.LittleEndian.PutUint64(bs, *u)
	return bs
}

func ToFloat32BitsPtr(bs []byte, v *uint32) {
	*v = binary.LittleEndian.Uint32(bs)
}

func FromFloat32BitsPtr(u *uint32, alloc func(size int) []byte) []byte {
	bs := alloc(4)
	binary.LittleEndian.PutUint32(bs, *u)
	return bs
}

func ToStringPtr(buf []byte, v *string) {
	*v = string(buf)
}

func FromStringPtr(s *string, alloc func(size int) []byte) []byte {
	bs := alloc(len(*s))
	copy(bs, *s)
	return bs
}
