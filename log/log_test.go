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
	"strings"
	"testing"

	. "katydid.org.za/go/parser-go/hedge"
	"katydid.org.za/go/parser-go/parse"
)

func TestFilename(t *testing.T) {
	n := NewStringNode("num", NewFloat64Node(3.14))

	var p parse.Parser = NewParser([]Node{n})
	buf := bytes.NewBuffer(nil)
	p = WrapParser(p, WithWriter(buf))
	p.Next()

	logs := string(buf.Bytes())
	t.Log(logs)
	if !strings.Contains(logs, "log_test.go") {
		t.Fatalf("logs do not contain file name from which printer was called: %s", logs)
	}
	if !strings.Contains(logs, "log_test.go:32") {
		t.Fatalf("logs do not contain line number from which printer was called: %s", logs)
	}
}
