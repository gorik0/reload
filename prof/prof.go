package main

func main() {

	for i := 0; i < 10; i++ {

	}

}

func mapString() {
	ma := make(map[int]string, 1)
	for i := range int(1000000) {
		//stri := rand.Intn(60000)
		ma[i] = string(rune(i))
	}
}

func mapInt() {
	ma := make(map[int]int, 1)
	for i := range int(1000000) {
		//stri := rand.Intn(60000)
		ma[i] = i
	}
}
