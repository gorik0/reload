package main

func main() {
	var rio = make([]uint8, 3)
	println(&rio)
	println(rio)

	println(&(rio[2]))
	rio = append(rio, rio...)
	println(&rio)
	println(rio)
	println(&(rio[2]))
}
