package main

import "fmt"

// 指针
func main() {
	a := 1
	test(&a)
	fmt.Println(a)

}
func test(v *int) {
	a := v
	b := &v
	c := **b
	*v += 1
	fmt.Println(a, b, c)
}
