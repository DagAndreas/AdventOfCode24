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
			num, _ := strconv.ParseInt(splt[i], 10, 32)
			int_list = append(int_list, int(num))

		}
		top_map = append(top_map, int_list)

		index++
	}
	var paths int
	result, paths := process_map(top_map)
	fmt.Println("result (star1): ", result)
	fmt.Println("paths: (star2)", paths)

}

func process_map(top_map [][]int) (int, int) {

	// first lets iterate and find a '0'
	start_tile := 0

	sum := 0
	met := 0
	for row := range top_map {
		for column := range top_map[row] {
			curr_tile := top_map[row][column]
			if curr_tile != start_tile {
				continue
			}

			results := make(map[string]bool)
			var found int
			results, found = traverse_map_from_pos(top_map, row, column, 0, results)
			met += found

			// convert results to a set
			sum += len(results)
			// count the length of set and add to sum
		}
	}

	return sum, met

}

func traverse_map_from_pos(top_map [][]int, row int, column int, current int, found map[string]bool) (map[string]bool, int) {
	end_value := 9

	if current == end_value {
		// create key
		a := strconv.Itoa(row) + "," + strconv.Itoa(column)
		found[a] = true

		return found, 1
	}

	// limits for indexing
	row_top_limit := 0
	row_bottom_limit := len(top_map) - 1
	column_left_limit := 0
	column_right_limit := len(top_map[0]) - 1

	found_paths := 0

	var met int
	if check_cell(row-1, row_top_limit, true) && top_map[row-1][column] == current+1 {
		found, met = traverse_map_from_pos(top_map, row-1, column, current+1, found)
		found_paths += met
	}

	if check_cell(row+1, row_bottom_limit, false) && top_map[row+1][column] == current+1 {
		found, met = traverse_map_from_pos(top_map, row+1, column, current+1, found)
		found_paths += met
	}

	if check_cell(column-1, column_left_limit, true) && top_map[row][column-1] == current+1 {
		found, met = traverse_map_from_pos(top_map, row, column-1, current+1, found)
		found_paths += met
	}

	if check_cell(column+1, column_right_limit, false) && top_map[row][column+1] == current+1 {
		found, met = traverse_map_from_pos(top_map, row, column+1, current+1, found)
		found_paths += met
	}

	return found, found_paths

}

func check_cell(value int, limit int, lowerlimit bool) bool {
	if lowerlimit {
		return value >= limit
	}

	return value <= limit
}
