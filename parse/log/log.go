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

package log

import (
	"bytes"

	"katydid.org.za/go/parser-go/hedge"
	"katydid.org.za/go/parser-go/parse"
)

// Logger is an interface for a type that is made to log debug info.
type Logger interface {
	Printf(format string, v ...any)
}

type l struct {
	name string
	p    parse.ParserWithInit
	l    Logger
}

// WrapParserWithInit returns a ParserWithInit that when called returns and logs the value returned by the argument parser to the argument logger.
func WrapParserWithInit(p parse.ParserWithInit, opts ...Option) parse.ParserWithInit {
	return &l{"parser", p, newLogger(newOptions(opts...))}
}

// WrapParser returns a Parser that when called returns and logs the value returned by the argument parser to the argument logger.
func WrapParser(p parse.Parser, opts ...Option) parse.Parser {
	return &l{"parser", parse.WithNoopInit(p), newLogger(newOptions(opts...))}
}

func (l *l) Init(buf []byte) {
	l.p.Init(buf)
	l.l.Printf("%s.Init(...)", l.name)
}

func (l *l) Next() (parse.Hint, error) {
	hint, err := l.p.Next()
	l.l.Printf("%s.Next() (%v, %v)", l.name, hint, err)
	return hint, err
}

func (l *l) Skip() error {
	err := l.p.Skip()
	l.l.Printf("%s.Skip() (%v)", l.name, err)
	return err
}

func (l *l) Token() (parse.Kind, []byte, error) {
	kind, val, err := l.p.Token()
	val = bytes.Clone(val)
	tok, tokerr := hedge.NewToken(kind, val, err)
	if tokerr != nil {
		l.l.Printf("%s.Token() (%v, %v, %v) hedge.NewToken() (%v)", l.name, kind, val, err, tokerr)
	} else {
		l.l.Printf("%s.Token() (%v, %v, %v)", l.name, kind, tok, err)
	}
	return kind, val, err
}
