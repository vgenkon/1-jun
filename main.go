package main

import "fmt"

func main() {
	var s1 []int
	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("s1", s1, len(s1), cap(s1))
	s2 := []int{10, 20, 30}
	fmt.Println("s2", s2, len(s2), cap(s2))
	s2 = append(s2, s1...)
	fmt.Println("s2", s2, len(s2), cap(s2))
	var s3 [][]int
	s3 = append(s3, s1, s2)
	fmt.Println("s3 [][]", s3, len(s3), cap(s3))
	s4 := make([]int, 5, 5)
	fmt.Println("s4", s4, len(s4), cap(s4))
	s4 = append(s4, 11)
	fmt.Println("s4.2", s4, len(s4), cap(s4))
	s5 := make([]int, 0, 5)
	fmt.Println("s5", s5, len(s5), cap(s5))
	s6 := s4
	s6[1] = 777
	fmt.Println("s6 s4", s6, s4, len(s6), cap(s6), len(s4), cap(s4))
	s4 = append(s4, []int{1, 2, 3, 4, 5, 6}...)
	s4[1] = 888
	fmt.Println("s6 s4", s6, s4, len(s6), cap(s6), len(s4), cap(s4))
	s7 := make([]int, len(s6), len(s6))
	copy(s7, s6)
	fmt.Println("s7", s7, len(s7), cap(s7))
	s8 := append(s4[:2], s1[5:]...)
	fmt.Println("s8", s8, len(s8), cap(s8))
	arr := [3]int{1, 2, 3}
	s9 := arr[:]
	arr[1] = 555
	fmt.Println("s9", s9, len(s9), cap(s9))

}
