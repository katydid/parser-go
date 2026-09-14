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
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"katydid.org.za/go/parser-go/hedge"
	"katydid.org.za/go/parser-go/parse"
)

type l struct {
	name        string
	parser      parse.Parser
	printer     printer
	delay       *time.Duration
	lineNumbers bool
	token       []byte
}

// WrapParserWithInit returns a ParserWithInit that when called returns and logs the value returned by the argument parser to the argument logger.
func WrapParserWithInit(p parse.ParserWithInit, opts ...Option) parse.ParserWithInit {
	o := newOptions(opts...)
	return &l{name: o.name, parser: p, printer: newPrinter(o.writer), delay: o.delay, lineNumbers: o.lineNumbers, token: nil}
}

// WrapParserWithReset returns a ParserWithReset that when called returns and logs the value returned by the argument parser to the argument logger.
func WrapParserWithReset(p parse.ParserWithReset, opts ...Option) parse.ParserWithReset {
	o := newOptions(opts...)
	return &l{name: o.name, parser: p, printer: newPrinter(o.writer), delay: o.delay, lineNumbers: o.lineNumbers, token: nil}
}

// WrapParser returns a Parser that when called returns and logs the value returned by the argument parser to the argument logger.
func WrapParser(p parse.Parser, opts ...Option) parse.Parser {
	o := newOptions(opts...)
	return &l{name: o.name, parser: p, printer: newPrinter(o.writer), delay: o.delay, lineNumbers: o.lineNumbers, token: nil}
}

func (l *l) Init(buf []byte) {
	l.parser.(parse.ParserWithInit).Init(buf)
	l.Printf("%s.Init(...)", l.name)
}

func (l *l) Reset() {
	l.parser.(parse.ParserWithReset).Reset()
	l.Printf("%s.Reset()", l.name)
}

func (l *l) Next() (parse.Hint, error) {
	hint, err := l.parser.Next()
	l.Printf("%s.Next() (%v, %v)", l.name, hint, err)
	return hint, err
}

func (l *l) Skip() error {
	err := l.parser.Skip()
	l.Printf("%s.Skip() (%v)", l.name, err)
	return err
}

func (l *l) Token() (parse.Kind, []byte, error) {
	kind, val, err := l.parser.Token()
	l.token = bytes.Clone(val)
	s := hedge.TokenString(kind, l.token)
	l.Printf("%s.Token() (%v, %v, %v)", l.name, kind, s, err)
	return kind, l.token, err
}

func (l *l) Printf(format string, v ...any) {
	lineNumber := ""
	if l.lineNumbers {
		lineNumber = getLineNumber()
	}
	l.printer.Printf(lineNumber+": "+format, v...)
	if l.delay != nil {
		time.Sleep(*l.delay)
	}
}

func getLineNumber() string {
	_, thisfile, _, ok := runtime.Caller(0)
	if !ok {
		return "<weirdlyunknown>:0"
	}
	i := 0
	for {
		i++
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			return "<unknown>:" + strconv.Itoa(i)
		}
		if file == thisfile {
			continue
		}
		_, name := filepath.Split(file)
		return name + ":" + strconv.Itoa(line)
	}
}
