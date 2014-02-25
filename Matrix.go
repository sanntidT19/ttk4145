package main

import (
	"fmt"
)

type Matrix struct {
	size int
	arr [][]int
}

func NewMatrix(n int) Matrix {
	m := new(Matrix)
	m.size = n
	
	m.arr = make([][]int,n)
	
	for i:=0;i<n;i++ {
		m.arr[i] = make([]int,n)
	}
	return *m
}

func (m Matrix) Print() {
	for i:=0;i<m.size;i++ {
		for j:=0;j<m.size;j++{
			fmt.Print(m.arr[i][j])
			fmt.Print(" ")
		}
		fmt.Println(" ")
	}
}

func (m Matrix) Set(row int, col int, val int) {
	m.arr[row][col] = val
}

func main() {
	N := 4
	m := NewMatrix(N)
	
	
	m.Set(0,1,7)
	
	m.Print()



}


