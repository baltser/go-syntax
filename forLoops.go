package main

import "fmt"

func main() {

	//традиционный цикл
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	//идиоматический Go
	i := 0
	for ok := true; ok; ok = i != 10 {
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	//цикл for, исползуемый как цикл while
	i1 := 0
	for {
		if i1 == 10 {
			break
		}
		fmt.Print(i1*i1, " ")
		i1++
	}
	fmt.Println()

	// это срез, но range работет и с массивами
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index: ", i, " value: ", v)
	}

	//пример с массивом
	x := [5]float64{98, 93, 77, 82, 83}

	var total float64
	for _, value := range x {
		total += value
	}
	fmt.Println(total/float64(len(x)), "Среднеарихметическое")

	x1 := make(map[string]int)
	x1["key"] = 10
	fmt.Println(x1)
}
