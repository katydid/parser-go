// Copyright 2015 Walter Schulze
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

// Logger is an interface for a type that is made to log debug info.
type Logger interface {
	Printf(format string, v ...any)
}

type logger struct {
	*options
}

func NewLogger(opts ...Option) Logger {
	options := newOptions(opts...)
	return &logger{options}
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

func (l *logger) Printf(format string, v ...any) {
	lineNumber := ""
	if l.lineNumbers {
		lineNumber = getLineNumber()
	}
	l.logger.Printf(lineNumber+": "+format, v...)
	if l.delay != nil {
		time.Sleep(*l.delay)
	}
}
