package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coord struct {
	row    int
	column int
}

func main() {
	file_path := "test.txt"
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var board1 []string
	for scanner.Scan() {
		board1 = append(board1, scanner.Text())
	}

	// create the board
	var board [][]rune
	for _, row := range board1 {
		var rune_lst []rune
		for _, c := range row {
			rune_lst = append(rune_lst, c)
		}
		board = append(board, rune_lst)

	}

	fmt.Println("Printing board:")
	for _, row := range board {
		fmt.Println("row =", row)
	}

	result := get_fence_cost(board)
	fmt.Println("result=", result)
}

func get_fence_cost(board [][]rune) int {
	// create a map of visited
	visited_tile := make(map[Coord]bool)
	sum := 0
	for i, row := range board {
		for j, _ := range row {

			key := Coord{
				row:    i,
				column: j,
			}
			_, ok := visited_tile[key]
			if ok {
				fmt.Println("already visited ", key)
				continue
			}
			fmt.Println("visiting coord=", key)

			// visited_tile[key] = true
			var result_sum int

			var found_region map[Coord]bool
			visited_tile, found_region = get_region(board, visited_tile, key)

			result_sum = fence_cost_for_region(found_region)
			sum += result_sum
		}
	}
	return sum
}

func get_region(board [][]rune, visited_tiles map[Coord]bool, cord Coord) (map[Coord]bool, map[Coord]bool) {
	current_region := make(map[Coord]bool)
	fmt.Println("getting region from ", cord)
	region_value := board[cord.row][cord.column]
	current_region, visited_tiles = traverse_from_cord(board, cord, current_region, visited_tiles, region_value)

	for key, _ := range current_region {
		visited_tiles[key] = true
	}

	return visited_tiles, current_region
}

func fence_cost_for_region(region map[Coord]bool) int {
	inner := 0
	perimeter := 0
	for key, _ := range region {
		inner++
		// check all sides if they are in region. If they are not, then inc fence

		perimeter += add_if_boundary_of_region(region, Coord{
			row:    key.row,
			column: key.column - 1,
		})

		perimeter += add_if_boundary_of_region(region, Coord{
			row:    key.row,
			column: key.column + 1,
		})

		perimeter += add_if_boundary_of_region(region, Coord{
			row:    key.row - 1,
			column: key.column,
		})

		perimeter += add_if_boundary_of_region(region, Coord{
			row:    key.row + 1,
			column: key.column,
		})
	}
	fmt.Println("inner =", inner, ", perimeter=", perimeter)
	return inner * perimeter
}

func add_if_boundary_of_region(region map[Coord]bool, cord Coord) int {
	_, ok := region[cord]
	if ok {
		return 0
	}
	return 1
}

func traverse_from_cord(board [][]rune, cord Coord, current_region map[Coord]bool, visited_tiles map[Coord]bool, reg_value rune) (map[Coord]bool, map[Coord]bool) {
	// check if coord is within boundary
	top_limit := 0
	bottom_limit := len(board) - 1
	left_limit := 0
	right_limit := len(board[1]) - 1

	_, ok := visited_tiles[cord]
	if ok {
		fmt.Println("already visited from previous run ", cord)
		return current_region, visited_tiles
	}

	_, ok = current_region[cord]
	if ok {
		fmt.Println("current region already visited this run ")
		return current_region, visited_tiles
	}

	row := cord.row
	if row > bottom_limit || row < top_limit {
		fmt.Println("row is out of scope ", cord)
		return current_region, visited_tiles
	}

	column := cord.column
	if column < left_limit || column > right_limit {
		fmt.Println("column is out of scope")
		return current_region, visited_tiles
	}

	// within boundaries
	cord_value := board[row][column]
	if cord_value != reg_value {
		println("Looking region ", string(reg_value), " but found ", string(cord_value))
		return current_region, visited_tiles
	}

	current_region[cord] = true
	visited_tiles[cord] = true
	left_cord := Coord{
		row:    row,
		column: column - 1,
	}
	current_region, visited_tiles = traverse_from_cord(board, left_cord, current_region, visited_tiles, reg_value)

	right_cord := Coord{
		row:    row,
		column: column + 1,
	}
	current_region, visited_tiles = traverse_from_cord(board, right_cord, current_region, visited_tiles, reg_value)

	up_cord := Coord{
		row:    row - 1,
		column: column,
	}
	current_region, visited_tiles = traverse_from_cord(board, up_cord, current_region, visited_tiles, reg_value)

	down_cord := Coord{
		row:    row + 1,
		column: column,
	}
	current_region, visited_tiles = traverse_from_cord(board, down_cord, current_region, visited_tiles, reg_value)

	return current_region, visited_tiles
}
