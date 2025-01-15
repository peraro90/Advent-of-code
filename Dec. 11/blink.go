package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type StoneKey struct {
	rock, blink int
}

func blink(rocks []string, depth int) int {
	mem := map[StoneKey]int{}

	var compute func(rock, depth int) int
	compute = func(rock, depth int) int {
		if depth == 0 {
			return 1
		} else if value, exists := mem[StoneKey{rock, depth}]; exists {
			return value
		}
		var count int
		if rock == 0 {
			count = compute(1, depth-1)
		} else if rockStr := strconv.Itoa(rock); len(rockStr)%2 == 0 {
			left, _ := strconv.Atoi(rockStr[:len(rockStr)/2])
			right, _ := strconv.Atoi(rockStr[len(rockStr)/2:])

			count += compute(left, depth-1) + compute(right, depth-1)
		} else {
			count += compute(rock*2024, depth-1)
		}
		mem[StoneKey{rock, depth}] = count
		return count
	}

	count := 0
	for _, rock := range rocks {
		rock_val, _ := strconv.Atoi(rock)
		count += compute(rock_val, depth)
	}

	return count
}

func main() {
	println("Sorting stones")

	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	stones := strings.Split(lines[0], " ")
	fmt.Println(stones)

	fmt.Println("number of stones after 25 blinks: ", blink(stones, 25))
	fmt.Println("number of stones after 75 blinks: ", blink(stones, 75))

}
