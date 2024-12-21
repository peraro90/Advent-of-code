package main

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	println("Starting wall hack")

	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	starting_pos_x := 0
	starting_pos_y := 0

	end_pos_x := 0
	end_pos_y := 0

	x_len := len(lines)
	y_len := len(lines[0])

	println(x_len)
	println(y_len)

	var comp_map [141][141]string

	// find start and end possitions
	for i, line := range lines {
		for j, letter := range line {
			if string(letter) == "S" {
				starting_pos_x = i
				starting_pos_y = j
			}
			if string(letter) == "E" {
				end_pos_x = i
				end_pos_y = j
			}
			comp_map[i][j] = string(letter)
			if string(letter) == "E" {
				comp_map[i][j] = "."
			}
		}
	}

	println("start pos ", starting_pos_x, ",", starting_pos_y)
	println("end pos ", end_pos_x, ",", end_pos_y)

	pos_x := starting_pos_x
	pos_y := starting_pos_y

	type coordinate struct {
		x, y int
	}

	/*
		create a list of movement through the map
	*/
	traversal_list := []coordinate{}
	finished := false
	for !finished {
		coordinates := coordinate{pos_x, pos_y}
		traversal_list = append(traversal_list, coordinates)
		if pos_x == end_pos_x && pos_y == end_pos_y {
			finished = true
			break
		}
		comp_map[pos_x][pos_y] = "X"
		if comp_map[pos_x+1][pos_y] == "." {
			pos_x = pos_x + 1
		} else if comp_map[pos_x-1][pos_y] == "." {
			pos_x = pos_x - 1
		} else if comp_map[pos_x][pos_y+1] == "." {
			pos_y = pos_y + 1
		} else if comp_map[pos_x][pos_y-1] == "." {
			pos_y = pos_y - 1
		}
	}

	old_cheat := 0
	cheats := 0

	/*
		to save at least 100 picoseconds we must be at least 101 steps ahead in the list
	*/
	for a := 0; a < len(traversal_list); a++ {
		for b := a + 101; b < len(traversal_list); b++ {
			distance_x := traversal_list[a].x - traversal_list[b].x
			if distance_x < 0 {
				distance_x = distance_x * -1
			}
			distance_y := traversal_list[a].y - traversal_list[b].y
			if distance_y < 0 {
				distance_y = distance_y * -1
			}
			total_distance := distance_x + distance_y

			if total_distance <= 20 {
				if b-a-total_distance >= 100 {
					cheats++
				}
			}
			if total_distance <= 2 {
				old_cheat++
			}

		}
	}

	println(old_cheat)
	println(cheats)

}
