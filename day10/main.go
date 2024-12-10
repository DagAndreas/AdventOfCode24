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
	filepath := "input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var top_map [][]int
	index := 0
	for scanner.Scan() {

		nums_text := scanner.Text()
		splt := strings.Split(nums_text, "")
		var int_list []int
		for i := range splt {
			fmt.Println(splt[i])
			num, _ := strconv.ParseInt(splt[i], 10, 32)
			int_list = append(int_list, int(num))

		}
		top_map = append(top_map, int_list)

		index++
	}

	result := process_map(top_map)
	fmt.Println("result: ", result)

}

func process_map(top_map [][]int) int {

	// first lets iterate and find a '0'
	start_tile := 0

	sum := 0
	for row := range top_map {
		for column := range top_map[row] {
			curr_tile := top_map[row][column]
			if curr_tile != start_tile {
				continue
			}

			fmt.Println("starting traversasl from ", row, ",", column)
			results := make(map[string]bool)
			results = traverse_map_from_pos(top_map, row, column, 0, results)
			fmt.Println(results)
			// convert results to a set

			sum += len(results)
			// count the length of set and add to sum
		}
	}

	return sum

}

func traverse_map_from_pos(top_map [][]int, row int, column int, current int, found map[string]bool) map[string]bool {
	end_value := 9
	fmt.Println("traversing from ", row, ", ", column, " with current ", current)

	if current == end_value {
		fmt.Println("found the 9! Returning")

		// a := ""
		// a = a + string(row) + "," + string(column)
		a := strconv.Itoa(row) + "," + strconv.Itoa(column)
		fmt.Println("a=", a)
		found[a] = true
		fmt.Println("added to map")

		return found
	}

	// limits for indexing
	row_top_limit := 0
	row_bottom_limit := len(top_map) - 1
	column_left_limit := 0
	column_right_limit := len(top_map[0]) - 1

	// found_paths := 0

	if check_cell(row-1, row_top_limit, true) && top_map[row-1][column] == current+1 {
		fmt.Println("checking above in ", row, ", ", column)
		found = traverse_map_from_pos(top_map, row-1, column, current+1, found)
	}

	if check_cell(row+1, row_bottom_limit, false) && top_map[row+1][column] == current+1 {
		fmt.Println("checking below in ", row, ", ", column)
		found = traverse_map_from_pos(top_map, row+1, column, current+1, found)
	}

	if check_cell(column-1, column_left_limit, true) && top_map[row][column-1] == current+1 {
		fmt.Println("checking left in ", row, ", ", column)
		found = traverse_map_from_pos(top_map, row, column-1, current+1, found)
	}

	if check_cell(column+1, column_right_limit, false) && top_map[row][column+1] == current+1 {
		fmt.Println("checking right in ", row, ", ", column)
		found = traverse_map_from_pos(top_map, row, column+1, current+1, found)

	}

	fmt.Println("found len found:", len(found), " paths")
	return found

}

func check_cell(value int, limit int, lowerlimit bool) bool {
	if lowerlimit {
		return value >= limit
	}

	return value <= limit
}
