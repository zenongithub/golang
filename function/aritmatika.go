package function

// Function with return value
func Penjumlahan(a int, b int) int {
	return a + b
}

func Pengurangan(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func Perkalian(a int, b int) int {
	return a * b
}

func Pembagian(a int, b int) int {
	if a > b {
		return a / b
	} else {
		return b / a
	}
}
