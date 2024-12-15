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

type Button struct {
	x int64
	y int64
}
type Prize = Button

func main() {
	input_path := "test.txt"
	file, err := os.Open(input_path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	total_tokens := int64(0)

	for scanner.Scan() {
		buttonA := parse_button(scanner.Text())
		scanner.Scan()
		buttonB := parse_button(scanner.Text())
		scanner.Scan()
		prize := parse_prize(scanner.Text())
		scanner.Scan()
		fmt.Println("a:", buttonA, " b:", buttonB, " p", prize)

		tokens_used := check_if_possible(buttonA, buttonB, prize)
		total_tokens += tokens_used
	}

	fmt.Println("result: ", total_tokens)
}

func check_if_possible(a Button, b Button, p Prize) int64 {
	max_a_x := p.x / a.x
	max_a_y := p.y / a.y
	max_a := int64(math.Min(float64(max_a_x), float64(max_a_y)))

	max_b_x := p.x / b.x
	max_b_y := p.y / b.y
	max_b := int64(math.Min(float64(max_b_x), float64(max_b_y)))

	min_cost := int64(math.MaxInt64) // Initialize to a large number
	found := false

	for i := int64(0); i <= max_b; i++ {
		for j := int64(0); j <= max_a; j++ {
			if a.x*j+b.x*i == p.x && a.y*j+b.y*i == p.y {
				cost := int64(3)*j + int64(1)*i
				if cost < min_cost {
					min_cost = cost
				}
				found = true
			}
		}
	}
	if found {
		return min_cost
	}
	return 0
}

func try_combination(a Button, i int64, b Button, j int64, p Prize) bool {
	return a.x*i+b.x*j == p.x && a.y*i+b.y*j == p.y
}

// Prize: X=5233, Y=14652
func parse_prize(line string) Prize {
	cords := strings.Split(line, ":")[1]
	vals := strings.Split(cords, ",")
	x_str := vals[0]
	x_value_str := strings.Split(x_str, "=")[1]
	x_value, _ := strconv.ParseInt(x_value_str, 10, 64)

	y_str := vals[1]
	y_value_str := strings.Split(y_str, "=")[1]
	y_value, _ := strconv.ParseInt(y_value_str, 10, 64)

	return Prize{
		x: x_value + 10000000000000,
		y: y_value + 10000000000000,
	}
}

// Button A: X+28, Y+65
func parse_button(line string) Button {
	cords := strings.Split(line, ":")[1] // X+28, Y+65
	vals := strings.Split(cords, ",")    // ["X+28", "Y+65"]
	x_str := vals[0]                     //x+28
	x_value_str := strings.Split(x_str, "+")[1]
	y_str := vals[1]
	y_value_str := strings.Split(y_str, "+")[1]

	x_value, _ := strconv.ParseInt(x_value_str, 10, 32)
	y_value, _ := strconv.ParseInt(y_value_str, 10, 32)

	return Button{
		x: x_value,
		y: y_value,
	}
}
