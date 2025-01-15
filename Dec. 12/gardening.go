package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type GardenPlot struct {
	area, perimiter, walls int
}

func find_area(garden_map [][]string, y int, x int, covered [][]int) [][]int {

	for _, c := range covered {
		if c[0] == y && c[1] == x {
			return covered
		}
	}

	covered = append(covered, []int{y, x})
	if y > 0 && garden_map[y-1][x] == garden_map[y][x] {
		covered = find_area(garden_map, y-1, x, covered)
	}
	if y < len(garden_map)-1 && garden_map[y+1][x] == garden_map[y][x] {
		covered = find_area(garden_map, y+1, x, covered)
	}
	if x > 0 && garden_map[y][x-1] == garden_map[y][x] {
		covered = find_area(garden_map, y, x-1, covered)
	}
	if x < len(garden_map[0])-1 && garden_map[y][x+1] == garden_map[y][x] {
		covered = find_area(garden_map, y, x+1, covered)
	}

	return covered
}

func find_plots(garden_map [][]string, y int, x int) (GardenPlot, [][]int) {

	var gardenplot GardenPlot
	var empty_slice [][]int
	area := find_area(garden_map, y, x, empty_slice)
	gardenplot.area = len(area)
	var top_fence, left_fence, right_fence, bottom_fence [][]int
	for _, a := range area {
		if a[0] == 0 || !slices.ContainsFunc(area, func(n []int) bool {
			return a[0]-1 == n[0] && a[1] == n[1]
		}) {
			gardenplot.perimiter++
			top_fence = append(top_fence, a)
		}
		if a[0] == len(garden_map)-1 || !slices.ContainsFunc(area, func(n []int) bool {
			return a[0]+1 == n[0] && a[1] == n[1]
		}) {
			gardenplot.perimiter++
			bottom_fence = append(bottom_fence, a)
		}
		if a[1] == 0 || !slices.ContainsFunc(area, func(n []int) bool {
			return a[0] == n[0] && a[1]-1 == n[1]
		}) {
			gardenplot.perimiter++
			left_fence = append(left_fence, a)
		}
		if a[1] == len(garden_map[0])-1 || !slices.ContainsFunc(area, func(n []int) bool {
			return a[0] == n[0] && a[1]+1 == n[1]
		}) {
			gardenplot.perimiter++
			right_fence = append(right_fence, a)
		}

	}

	slices.SortFunc(top_fence, func(a, b []int) int {
		if a[0] == b[0] {
			return cmp.Compare(a[1], b[1])
		}
		return cmp.Compare(a[0], b[0])
	})
	slices.SortFunc(bottom_fence, func(a, b []int) int {
		if a[0] == b[0] {
			return cmp.Compare(a[1], b[1])
		}
		return cmp.Compare(a[0], b[0])
	})

	slices.SortFunc(left_fence, func(a, b []int) int {
		if a[1] == b[1] {
			return cmp.Compare(a[0], b[0])
		}
		return cmp.Compare(a[1], b[1])
	})
	slices.SortFunc(right_fence, func(a, b []int) int {
		if a[1] == b[1] {
			return cmp.Compare(a[0], b[0])
		}
		return cmp.Compare(a[1], b[1])
	})

	walls := 0

	fmt.Println(top_fence)
	for z, fence := range top_fence {
		if z == 0 {
			walls++
			continue
		}
		if fence[0] != top_fence[z-1][0] {
			walls++
			continue
		}
		if fence[1] == top_fence[z-1][1]+1 {
			continue
		} else {
			walls++
			continue
		}
	}

	for z, fence := range bottom_fence {
		if z == 0 {
			walls++
			continue
		}
		if fence[0] != bottom_fence[z-1][0] {
			walls++
			continue
		}
		if fence[1] == bottom_fence[z-1][1]+1 {
			continue
		} else {
			walls++
			continue
		}
	}

	for z, fence := range left_fence {
		if z == 0 {
			walls++
			continue
		}
		if fence[1] != left_fence[z-1][1] {
			walls++
			continue
		}
		if fence[0] == left_fence[z-1][0]+1 {
			continue
		} else {
			walls++
			continue
		}
	}

	for z, fence := range right_fence {
		if z == 0 {
			walls++
			continue
		}
		if fence[1] != right_fence[z-1][1] {
			walls++
			continue
		}
		if fence[0] == right_fence[z-1][0]+1 {
			continue
		} else {
			walls++
			continue
		}
	}

	gardenplot.walls = walls

	return gardenplot, area
}

func main() {
	println("Fancing gardens")

	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	var lines []string
	var gardenplots []GardenPlot
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var garden_map [][]string
	for _, line := range lines {
		var temp_map []string
		for _, letter := range line {
			temp_map = append(temp_map, string(letter))
		}
		garden_map = append(garden_map, temp_map)
	}

	var covered_plots [][]int
	for y := 0; y <= len(garden_map)-1; y++ {
		for x := 0; x <= len(garden_map[0])-1; x++ {
			covered := false
			for _, plot := range covered_plots {
				if plot[0] == y && plot[1] == x {
					covered = true
					break
				}
			}

			if covered {
				continue
			}

			gardenplot, covered_area := find_plots(garden_map, y, x)
			covered_plots = append(covered_plots, covered_area...)
			fmt.Println("Area, perimiter for ", garden_map[y][x], ": ", y, ",", x, " : ", gardenplot)
			gardenplots = append(gardenplots, gardenplot)
		}
	}

	//fmt.Println(garden_map)
	//fmt.Println(gardenplots)

	cost := 0
	wallcost := 0
	for _, gardenplot := range gardenplots {
		cost += gardenplot.area * gardenplot.perimiter
		//fmt.Println("Fence cost: ", gardenplot.area, " * ", gardenplot.perimiter, " = ", cost)
		wallcost += gardenplot.area * gardenplot.walls
	}
	fmt.Println(cost)
	fmt.Println(wallcost)
}
