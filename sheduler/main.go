package main

import "fmt"

//go:noinline
func main() {

	var a = "HELO"
	println("MAIN:::")
	println(&a)

	foo(&a)
	go println("main")
	fmt.Println(a)
}

func foo(a *string) {
	println("MAIN:::")
	println(a)
	*a = "egor"
	println("MAIN:::")
	println(a)

	fmt.Println(a)

}
