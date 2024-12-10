package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := "test.txt"
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
		fmt.Println("backwards: ", disk[i], " i=", i)
		cell_value := disk[i]

		if cell_value >= biggest_or_last_moved {
			fmt.Println("Cellvalue is the biggest (", biggest_or_last_moved, "). Let's not move it again.")
			continue
		}

		empty_cell := -1
		if cell_value == empty_cell {
			fmt.Println("empty cell")
			continue
		}

		// we found a chunk of memory
		// let's count how big the memory is

		// look_for_next_value := true
		chunk_sizes := 0
		for j := i; j > 0; j-- {
			fmt.Println("checking for size. j=", j)
			if j == 0 && disk[j] != empty_cell {
				fmt.Println("j reached bottom. breaking")
				return disk
			}

			previous_cell := disk[j]
			fmt.Println("checking for chunk: ", previous_cell)
			fmt.Println("prev:", previous_cell, "vs cell:", cell_value)
			if previous_cell != cell_value {
				break
			}

			fmt.Println("increasing chunksize to : ", chunk_sizes+1)
			chunk_sizes++
		}

		fmt.Println("found a chunk of ", cell_value, " that is ", chunk_sizes, "long")
		first_index_of_memory := i - chunk_sizes + 1
		fmt.Println("chunk begins at i=", first_index_of_memory)

		// now we need to look forward to find a continous memory addess for it.
		for j := 0; j < first_index_of_memory; j++ {
			fmt.Println("starting from the bottom. j=", j)
			not_empty_cell := disk[j] != empty_cell
			if not_empty_cell {
				fmt.Println(disk[j], " is not empty")
				continue
			}

			// we found an empty cell. Let's look ahead and see if it has enough memory for the whole chunk
			// begin on j.
			// look ahead memorysize chunk times
			fmt.Println("Before looking for space")
			found_enough_space := true
			fmt.Println("checking for enough empty cells from ", j)
			for x := j; x < chunk_sizes; x++ {
				fmt.Println("x=", x)
				if disk[x] != empty_cell {
					fmt.Println("Not enough space for the chunk at ", j)
					fmt.Println("Need ", chunk_sizes, " but found only ", x)
					found_enough_space = false
					break
				}
				fmt.Println("nice. x was empty. Lets check next if we need more. x:", x, " and chunk_indexes: ", chunk_sizes)

			}
			fmt.Println("after looking for space")
			if !found_enough_space {
				fmt.Println("not enough space from index ", j)
				continue
			}

			// lets clear the memory chunk
			fmt.Println("first index: ", first_index_of_memory)
			fmt.Println("i: ", i)
			for x := first_index_of_memory; x <= i; x++ {

				fmt.Println("x:", x, " : ", i, " mem_size_chunk:", chunk_sizes)
				fmt.Println("setting value ", disk[x], " to -1 ")
				disk[x] = -1
				fmt.Println("disk x =", disk[x])
			}

			// now lets shift i down mem_size
			// i -= chunk_sizes

			// lets copy the cell value into the next memorychunk spots
			for x := j; x < chunk_sizes+j; x++ {
				fmt.Println("copying value. Clearing ", disk[x], " over to index ", x, " with value ", cell_value)
				disk[x] = cell_value
			}
			fmt.Println("")
			i = first_index_of_memory + 1

			biggest_or_last_moved = cell_value //update to stop moving smaller chunks of the large one
			break
		}
		// TODO: understand how to always count down how many chunks found

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
