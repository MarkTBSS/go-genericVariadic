package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func AddInt(a, b int) int           { return a + b }
func AddFloat(a, b float64) float64 { return a + b }

// Generic function
func AddGeneric[T Number](a, b T) T { return a + b }

// Variadic function
func AddVariadic[T Number](nums ...T) T {
	sum := 0
	for _, num := range nums {
		sum += int(num)
	}
	return T(sum)
}

func main() {
	fmt.Println(AddInt(1, 2))
	fmt.Println(AddFloat(1.1, 2.2))
	fmt.Println(AddGeneric(3, 4))
	fmt.Println(AddGeneric(3.3, 4.4))
	fmt.Println(AddVariadic(5, 6, 7))
	fmt.Println(AddVariadic(8.8, 9.9, 10.10, 11.11))
}
