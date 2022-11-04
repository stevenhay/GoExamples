package main

import "fmt"

type MyStruct struct{}
type MySecondStruct struct {
	Num    int
	Str    string
	MyBool bool

	// Can have multiple declarations on one line
	Str1, Str2 string
}

// we'll go over actually parsing json later
type MyStructWithTags struct {
	AccountID   string `json:"account_id"`
	AccountName string `json:"account_name,omitempty" otherfmt:"acc_nm"`
}

type StructWithMethods struct {
	Value string
}

// this function can only be called via a StructWithMethods variable
func (s StructWithMethods) GetValue() string {
	return s.Value
}

func (s StructWithMethods) SetValue(val string) {
	s.Value = val
}

func (s *StructWithMethods) ReallySetValue(val string) {
	s.Value = val
}

type MyInt int

// some people call these "methods" and the ones not assigned to a specific type
// "functions" but I use the term method/function interchangibly
func (i *MyInt) AddOne() {
	*i++
}

func main() {
	// we don't need to specify any of the members
	// when declaring the struct variable
	s := MySecondStruct{}
	s.MyBool = true

	// but if we want to, we can do it like this
	s = MySecondStruct{
		MyBool: true, // this comma is mandatory
	}

	// or
	s = MySecondStruct{MyBool: true}

	// or
	// however this relies on the order of members given here
	// matching the order declared within the struct
	s = MySecondStruct{2, "hello", true, "world", "123"}

	s2 := MyStructWithTags{AccountID: "123", AccountName: "ABC"}
	s3 := MyStructWithTags{AccountID: "345"}

	// this assignment is the same as copying all fields from s2 to se
	s3 = s2
	fmt.Printf("%v @ %p\n", s2, &s2)
	fmt.Printf("%v @ %p\n", s3, &s3)

	// dereferecing receives are implicit
	// (pointers still covered later)
	sptr := &MySecondStruct{MyBool: true}
	fmt.Printf("%t\n", sptr.MyBool)
	fmt.Printf("%t\n", (*sptr).MyBool) // no need to explicitly deference sptr

	// GetValue() // undeclared value
	s4 := StructWithMethods{Value: "Hello!"}
	fmt.Println(s4.GetValue())

	s4.SetValue("World!")
	fmt.Println(s4.GetValue())

	s4.ReallySetValue("World!")
	fmt.Println(s4.GetValue())

	i1 := MyInt(1)
	i1.AddOne()
	fmt.Printf("%d\n", i1)
}
