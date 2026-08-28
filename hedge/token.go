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
	"fmt"
	"math"
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
	case parse.UnknownKind:
		return true
	}
	panic("unreachable")
}

func (t Token) String() string {
	switch t.kind {
	case parse.NullKind:
		return fmt.Sprintf("%v", nil)
	case parse.FalseKind:
		return fmt.Sprintf("%v", false)
	case parse.TrueKind:
		return fmt.Sprintf("%v", true)
	case parse.BytesKind:
		return fmt.Sprintf("%v", t.b)
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
	case parse.UnknownKind:
		return "<unknown kind>"
	}
	panic("unreachable")
}

func NewToken(kind parse.Kind, b []byte, err error) (Token, error) {
	if err != nil {
		return Token{}, err
	}
	t := &Token{kind: kind, b: b}
	switch kind {
	case parse.UnknownKind:
		return *t, nil
	case parse.NullKind:
		return *t, nil
	case parse.FalseKind:
		return *t, nil
	case parse.TrueKind:
		return *t, nil
	case parse.BytesKind:
		return *t, nil
	case parse.StringKind:
		t.s = cast.ToString(t.b)
		return *t, nil
	case parse.Int64Kind:
		cast.ToInt64Ptr(t.b, &t.i)
		return *t, nil
	case parse.Float64Kind:
		cast.ToFloat64BitsPtr(t.b, &t.u)
		return *t, nil
	case parse.DecimalKind:
		t.s = cast.ToString(t.b)
		return *t, nil
	case parse.NanosecondsKind:
		t.i = cast.ToInt64(t.b)
		return *t, nil
	case parse.DateTimeKind:
		t.s = cast.ToString(t.b)
		return *t, nil
	case parse.TagKind:
		t.s = cast.ToString(t.b)
		return *t, nil
	}
	panic("unreachable")
}

func (t Token) Token(alloc func(size int) []byte) (parse.Kind, []byte, error) {
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
		return parse.StringKind, cast.FromString(t.s, alloc), nil
	case parse.Int64Kind:
		return parse.Int64Kind, cast.FromInt64Ptr(&t.i, alloc), nil
	case parse.Float64Kind:
		kind, val := parse.Float64Kind, cast.FromFloat64BitsPtr(&t.u, alloc)
		return kind, val, nil
	case parse.DecimalKind:
		return parse.DecimalKind, cast.FromString(t.s, alloc), nil
	case parse.NanosecondsKind:
		return parse.NanosecondsKind, cast.FromInt64Ptr(&t.i, alloc), nil
	case parse.DateTimeKind:
		return parse.DateTimeKind, cast.FromString(t.s, alloc), nil
	case parse.TagKind:
		return parse.TagKind, cast.FromString(t.s, alloc), nil
	}
	panic("unreachable")
}

func NewUknownToken() Token {
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
		b:    b,
	}
}

func NewStringToken(s string) Token {
	return Token{
		kind: parse.StringKind,
		s:    s,
	}
}

func NewInt64Token(i int64) Token {
	return Token{
		kind: parse.Int64Kind,
		i:    i,
	}
}

func NewFloat64Token(f float64) Token {
	return Token{
		kind: parse.Float64Kind,
		u:    math.Float64bits(f),
	}
}

func NewDecimalKind(d string) Token {
	return Token{
		kind: parse.DecimalKind,
		s:    d,
	}
}

func NewNanosecondsToken(n int64) Token {
	return Token{
		kind: parse.NanosecondsKind,
		i:    n,
	}
}

func NewDateTimeToken(t time.Time) Token {
	return Token{
		kind: parse.DateTimeKind,
		s:    t.Format(time.RFC3339Nano),
	}
}

func NewTagToken(t string) Token {
	return Token{
		kind: parse.TagKind,
		s:    t,
	}
}
