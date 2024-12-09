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

	disk_shifted := left_shift_disk(disk_ints)

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

func left_shift_disk(disk []int) []int {
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
			println("id is now", id)
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
