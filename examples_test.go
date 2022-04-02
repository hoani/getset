package getset_test

import (
	"fmt"

	"github.com/hoani/getset"
)

// A set holds a group of unique values.
func Example() {
	mySet := getset.New("a", "b", "b", "c")
	fmt.Println(len(mySet))
	// Output: 3
}

// Create new sets of various types.
func ExampleNew() {
	floatSet := getset.New(1.0, 2.5, 4.75, 5.5)
	intSet := getset.New(-1, 0, 1, 3, 5)
	boolSet := getset.New(true)

	fmt.Println(len(floatSet))
	fmt.Println(len(intSet))
	fmt.Println(len(boolSet))
	// Output:
	// 4
	// 5
	// 1
}

func ExampleSet_Has() {
	mySet := getset.New("a", "c")

	fmt.Println(mySet.Has("a"))
	fmt.Println(mySet.Has("b"))
	// Output:
	// true
	// false
}

func ExampleSet_Insert() {
	mySet := getset.New(1, 3)
	fmt.Println(mySet.Has(2))

	mySet.Insert(2)
	fmt.Println(mySet.Has(2))
	// Output:
	// false
	// true
}

func ExampleSet_ToArray() {
	mySet := getset.New(1, 3, 5)
	myArr := mySet.ToArray()

	fmt.Println(len(myArr))
	// Output:
	// 3
}
