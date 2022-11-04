package main

import "fmt"

func foo() {
	bar(1, 2, 3, "a", "b", "c", true, "d", "e") // can use methods before declaration
}

func main() {
	inline := func(x string) (string, string, string) {
		return "inline", "function", x
	}

	s1, s2, s3 := inline("x")
	fmt.Printf("%s, %s, %s\n", s1, s2, s3)

	// inline functions can be called immediately
	a, b := func() (int, int) {
		return 1, 2
	}()
	fmt.Printf("%d, %d\n", a, b)

	ans := sum(5, 5)
	fmt.Printf("5 + 5 = %d\n", ans)

	deferCalls()
}

// parameters will take the datatype of the next typed parameter
// or you can type them all
func bar(a, b, c int, d, e, f string, g bool, h string, i string) {
}

func baz() (string, int, bool) {
	// functions can return multiple values
	return "Hello", 12, true
}

// you can also specify named returns
func abc(a int) (x, y int) {
	x = a * 10
	y = a * 100
	return
}

// errors will be covered more later
// but functions can return errors
func err() error {
	return fmt.Errorf("my error")
}
