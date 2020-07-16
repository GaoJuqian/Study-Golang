package main

import "fmt"

func main() {
	sum, _ := sum(1, 2, 4, 57, 235, 21)
	fmt.Println(sum)

	test := a(1, 2, b)
	fmt.Println(test)

}

func sum(a ...int) (sum int, err error) {
	for _, v := range a {
		sum += v
	}
	return
}

func a(x, y int, b func(int, int) int) int {
	return b(x, y)
}
func b(x, y int) int {
	return x * y
}
