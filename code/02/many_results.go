package main

import "fmt"

func main() {
	fmt.Println("Резултатът от sumAndCount(2, 3, 4, 5) е")
	fmt.Println(sumAndCount(2, 3, 4, 5))
}

func sumAndCount(args ...int) (int, int) {
	result := 0
	for _, v := range args {
		result += v
	}

	return result, len(args)
}
