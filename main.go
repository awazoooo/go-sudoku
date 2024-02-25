package main

import (
	"fmt"
	"os"
	"strings"
)

type Sudoku struct {
	board       []int
	solved      bool
	width       int
	height      int
	blockWidth  int
	blockHeight int
}

func (s Sudoku) getBoard() []int {
	return s.board
}

func (s Sudoku) getRow(num int) []int {
	width := s.width
	return s.board[width*num : width*(num+1)]
}

func (s Sudoku) getColumn(num int) []int {
	values := []int{}
	for i := 0; i < s.height; i++ {
		values = append(values, s.getRow(i)[num])
	}
	return values
}

func (s Sudoku) getBlock(num int) []int {
	values := []int{}
	rowBlocks := s.width / s.blockWidth
	columnBlocks := s.height / s.blockHeight
	startCellId := (num/rowBlocks)*s.width*columnBlocks + (num%rowBlocks)*rowBlocks
	for y := 0; y < columnBlocks; y++ {
		for x := startCellId; x < startCellId+rowBlocks; x++ {
			values = append(values, s.board[x+y*s.width])
		}
	}
	return values
}

func (s Sudoku) checkSolved() {
	expects = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for y := 0; y < s.blockHeight; y++ {
		// ソートして比較する
		values := s.getRow(y)
		s.width
	}
}

func read(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file")
		panic(err)
	}
	data := make([]byte, 1024)
	count, err := f.Read(data)
	if err != nil {
		fmt.Println("Failed to read data")
		panic(err)
	}
	fmt.Printf("read %d bytes:\n", count)
	//fmt.Println(string(data[:count]))

	// 数値に変換して配列に詰め替える
	arr := []int{}
	for _, line := range strings.Split(string(data[:count]), "\r\n") {
		for _, v := range line {
			arr = append(arr, int(v-'0'))
		}
	}

	s := Sudoku{
		arr,
		false,
		9,
		9,
		3,
		3,
	}
	fmt.Println(s.getRow(0))
	fmt.Println(s.getColumn(0))
	fmt.Println(s.getBlock(4))
	defer f.Close()
}

func main() {
	fmt.Println("Read board")
	read("./board.txt")
}
