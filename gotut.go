package main

import (
	"fmt"
	"math"
	"math/rand"
)

// 0. `go run gotut.go` to run the file
// 1. capital letter means it will be exported/ public
// 2. main function will always run
// 3. `go doc fmt.Printf` to see the documentation for the function.
//   - installed godoc first `go install golang.org/x/tools/cmd/godoc@latest`
// 4. need to import to the specific file itself. I.e math/rand
// 5. need to specify the seed, if not rand.Intn(100) will always return 81

func main1() {
	fmt.Println("Welcome to Go!")
	fmt.Println("The square root of 4 is", math.Sqrt(4))
	fmt.Println("A number from 1-99", rand.Intn(100))
}

func foo() {
	fmt.Println("I'm in foo")
}

// ----------- Part 3 -------------------
// 6. var num1, num2 float64 = 5.6, 9.5 can be written as num1, num2 := 5.6, 9.5
func add(x, y float64) float64 {
	return x + y
}

// 7. you need to speify every return type (string, string)
func multiple(a, b string) (string, string) {
	return a, b
}

func main3() {
	//var num1, num2 float64 = 5.6, 9.5
	num1, num2 := 5.6, 9.5 // default is 64 bit float
	str1, str2 := "Hello", "World"

	fmt.Println(add(num1, num2))
	fmt.Println(multiple(str1, str2))

	/* 8. Type conversion
	var a int = 62
	var b float64 = float64(a)

	x := a // x will be type int
	*/
}

// ----------- Part 4 Pointers -------------------
// 9. &x will give you the memory address of x
// 10. *a will give you the value of the memory address
func main() {
	x := 15 // int
	a := &x // memory address of x
	fmt.Println(a) // 0x140000140c0
	*a = 5 // change the value at the memory address
	fmt.Println(x) // 5
	*a = *a**a // 5*5
	fmt.Println(x) // 25
	fmt.Println(*a) // 25
}
