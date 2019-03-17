package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)
type err_t struct {
	when time.Time
	what string
}
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
func printOpts2(board [][]int, i int, j int) {
	fmt.Println("Enter the number to swap with the blank space in order to organize the numbers from 1 to ",cap(board)* cap(board) - 1)
	if i != 0 {
		fmt.Println("type ", board[i-1][j]," to move")
	}
	if i != cap(board)-1 {
		fmt.Println("type ", board[i+1][j]," to move")
	}
	if j != 0 {
		fmt.Println("type ", board[i][j-1]," to move")
	}
	if j != cap(board[i])-1 {
		fmt.Println("type ", board[i][j+1]," to move")
	}
}
func printOpts(board [][]int, i int, j int) {
	fmt.Println("Write the direction to move the blank in order to organize the numbers from 1 to ",cap(board)* cap(board) - 1)
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
func getMotion2(board [][]int, i int, j int) (ii int, jj int){
	ii = i
	jj = j
	reader := bufio.NewReader(os.Stdin)
	valstr, _ := reader.ReadString('\n')
	valstr = strings.ToLower(valstr[:len(valstr)-1])
	val, err := strconv.ParseInt(valstr,10,8)
	if err!= nil {
		fmt.Println(err.Error())
		return
	}
	if i > 0 && int(val) == board[i - 1][j] {
		if i != 0 {
			ii = i - 1
		} else {
			fmt.Println("Cant move Up try again")
		}
	} else {
		if i < cap(board)-1 && int(val) == board[i + 1][j] {
			if i != cap(board)-1 {
				ii = i + 1
			} else {
				fmt.Println("Cant move Down try again")
			}
		} else {
			if j > 0 && int(val) == board[i][j - 1] {
				if j != 0 {
					jj = j - 1
				} else {
					fmt.Println("Cant move Left try again")
				}
			} else {
				if j < cap(board[i]) - 1 && int(val) == board[i][j + 1] {
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
	return
}
func getMotion(board [][]int, i int, j int) (ii int, jj int, err err_t) {
	ii = i
	jj = j

	reader := bufio.NewReader(os.Stdin)
	dir, _ := reader.ReadString('\n')
	dir = strings.ToLower(dir[:len(dir)-1])
	if dir == "up" || dir == "u" {
		if i != 0 {
			ii = i - 1
		} else {
			fmt.Println("Cant move Up try again")
		}
	} else {
		if dir == "down" || dir == "d" {
			if i != cap(board)-1 {
				ii = i + 1
			} else {
				fmt.Println("Cant move Down try again")
			}
		} else {
			if dir == "left" || dir == "l" {
				if j != 0 {
					jj = j - 1
				} else {
					fmt.Println("Cant move Left try again")
				}
			} else {
				if dir == "right" || dir == "r"{
					if j != cap(board[i])-1 {
						jj = j + 1
					} else {
						fmt.Println("Cant move Right try again")
					}
				} else {
					if dir == "exit" {
						return 0, 0, err_t{time.Now(),"exit"}
					}else {
						fmt.Println("bad input, please try again")
					}
				}
			}
		}
	}
	return ii, jj, err
}

func swapBoard(board [][]int, i int, j int, ii int, jj int) {
	temp := board[i][j]
	board[i][j] = board[ii][jj]
	board[ii][jj] = temp

}
func main() {
	fmt.Println("Please enter the size of the board:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	size, err := strconv.ParseInt(input,10,8)
	if err!= nil {
		fmt.Println(err.Error())
		return
	}
	board := createBoard(int(size))
	for ; checkBoard(board) == false; {
		i, j := printBoard(board)
		printOpts2(board, i, j)
		ii, jj:= getMotion2(board, i, j)
		//printOpts(board, i, j)
		//ii, jj, _:= getMotion(board, i, j)
		swapBoard(board, i, j, ii, jj)
	}
	fmt.Println("Well Done yuo won!")
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
