package main

import "fmt"

type MyInterface interface {
	Hello(name string)
	GetString() string
	NoNamedVariables(string, int) bool
}

type SecondInterface interface {
	Foo(string) string
}

type ThirdInterface interface {
	Bar(string) string
	SecondInterface
}

type StructFoo struct{}

func (s StructFoo) Foo(x string) string {
	return "Foo"
}

type MyString string

func (s MyString) Foo(x string) string {
	return "Foo"
}

func (s MyString) Bar(x string) string {
	return "Bar"
}

// ---

type Bar interface {
	Bar(string) string
}

type HelloBar struct{}

func (HelloBar) Bar(x string) string {
	return fmt.Sprintf("Hello, %s", x)
}

type GoodbyeBar struct{}

func (GoodbyeBar) Bar(x string) string {
	return fmt.Sprintf("Goodbye, %s", x)
}

func DoSomething(action Bar, name string) string {
	return action.Bar(name)
}

func main() {
	var sif SecondInterface = StructFoo{}
	fmt.Println(sif.Foo("x"))

	// var tif ThirdInterface = StructFoo{} // does not compile!

	var sif2 SecondInterface = MyString("Hello!")
	var tif2 ThirdInterface = MyString("World!")

	sif2.Foo("x")
	tif2.Foo("x")
	tif2.Bar("x")

	fmt.Println(DoSomething(HelloBar{}, "World"))
	fmt.Println(DoSomething(GoodbyeBar{}, "World"))

	// the compiler stores type information
	// of interfaces
	var x interface{} = 666

	// we can type assert
	i, ok := x.(int)
	fmt.Println(i, ok)

	f, ok := x.(float64)
	fmt.Println(f, ok)

	b, ok := x.(bool)
	fmt.Println(b, ok)

	// if you do not include the second safety variable
	// Go will panic if type assert fails
	i = x.(int)
	// b = x.(bool) // panic

	// can also assert things as types
	var h interface{} = HelloBar{}
	gb, ok := h.(Bar)
	fmt.Println(gb, ok)

	s := []interface{}{1, true, "hello", false, "bar", GoodbyeBar{}, 2.32}
	for _, x := range s {
		switch v := x.(type) {
		case int:
			fmt.Printf("int: %d\n", v)
		case bool:
			fmt.Printf("bool: %v\n", v)
		case string:
			fmt.Printf("string: %s\n", v)
		case GoodbyeBar:
			fmt.Printf("GoodbyeBar: %v\n", v)
		default:
			fmt.Printf("unknown type %T", v)
		}
	}
}
