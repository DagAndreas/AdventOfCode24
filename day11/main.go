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

func add_to_map(m map[stone]stone, key stone) map[stone]stone {
	// check if value exists
	val, ok := m[key]
	if ok {
		m[key] = val + 1
		return m
	}

	m[key] = 1
	return m

}

func add_mult_to_map(m map[stone]stone, key stone, values stone) map[stone]stone {
	val, ok := m[key]
	if ok {
		m[key] = val + values
	} else {
		m[key] = values
	}
	return m
}

// todo: create a hashmap
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
	// var start_nums []stone
	start_nums := make(map[stone]stone)
	for i := range splt {
		value, _ := strconv.ParseInt(splt[i], 10, 64)
		start_nums = add_to_map(start_nums, uint64(value))
	}

	result := find_all_stones(start_nums, 75)
	fmt.Println("result: ", result)

}

func find_all_stones(stones_map map[stone]stone, iterations int) stone {
	for i := range iterations {
		stones_map = blink(stones_map)
		fmt.Println("starting iteration ", i)
	}

	// fmt.Println(stones_list)
	// sum up the values from the map
	var sum stone
	for _, value := range stones_map {
		sum += value
	}

	return sum
}

func blink(stones_map map[stone]stone) map[stone]stone {
	temp_map := make(map[stone]stone)
	for key, val := range stones_map {

		// rule 1
		if key == 0 {
			temp_map = add_mult_to_map(temp_map, 1, val)
			continue
		}

		amount_of_digits := 1
		digits_value := key
		for digits_value > 9 {
			digits_value = digits_value / 10
			amount_of_digits++
		}

		// rule 2
		is_even_digits := amount_of_digits%2 == 0
		if is_even_digits {
			// split into two.
			// first get the first two numbers
			first_num, second_num := split_even(key, amount_of_digits)
			temp_map = add_mult_to_map(temp_map, first_num, val)
			temp_map = add_mult_to_map(temp_map, second_num, val)
			continue
		}

		if !is_even_digits {
			multiply_val := 2024
			new_key := key * stone(multiply_val)
			temp_map = add_mult_to_map(temp_map, new_key, val)
		}
	}

	return temp_map
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
// star 2: 229682160383225
