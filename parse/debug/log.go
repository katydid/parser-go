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
	"time"

	"github.com/katydid/parser-go/log"
	"github.com/katydid/parser-go/parse"
)

// Logger is an interface for a type that is made to log debug info.
type Logger interface {
	Printf(format string, v ...any)
}

// NewLineLogger returns a logger that logs the line at which the Printf method was called to stderr.
func NewLineLogger() Logger {
	return log.NewLogger(log.WithLineNumbers())
}

// NewDelayLogger returns a logger that sleeps after every log.
// This is useful for debugging infinite loops.
func NewDelayLogger(delay time.Duration) Logger {
	return log.NewLogger(log.WithLineNumbers(), log.WithDelay(delay))
}

type l struct {
	name   string
	p      parse.ParserWithInit
	l      Logger
	copies int
}

// NewLogger returns a parser that when called returns and logs the value returned by the argument parser to the argument logger.
func NewLogger(p parse.ParserWithInit, logger Logger) parse.ParserWithInit {
	return &l{"parser", p, logger, 0}
}

func (l *l) Init(buf []byte) {
	l.p.Init(buf)
	l.l.Printf(l.name + ".Init(...)")
}

func (l *l) Next() (parse.Hint, error) {
	hint, err := l.p.Next()
	l.l.Printf(l.name+".Next() (%v, %v)", hint, err)
	return hint, err
}

func (l *l) Skip() error {
	err := l.p.Skip()
	l.l.Printf(l.name+".Skip() (%v)", err)
	return err
}

func (l *l) Token() (parse.Kind, []byte, error) {
	kind, val, err := l.p.Token()
	value, valerr := parse.GetValue(l.p)
	if valerr != nil {
		l.l.Printf(l.name+".Token() (%v, %v, %v) parse.GetValue() (%v)", kind, val, err, valerr)
	} else {
		l.l.Printf(l.name+".Token() (%v, %v, %v)", kind, value, err)
	}
	return kind, val, err
}
