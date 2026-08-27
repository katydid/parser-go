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

package debug

import (
	"katydid.org.za/go/parser-go/cast"
	"katydid.org.za/go/parser-go/parse"
)

// These functions are only for debugging purposes, so does not need to be optimized with a pool.
func alloc(size int) []byte {
	return make([]byte, size)
}

type doubleToken struct {
	v float64
}

// NewDoubleToken wraps a native go type into a parse.Token.
func NewDoubleToken(v float64) parse.Token {
	return &doubleToken{v}
}

func (v *doubleToken) Token() (parse.Kind, []byte, error) {
	return parse.Float64Kind, cast.FromFloat64(v.v, alloc), nil
}

type intToken struct {
	v int64
}

// NewIntToken wraps a native go type into a parse.Token.
func NewIntToken(v int64) parse.Token {
	return &intToken{v}
}

func (v *intToken) Token() (parse.Kind, []byte, error) {
	return parse.Float64Kind, cast.FromInt64Ptr(&v.v, alloc), nil
}

type uintToken struct {
	v uint64
}

// NewUintToken wraps a native go type into a parse.Token.
func NewUintToken(v uint64) parse.Token {
	return &uintToken{v}
}

func (v *uintToken) Token() (parse.Kind, []byte, error) {
	return parse.Float64Kind, cast.FromUint64(v.v, alloc), nil
}

type boolToken struct {
	v bool
}

// NewBoolToken wraps a native go type into a parse.Token.
func NewBoolToken(v bool) parse.Token {
	return &boolToken{v}
}

func (v *boolToken) Token() (parse.Kind, []byte, error) {
	if v.v {
		return parse.TrueKind, nil, nil
	}
	return parse.FalseKind, nil, nil
}

type stringToken struct {
	v string
}

// NewStringToken wraps a native go type into a parse.Token.
func NewStringToken(v string) parse.Token {
	return &stringToken{v}
}

func (v *stringToken) Token() (parse.Kind, []byte, error) {
	return parse.StringKind, cast.FromString(v.v, alloc), nil
}

type bytesToken struct {
	v []byte
}

// NewBytesToken wraps a native go type into a parse.Token.
func NewBytesToken(v []byte) parse.Token {
	return &bytesToken{v}
}

func (v *bytesToken) Token() (parse.Kind, []byte, error) {
	return parse.BytesKind, v.v, nil
}
