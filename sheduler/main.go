package main

//go:noinline
func main() {
	var v int = 0
	//fmt.Println(v)
	foo(&v)
}

//go:noinline
func foo(a ...any) {
	foo2(a)

}
func foo2(g ...any) interface{} {
	g
	return g

}
