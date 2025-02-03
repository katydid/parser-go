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

package debug

import (
	"math/rand"
	"time"
)

// NewRand returns a random integer generator, that can be used with RandomWalk.
// Its seed is the current time.
func NewRand() Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Rand is a subset of the interface that is implemented by math/rand representing only the methods used by the RandomWalk function.
type Rand interface {
	Intn(n int) int
}
