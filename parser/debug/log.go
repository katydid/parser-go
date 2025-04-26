//  Copyright 2015 Walter Schulze
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
	"github.com/katydid/parser-go/parser"
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
	s      parser.Interface
	l      Logger
	copies int
}

// NewLogger returns a parser that when called returns and logs the value returned by the argument parser to the argument logger.
func NewLogger(s parser.Interface, logger Logger) parser.Interface {
	return &l{"parser", s, logger, 0}
}

func (l *l) Double() (float64, error) {
	v, err := l.s.Double()
	l.l.Printf(l.name+".Double() (%v, %v)", v, err)
	return v, err
}

func (l *l) Int() (int64, error) {
	v, err := l.s.Int()
	l.l.Printf(l.name+".Int() (%v, %v)", v, err)
	return v, err
}

func (l *l) Uint() (uint64, error) {
	v, err := l.s.Uint()
	l.l.Printf(l.name+".Uint() (%v, %v)", v, err)
	return v, err
}

func (l *l) Bool() (bool, error) {
	v, err := l.s.Bool()
	l.l.Printf(l.name+".Bool() (%v, %v)", v, err)
	return v, err
}

func (l *l) String() (string, error) {
	v, err := l.s.String()
	l.l.Printf(l.name+".String() (%v, %v)", v, err)
	return v, err
}

func (l *l) Bytes() ([]byte, error) {
	v, err := l.s.Bytes()
	l.l.Printf(l.name+".Bytes() (%v, %v)", v, err)
	return v, err
}

func (l *l) Next() error {
	err := l.s.Next()
	l.l.Printf(l.name+".Next() (%v)", err)
	return err
}

func (l *l) IsLeaf() bool {
	v := l.s.IsLeaf()
	l.l.Printf(l.name+".IsLeaf() (%v)", v)
	return v
}

func (l *l) Up() {
	l.s.Up()
	l.l.Printf(l.name + ".Up()")
}

func (l *l) Down() {
	l.s.Down()
	l.l.Printf(l.name + ".Down()")
}
