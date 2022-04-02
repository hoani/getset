package getset

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Set_String(t *testing.T) {
	set := New("a", "A", "b", "C")
	testCases := []struct {
		item     string
		expected bool
	}{
		{item: "a", expected: true},
		{item: "A", expected: true},
		{item: "b", expected: true},
		{item: "B", expected: false},
		{item: "c", expected: false},
		{item: "C", expected: true},
		{item: "d", expected: false},
		{item: "ab", expected: false},
		{item: "", expected: false},
	}
	for _, tc := range testCases {
		name := fmt.Sprintf("check %s", tc.item)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, set.Has(tc.item))
		})
	}
}

func Test_Set_Int(t *testing.T) {
	set := New(-1, 0, 1, 100)
	testCases := []struct {
		item     int
		expected bool
	}{
		{item: -100, expected: false},
		{item: -50, expected: false},
		{item: -1, expected: true},
		{item: 0, expected: true},
		{item: 1, expected: true},
		{item: 50, expected: false},
		{item: 100, expected: true},
	}
	for _, tc := range testCases {
		name := fmt.Sprintf("check %d", tc.item)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, set.Has(tc.item))
		})
	}
}

func Test_Set_Float64(t *testing.T) {
	set := New(-1.0, 0.5, 1.9, math.Inf(1), math.Inf(-1), math.NaN())
	testCases := []struct {
		item     float64
		expected bool
	}{
		{item: -1.0, expected: true},
		{item: 0.5, expected: true},
		{item: 1.9, expected: true},
		{item: math.Inf(1), expected: true},
		{item: math.Inf(-1), expected: true},
		{item: math.NaN(), expected: false}, // NaN cannot be compared, up to the user not to make this mistake.
		{item: -0.5, expected: false},
		{item: 100.0, expected: false},
	}
	for _, tc := range testCases {
		name := fmt.Sprintf("check %.2f", tc.item)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, set.Has(tc.item))
		})
	}
}

func Test_Set_Custom(t *testing.T) {
	type customType struct {
		i int
		s string
	}
	set := New(
		customType{i: -1, s: ""},
		customType{i: 0, s: "ab"},
		customType{i: 1, s: "1"},
		customType{i: 1, s: "a"},
	)
	testCases := []struct {
		item     customType
		expected bool
	}{
		{item: customType{i: -1, s: ""}, expected: true},
		{item: customType{i: -1, s: "a"}, expected: false},
		{item: customType{i: -2, s: ""}, expected: false},
		{item: customType{i: 0, s: "ab"}, expected: true},
		{item: customType{i: 0, s: "a"}, expected: false},
		{item: customType{i: 0, s: "ba"}, expected: false},
		{item: customType{i: 1, s: "1"}, expected: true},
		{item: customType{i: 1, s: "1.1"}, expected: false},
		{item: customType{i: 1, s: "a"}, expected: true},
	}
	for _, tc := range testCases {
		name := fmt.Sprintf("check %v", tc.item)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, set.Has(tc.item))
		})
	}
}

func Test_Set_Length(t *testing.T) {
	type customType struct {
		i int
		f float32
	}

	testCases := []struct {
		name     string
		set      interface{}
		expected int
	}{
		{
			name:     "empty set",
			set:      New[string](),
			expected: 0,
		},
		{
			name:     "set with one item",
			set:      New("a"),
			expected: 1,
		},
		{
			name:     "set with two items",
			set:      New("a", "ab"),
			expected: 2,
		},
		{
			name:     "integer set",
			set:      New(-1, 0, 1, 100),
			expected: 4,
		},
		{
			name:     "set of custom types",
			set:      New(customType{i: 1, f: 2.5}, customType{i: 3}, customType{f: -0.5}),
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Len(t, tc.set, tc.expected)
		})
	}
}

func Test_Set_Delete_String(t *testing.T) {
	testCases := []struct {
		name     string
		items    []string
		expected int
	}{
		{
			name:     "delete known item",
			items:    []string{"a"},
			expected: 3,
		},
		{
			name:     "delete unknown item",
			items:    []string{"A"},
			expected: 4,
		},
		{
			name:     "delete all items",
			items:    []string{"a", "ab", "abc", "abcd"},
			expected: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set := New("a", "ab", "abc", "abcd")
			for _, item := range tc.items {
				delete(set, item)
			}
			assert.Len(t, set, tc.expected)
		})
	}
}

func Test_Set_Delete_CustomStruct(t *testing.T) {
	type customType struct {
		i int
		f float64
	}

	testCases := []struct {
		name     string
		items    []customType
		expected int
	}{
		{
			name:     "delete known item",
			items:    []customType{{i: 3}},
			expected: 2,
		},
		{
			name:     "delete unknown item",
			items:    []customType{{i: 3, f: -0.5}},
			expected: 3,
		},
		{
			name:     "delete all items",
			items:    []customType{{i: 1, f: 2.5}, {i: 3}, {f: -0.5}},
			expected: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set := New([]customType{{i: 1, f: 2.5}, {i: 3}, {f: -0.5}}...)
			for _, item := range tc.items {
				delete(set, item)
			}
			assert.Len(t, set, tc.expected)
		})
	}
}

func Test_Set_Insert_String(t *testing.T) {
	testCases := []struct {
		name     string
		item     string
		expected int
	}{
		{
			name:     "insert known item",
			item:     "a",
			expected: 2,
		},
		{
			name:     "insert unknown item",
			item:     "A",
			expected: 3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set := New("a", "b")

			set.Insert(tc.item)
			assert.True(t, set.Has(tc.item))
			assert.Len(t, set, tc.expected)
		})
	}
}

func Test_Set_Insert_Integer(t *testing.T) {
	testCases := []struct {
		name     string
		item     int
		expected int
	}{
		{
			name:     "insert known item",
			item:     0,
			expected: 2,
		},
		{
			name:     "insert unknown item",
			item:     -5,
			expected: 3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set := New(0, 5)

			set.Insert(tc.item)
			assert.True(t, set.Has(tc.item))
			assert.Len(t, set, tc.expected)
		})
	}
}

func Test_Set_Slice_Int(t *testing.T) {
	testCases := []struct {
		name  string
		items []int
	}{
		{
			name:  "empty",
			items: []int{},
		},
		{
			name:  "single",
			items: []int{5},
		},
		{
			name:  "many",
			items: []int{5, -5, 10, -10, 15, -15},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set := New(tc.items...)

			result := set.ToArray()
			assert.ElementsMatch(t, tc.items, result)
		})
	}
}
