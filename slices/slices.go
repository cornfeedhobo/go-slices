// Copyright © 2016 cornfeedhobo <cornfeedhobo@fuzzlabs.org>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package slices

import "reflect"

// Returns true if slice containes elem
//
// Disclaimer: Panic on unexported struct fields
func Contains(slice interface{}, elem interface{}) bool {
	val := reflect.ValueOf(slice)
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			if val.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}

// Returns slice index of elem if elem exists
// or -1 if unable to find elem.
//
// Disclaimer: Panic on unexported struct fields
func IndexOf(slice interface{}, elem interface{}) int {
	val := reflect.ValueOf(slice)
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			if val.Index(i).Interface() == elem {
				return i
			}
		}
	}
	return -1
}

// Returns the next index, wrapping to 0 if needed
// or -1 if not a slice
func IndexNext(slice interface{}, index int) int {
	val := reflect.ValueOf(slice)
	if val.Kind() == reflect.Slice {
		if index + 1 == val.Len() {
			return 0
		}
		return index + 1
	}
	return -1
}
