package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("test.txt")
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
	guard_visited_map := make(map[string]Direction)
	var loop bool
	for is_guard_on_board(board) {
		// let's get the position and direction
		board, guard_visited_map, loop = guard_update(board, guard_visited_map)
		if loop {
			fmt.Println("guard looped board!!!\n ")
			break
		}

	}
	fmt.Println("done checking board.")
	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}

	sum := count_x_on_board(board)
	fmt.Println(sum)
	// we want to iterate over the entire board.

	guard_looping_forever_count := 0
	for i, row := range board {
		for j, char := range row {
			if char == 'X' {
				fmt.Println("replacing on (", i, ",", j, ") with #")

				// replace char
				new_map := make(map[string]Direction)
				new_row := replace_char_in_string(row, j, '#')
				new_board := replace_line_on_board(board, new_row, i)
				var looped bool
				fmt.Println("lets see if the guard is on the board:", is_guard_on_board(new_board))
				for is_guard_on_board(new_board) {
					new_board, new_map, looped = guard_update(new_board, new_map)
					fmt.Println("new map", new_map)

					if looped {
						guard_looping_forever_count++
						fmt.Println("guard DID loop")
						break
					}
				}
				for _, row := range new_board {
					fmt.Println(row)
				}
			} else {
			}
		}
	}
	fmt.Println("there are ", guard_looping_forever_count, "spots to place an obstacle for the guard to make him loop")
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

func guard_update(board []string, guard_visited_map map[string]Direction) ([]string, map[string]Direction, bool) {
	left_limit := 0
	top_limit := 0
	bottom_limit := len(board) - 1
	right_limit := len(board[0]) // can assume that the board is symmetrically rect

	for i := 0; i < len(board); i++ {
		row := board[i]

		pos, direction := guard_is_in_row(row)
		if direction == None {
			fmt.Println("no direction. Not where the guard is. in pos=", pos, ". row=", row, " dir=", direction)
			continue
		} else {
			fmt.Println("guard facing ", direction, " in ", i, ", ", pos)
		}
		// fmt.Println("update in ", i, ", ", pos)

		// let's check if the guard has been here before.
		// create the key

		// check if key in map
		key := strconv.Itoa(i) + "," + strconv.Itoa(pos) + "," + strconv.Itoa(int(direction))
		_, ok := guard_visited_map[key]
		fmt.Println("key=", key, " and ok=", ok)
		if ok {
			return board, guard_visited_map, true
		} else {
			guard_visited_map[key] = direction
		}

		//

		//guard is in row
		// mark current position
		new_row := replace_char_in_string(row, pos, 'X')
		board = replace_line_on_board(board, new_row, i)
		row = board[i]

		// let's mark the next position of guard
		switch direction {
		case UP:
			if check_for_obstacle_ahead(board, i-1, pos) {
				// check for obstacle on the next corner.
				if check_for_obstacle_ahead(board, i, pos+1) {
					// turn 180 degrees back
					new_row := replace_char_in_string(row, pos, 'v')
					board = replace_line_on_board(board, new_row, i)
				} else {
					new_row := replace_char_in_string(row, pos+1, '>')
					board = replace_line_on_board(board, new_row, i)
				}
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
				if check_for_obstacle_ahead(board, i, pos-1) {
					new_row := replace_char_in_string(row, pos, '^')
					board = replace_line_on_board(board, new_row, i)
				} else {
					new_row := replace_char_in_string(row, pos-1, '<')
					board = replace_line_on_board(board, new_row, i)
				}
				break
			}
			if i < bottom_limit {
				new_row := replace_char_in_string(board[i+1], pos, 'v')
				board = replace_line_on_board(board, new_row, i+1)
			}
		case RIGHT:
			if check_for_obstacle_ahead(board, i, pos+1) {
				// will never be on the bottom-row and walk right. Against rules
				if check_for_obstacle_ahead(board, i+1, pos) {
					new_row := replace_char_in_string(row, pos, '<')
					board = replace_line_on_board(board, new_row, i)
				} else {
					new_row := replace_char_in_string(board[i+1], pos, 'v')
					board = replace_line_on_board(board, new_row, i+1)
				}
				break
			}
			// no obstacle
			if i < right_limit {
				new_row := replace_char_in_string(board[i], pos+1, '>')
				board = replace_line_on_board(board, new_row, i)
			}
		case LEFT:
			if check_for_obstacle_ahead(board, i, pos-1) {
				if check_for_obstacle_ahead(board, i-1, pos) {
					new_row := replace_char_in_string(row, pos, '>')
					board = replace_line_on_board(board, new_row, i)
				} else {
					new_row := replace_char_in_string(board[i-1], pos, '^')
					board = replace_line_on_board(board, new_row, i-1)
				}
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

	return board, guard_visited_map, false
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
