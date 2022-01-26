package main

import "fmt"

func filter(source []int, condition func(int) bool) []int {
	var result []int

	for _, value := range source {
		isValid := condition(value)

		if isValid {
			result = append(result, value)
		}
	}

	return result
}

func isPar(value int) bool {
	mod := value % 2
	return mod == 0
}

func addValue(reference *int, value int) {
	*reference += value
}

func main() {
	values := []int {1, 9, 3, 5, 10, 30, 2, 28, 3, 17, 98, 100}
	parValues := filter(values, isPar)

	var value int = 1
	var valuePointer *int = nil

	valuePointer = &value

	fmt.Println(parValues)

	addValue(valuePointer, 99)

	fmt.Println("Valor:", *valuePointer)
}
