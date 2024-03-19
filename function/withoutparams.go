package function

import "fmt"

// Function without param
func Segitiga()  {
	a := int(7)
	t := int(10)

	luas_segitiga := a * t / 2

	fmt.Println("Luas Segitiga", luas_segitiga)
}

func Call()  {
	inc := func (a int, b int) int {
		return a + b
	}

	result := inc(5 , 10)

	fmt.Println("Hasil Penjumlahan : ", result)
}