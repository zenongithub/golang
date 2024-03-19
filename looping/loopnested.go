package looping

import "fmt"

func Loopnested()  {
	for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            fmt.Printf("Hello Go! i: %d, j: %d\n", i, j)
        }
    }
}