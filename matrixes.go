package main

import (
	"fmt"
	"math/rand"
	"time"
)

func AddMatrix(matrix1 [][]int, matrix2 [][]int) [][]int {
	result := make([][]int, len(matrix1))
	for i, a := range matrix1 {
		for j, _ := range a {
			result[i] = append(result[i], matrix1[i][j]+matrix2[i][j])
		}
	}
	return result
}

func SubMatrix(matrix1 [][]int, matrix2 [][]int) [][]int {
	result := make([][]int, len(matrix1))
	for i, a := range matrix1 {
		for j, _ := range a {
			result[i] = append(result[i], matrix1[i][j]-matrix2[i][j])
		}
	}
	return result
}

func populateRandomValues(size int) [][]int {

	m := make([][]int, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			m[i] = append(m[i], rand.Intn(10)-rand.Intn(9))
		}
	}
	return m
}

func main() {
	rand.Seed(time.Now().Unix())
	var size int
	fmt.Println("Enter size of the square matrix: ")
	fmt.Scanln(&size)
	x1 := populateRandomValues(size)
	x2 := populateRandomValues(size)

	fmt.Println("matrix1:", x1)
	fmt.Println("matrix2:", x2)

	fmt.Println("ADD: matrix1 + matrix2: ", AddMatrix(x1, x2))
	fmt.Println("SUB: matrix1 - matrix2: ", SubMatrix(x1, x2))	
}
