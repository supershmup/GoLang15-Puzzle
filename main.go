package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func printBoard(board [][]int) (ii int, jj int) {
	fmt.Printf("The board is:\n")

	for i := 0; i < cap(board); i++ {
		fmt.Printf("[")
		for j := 0; j < cap(board[i]); j++ {
			if board[i][j] == cap(board)*cap(board[i]) {
				fmt.Printf("    ")
				ii = i
				jj = j
			} else {
				fmt.Printf("%3d ", board[i][j])
			}
		}
		fmt.Printf("]\n")
	}
	return ii, jj
}

func checkBoard(board [][]int) bool {
	for i := 0; i < cap(board); i++ {
		for j := 0; j < cap(board[i]); j++ {
			if board[i][j] != i*cap(board)+j+1 {
				return false
			}
		}
	}
	return true
}

func createBoard(size int) [][]int {
	vals := make([]bool, size*size)
	rand.Seed(int64(time.Now().Nanosecond()))
	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			board[i][j] = -1
			temp := rand.Intn(size * size)
			if vals[temp] {
				j--
				continue
			}
			vals[temp] = true
			board[i][j] = temp + 1
		}
	}
	return board
}

func updateBoard(board [][]int, dir int) {

}
func printOpts(board [][]int, i int, j int) {
	fmt.Println("Use The keys to move the blank in order to organize the numbers from 1 to 15")
	if i != 0 {
		fmt.Println("press Up to move Up")
	}
	if i != cap(board)-1 {
		fmt.Println("press Down to move Down")
	}
	if j != 0 {
		fmt.Println("press Left to move Left")
	}
	if j != cap(board[i])-1 {
		fmt.Println("press Right to move Right")
	}
}

func getMotion(board [][]int, i int, j int) (ii int, jj int) {
	ii = i
	jj = j
	reader := bufio.NewReader(os.Stdin)
	dir, _ := reader.ReadString('\n')
	dir = strings.ToLower(dir[:len(dir)-1])
	if dir == "up" {
		if i != 0 {
			ii = i - 1
		} else {
			fmt.Println("Cant move Up try again")
		}
	} else {
		if dir == "down" {
			if i != cap(board)-1 {
				ii = i + 1
			} else {
				fmt.Println("Cant move Down try again")
			}
		} else {
			if dir == "left" {
				if j != 0 {
					jj = j - 1
				} else {
					fmt.Println("Cant move Left try again")
				}
			} else {
				if dir == "right" {
					if j != cap(board[i])-1 {
						jj = j + 1
					} else {
						fmt.Println("Cant move Right try again")
					}
				} else {
					fmt.Println("bad input, please try again")
				}
			}
		}
	}
	return ii, jj
}

func swapBoard(board [][]int, i int, j int, ii int, jj int) {
	temp := board[i][j]
	board[i][j] = board[ii][jj]
	board[ii][jj] = temp

}
func main() {
	board := createBoard(4)
	for ; checkBoard(board) == false; {
		i, j := printBoard(board)
		printOpts(board, i, j)
		ii, jj := getMotion(board, i, j)
		swapBoard(board, i, j, ii, jj)
	}
}
