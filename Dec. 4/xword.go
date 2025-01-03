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
	println("solving crossword puzzle")

	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var x_map [][]int
	var a_map [][]int
	//sum := 0
	for i, line := range lines {
		//println(line)
		for j, letter := range line {
			if string(letter) == "X" {
				x_map = append(x_map, []int{i, j})
			}
			if string(letter) == "A" {
				a_map = append(a_map, []int{i, j})
			}
		}
	}
	total := 0
	total_up := 0
	total_down := 0
	total_right := 0
	total_left := 0
	total_up_left := 0
	total_up_right := 0
	total_down_right := 0
	total_down_left := 0
	for _, coordinate := range x_map {
		x := coordinate[0]
		y := coordinate[1]

		up := x-3 >= 0
		down := x+3 < len(lines)
		left := y-3 >= 0
		right := y+3 < len(lines[0])

		if up && string(lines[x-1][y])+string(lines[x-2][y])+string(lines[x-3][y]) == "MAS" {
			total++
			total_up++
		}
		if down && string(lines[x+1][y])+string(lines[x+2][y])+string(lines[x+3][y]) == "MAS" {
			total++
			total_down++
		}
		if left && string(lines[x][y-1])+string(lines[x][y-2])+string(lines[x][y-3]) == "MAS" {
			total++
			total_left++
		}
		if right && string(lines[x][y+1])+string(lines[x][y+2])+string(lines[x][y+3]) == "MAS" {
			total++
			total_right++
		}
		if up && right && string(lines[x-1][y+1])+string(lines[x-2][y+2])+string(lines[x-3][y+3]) == "MAS" {
			total++
			total_up_right++
		}
		if down && right && string(lines[x+1][y+1])+string(lines[x+2][y+2])+string(lines[x+3][y+3]) == "MAS" {
			total++
			total_down_right++
		}
		if up && left && string(lines[x-1][y-1])+string(lines[x-2][y-2])+string(lines[x-3][y-3]) == "MAS" {
			total++
			total_up_left++
		}
		if down && left && string(lines[x+1][y-1])+string(lines[x+2][y-2])+string(lines[x+3][y-3]) == "MAS" {
			total++
			total_down_left++
		}

	}

	println("UP:         ", total_up)
	println("DOWN:       ", total_down)
	println("LEFT:       ", total_left)
	println("RIGHT:      ", total_right)
	println("UP_RIGHT:   ", total_up_right)
	println("DOWN_RIGHT: ", total_down_right)
	println("UP_LEFT:    ", total_up_left)
	println("DOWN_LEFT:  ", total_down_left)

	println("XMAS: ", total)
	new_total := 0
	for _, coordinate := range a_map {
		x := coordinate[0]
		y := coordinate[1]

		if x == 0 || x >= len(lines)-1 {
			continue
		}
		if y == 0 || y >= len(lines[0])-1 {
			continue
		}

		if string(lines[x+1][y-1]) == "M" && string(lines[x+1][y+1]) == "M" && string(lines[x-1][y-1]) == "S" && string(lines[x-1][y+1]) == "S" {
			new_total++
		}
		if string(lines[x+1][y-1]) == "M" && string(lines[x-1][y-1]) == "M" && string(lines[x-1][y+1]) == "S" && string(lines[x+1][y+1]) == "S" {
			new_total++
		}
		if string(lines[x-1][y-1]) == "M" && string(lines[x-1][y+1]) == "M" && string(lines[x+1][y-1]) == "S" && string(lines[x+1][y+1]) == "S" {
			new_total++
		}
		if string(lines[x-1][y+1]) == "M" && string(lines[x+1][y+1]) == "M" && string(lines[x+1][y-1]) == "S" && string(lines[x-1][y-1]) == "S" {
			new_total++
		}

	}

	println("X_MAS:", new_total)

}
