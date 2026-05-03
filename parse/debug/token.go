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
	"github.com/katydid/parser-go/cast"
	"github.com/katydid/parser-go/parse"
)

// These functions are only for debugging purposes, so does not need to be optimized with a pool.
func alloc(size int) []byte {
	return make([]byte, size)
}

type doubleValue struct {
	v float64
}

// NewDoubleValue wraps a native go type into a parse.Token.
func NewDoubleValue(v float64) parse.Token {
	return &doubleValue{v}
}

func (v *doubleValue) Token() (parse.Kind, []byte, error) {
	return parse.Float64Kind, cast.FromFloat64(v.v, alloc), nil
}

type intValue struct {
	v int64
}

// NewIntValue wraps a native go type into a parse.Token.
func NewIntValue(v int64) parse.Token {
	return &intValue{v}
}

func (v *intValue) Token() (parse.Kind, []byte, error) {
	return parse.Float64Kind, cast.FromInt64(v.v, alloc), nil
}

type uintValue struct {
	v uint64
}

// NewUintValue wraps a native go type into a parse.Token.
func NewUintValue(v uint64) parse.Token {
	return &uintValue{v}
}

func (v *uintValue) Token() (parse.Kind, []byte, error) {
	return parse.Float64Kind, cast.FromUint64(v.v, alloc), nil
}

type boolValue struct {
	v bool
}

// NewBoolValue wraps a native go type into a parse.Token.
func NewBoolValue(v bool) parse.Token {
	return &boolValue{v}
}

func (v *boolValue) Token() (parse.Kind, []byte, error) {
	if v.v {
		return parse.TrueKind, nil, nil
	}
	return parse.FalseKind, nil, nil
}

type stringValue struct {
	v string
}

// NewStringValue wraps a native go type into a parse.Token.
func NewStringValue(v string) parse.Token {
	return &stringValue{v}
}

func (v *stringValue) Token() (parse.Kind, []byte, error) {
	return parse.StringKind, cast.FromString(v.v, alloc), nil
}

type bytesValue struct {
	v []byte
}

// NewBytesValue wraps a native go type into a parse.Token.
func NewBytesValue(v []byte) parse.Token {
	return &bytesValue{v}
}

func (v *bytesValue) Token() (parse.Kind, []byte, error) {
	return parse.BytesKind, v.v, nil
}
