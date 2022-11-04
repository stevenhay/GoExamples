package main

import "fmt"

// a pointer is simply a variable which holds a memory location instead of a value
var (
	aPtr          *int   // a variable that holds a memory location which, at that memory location, is an int.
	twoLevelPtr   **int  // a two level pointer, points at a memory location which points at an int
	threeLevelPtr ***int // you get the point, btw the zero value of a poiner is nil
)

var iPtr *int = new(int) // new(T) will create a zero-valued T and return a pointer to it

var (
	myInt1   int  = 5
	myIntPtr *int = &myInt1 // or use the &var notation, &t returns *T
)

func main() {
	i := 1337
	fmt.Printf("i: %d\n", i)
	fmt.Printf("&i: %p\n", &i)
	fmt.Println("")

	fmt.Println("Change Int")
	changeInt(&i, 1000)
	fmt.Printf("i: %d\n", i)
	fmt.Printf("&i: %p\n", &i)
	fmt.Println("")

	// as expected, == on pointers check they point to the same address
	// not that the underlying value is the same
	//
	// note: these comparisons can only work if the types of the pointers
	// are the same, or one pointer can be implicitly converted to the other
	// or one of them is nil
	x := &i
	y := &i
	fmt.Printf("x = %p, *x = %d\n", x, *x)
	fmt.Printf("y = %p, *y = %d\n", y, *y)
	fmt.Printf("%p == %p: %t\n", x, y, x == y)
	fmt.Println("")

	a := *x
	b := &a
	fmt.Printf("x = %p, *x = %d\n", x, *x)
	fmt.Printf("b = %p, *b = %d\n", b, *b)
	fmt.Printf("%p == %p: %t\n", x, b, x == b)

	p := &i

	// these are the same, * has higher precedence than ++ or --
	// note: you *cannot* do arithmetic operations on pointers
	// p++ or p-- will not compile
	*p++
	(*p)++
}

func changeInt(ptr *int, newVal int) {
	// can't directly update the value pointed to by this pointer without dereferencing first
	// this would attempt to update the pointer itself and would not compile
	// ptr = newVal

	*ptr = newVal // *t dereferences the pointer so we can use the underlying value and type

	// remember that even though we're passing a pointer, we're passing this pointer value by
	// reference, so any updates to this pointer value wont be reflected anywhere else, in order
	// to update the pointer value we'd need to pass **int
	ptr = nil
}

func changeIntLoc(ptr **int, newPtr *int) {
	// double pointers in go are rare, and for our usecase we'll probably never use them
	// but this is one use case
	*ptr = newPtr
}
