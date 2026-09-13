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

// RelaxNGable is an extra method for a Parser that distinguishes between elements and attributes.
// This allows the Parser to be tagged to handle RelaxNG's `<element ...>` and `<attribute ...>` operators.
type RelaxNGable interface {
	// RelaxNGType returns a type that distinguishes between elements and attributes, after Next returned a FieldHint.
	RelaxNGType() RelaxNGType
}

type RelaxNGType byte

const RelaxNGTypeUnknown = RelaxNGType(0)

const RelaxNGTypeElem = RelaxNGType('e')

const RelaxNGTypeAttr = RelaxNGType('a')
