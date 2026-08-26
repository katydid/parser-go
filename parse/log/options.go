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
	"time"

	"katydid.org.za/go/parser-go/logger"
)

type options struct {
	lineNumbers bool
	delay       *time.Duration
}

type Option func(*options)

func newDefaultOptions() *options {
	return &options{
		lineNumbers: true,
		delay:       nil,
	}
}

func WithDelay(dur time.Duration) Option {
	return func(o *options) {
		o.delay = &dur
	}
}

func WithoutLineNumbers() Option {
	return func(o *options) {
		o.lineNumbers = false
	}
}

func newOptions(opts ...Option) *options {
	o := newDefaultOptions()
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func newLogger(o *options) Logger {
	logopts := []logger.Option{}
	if o.lineNumbers {
		logopts = append(logopts, logger.WithLineNumbers())
	}
	if o.delay != nil {
		logopts = append(logopts, logger.WithDelay(*o.delay))
	}
	return logger.New(logopts...)
}
