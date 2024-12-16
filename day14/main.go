package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Guard struct {
	pos_x int
	pos_y int
	mov_x int
	mov_y int
}

func main() {
	input_path := "test.txt"
	file, err := os.Open(input_path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var guard_list []Guard
	for scanner.Scan() {
		guard_list = append(guard_list, parse_guard(scanner.Text()))
	}
	fmt.Println("guards before sim", guard_list)
	guard_list = run_simulation(guard_list)
	fmt.Println("guards after sim ", guard_list)

	result := calculate_quadrant_sum(guard_list)
	fmt.Println("result =", result)

}

func calculate_quadrant_sum(guards []Guard) int64 {
	fmt.Println("q1")
	q1 := count_inside_quadrant(0, VERTICAL_LENGTH/2-1, 0, HORISONTAL_LENGTH/2-1, guards)
	fmt.Println("q2")
	q2 := count_inside_quadrant(0, VERTICAL_LENGTH/2-1, HORISONTAL_LENGTH/2+1, HORISONTAL_LENGTH, guards)
	fmt.Println("q3")
	q3 := count_inside_quadrant(VERTICAL_LENGTH/2+1, VERTICAL_LENGTH, 0, HORISONTAL_LENGTH/2-1, guards)
	fmt.Println("q4")
	q4 := count_inside_quadrant(VERTICAL_LENGTH/2+1, VERTICAL_LENGTH, HORISONTAL_LENGTH/2+1, HORISONTAL_LENGTH, guards)

	fmt.Println("q1:", q1, " q2:", q2, " q3", q3, " q4:", q4)
	return int64(q1) * int64(q2) * int64(q3) * int64(q4)

}

func count_inside_quadrant(top int, bot int, left int, right int, guards []Guard) int {
	fmt.Println("checking inside q:")
	fmt.Println("top:", top, " bot:", bot, " left:", left, " right:", right)
	total := 0
	for _, guard := range guards {
		fmt.Println("g posx", guard.pos_x)
		if !(guard.pos_x >= left && guard.pos_x <= right) {
			fmt.Println("Guard not inside x")
			continue
		}
		fmt.Println("guard inside x")
		fmt.Println("g posy:", guard.pos_y)
		if !(guard.pos_y >= top && guard.pos_y <= bot) {
			fmt.Println("posy>=left =", guard.pos_y >= left)
			fmt.Println("guard.pos_y <= right", guard.pos_y <= right)
			fmt.Println("guard.pos_y >= left && guard.pos_y <= right", guard.pos_y >= left && guard.pos_y <= right)

			fmt.Println("guard not inside y")
			continue
		}
		fmt.Println("guard inside this q.")
		total++
	}

	return total
}

func run_simulation(guards []Guard) []Guard {
	runs := 100
	for range runs {
		for i, guard := range guards {
			guards[i] = update_guard(guard)
		}
	}
	return guards
}

const HORISONTAL_LENGTH = 101 // 0 - 100 = 101 digits
const VERTICAL_LENGTH = 103   // 0-102 = 103 positions
// trying 0, 1, 2. Expecting about 1/3 in above, 1/3 below, and 1/3 gone. Not what i am seeing though.

func update_guard(guard Guard) Guard {
	guard.pos_x = (guard.pos_x + guard.mov_x + HORISONTAL_LENGTH) % HORISONTAL_LENGTH
	guard.pos_y = (guard.pos_y + guard.mov_y + VERTICAL_LENGTH) % VERTICAL_LENGTH
	return guard
}

func parse_guard(line string) Guard {
	// p=0,4 v=3,-3
	middle_split := strings.Split(line, " ")
	p_position := strings.Split(middle_split[0], "=")
	p_position = strings.Split(p_position[1], ",")
	px, _ := strconv.ParseInt(p_position[0], 10, 32)
	py, _ := strconv.ParseInt(p_position[1], 10, 32)

	v_position := strings.Split(middle_split[1], "=")
	v_position = strings.Split(v_position[1], ",")
	vx, _ := strconv.ParseInt(v_position[0], 10, 32)
	vy, _ := strconv.ParseInt(v_position[1], 10, 32)

	fmt.Println("posx", px)
	fmt.Println("posy", py)
	fmt.Println("vosx", vx)
	fmt.Println("vosy", vy)
	return Guard{
		pos_x: int(px),
		pos_y: int(py),
		mov_x: int(vx),
		mov_y: int(vy),
	}

}

// star1 ans > 215175360 > 213992460
