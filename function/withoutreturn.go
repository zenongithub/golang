package function

import "fmt"

// Function without return value
func Pembagiannoreturn(a int, b int) {
	if a > b {
		fmt.Println("Hasil Pembagian : ", a/b )
	} else {
		fmt.Println("Hasil Pembagian : ", b/a )
	}
}