package looping

import "fmt"

// break
func Loopcondition1()  {
	for i := 0; i < 10; i++ {
		if i == 8 {
            break
        }
        fmt.Println("Hello Go!", i+1)
	}
}

// continue
func Loopcondition2()  {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
            continue
        }
        fmt.Println("Hello Go!", i+1)
	}
}

func Loopcondition3() {
    i := 0 
    for i < 5 {
        fmt.Println("Hello Go!", i+ 1)
        i++
    }
}