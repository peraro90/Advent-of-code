package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	up    = iota
	down  = iota
	right = iota
	left  = iota
)

func main() {
	println("avoiding time paradoxes")

	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fresh_map := make([]string, len(lines))
	copy(fresh_map, lines)

	start_pos := []int{0, 0}
	for y, line := range lines {
		for x, letter := range line {
			if string(letter) == "^" {
				start_pos = []int{y, x}
			}
		}
	}

	position := make([]int, len(start_pos))
	copy(position, start_pos)
	//new_position := start_pos
	direction := up

	for {
		temp_string := []byte(lines[position[0]])
		temp_string[position[1]] = 'X'
		lines[position[0]] = string(temp_string)

		if direction == up {
			if position[0] == 0 {
				fmt.Println("Moving of the board upwards, position: ", position[1], ",", position[0])
				break
			}
			if string(lines[position[0]-1][position[1]]) == "#" {
				direction = right
				continue
			}
			position[0] = position[0] - 1
		}

		if direction == down {
			if position[0] == len(lines)-1 {
				fmt.Println("Moving of the board downwards, position: ", position[1], ",", position[0])
				break
			}
			if string(lines[position[0]+1][position[1]]) == "#" {
				direction = left
				continue
			}
			position[0] = position[0] + 1
		}

		if direction == right {
			if position[1] == len(lines[0])-1 {
				fmt.Println("Moving of the board rightwards, position: ", position[1], ",", position[0])
				break
			}
			if string(lines[position[0]][position[1]+1]) == "#" {
				direction = down
				continue
			}
			position[1] = position[1] + 1
		}

		if direction == left {
			if position[1] == 0 {
				fmt.Println("Moving of the board leftwards, position: ", position[1], ",", position[0])
				break
			}
			if string(lines[position[0]][position[1]-1]) == "#" {
				direction = up
				continue
			}
			position[1] = position[1] - 1
		}

	}
	sum := 0
	var possible_obstacle_postitions [][]int
	for y, line := range lines {
		for x, letter := range line {
			if string(letter) == "X" {
				sum += 1
				possible_obstacle_postitions = append(possible_obstacle_postitions, []int{y, x})
			}
		}
	}

	infinite_loops := 0
	for _, obs_pos := range possible_obstacle_postitions {
		if start_pos[0] == obs_pos[0] && start_pos[1] == obs_pos[1] {
			continue
		}
		direction = up
		obstacle_map := make([]string, len(lines))
		copy(obstacle_map, fresh_map)
		fmt.Println("Putting object in area ", obs_pos[0], ",", obs_pos[1])
		copy(position, start_pos)
		temp_string := []byte(obstacle_map[obs_pos[0]])
		temp_string[obs_pos[1]] = '#'
		obstacle_map[obs_pos[0]] = string(temp_string)

		for i := 0; ; i++ {
			temp_byte := []byte(obstacle_map[position[0]])
			switch d := direction; d {
			case up:
				temp_byte[position[1]] = '^'
			case down:
				temp_byte[position[1]] = 'V'
			case right:
				temp_byte[position[1]] = '>'
			case left:
				temp_byte[position[1]] = '<'
			}
			obstacle_map[position[0]] = string(temp_byte)

			if direction == up {
				if position[0] == 0 {
					fmt.Println("Moving of the board upwards, position: ", position[1], ",", position[0])
					break
				}
				if string(obstacle_map[position[0]-1][position[1]]) == "#" {
					direction = right
					continue
				}
				if string(obstacle_map[position[0]-1][position[1]]) == "^" {
					infinite_loops++
					fmt.Println("INFINITE LOOP DETECTED")
					break
				}
				position[0] = position[0] - 1
			}

			if direction == down {
				if position[0] == len(obstacle_map)-1 {
					fmt.Println("Moving of the board downwards, position: ", position[1], ",", position[0])
					break
				}
				if string(obstacle_map[position[0]+1][position[1]]) == "#" {
					direction = left
					continue
				}
				if string(obstacle_map[position[0]+1][position[1]]) == "V" {
					infinite_loops++
					fmt.Println("INFINITE LOOP DETECTED")
					break
				}
				position[0] = position[0] + 1
			}

			if direction == right {
				if position[1] == len(obstacle_map[0])-1 {
					fmt.Println("Moving of the board rightwards, position: ", position[1], ",", position[0])
					break
				}
				if string(obstacle_map[position[0]][position[1]+1]) == "#" {
					direction = down
					continue
				}
				if string(obstacle_map[position[0]][position[1]+1]) == ">" {
					infinite_loops++
					fmt.Println("INFINITE LOOP DETECTED")
					break
				}
				position[1] = position[1] + 1
			}

			if direction == left {
				if position[1] == 0 {
					fmt.Println("Moving of the board leftwards, position: ", position[1], ",", position[0])
					break
				}
				if string(obstacle_map[position[0]][position[1]-1]) == "#" {
					direction = up
					continue
				}
				if string(obstacle_map[position[0]][position[1]-1]) == "<" {
					infinite_loops++
					fmt.Println("INFINITE LOOP DETECTED")
					break
				}
				position[1] = position[1] - 1
			}

			// some infinite loops were undetectable, check if guard have walked more than steps than tiles on the map
			if i > len(obstacle_map)*len(obstacle_map[0]) {
				infinite_loops++
				fmt.Println("INFINITE LOOP DETECTED")
				break
			}
		}

	}

	fmt.Println("SUM: ", sum)
	fmt.Println("Infinite loops: ", infinite_loops)

}
