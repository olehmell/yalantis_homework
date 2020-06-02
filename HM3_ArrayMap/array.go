package main

import "fmt"

func arrayFromMatrix(matrix *[][]int) (res []int) {
	lenght := len(*matrix)
	x, y := -1, 0
	add := lenght
	dir := 1
	for add != 0 {
		for i := 0; i != add; i++ {
			x = x + dir
			res = append(res, (*matrix)[y][x])
		}
		add--
		for i := 0; i != add; i++ {
			y = y + dir
			res = append(res, (*matrix)[y][x])
		}
		dir = -dir
	}
	return
}

func main() {
	matrix0 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix1 := [][]int{{1, 2, 3, 1}, {4, 5, 6, 4}, {7, 8, 9, 7}, {7, 8, 9, 7}}
	fmt.Println(arrayFromMatrix(&matrix1))
	fmt.Println(arrayFromMatrix(&matrix0))
}
