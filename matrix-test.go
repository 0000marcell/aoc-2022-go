package main

import (
	"fmt"
)

// func populateRandomValues(size int) [][]int {

	
// 	// for i := 0; i < size; i++ {
// 	// 	for j := 0; j < size; j++ {
// 	// 		m[i] = append(m[i], rand.Intn(10)-rand.Intn(9))
// 	// 	}
// 	// }
// 	// return m
// }

func main() {
  m := make([][]int, 10)
  m[0] = append(m[0], 2)
  m[9] = append(m[9], 99)
  fmt.Println(m[9][0])
	// rand.Seed(time.Now().Unix())
	// var size int
	// fmt.Println("Enter size of the square matrix: ")
	// fmt.Scanln(&size)
	// x1 := populateRandomValues(size)
	// x2 := populateRandomValues(size)

	// fmt.Println("matrix1:", x1)
	// fmt.Println("matrix2:", x2)

	// fmt.Println("ADD: matrix1 + matrix2: ", AddMatrix(x1, x2))
	// fmt.Println("SUB: matrix1 - matrix2: ", SubMatrix(x1, x2))	
}
