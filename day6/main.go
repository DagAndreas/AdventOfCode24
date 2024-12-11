package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var board []string

	for scanner.Scan() {
		board = append(board, scanner.Text())
	}

	// now we can begin mapping out what we want to happen
	guard_on_board := is_guard_on_board(board)
	for guard_on_board {
		// let's get the position and direction
		for i := 0; i < len(board); i++ {
			// fmt.Println(board[i])
		}
		board = guard_update(board)
		// fmt.Println("Updated board")
		guard_on_board = is_guard_on_board(board)
	}
	fmt.Println("done checking board.")
	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}

	sum := count_x_on_board(board)
	fmt.Println(sum)
}

type Direction int

// 'enum' in go
const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	None
)

func guard_update(board []string) []string {
	left_limit := 0
	top_limit := 0
	bottom_limit := len(board) - 1
	right_limit := len(board[0]) // can assume that the board is symmetrically rect

	for i := 0; i < len(board); i++ {
		row := board[i]

		// guard going up
		pos, direction := guard_is_in_row(row)
		if direction == None {
			continue
		}

		//guard is in row
		// mark current position
		new_row := replace_char_in_string(row, pos, 'X')
		board = replace_line_on_board(board, new_row, i)
		row = board[i]

		// let's mark the next position of guard
		switch direction {
		case UP:
			if check_for_obstacle_ahead(board, i-1, pos) {
				new_row := replace_char_in_string(row, pos+1, '>')
				board = replace_line_on_board(board, new_row, i)
				break
			}
			if i > top_limit {
				new_row := replace_char_in_string(board[i-1], pos, '^') // one row up, same column. icon: ^
				board = replace_line_on_board(board, new_row, i-1)      // swap the new row into the board.

				for j := 0; j < len(board); j++ {
				}
			}
		case DOWN:
			if check_for_obstacle_ahead(board, i+1, pos) {
				new_row := replace_char_in_string(row, pos-1, '<')
				board = replace_line_on_board(board, new_row, i)
				break
			}
			if i < bottom_limit {
				new_row := replace_char_in_string(board[i+1], pos, 'v')
				board = replace_line_on_board(board, new_row, i+1)
			}
		case RIGHT:
			if check_for_obstacle_ahead(board, i, pos+1) {
				// will never be on the bottom-row and walk right. Against rules
				new_row := replace_char_in_string(board[i+1], pos, 'v')
				board = replace_line_on_board(board, new_row, i+1)
				break
			}
			// no obstacle
			if i < right_limit {
				new_row := replace_char_in_string(board[i], pos+1, '>')
				board = replace_line_on_board(board, new_row, i)
			}
		case LEFT:
			if check_for_obstacle_ahead(board, i, pos-1) {
				new_row := replace_char_in_string(board[i-1], pos, '^')
				board = replace_line_on_board(board, new_row, i-1)
				break
			}
			// no obstacle
			// if not out of bounds, write the new position
			if i >= left_limit {
				new_row := replace_char_in_string(board[i], pos-1, '<')
				board = replace_line_on_board(board, new_row, i)
			}
		}
		// break loop
		break
	}

	return board
}

func count_x_on_board(board []string) int {
	count := 0
	for i := 0; i < len(board); i++ {
		count += strings.Count(board[i], "X")
	}
	return count
}

func check_for_obstacle_ahead(board []string, row_index int, column_index int) bool {
	if row_index < 0 || row_index >= len(board) {
		return false
	}
	row := board[row_index]
	if column_index < 0 || column_index >= len(row) {
		return false
	}

	char := string(row[column_index])
	obstacle := "#"
	return char == obstacle
}

func replace_line_on_board(board []string, line string, index int) []string {
	var new_board []string
	new_board = append(new_board, board[:index]...)
	new_board = append(new_board, line)
	if index < len(board) {
		// only if there is more to append at the end
		new_board = append(new_board, board[index+1:]...)
	}

	return new_board
}

func replace_char_in_string(s string, index int, char rune) string {
	if index < 0 || index >= len(s) {
		// out of range. just return string
		// help with out of bound for guard
		return s
	}

	runes := []rune(s)
	runes[index] = char
	return string(runes)
}

func guard_is_in_row(current_row string) (int, Direction) {
	if strings.Contains(current_row, "v") {
		// fmt.Println("guard going down")
		index := strings.Index(current_row, "v")
		return index, DOWN
	}
	if strings.Contains(current_row, "<") {
		// fmt.Println("goard going left")
		index := strings.Index(current_row, "<")
		return index, LEFT
	}
	if strings.Contains(current_row, "^") {
		// fmt.Println("guard going up")
		index := strings.Index(current_row, "^")
		return index, UP
	}
	if strings.Contains(current_row, ">") {
		// fmt.Println("goard going right")
		index := strings.Index(current_row, ">")
		return index, RIGHT
	}

	return -1, None
}

func is_guard_on_board(board []string) bool {
	for i := 0; i < len(board); i++ {
		current_row := board[i]
		if strings.Contains(current_row, "v") {
			// fmt.Println("guard going down")
			return true
		}
		if strings.Contains(current_row, "<") {
			// fmt.Println("goard going left")
			return true
		}
		if strings.Contains(current_row, "^") {
			// fmt.Println("guard going up")
			return true
		}
		if strings.Contains(current_row, ">") {
			// fmt.Println("goard going right")
			return true
		}
	}

	return false
}
