package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type move_t struct {
	i   int
	j   int
	val int
}

func errChk(errCond bool) (b bool) {
	errStr := []string{"Wait, What?", "Why did you do that?", "what are you? QA?", "Hello, what are you testing me?",
		"Is it on purpose?", "Come on!"}
	if errCond {
		pc, fn, line, _ := runtime.Caller(1)

		fmt.Println("[error]", errStr[rand.Intn(cap(errStr))])
		fmt.Printf("[error] in %s[%s:%d]\n", runtime.FuncForPC(pc).Name(), fn, line)
		b = true
	}
	return
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
func validateBoard(board [][]int, size int) (err int) {
	if errChk(cap(board) != size) {
		return 1
	}
	if errChk(size < 1) {
		return 1
	}
	for i := range board {
		if errChk(cap(board) != cap(board[i])) {
			return 1
		}
	}
	values := make([]bool, size*size)
	for _, row := range board {
		for _, value := range row {
			if errChk(value > size*size || value < 1) {
				fmt.Println("value =", value)
				return 1
			}
			if errChk(values[value-1]) {
				return 1
			}
			values[value-1] = true
		}
	}
	return 0
}
func createBoard(size int) ([][]int, /*err*/ int) {
	values := make([]bool, size*size)
	rand.Seed(int64(time.Now().Nanosecond()))
	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
		for j := 0; j < size; {
			temp := rand.Intn(size * size)
			if values[temp] {
				continue
			}
			values[temp] = true
			board[i][j] = temp + 1
			j++
		}
	}
	return board, 0
}

func updateBoard(board [][]int, dir int) {

}
func getOpts(board [][]int, i int, j int) (move []move_t, err int) {

	if i != 0 {
		move = append(move, move_t{i: i - 1, j: j, val: board[i-1][j]})
	}
	if i != cap(board)-1 {
		move = append(move, move_t{i: i + 1, j: j, val: board[i+1][j]})
	}
	if j != 0 {
		move = append(move, move_t{i: i, j: j - 1, val: board[i][j-1]})
	}
	if j != cap(board[i])-1 {
		move = append(move, move_t{i: i, j: j + 1, val: board[i][j+1]})
	}
	return
}

func printOpts(values []int, size int) (err int) {
	fmt.Println("Enter the number to swap with the blank space in order to organize the numbers from 1 to ", size*size-1)
	for _, value := range values {
		fmt.Println("type ", value, " to move ", value, " to the blank space")
	}
	return 0
}

func getPlayerMove(values []int) (idx int, err int) {
	reader := bufio.NewReader(os.Stdin)
	valueStr, _ := reader.ReadString('\n')
	valueStr = strings.ToLower(valueStr[:len(valueStr)-1])
	val, error := strconv.ParseInt(valueStr, 10, 8)
	if error != nil {
		fmt.Println(error.Error())
		return 0, 1
	}
	for i, value := range values {
		if errChk(int(val) == value) {
			return i, 0
		}
	}
	return 0, 1
}

func swapBoard(board [][]int, i int, j int, move move_t) (err int) {
	if errChk(i >= cap(board) || i < 0) {
		return 1
	}
	if errChk(j >= cap(board[i]) || j < 0) {
		return 1
	}
	if errChk(move.i >= cap(board) || move.i < 0) {
		return 1
	}
	if errChk(move.j >= cap(board[move.i]) || move.j < 0) {
		return 1
	}
	if errChk(move.val != board[move.i][move.j]) {
		return 1
	}
	if errChk(cap(board)*cap(board[i]) != board[i][j]) {
		return 1
	}
	if errChk(int(math.Abs(float64(i-move.i))) != 1 && int(math.Abs(float64(j-move.j))) != 1) {
		return 1
	}
	temp := board[i][j]
	board[i][j] = board[move.i][move.j]
	board[move.i][move.j] = temp
	return 0
}

func play() (err int) {
	fmt.Println("Please enter the size of the board:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	size, error := strconv.ParseInt(input, 10, 8)
	if error != nil {
		fmt.Println(error.Error())
		return -1
	}
	board, err := createBoard(int(size))
	if err == 1 {
		fmt.Println("error")
		return -1
	}
	err = validateBoard(board, int(size))
	if errChk(err == 1) {
		return -1
	}
	plays := 0
	for checkBoard(board) == false {
		i, j := printBoard(board)
		moves, err := getOpts(board, i, j)
		if errChk(err == 1) {
			fmt.Println("No available moves")
			return -1
		}
		var values []int
		for _, move := range moves {
			values = append(values, move.val)
		}
		err = printOpts(values, int(size))
		if errChk(err == 1) {
			return -1
		}
		idx, err := getPlayerMove(values)
		if err == 1 {
			fmt.Println("error: Bad move Option, try again!")
			continue
		}

		err = swapBoard(board, i, j, moves[idx])
		if errChk(err == 1) {
			return -1
		} else {
			plays++
		}
	}
	fmt.Println("Well Done you won in ", plays, "!")
	return 0
}
func test() (err int) {
	sizes := []int{-1, 3, 0, 3, 3, 3, 4, 4, 5}
	expRet := []int{1, 0, 1, 1, 1, 1, 0, 0, 1}
	boards := make([][][]int, cap(sizes))
	boards[0] = nil
	boards[1] = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	boards[2] = nil
	boards[3] = [][]int{{1, 22, 3}, {4, 5, 6}, {7, 8, 9}}
	boards[4] = [][]int{{1, 2, 2}, {4, 5, 6}, {7, 8, 9}}
	boards[5] = [][]int{{-1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	boards[6] = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	boards[7] = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 16}, {13, 14, 15, 12}}
	moves := make([]move_t, 5)
	moves[0] = move_t{i: 0, j: 0, val: 0}
	moves[1] = move_t{i: -1, j: 0, val: 0}
	moves[2] = move_t{i: 3, j: 0, val: 12}
	moves[3] = move_t{i: 0, j: 3, val: 12}
	moves[4] = move_t{i: 3, j: 3, val: 12}
	for idx := range sizes {
		fmt.Println("******************************************************************************")
		fmt.Println("TEST number: ", idx, " Starting")
		ret := validateBoard(boards[idx], sizes[idx])
		if ret != expRet[idx] {
			fmt.Println("TEST number: ", idx, " Failed")
			fmt.Println("******************************************************************************")
			return 1
		} else {
			if expRet[idx] == 0 {
				for plays := 0; checkBoard(boards[idx]) == false; plays++ {
					i, j := printBoard(boards[idx])
					moves_, err := getOpts(boards[idx], i, j)
					if errChk(err == 1) {
						fmt.Println("No available moves")
						fmt.Println("TEST number: ", idx, " Failed")
						fmt.Println("******************************************************************************")
						return -1
					}
					var values []int
					for _, move := range moves_ {
						values = append(values, move.val)
					}
					err = printOpts(values, int(sizes[idx]))
					if errChk(err == 1) {
						fmt.Println("TEST number: ", idx, " Failed")
						fmt.Println("******************************************************************************")
						return -1
					}
					fmt.Println("plays = ", plays)
					err = swapBoard(boards[idx], i, j, moves[plays])
					if err == 1 {
						fmt.Println("Bad move")
						//return -1
					}
				}
			}
		}
		fmt.Println("TEST number: ", idx, " Success")
		fmt.Println("******************************************************************************")
	}
	fmt.Println("Very Good Tests Are Ok")
	return 0
}
func main() {
	err := 0
	for {
		fmt.Println("Please type \"play\" to play or \"test\" to test:")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]

		switch input {
		case "test":
			err = test()
		case "play":
			err = play()
		default:
			fmt.Println("bad input: ", input, " try again")
			break
		}
		if err == 1 {
			fmt.Println("Error")
			os.Exit(err)
		}
	}


	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
