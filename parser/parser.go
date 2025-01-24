//  Copyright 2013 Walter Schulze
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

// Package parser represents the parser.Interface and some errors for implementing parsers.
package parser

import "fmt"

// A type conforming to the parser.Interface interface, abstracts away the implementation details of a parser.
type Interface interface {
	// Next skips to the next field or item in an array
	Next() error
	// IsLeaf is true if this is a value and you cannot traverse down anymore
	IsLeaf() bool
	// Down traverses down into a field value, which could be another message or array. Next must always be called after Down.
	Down()
	// Up traverses up out of a field value and back to the field's next sibling. Next must always be called after Up.
	Up()
	// Value is a collection of possible values that the field might have.
	Value
}

// A type confirming to the parser.Value interface, repesents one native value, tree node label (field name) or some repesentation a node label.
// Typically only one of the methods returns a value without an error, but more than one method can return without an error.
// For example a positive json number can return an errorless value for the Double, Int and Uint methods.
type Value interface {
	// String returns the string value if it is a string type or an error if it is not a string.
	String() (string, error)
	// Double returns the float64 value if it is a double type or an error if it is not a double.
	Double() (float64, error)
	// Int returns the int64 value if it is an integer type or an error if it is not an integer.
	Int() (int64, error)
	// Uint returns the uint64 value if it is an unsinged integer type or an error if it is not an unsinged integer.
	Uint() (uint64, error)
	// Bool returns the bool value if it is a boolean type or an error if it is not a boolean.
	Bool() (bool, error)
	// Bytes returns a byte slice value if it is a bytes type or an error if it is not bytes.
	Bytes() ([]byte, error)
}

// Sprint returns a value printed as a string.
func Sprint(value Value) string {
	return fmt.Sprintf("%#v", getValue(value))
}

func getValue(value Value) interface{} {
	var v interface{}
	var err error
	v, err = value.Bool()
	if err == nil {
		return v
	}
	v, err = value.Bytes()
	if err == nil {
		return v
	}
	v, err = value.String()
	if err == nil {
		return v
	}
	v, err = value.Int()
	if err == nil {
		return v
	}
	v, err = value.Uint()
	if err == nil {
		return v
	}
	v, err = value.Double()
	if err == nil {
		return v
	}
	return nil
}
