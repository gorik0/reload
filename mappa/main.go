package main

import "fmt"

func main() {

	rio := [5]int{1, 2, 3, 23, 5}

	println(&rio)
	println(&(rio[0]))

	fmt.Println(rio)
	inRIO(&rio)
	fmt.Println(rio)
}

func inRIO(rio *[5]int) {

	println(&rio)
	rio[0] = 999999

}
