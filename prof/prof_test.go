package prof

import "testing"

//var hardString = makeString()

func makeString() string {
	a := "egsd"
	for i := 0; i < 1000000; i++ {
		a = a + "dfdsf  d"
		if i == 20000 {
			a += "egor"
		}
	}
	return a

}
func Benchmark___Reg_Compile(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Reg_Compile("egor", "find in string name egor so on ")
	}
}
func Benchmark___Reg_NoCompile(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Reg_NoCompile("egor", "find in string name egor so on ")
	}
}
func Benchmark___No_Reg(b *testing.B) {

	for i := 0; i < b.N; i++ {
		No_Reg("egor", "find in string name egor so on ")
	}
}
