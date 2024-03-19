package looping

import (
	"fmt"
	"time"
)

func Loopinfinite()  {
	for {
        fmt.Println("Hello Go!")

        time.Sleep(1 * time.Second) // ini adalah fungsi yang dibuat agar program delay selama 1 detik
    }
}