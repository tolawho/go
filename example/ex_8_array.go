package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	var b []int
	fmt.Println(b)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	// Mảng hai chiều
	var twoD [2][3]int
	fmt.Println("emp:", twoD)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	/**
	Khai báo mảng với giá trị mặc định
	 */
	var c = [5]int{1, 2, 3, 4, 6}
	fmt.Println("dcl:", c)
	fmt.Println(c[4]) // lấy giá trị phần tử cuối

}