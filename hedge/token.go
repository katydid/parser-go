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

package hedge

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"

	"katydid.org.za/go/parser-go/cast"
	"katydid.org.za/go/parser-go/parse"
)

type Token struct {
	kind parse.Kind
	i    int64
	b    []byte
	s    string
	u    uint64
}

func (t Token) Equal(u Token) bool {
	if t.kind != u.kind {
		return false
	}
	switch t.kind {
	case parse.UnknownKind:
		return true
	case parse.NullKind:
		return true
	case parse.FalseKind:
		return true
	case parse.TrueKind:
		return true
	case parse.BytesKind:
		return bytes.Equal(t.b, u.b)
	case parse.StringKind:
		return t.s == u.s
	case parse.Int64Kind:
		return t.i == u.i
	case parse.Float64Kind:
		return t.u == u.u
	case parse.DecimalKind:
		return t.s == u.s
	case parse.NanosecondsKind:
		return t.i == u.i
	case parse.DateTimeKind:
		return t.s == u.s
	case parse.TagKind:
		return t.s == u.s
	}
	panic("unreachable")
}

func (t Token) String() string {
	switch t.kind {
	case parse.UnknownKind:
		return "<unknown kind>"
	case parse.NullKind:
		return fmt.Sprintf("%v", nil)
	case parse.FalseKind:
		return fmt.Sprintf("%v", false)
	case parse.TrueKind:
		return fmt.Sprintf("%v", true)
	case parse.BytesKind:
		return base64.StdEncoding.EncodeToString(t.b)
	case parse.StringKind:
		return fmt.Sprintf("%v", t.s)
	case parse.Int64Kind:
		return fmt.Sprintf("%v", t.i)
	case parse.Float64Kind:
		return fmt.Sprintf("%v", math.Float64frombits(t.u))
	case parse.DecimalKind:
		return fmt.Sprintf("%v", t.s)
	case parse.NanosecondsKind:
		return fmt.Sprintf("%v", t.i)
	case parse.DateTimeKind:
		return fmt.Sprintf("%v", t.s)
	case parse.TagKind:
		return fmt.Sprintf("%v", t.s)
	}
	panic("unreachable")
}

func TokenString(kind parse.Kind, bs []byte) string {
	switch kind {
	case parse.UnknownKind:
		return "<unknown kind>"
	case parse.NullKind:
		return fmt.Sprintf("%v", nil)
	case parse.FalseKind:
		return fmt.Sprintf("%v", false)
	case parse.TrueKind:
		return fmt.Sprintf("%v", true)
	case parse.BytesKind:
		return base64.StdEncoding.EncodeToString(bs)
	case parse.StringKind:
		return fmt.Sprintf("%v", string(bs))
	case parse.Int64Kind:
		return fmt.Sprintf("%v", int64(binary.LittleEndian.Uint64(bs)))
	case parse.Float64Kind:
		return fmt.Sprintf("%v", math.Float64frombits(binary.LittleEndian.Uint64(bs)))
	case parse.DecimalKind:
		return fmt.Sprintf("%v", string(bs))
	case parse.NanosecondsKind:
		return fmt.Sprintf("%v", int64(binary.LittleEndian.Uint64(bs)))
	case parse.DateTimeKind:
		return fmt.Sprintf("%v", string(bs))
	case parse.TagKind:
		return fmt.Sprintf("%v", string(bs))
	}
	panic("unreachable")
}

type jsonToken struct {
	Kind  parse.Kind
	Value string `json:",omitempty"`
}

func (t Token) MarshalJSON() ([]byte, error) {
	j := &jsonToken{
		Kind: t.kind,
	}
	switch t.kind {
	case parse.UnknownKind:
	case parse.NullKind:
	case parse.FalseKind:
	case parse.TrueKind:
	case parse.BytesKind:
		j.Value = base64.StdEncoding.EncodeToString(t.b)
	case parse.StringKind:
		j.Value = t.s
	case parse.Int64Kind:
		j.Value = fmt.Sprintf("%d", t.i)
	case parse.Float64Kind:
		j.Value = fmt.Sprintf("%f", math.Float64frombits(t.u))
	case parse.DecimalKind:
		j.Value = t.s
	case parse.NanosecondsKind:
		j.Value = fmt.Sprintf("%d", t.i)
	case parse.DateTimeKind:
		j.Value = t.s
	case parse.TagKind:
		j.Value = t.s
	default:
		return nil, fmt.Errorf("unknown kind %v", j.Kind)
	}
	return json.Marshal(j)
}

func (t *Token) UnmarshalJSON(data []byte) error {
	j := &jsonToken{}
	if err := json.Unmarshal(data, j); err != nil {
		return err
	}
	t.kind = j.Kind
	switch j.Kind {
	case parse.UnknownKind:
	case parse.NullKind:
	case parse.FalseKind:
	case parse.TrueKind:
	case parse.BytesKind:
		bs, err := base64.StdEncoding.DecodeString(j.Value)
		if err != nil {
			return err
		}
		t.b = bs
	case parse.StringKind:
		t.s = j.Value
	case parse.Int64Kind:
		i, err := strconv.ParseInt(j.Value, 10, 64)
		if err != nil {
			return err
		}
		t.i = i
	case parse.Float64Kind:
		f, err := strconv.ParseFloat(j.Value, 64)
		if err != nil {
			return err
		}
		t.u = math.Float64bits(f)
	case parse.DecimalKind:
		t.s = j.Value
	case parse.NanosecondsKind:
		i, err := strconv.ParseInt(j.Value, 10, 64)
		if err != nil {
			return err
		}
		t.i = i
	case parse.DateTimeKind:
		t.s = j.Value
	case parse.TagKind:
		t.s = j.Value
	default:
		return fmt.Errorf("unknown kind %v", j.Kind)
	}

	return nil
}

func NewToken(kind parse.Kind, b []byte, err error) (Token, error) {
	if err != nil {
		return NewUnknownToken(), err
	}
	t := &Token{kind: kind, b: b}
	switch kind {
	case parse.UnknownKind:
	case parse.NullKind:
	case parse.FalseKind:
	case parse.TrueKind:
	case parse.BytesKind:
	case parse.StringKind:
		cast.ToStringPtr(t.b, &t.s)
	case parse.Int64Kind:
		cast.ToInt64Ptr(t.b, &t.i)
	case parse.Float64Kind:
		cast.ToFloat64BitsPtr(t.b, &t.u)
	case parse.DecimalKind:
		cast.ToStringPtr(t.b, &t.s)
	case parse.NanosecondsKind:
		cast.ToInt64Ptr(t.b, &t.i)
	case parse.DateTimeKind:
		cast.ToStringPtr(t.b, &t.s)
	case parse.TagKind:
		cast.ToStringPtr(t.b, &t.s)
	default:
		panic("unreachable")
	}
	return *t, nil
}

func (t *Token) Token(alloc func(size int) []byte) (parse.Kind, []byte, error) {
	switch t.kind {
	case parse.UnknownKind:
		return parse.UnknownKind, nil, nil
	case parse.NullKind:
		return parse.NullKind, nil, nil
	case parse.FalseKind:
		return parse.FalseKind, nil, nil
	case parse.TrueKind:
		return parse.TrueKind, nil, nil
	case parse.BytesKind:
		return parse.BytesKind, t.b, nil
	case parse.StringKind:
		return parse.StringKind, t.b, nil
	case parse.Int64Kind:
		return parse.Int64Kind, t.b, nil
	case parse.Float64Kind:
		return parse.Float64Kind, t.b, nil
	case parse.DecimalKind:
		return parse.DecimalKind, t.b, nil
	case parse.NanosecondsKind:
		return parse.NanosecondsKind, t.b, nil
	case parse.DateTimeKind:
		return parse.DateTimeKind, t.b, nil
	case parse.TagKind:
		return parse.TagKind, t.b, nil
	}
	panic("unreachable")
}

func NewUnknownToken() Token {
	return Token{
		kind: parse.UnknownKind,
	}
}

func NewNullToken() Token {
	return Token{
		kind: parse.NullKind,
	}
}

func NewFalseToken() Token {
	return Token{
		kind: parse.FalseKind,
	}
}

func NewTrueToken() Token {
	return Token{
		kind: parse.TrueKind,
	}
}

func NewBytesToken(b []byte) Token {
	return Token{
		kind: parse.BytesKind,
		b:    bytes.Clone(b),
	}
}

func NewStringToken(s string) Token {
	t := &Token{
		kind: parse.StringKind,
		s:    s,
	}
	t.b = []byte(t.s)
	return *t
}

func NewInt64Token(i int64) Token {
	t := &Token{
		kind: parse.Int64Kind,
		i:    i,
	}
	t.b = make([]byte, 8)
	binary.LittleEndian.PutUint64(t.b, uint64(t.i))
	return *t
}

func NewFloat64Token(f float64) Token {
	t := &Token{
		kind: parse.Float64Kind,
		u:    math.Float64bits(f),
	}
	t.b = make([]byte, 8)
	binary.LittleEndian.PutUint64(t.b, t.u)
	return *t
}

func NewDecimalToken(d string) Token {
	t := &Token{
		kind: parse.DecimalKind,
		s:    d,
	}
	t.b = []byte(t.s)
	return *t
}

func NewNanosecondsToken(n int64) Token {
	t := &Token{
		kind: parse.NanosecondsKind,
		i:    n,
	}
	t.b = make([]byte, 8)
	binary.LittleEndian.PutUint64(t.b, uint64(t.i))
	return *t
}

func NewDateTimeToken(v time.Time) Token {
	t := &Token{
		kind: parse.DateTimeKind,
		s:    v.Format(time.RFC3339Nano),
	}
	t.b = []byte(t.s)
	return *t
}

func NewTagToken(v string) Token {
	t := &Token{
		kind: parse.TagKind,
		s:    v,
	}
	t.b = []byte(t.s)
	return *t
}
