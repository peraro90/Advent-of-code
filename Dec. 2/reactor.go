package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func is_unsafe(levels []int) bool {

	increasing := true
	if levels[0] > levels[1] {
		increasing = false
	}
	if levels[0] == levels[1] {
		return true

	}

	if increasing {
		for i := 0; i < len(levels)-1; i++ {
			if !(levels[i+1]-levels[i] >= 1 && levels[i+1]-levels[i] <= 3) {
				return true
			}
		}
	} else {
		for i := 0; i < len(levels)-1; i++ {
			if !(levels[i]-levels[i+1] >= 1 && levels[i]-levels[i+1] <= 3) {
				return true
			}
		}
	}
	return false
}

func main() {
	println("Checking nuclear reindeer reactor")

	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	unsafe := 0
	unsafe_levels := [][]int{}
	for _, line := range lines {
		reactor_levels := strings.Split(line, " ")
		levels := []int{}
		for _, value := range reactor_levels {
			val, _ := strconv.Atoi(value)
			levels = append(levels, val)
		}
		if is_unsafe(levels) {
			unsafe++
			unsafe_levels = append(unsafe_levels, levels)
		}
	}
	unsafe_before_override := unsafe

	for _, unsafe_level := range unsafe_levels {
		for i := range unsafe_level {
			temp_level_top := unsafe_level[:i]
			temp_level_bottom := unsafe_level[i+1:]
			var temp_level []int
			if len(temp_level_top) > 0 {
				temp_level = append(temp_level, temp_level_top...)
			}

			if len(temp_level_bottom) > 0 {
				temp_level = append(temp_level, temp_level_bottom...)
			}

			if !is_unsafe(temp_level) {
				unsafe--
				break
			}
		}
	}

	total := len(lines)
	safe_before_override := total - unsafe_before_override
	safe := total - unsafe

	println("SAFE: ", safe_before_override, "/", total)
	println("USING OVERRIDE, SAFE: ", safe, "/", total)

}
