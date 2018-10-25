package main

import "fmt"

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func rotate(a []int, shift int) {
	reverse(a[:shift])
	reverse(a[shift:])
	reverse(a[:])

}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%v\n", a)
	rotate(a, 1)
	fmt.Printf("%v\n", a)
}
