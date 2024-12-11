package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stone = uint64

func main() {
	path_name := "test.txt"
	file, err := os.Open(path_name)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var input string
	for scanner.Scan() {
		input = scanner.Text()
	}

	// lets create the input into a list of integers
	splt := strings.Split(input, " ")
	var start_nums []stone
	for i := range splt {
		value, _ := strconv.ParseInt(splt[i], 10, 64)
		start_nums = append(start_nums, uint64(value))
	}

	result := find_all_stones(start_nums, 25)
	fmt.Println("result: ", result)
}

func find_all_stones(stones_list []stone, iterations int) int {
	for i := range iterations {
		stones_list = blink(stones_list)

		largest := stone(0)
		for j := range stones_list {
			if stones_list[j] > largest {
				largest = stones_list[j]
			}
		}
		fmt.Println("largest stone for iteration ", i, " is ", largest)
	}
	fmt.Println(stones_list)
	return len(stones_list)
}

func blink(stones_list []stone) []stone {
	var temp_lst []stone
	for i := range stones_list {

		curr_value := stones_list[i]

		// rule 1
		if curr_value == 0 {
			temp_lst = append(temp_lst, 1)
			continue
		}

		amount_of_digits := 1
		digits_value := stone(stones_list[i])
		for digits_value > 9 {
			digits_value = digits_value / 10
			amount_of_digits++
		}

		// rule 2
		is_even_digits := amount_of_digits%2 == 0
		if is_even_digits {
			// split into two.
			// first get the first two numbers
			first_num, second_num := split_even(curr_value, amount_of_digits)
			temp_lst = append(temp_lst, first_num)
			temp_lst = append(temp_lst, second_num)
			continue
		}

		if !is_even_digits {
			multiply_val := 2024
			new_val := curr_value * stone(multiply_val)
			temp_lst = append(temp_lst, new_val)
		}
	}

	return temp_lst
}

func split_even(value stone, digits int) (stone, stone) {
	var first_num stone
	var second_num stone

	// lets gets the first value
	// first_float := float64(value)
	first_float := stone(value)
	for range digits / 2 {

		first_float = first_float / 10

	}
	first_num = stone(first_float)

	// now to find the second value.
	// str_value := strconv.Itoa(int(value))
	str_value := strconv.FormatUint(value, 10)

	second_num_builder := int64(0)
	for i := digits / 2; i < len(str_value); i++ {
		str_val := string(str_value[i])
		current_to_int, _ := strconv.ParseInt(str_val, 10, 64)
		second_num_builder = second_num_builder*10 + current_to_int

	}
	second_num = stone(second_num_builder)

	return first_num, second_num
}

// star 1: 193899
