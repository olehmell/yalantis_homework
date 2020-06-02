package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func genarate2dArray() (res [][]int) {
	length := 5
	max := 50

	for i := 0; i < length; i++ {
		tmp := []int{}
		for j := 0; j < length; j++ {
			tmp = append(tmp, rand.Intn(max))
		}
		res = append(res, tmp)
	}
	return
}

func TestKoans(t *testing.T) {
	matrix := genarate2dArray()
	fmt.Println("Input:")
	for i := range matrix {
		fmt.Println(matrix[i])
	}
	res := arrayFromMatrix(&matrix)
	fmt.Printf("\nOutput: %v", res)
}
