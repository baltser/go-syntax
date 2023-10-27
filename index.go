package main

import (
	"fmt"
)

func small(arr []int) int {
	out := arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] < out {
			out = arr[i]
		}
	}
	return out
}
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	Ints := []int{1, 2, 3}
	Strings := []string{"one", "two", "three"}
	Print(Ints)
	Print(Strings)
	/****************/
	arr := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	a := make([]int, 3, 9)

	x := [6]string{"a", "b", "c", "d", "e", "f"}

	arr1 := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	fmt.Println(arr[4])
	fmt.Println(len(a))
	fmt.Println(x[2:5])

	fmt.Println(small(arr1))

}
