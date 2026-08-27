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
	"encoding/binary"
	"math"
)

func ToInt64(bs []byte) int64 {
	res := make([]byte, len(bs))
	copy(res, bs)
	return int64(binary.LittleEndian.Uint64(res))
}

func FromInt64(i int64) []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(i))
	return bs
}

func ToInt32(bs []byte) int32 {
	res := make([]byte, len(bs))
	copy(res, bs)
	return int32(binary.LittleEndian.Uint32(res))
}

func FromInt32(i int32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(i))
	return bs
}

func ToUint64(bs []byte) uint64 {
	res := make([]byte, len(bs))
	copy(res, bs)
	return binary.LittleEndian.Uint64(res)
}

func FromUint64(i uint64) []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(i))
	return bs
}

func ToUint32(bs []byte) uint32 {
	res := make([]byte, len(bs))
	copy(res, bs)
	return binary.LittleEndian.Uint32(res)
}

func FromUint32(i uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(i))
	return bs
}

func ToFloat64(bs []byte) float64 {
	res := make([]byte, len(bs))
	copy(res, bs)
	u := binary.LittleEndian.Uint64(res)
	return math.Float64frombits(u)
}

func FromFloat64(f float64) []byte {
	bs := make([]byte, 8)
	u := math.Float64bits(f)
	binary.LittleEndian.PutUint64(bs, u)
	return bs
}

func ToFloat32(bs []byte) float32 {
	res := make([]byte, len(bs))
	copy(res, bs)
	u := binary.LittleEndian.Uint32(res)
	return math.Float32frombits(u)
}

func FromFloat32(f float32) []byte {
	bs := make([]byte, 4)
	u := math.Float32bits(f)
	binary.LittleEndian.PutUint32(bs, u)
	return bs
}

func ToString(buf []byte) string {
	return string(buf)
}

func FromString(s string) []byte {
	bs := make([]byte, len(s))
	copy(bs, s)
	return bs
}

func Bytes(buf []byte) []byte {
	res := make([]byte, len(buf))
	copy(res, buf)
	return res
}
