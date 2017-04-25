package main

import "fmt"

func main() {
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			fmt.Println(j, "is even")
		} else {
			fmt.Println(j, "is odd")
		}

		if j%4 == 0 {
			fmt.Println(j, "is devisible by 4")
		}

		if j > 10 {
			fmt.Println(j, "is negative and has miltiple digits")
		} else {
			fmt.Println(j, "is negative and has 1 digit")
		}
	}
}
