package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	var amax int
	for _, value := range array {
		if value > amax {
			amax = value
		}
	}
	fmt.Println(amax)
}
