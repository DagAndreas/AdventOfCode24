package main

import (
	"bufio"
	"fmt"
	"log"
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

// fuck Gauss-Jordan. We ballin with Cramer's Rule
// so to Professor Dave Explains
func check_if_possible(a Button, b Button, p Prize) int64 {
	// det(A) = ad - bc for any 2x2 matrix.
	// given the matrix:
	// [a.x, a.y
	//	b.x, b.y
	//]

	// x1 + 3x2 = 5
	// 2x1 + 2x2 = 6

	// a.x + b.x = p.x
	// a.y + b.y = p.y

	// |A|
	determanent := a.x*b.y - a.y*b.x
	if determanent == 0 {
		return 0
	}

	// now to set up the different matrixes we use with det

	// matrix on top:
	// |A1| = p.x, b.x
	//		  p.y, b.y
	determanent_of_b := p.x*b.y - p.y*b.x

	// solve x1 = |A1|
	//           ------
	//            |A|
	x1 := determanent_of_b / determanent
	if determanent_of_b%determanent != 0 {
		return 0 // no solution
	}

	// now we replace the second column, because we are solving for the second variable.
	// |A2| = a.x, p.x
	//	      a.y, p.y
	determanent_of_a := a.x*p.y - a.y*p.x
	x2 := determanent_of_a / determanent
	if determanent_of_a%determanent != 0 {
		return 0
	}

	// solution, x1 a buttons and x2 b buttons.
	cost := int64(x1*3 + x2*1)
	if cost <= 0 {
		return 0
	}

	return cost
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
