// Code generated by go generate; DO NOT EDIT.
package main

var pieTemplates = map[string]string{
	"All": `package functions

// All will return true if all callbacks return true. It follows the same logic
// as the all() function in Python.
//
// If the list is empty then true is always returned.
func (ss SliceType) All(fn func(value ElementType) bool) bool {
	for _, value := range ss {
		if !fn(value) {
			return false
		}
	}

	return true
}
`,
	"Any": `package functions

// Any will return true if any callbacks return true. It follows the same logic
// as the any() function in Python.
//
// If the list is empty then false is always returned.
func (ss SliceType) Any(fn func(value ElementType) bool) bool {
	for _, value := range ss {
		if fn(value) {
			return true
		}
	}

	return false
}
`,
	"Append": `package functions

// Append will return a new slice with the elements appended to the end. It is a
// wrapper for the internal append(). It is offered as a function so that it can
// more easily chained.
//
// It is acceptable to provide zero arguments.
func (ss SliceType) Append(elements ...ElementType) SliceType {
	return append(ss, elements...)
}
`,
	"AreSorted": `package functions

import (
	"sort"
)

// AreSorted will return true if the slice is already sorted. It is a wrapper
// for sort.SliceTypeAreSorted.
func (ss SliceType) AreSorted() bool {
	return sort.SliceIsSorted(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
}
`,
	"AreUnique": `package functions

// AreUnique will return true if the slice contains elements that are all
// different (unique) from each other.
func (ss SliceType) AreUnique() bool {
	return ss.Unique().Len() == ss.Len()
}
`,
	"Average": `package functions

// Average is the average of all of the elements, or zero if there are no
// elements.
func (ss SliceType) Average() float64 {
	if l := ElementType(len(ss)); l > 0 {
		return float64(ss.Sum()) / float64(l)
	}

	return 0
}
`,
	"Bottom": `package functions

// Bottom will return n elements from bottom
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss SliceType) Bottom(n int) (top SliceType) {
	var lastIndex = len(ss) - 1
	for i := lastIndex; i > -1 && n > 0; i-- {
		top = append(top, ss[i])
		n--
	}

	return
}
`,
	"Contains": `package functions

// Contains returns true if the element exists in the slice.
//
// When using slices of pointers it will only compare by address, not value.
func (ss SliceType) Contains(lookingFor ElementType) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}
`,
	"Extend": `package functions

// Extend will return a new slice with the slices of elements appended to the
// end.
//
// It is acceptable to provide zero arguments.
func (ss SliceType) Extend(slices ...SliceType) (ss2 SliceType) {
	ss2 = ss

	for _, slice := range slices {
		ss2 = ss2.Append(slice...)
	}

	return ss2
}
`,
	"First": `package functions

// First returns the first element, or zero. Also see FirstOr().
func (ss SliceType) First() ElementType {
	return ss.FirstOr(ElementZeroValue)
}
`,
	"FirstOr": `package functions

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss SliceType) FirstOr(defaultValue ElementType) ElementType {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}
`,
	"JSONString": `package functions

import (
	"encoding/json"
)

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss SliceType) JSONString() string {
	if ss == nil {
		return "[]"
	}

	// An error should not be possible.
	data, _ := json.Marshal(ss)

	return string(data)
}
`,
	"Join": `package functions

// Join returns a string from joining each of the elements.
func (ss StringSliceType) Join(glue string) (s string) {
	for i, element := range ss {
		if i > 0 {
			s += glue
		}

		s += string(element)
	}

	return s
}
`,
	"Keys": `package functions

// Keys returns the keys in the map. All of the items will be unique.
//
// Due to Go's randomization of iterating maps the order is not deterministic.
func (m MapType) Keys() KeySliceType {
	// Avoid allocation
	l := len(m)
	if l == 0 {
		return nil
	}

	i := 0
	keys := make(KeySliceType, len(m))
	for key := range m {
		keys[i] = key
		i++
	}

	return keys
}
`,
	"Last": `package functions

// Last returns the last element, or zero. Also see LastOr().
func (ss SliceType) Last() ElementType {
	return ss.LastOr(ElementZeroValue)
}
`,
	"LastOr": `package functions

// LastOr returns the last element or a default value if there are no elements.
func (ss SliceType) LastOr(defaultValue ElementType) ElementType {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}
`,
	"Len": `package functions

// Len returns the number of elements.
func (ss SliceType) Len() int {
	return len(ss)
}
`,
	"Max": `package functions

// Max is the maximum value, or zero.
func (ss SliceType) Max() (max ElementType) {
	if len(ss) == 0 {
		return
	}

	max = ss[0]
	for _, s := range ss {
		if s > max {
			max = s
		}
	}

	return
}
`,
	"Min": `package functions

// Min is the minimum value, or zero.
func (ss SliceType) Min() (min ElementType) {
	if len(ss) == 0 {
		return
	}

	min = ss[0]
	for _, s := range ss {
		if s < min {
			min = s
		}
	}

	return
}
`,
	"Reverse": `package functions

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//   ss.Sort().Reverse()
//
func (ss SliceType) Reverse() SliceType {
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]ElementType, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}
`,
	"Select": `package functions

// Select will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// Unselect works in the opposite way as Select.
func (ss SliceType) Select(condition func(ElementType) bool) (ss2 SliceType) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}
`,
	"Shuffle": `package functions

import (
	"github.com/elliotchance/pie/pie/util"
	"math/rand"
)

// Shuffle returns shuffled slice by your rand.Source
func (ss SliceType) Shuffle(source rand.Source) SliceType {
	n := len(ss)

	// Avoid the extra allocation.
	if n < 2 {
		return ss
	}

	// go 1.10+ provides rnd.Shuffle. However, to support older versions we copy
	// the algorithm directly from the go source: src/math/rand/rand.go below,
	// with some adjustments:
	shuffled := make([]ElementType, n)
	copy(shuffled, ss)

	rnd := rand.New(source)

	util.Shuffle(rnd, n, func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}
`,
	"Sort": `package functions

import (
	"sort"
)

// Sort works similar to sort.SliceType(). However, unlike sort.SliceType the
// slice returned will be reallocated as to not modify the input slice.
//
// See Reverse() and AreSorted().
func (ss SliceType) Sort() SliceType {
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]ElementType, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}
`,
	"Sum": `package functions

// Sum is the sum of all of the elements.
func (ss SliceType) Sum() (sum ElementType) {
	for _, s := range ss {
		sum += s
	}

	return
}
`,
	"ToStrings": `package functions

import (
	"github.com/elliotchance/pie/pie"
)

// ToStrings transforms each element to a string.
func (ss SliceType) ToStrings(transform func(ElementType) string) pie.Strings {
	l := len(ss)

	// Avoid the allocation.
	if l == 0 {
		return nil
	}

	result := make(pie.Strings, l)
	for i := 0; i < l; i++ {
		result[i] = transform(ss[i])
	}

	return result
}
`,
	"Top": `package functions

// Top will return n elements from head of the slice
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss SliceType) Top(n int) (top SliceType) {
	for i := 0; i < len(ss) && n > 0; i++ {
		top = append(top, ss[i])
		n--
	}

	return
}
`,
	"Transform": `package functions

// Transform will return a new slice where each element has been transformed.
// The number of element returned will always be the same as the input.
//
// Be careful when using this with slices of pointers. If you modify the input
// value it will affect the original slice. Be sure to return a new allocated
// object or deep copy the existing one.
func (ss SliceType) Transform(fn func(ElementType) ElementType) (ss2 SliceType) {
	if ss == nil {
		return nil
	}

	ss2 = make([]ElementType, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}
`,
	"Unique": `package functions

// Unique returns a new slice with all of the unique values.
//
// The items will be returned in a randomized order, even with the same input.
//
// The number of items returned may be the same as the input or less. It will
// never return zero items unless then input slice has zero items.
//
// A slice with zero elements is considered to be unique.
//
// See AreUnique().
func (ss SliceType) Unique() SliceType {
	// Avoid the allocation. If there is one element or less it is already
	// unique.
	if len(ss) < 2 {
		return ss
	}

	values := map[ElementType]struct{}{}

	for _, value := range ss {
		values[value] = struct{}{}
	}

	var uniqueValues SliceType
	for value := range values {
		uniqueValues = append(uniqueValues, value)
	}

	return uniqueValues
}
`,
	"Unselect": `package functions

// Unselect works the same as Select, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss SliceType) Unselect(condition func(ElementType) bool) (ss2 SliceType) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}
`,
	"Values": `package functions

// Values returns the values in the map.
//
// Due to Go's randomization of iterating maps the order is not deterministic.
func (m MapType) Values() []ElementType {
	// Avoid allocation
	l := len(m)
	if l == 0 {
		return nil
	}

	i := 0
	keys := make([]ElementType, len(m))
	for _, value := range m {
		keys[i] = value
		i++
	}

	return keys
}
`,
}
