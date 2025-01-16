package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Robot struct {
	position []int
	velosity []int
}

func main() {
	println("Evading robot rabits")
	tile_width := 101
	mid_x := int(tile_width / 2)
	tile_height := 103
	mid_y := int(tile_height / 2)
	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var robots []Robot
	for _, line := range lines {
		p := strings.Split(line, " ")[0]
		v := strings.Split(line, " ")[1]

		var temp_robot Robot

		x, _ := strconv.Atoi(strings.Split(strings.Split(p, "=")[1], ",")[0])
		y, _ := strconv.Atoi(strings.Split(strings.Split(p, "=")[1], ",")[1])
		temp_robot.position = append(temp_robot.position, x, y)
		x, _ = strconv.Atoi(strings.Split(strings.Split(v, "=")[1], ",")[0])
		y, _ = strconv.Atoi(strings.Split(strings.Split(v, "=")[1], ",")[1])
		temp_robot.velosity = append(temp_robot.velosity, x, y)
		robots = append(robots, temp_robot)
	}

	var positions [][]int
	for _, robot := range robots {
		if robot.velosity[0] < 0 {
			robot.velosity[0] += tile_width
		}
		if robot.velosity[1] < 0 {
			robot.velosity[1] += tile_height
		}

		x := (robot.position[0] + (robot.velosity[0] * 100)) % (tile_width)
		y := (robot.position[1] + (robot.velosity[1] * 100)) % (tile_height)

		final_position := []int{x, y}
		fmt.Println(final_position)
		positions = append(positions, final_position)
	}

	fmt.Println(positions)
	tl, tr, bl, br := 0, 0, 0, 0

	for _, position := range positions {
		if position[0] == mid_x || position[1] == mid_y {
			continue
		}

		if position[0] < mid_x {
			if position[1] < mid_y {
				tl++
			} else {
				bl++
			}
		} else {
			if position[1] < mid_y {
				tr++
			} else {
				br++
			}
		}
	}

	fmt.Println(tl * tr * bl * br)
	safety := tl * tr * bl * br
	iteration := 100
	i := 0
	for i = 1; i < tile_height*tile_width; i++ {
		tl, tr, bl, br = 0, 0, 0, 0
		positions = [][]int{}
		var drawing [][]int
		for range tile_height {
			drawing = append(drawing, slices.Repeat([]int{0}, tile_width))
		}
		for _, robot := range robots {

			if robot.velosity[0] < 0 {
				robot.velosity[0] += tile_width
			}
			if robot.velosity[1] < 0 {
				robot.velosity[1] += tile_height
			}

			x := (robot.position[0] + (robot.velosity[0] * i)) % (tile_width)
			y := (robot.position[1] + (robot.velosity[1] * i)) % (tile_height)

			final_position := []int{x, y}
			drawing[y][x] = drawing[y][x] + 1
			positions = append(positions, final_position)
		}

		for _, position := range positions {
			if position[0] == mid_x || position[1] == mid_y {
				continue
			}

			if position[0] < mid_x {
				if position[1] < mid_y {
					tl++
				} else {
					bl++
				}
			} else {
				if position[1] < mid_y {
					tr++
				} else {
					br++
				}
			}
		}
		temp_safety := tl * tr * bl * br
		if temp_safety < safety {
			safety = temp_safety
			iteration = i
			fmt.Println("Lowest current safety = iter: ", i, " score: ", safety)
		}

	}
	fmt.Println(iteration)

}
