package util

import (
	"fmt"
)

type Matrix struct {
	sizeX int
	sizeY int
	arr [][]int
}

func NewMatrix(X int, Y int) Matrix {
	m := new(Matrix)
	m.sizeX = X
	m.sizeY = Y
	
	m.arr = make([][]int,X)
	
	for i:=0;i<X;i++ {
		m.arr[i] = make([]int,Y)
	}
	return *m
}

func (m Matrix) Print() {
	for i:=0;i<m.sizeX;i++ {
		for j:=0;j<m.sizeY;j++{
			fmt.Print(m.arr[i][j])
			fmt.Print(" ")
		}
		fmt.Println(" ")
	}
}

func (m Matrix) Set(row int, col int, val int) {
	m.arr[row][col] = val
}

func (m Matrix) Get(row int, col int) int {
	return m.arr[row][col]
}
