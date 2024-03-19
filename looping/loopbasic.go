package looping

import "fmt"

// loop basic
func Loop1basic1()  {
	for i := 0; i <= 5; i++ {
		fmt.Println("Nilai i = ", i)
	}
}

func Loopbasic2()  {
	for i := 50; i > 0; i -= 10 {
		fmt.Println(i)
	}
}