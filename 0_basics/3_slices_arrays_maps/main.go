package main

import (
	"fmt"
)

var ArraysAreSized [3]int = [3]int{1, 2, 3}
var IndexArrayDecl [3]int = [3]int{2: 3, 0: 1, 1: 2}
var MixedArrayDecl [3]int = [3]int{2: 3, 0: 1, 2}

var CompilerDoTheWorkPls = [...]int{1, 2, 3, 4, 5, 6}

// Arrays don't have to be primitive values
var ArrayOfSlices [3][]int = [3][]int{[]int{1}, {2, 3, 4}, {5, 6, 7, 8, 9, 0}}

type Employee struct {
	ID         int
	Name       string
	Department string
}

var ArrayOfEmployees [2]Employee = [2]Employee{Employee{ID: 0, Name: "Name1", Department: "Department1"}, Employee{ID: 1, Name: "Name2", Department: "Dep2"}}

// The only difference in array and slice declaration is the inclusion
// or exclusion of a size
var SlicesAreNotSized []int = []int{1, 2, 3, 4, 5}
var SliceOfEmployees []Employee = []Employee{}  // as many employees as you'd like, leaving it empty because I'm lazy
var SliceOfPointers []*Employee = []*Employee{} // pointers covered later

var MapsAreHashtabes map[string]string = map[string]string{"key1": "val1", "key2": "val2"}

// map keys don't have to be primitives either
// this is a stupid example but it gets the point across
var MapWithEmployeeAsKey map[Employee]*Employee
var MapMapMap map[string]map[int]map[Employee]string = map[string]map[int]map[Employee]string{
	"String1": map[int]map[Employee]string{
		1: map[Employee]string{
			Employee{ID: 0, Name: "Name", Department: "Dep1"}: "Wtf",
		},
	},
	"String2": {1: {{ID: 0, Name: "Name", Department: "Dep1"}: "Wtf"}},
}

// All three are addressable
var (
	ArrPtr *[3]int            = &[3]int{0, 1, 2}
	SlcPtr *[]int             = &[]int{0, 1, 2}
	MapPtr *map[string]string = &map[string]string{"A": "B"}
)

func arrays() {
	var a1 = [3]int{1, 2, 3}
	var a2 = [3]int{1, 2, 3}

	// arrays can be compared
	// and yes, it compares the elements, not just
	// if they're the same object (java pls)
	_ = a1 == a2

	// can use the buildin methods len and cap
	_ = len(a1) // 3
	_ = cap(a1) // 3

	// getting values from slices and arrays are the same
	// as java
	_ = a1[0] // 1
	_ = a2[1] // 2

	// so is setting
	a1[0] = 10
	a2[0] = 10
}

func slices() {
	s1 := []int{1, 2, 3}
	_ /* s2 */ = []int{1, 2, 3}

	// slices cannot be compared except to nil
	// var sEq = s1 == s2 // NOT OK
	_ = s1 == nil // OK

	// appending to slice
	// this will extend the size of the underlying array
	// wont go into details here as this is the basics but the new
	// cap will be at least double the last one
	s1 = append(s1, 4)
	fmt.Printf("%v\n", s1)

	// there's no way to "delete" from a slice, instead
	// make a new slice by subslicing the original
	// can also overwrite the original slice
	indexToDelete := 1
	newSlice := append(s1[:indexToDelete], s1[indexToDelete+1:]...)
	// s1 = append(s1[:indexToDelete], s1[indexToDelete+1:]...) // to overwrite original
	fmt.Println(newSlice)

	// as for subslicing, the new slice uses the same underlying array
	// as the original slice, but points at the selected data
	//
	// slicing works as such: slice[low:high] where [low, high)
	// two-value slicing is shorthand for slice[low:high:cap(slice)]
	ss := s1[1:]
	ss[0] = 99
	fmt.Println(ss, s1)

	// cloning a slice
	cloned := make([]int, len(s1))
	copy(cloned, s1)

	fmt.Printf("original: %v, new: %v\n", s1, cloned)
	fmt.Printf("original: %p, new: %p\n", &(*[4]int)(s1)[0], &(*[4]int)(cloned)[0])
	fmt.Printf("original == new?: %t\n", *(*[4]int)(s1) == *(*[4]int)(cloned))
}

func maps() {
	m1 := map[string]int{"A": 65, "B": 66}
	_ /* m2 */ = map[string]int{"A": 65, "B": 66}

	// maps cannot be compared except to nil
	// var eq = m1 == m2 // NOT OK
	_ = m1 == nil //OK

	// getting from maps is different
	elem, exists := m1["B"]
	if exists {
		fmt.Printf("elem exists: %d\n", elem)
	}

	if elem, exists := m1["A"]; exists {
		fmt.Printf("elem exists: %d\n", elem)
	}

	// map assignment is by reference, not copy
	mcopy := m1
	mcopy["C"] = 67
	fmt.Println(mcopy, m1)

	// delete from map
	delete(mcopy, "C")
	fmt.Println(mcopy, m1)
}

func usingMake() {
	// maps
	// without specifying a size, the make command for maps will
	// allocate enough memory for a small number of elements
	// (compiler dependant)
	_ = make(map[string]string)

	// but we can specify the number of elements this map should
	// be able to take before having to re-allocate memory
	_ = make(map[string]string, 100)

	// slices
	// make(T, len, cap)
	// make a new slice and initalise elements = length
	// but capacity of cap
	_ = make([]string, 0, 100)

	// make(T, len)
	_ = make([]string, 0)
}

func main() {
	arrays()
	slices()
	maps()
	usingMake()
}
