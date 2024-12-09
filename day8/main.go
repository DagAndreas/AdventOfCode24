package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// read file
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

	// process the boards and find all negative nodes
	markboard := find_negative_nodes(board)

	// find how many '#'s in markboard
	sum := 0
	for i := 0; i < len(markboard); i++ {
		for j := 0; j < len(markboard[i]); j++ {
			if markboard[i][j] {
				sum++
			}
		}
	}

	fmt.Println("result is ", sum)

}

func find_negative_nodes(board []string) [][]bool {
	// loop through each element and check if they are not '.'
	// if they are a char, then we look for other characters of the same
	// type and mark the opposite side
	rows := len(board)
	columns := len(board[0])
	markboard := make([][]bool, rows)
	for i := range markboard {
		markboard[i] = make([]bool, columns)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			var char_rune byte
			char_rune = board[i][j]

			dot_ascii := 46
			if int(char_rune) != dot_ascii {
				// mark opposite
				fmt.Println("found antenna: ", i, ",", j)
				markboard = antenna_found(i, j, char_rune, board, markboard)
			}
		}
	}

	for i := range markboard {
		fmt.Println(markboard[i])
	}
	return markboard
}

func antenna_found(row int, column int, antenna_type byte, board []string, markboard [][]bool) [][]bool {
	// iterate the board again and check for '{antenna_type}'

	// limits, valid indexes
	left_limit := 0
	top_limit := 0
	bottom_limit := len(board) - 1
	right_limit := len(board[row]) - 1

	fmt.Println("left_limit", left_limit)
	fmt.Println("rightlimit", right_limit)
	fmt.Println("bottom limit", bottom_limit)
	fmt.Println("top limit", top_limit)

	fmt.Println("len ", len(markboard))
	fmt.Println("len []", len(markboard[row]))

	// iterate and look for more.
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[row]); j++ {
			// check for value
			if i == row && j == column {
				continue
			}

			// we found same type
			if board[i][j] == antenna_type {
				// we know x,y and i,j. Now we can find the difference.
				fmt.Println("found matching antenna at: ", i, ",", j)

				// positive node
				add_neg_node_x := i*2 - row
				add_neg_node_y := j*2 - column
				// check for limits and mark
				fmt.Println("add row ", add_neg_node_x, " bottom ", bottom_limit)
				fmt.Println("add col ", add_neg_node_y, " top ", top_limit)

				within_bounds := add_neg_node_x <= bottom_limit && add_neg_node_x >= top_limit && add_neg_node_y >= left_limit && add_neg_node_y <= right_limit

				for within_bounds {
					fmt.Println("marking ", add_neg_node_x, ",", add_neg_node_y, " as negative pos node")
					markboard[add_neg_node_x][add_neg_node_y] = true
					add_neg_node_x += i - row
					add_neg_node_y += j - column
					fmt.Println("add row ", add_neg_node_x, " bottom ", bottom_limit)
					fmt.Println("add col ", add_neg_node_y, " top ", top_limit)

					within_bounds = add_neg_node_x <= bottom_limit && add_neg_node_x >= top_limit && add_neg_node_y >= left_limit && add_neg_node_y <= right_limit
					if !within_bounds {
						fmt.Println("out of bounds. Breaking")
					}
				}

				markboard[i][j] = true
				markboard[row][column] = true
			}
		}
	}

	return markboard
}

// star 1: 249
// star 2: 905
