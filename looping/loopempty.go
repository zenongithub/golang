package looping

import "fmt"

func Loopempty()  {
	var i = 0

    for {
        if i >= 5 { // condition statement
            break // menghentikan looping jika kondisi 'true'
        }

        fmt.Println("Hello Go!", i + 1)

        i++
    }
}