package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var elevation [][]int

func find_peak(y int, x int, current_score int, current_rating int, taken_peaks [][]int) (int, int, [][]int) {
	if elevation[y][x] == 9 {
		for _, peak := range taken_peaks {
			if y == peak[0] && x == peak[1] {
				return current_score, current_rating + 1, taken_peaks
			}
		}
		return current_score + 1, current_rating + 1, append(taken_peaks, []int{y, x})
	}
	if y > 0 && elevation[y-1][x] == elevation[y][x]+1 {
		current_score, current_rating, taken_peaks = find_peak(y-1, x, current_score, current_rating, taken_peaks)
	}
	if y < len(elevation)-1 && elevation[y+1][x] == elevation[y][x]+1 {
		current_score, current_rating, taken_peaks = find_peak(y+1, x, current_score, current_rating, taken_peaks)
	}
	if x > 0 && elevation[y][x-1] == elevation[y][x]+1 {
		current_score, current_rating, taken_peaks = find_peak(y, x-1, current_score, current_rating, taken_peaks)
	}
	if x < len(elevation[0])-1 && elevation[y][x+1] == elevation[y][x]+1 {
		current_score, current_rating, taken_peaks = find_peak(y, x+1, current_score, current_rating, taken_peaks)
	}
	return current_score, current_rating, taken_peaks
}

func main() {
	println("Finding the best hiking trails")

	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var start_pos [][]int
	for y, numbers := range lines {
		var temp_elev []int
		for x, number := range numbers {
			elev, _ := strconv.Atoi(string(number))
			temp_elev = append(temp_elev, elev)
			if elev == 0 {
				start_pos = append(start_pos, []int{y, x})
			}
		}
		elevation = append(elevation, temp_elev)
	}
	fmt.Println(elevation)
	fmt.Println(start_pos)

	score := 0
	rating := 0
	var empty_list [][]int
	for _, sp := range start_pos {
		temp_score, temp_rating, _ := find_peak(sp[0], sp[1], 0, 0, empty_list)
		score += temp_score
		rating += temp_rating
	}

	fmt.Println("Score: ", score)
	fmt.Println("Rating: ", rating)

}
