package main
// Проблема старого кода в том, что мы брали слайс от большой строки,
// указатель на базовый массив которого оставался таким же, как и у основной строки.

var justString string

func createHugeString(size int) string {
	return string(make([]rune, size))
}

func getSlice(v string, from, to uint64) string {
	// Создаем слайс рун от базового масссива
	data := []rune(v)

	// Создаем слайс рун нужного нам размера,
	// чтобы не прибегать к постоянному вызову append()
	slice := make([]rune, to)
	
	// Итерируемся по основному слайсу, добавляя в новый необходимые элементы 
	for i := from; i < to; i++ {
		slice[i] = data[i]
	}
	
	// Возвращаем слайс с новым указателем 
	return string(slice)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = getSlice(v, 0, 100)
}

func main() {
	someFunc()
}
