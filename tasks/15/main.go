package main

var justString string

func createHugeString(size int) string {
	return string(make([]rune, size))
}

func getSlice(v string, from, to uint64) string {
	data := []rune(v)
	slice := make([]rune, to)
	
	for i := from; i < to; i++ {
		slice[i] = data[i]
	}
	
	return string(slice)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
