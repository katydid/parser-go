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

package parse

import (
	"github.com/katydid/parser-go/log"
)

type l struct {
	name string
	p    ParserWithInit
	l    log.Logger
}

// NewLogger returns a parser that when called returns and logs the value returned.
// This is only to be used for debugging purposes.
func NewLogger(parser ParserWithInit, opts ...log.Option) ParserWithInit {
	return &l{"p", parser, log.NewLogger(opts...)}
}

func (l *l) Init(buf []byte) {
	l.p.Init(buf)
	l.l.Printf(l.name + ".Init(...)")
}

func (l *l) Skip() error {
	err := l.p.Skip()
	l.l.Printf(l.name+".Skip() (%v)", err)
	return err
}

func (l *l) Next() (Hint, error) {
	v, err := l.p.Next()
	l.l.Printf(l.name+".Next() (%v, %v)", v, err)
	return v, err
}

func (l *l) Token() (Kind, []byte, error) {
	k, bs, err := l.p.Token()
	l.l.Printf(l.name+".Token() (%v, %v, %v)", k, bs, err)
	return k, bs, err
}
