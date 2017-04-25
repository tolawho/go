package main

import "fmt"

func main() {
	i:= 1
	for i < 10 {
		fmt.Println("I", i)
		i++
	}

	for j := 1; j <= 10; j++ {
		fmt.Println("J:", j)
	}
	a := 0
	for {
		fmt.Println("A: ", a)
		a++
		if a == 10 {
			break
		}
	}

	for n:=0; n<=10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("N:", n)
	}
}
