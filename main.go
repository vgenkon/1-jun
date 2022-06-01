package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 5, 6, 7}
	slice = append(slice[:4], slice[5:]...)
	fmt.Println(slice)
}
