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

package relaxng

type state struct {
	kind stateKind
}

type stateKind byte

const startState = stateKind(0)

const enterState = stateKind('{')

const unknownFieldState = stateKind('u')

const valueState = stateKind('v')

const endState = stateKind('$')

// attr states

const attrStartState = stateKind('a') // emit "attrs" tag

const attrStartedState = stateKind('b') // emit enter hint

const attrFieldState = stateKind('c') // emit attr field

const attrValueState = stateKind('d') // emit attr value

const attrLeaveState = stateKind('e') // emit leave hint

// elem states

const elemStartState = stateKind('A') // emit "elems" tag

const elemStartedState = stateKind('B') // emit enter hint

const elemFirstState = stateKind('C') // emit first elem field
