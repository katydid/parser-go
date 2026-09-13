// Copyright 2026 Walter Schulze
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

package relaxng

// Option is used set options when creating a new JSON Parser.
type Option func(*tagger)

// WithTags tags
// 1. elements with an `elems` tag, for example `<a/><b/>` is parsed as `{"elems": {"a": {}, "b": {}}}`.
// 2. attributes with an `attrs` tag, for example `<A a="b" c="d"/>` is parsed as `{"elems": {"A": {"attrs": {"a": "b", "c": "d"}}}}`.
func WithTags() Option {
	return func(t *tagger) {
		t.tag = true
	}
}

// WithAllocator replaces the default `func(size int) []byte { return make([]byte, size) }` allocator
// with a different allocator function.
// Usually an allocator that uses a pool.
func WithAllocator(alloc func(int) []byte) Option {
	return func(t *tagger) {
		t.alloc = alloc
	}
}
