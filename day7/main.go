package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int64
	sum = 0
	for scanner.Scan() {
		sum += is_valid_equation(scanner.Text())
	}

	fmt.Println("final sum is ", sum)
}

func is_valid_equation(line string) int64 {
	res := strings.Split(line, ":")
	target, _ := strconv.ParseInt(res[0], 10, 64)

	num_str := strings.Split(res[1], " ")
	var nums []int64
	for i := 0; i < len(num_str); i++ {
		target, _ := strconv.ParseInt(num_str[i], 10, 64)
		nums = append(nums, target)

	}

	if recursive_check(nums, target, 0) {
		return target
	}
	return 0
}

func get_digits_of_number(num int64) int {
	if num == 0 {
		return 0
	}
	if num < 0 {
		num = -num
	}

	digits := int(math.Log10(float64(num)) + 1)
	// fmt.Printf("there are %d digits in %d\n", digits, num)
	return digits
}

func concat_two_ints(a int64, b int64) int64 {
	exp := get_digits_of_number(b)
	ten_pushed := int64(math.Pow10(exp))
	// fmt.Printf("i need to times with %d\n", ten_pushed)

	res := a*int64(ten_pushed) + b
	// fmt.Println(a, " concated with ", b, " becomes ", res)
	return res
}

func recursive_check(nums []int64, target int64, sum int64) bool {
	if len(nums) == 1 {
		// only return self. No operation
		var ans bool
		ans = sum*nums[0] == target || sum+nums[0] == target ||
			// part 2:
			concat_two_ints(sum, nums[0]) == target
		return ans
	}

	plus_rec := recursive_check(nums[1:], target, sum+nums[0])
	mult_rec := plus_rec || recursive_check(nums[1:], target, sum*nums[0])
	//part 2:
	concatrec := mult_rec || recursive_check(nums[1:], target, concat_two_ints(sum, nums[0]))
	return concatrec

}

// star 1: 3245122495150
// star 2: 105517128211543
