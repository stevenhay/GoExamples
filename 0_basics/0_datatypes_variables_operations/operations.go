package main

func main() {
	a := 5
	_ = a + 2
	_ = a - 2
	_ = a * a
	_ = a / 2
	_ = a % 2

	_ = +a
	_ = -a
	_ = ^a // bitwise NOTa

	a++
	a--

	// n := a++
	// a++ is not an expression in Go
	// ++a doesn't exist
	// --a doesn't exist

	b0 := 0b1010
	b1 := 0b0101

	_ = b0 & b1  // bitwise AND
	_ = b0 | b1  // bitwise OR
	_ = b0 ^ b1  // bitwise XOR
	_ = b0 &^ b1 // bitwise CLEAR
	_ = b0 << 1  // bitwise LEFT SHIFT
	_ = b1 >> 1  // bitwise RIGHT SHIFT

	s := "hello " + "world" // string concat
	s += " 123"

	_ = !true
	_ = true && false
	_ = true || false

	_ = 1 == 2
	_ = 1 != 2

	_ = 1 >= 2
	_ = 1 > 2

	_ = 1 <= 2
	_ = 1 < 2
}
