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
	// ans := find_checksum(disk_shifted)
	// fmt.Println("result: ", ans)
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

func left_shift_disk_part2(disk []int) []int {
	// left_pointer := 0
	// right_pointer := len(disk) - 1

	// look at the end and find a chunk of data.

	for i := len(disk) - 1; i > 0; i-- {

		// backwards indexing
		fmt.Println("backwards: ", disk[i])
		cell_value := disk[i]

		empty_cell := -1
		if cell_value == empty_cell {
			fmt.Println("empty cell")
			continue
		}

		// we found a chunk of memory
		// let's count how big the memory is

		// look_for_next_value := true
		memory_size_chunk := 1 // we already know this is one
		for j := i - 1; i > 0; j-- {
			if j == 1 {
				break
			}
			previous_cell := disk[j]
			fmt.Println("checking for chunk: ", previous_cell)
			if previous_cell != cell_value {
				break
			}
			memory_size_chunk++
		}

		fmt.Println("found a chunk of ", cell_value, " that is ", memory_size_chunk, "long")
		first_index_of_memory := i - memory_size_chunk
		fmt.Println("chunk begins at i=", first_index_of_memory)

		// now we need to look forward to find a continous memory addess for it.
		for j := 0; j < first_index_of_memory; j++ {

			not_empty_cell := disk[j] != empty_cell
			if not_empty_cell {
				fmt.Println(j, " is not empty")
				continue
			}

			// we found an empty cell. Let's look ahead and see if it has enough memory for the whole chunk
			// begin on j.
			// look ahead memorysize chunk times
			fmt.Println("Before looking for space")
			found_enough_space := true
			for x := j; x < memory_size_chunk; x++ {
				if disk[x] != empty_cell {
					fmt.Println("Not enough space for the chunk at ", j)
					fmt.Println("Need ", memory_size_chunk, " but found only ", x)
					found_enough_space = false
					break
				}
			}
			fmt.Println("after looking for space")
			if !found_enough_space {
				fmt.Println("not enough space from index ", j)
				continue
			}

			// lets copy the cell value into the next memorychunk spots
			for x := j; x < memory_size_chunk+j; x++ {
				fmt.Println("copying value. Clearing ", disk[x], " over to index ", x, " with value ", cell_value)
				disk[x] = cell_value
			}

			// lets clear the memory chunk
			for x := i - memory_size_chunk; x < i; x-- {
				fmt.Println("x:", x, " : ", i, " mem_size_chunk:", memory_size_chunk)
				fmt.Println("setting value ", disk[x], " to -1 ")
				disk[x] = -1
			}

			// now lets shift i down mem_size
			i -= memory_size_chunk
			break

		}

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
