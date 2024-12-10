package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := "input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	disk_ints := format_disk_map(scanner.Text())
	fmt.Println("disk_ints: ", disk_ints)

	disk_shifted := left_shift_disk_part2(disk_ints)
	fmt.Println(disk_shifted)
	ans := find_checksum(disk_shifted)
	fmt.Println("result: ", ans)
}

func find_checksum(disk []int) int64 {
	var sum int64
	for i := range len(disk) {
		if disk[i] == -1 {
			continue
		}

		// not empty
		sum += int64(disk[i] * i)
	}

	return sum
}

func left_shift_disk_part1(disk []int) []int {
	left_pointer := 0
	right_pointer := len(disk) - 1

	for left_pointer < right_pointer {
		left_val := disk[left_pointer]
		right_val := disk[right_pointer]

		if left_val != -1 {
			left_pointer++
			continue
		}

		// left pointer is .
		if right_val == -1 {
			right_pointer--
			continue
		}

		// left pointer is . and rightpointer is a an id
		disk[left_pointer] = disk[right_pointer]
		disk[right_pointer] = -1 // avoid dataloss
	}

	return disk
}

func find_biggest_id_start(disk []int) int {
	empty_cell := -1
	biggest := 0
	for i := len(disk) - 1; i > 0; i-- {
		if disk[i] != empty_cell {
			biggest = disk[i]
			break
		}
	}
	return biggest
}

func left_shift_disk_part2(disk []int) []int {
	// left_pointer := 0
	// right_pointer := len(disk) - 1

	// look at the end and find a chunk of data.

	// to stop pushing the bigger parts backwards. we need to keep track of the largest id we've parsed and stop moving any id under it.
	// we can do this because the id's are in asc order.
	biggest_or_last_moved := find_biggest_id_start(disk) + 1

	for i := len(disk) - 1; i > 0; i-- {

		// backwards indexing
		cell_value := disk[i]

		if cell_value >= biggest_or_last_moved {
			continue
		}

		empty_cell := -1
		if cell_value == empty_cell {
			continue
		}

		// we found a chunk of memory
		// let's count how big the memory is

		// look_for_next_value := true
		chunk_size := 1
		for j := i - 1; j > 0; j-- {
			if j == 0 && disk[j] != empty_cell {
				return disk
			}

			previous_cell := disk[j]
			if previous_cell != cell_value {
				break
			}

			chunk_size++
		}

		// now that we know the chunk size.
		free_space_index_start := -1
		for j := 0; j < i-chunk_size+1; j++ {
			if disk[j] != empty_cell {
				continue
			}

			found_enough_space := true
			for free_space_count := 0; free_space_count < chunk_size; free_space_count++ {
				// should count if chunksize = 1
				if disk[j+free_space_count] == empty_cell {
					found_enough_space = found_enough_space && true
				} else {
					found_enough_space = false
					biggest_or_last_moved = cell_value
					break
				}

			}

			if found_enough_space {
				free_space_index_start = j
				break
			}
		}

		if free_space_index_start == -1 {
			continue
		}

		// we found the free space index for the entire chunk
		for j := 0; j < chunk_size; j++ {
			// set the first value
			disk[free_space_index_start+j] = cell_value
			disk[i-j] = -1

		}

		biggest_or_last_moved = cell_value

	}

	return disk
}

func format_disk_map(s string) []int {
	even := true
	ascii_ofset := 48
	id := 0
	var intslice []int

	for i := range len(s) {
		value := s[i] - byte(ascii_ofset)
		// continue
		if value == 0 {
			even = !even
			continue
		}

		if even {
			for range value {
				intslice = append(intslice, id)
			}
			id++
		} else {
			for range value {
				intslice = append(intslice, -1) // -1 will be my symbol for '.'
			}
		}
		even = !even

	}
	return intslice
}
