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

func main() {
	println("Locating Easterbunny black ops")

	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var antenna_pos [][]int
	var antinodes [][]int
	var harmonics [][]int
	for y, line := range lines {
		for x, letter := range line {
			if string(letter) != "." {
				antenna_pos = append(antenna_pos, []int{y, x})
			}
		}
	}

	for i, pos := range antenna_pos {
		for _, pos2 := range antenna_pos[i+1:] {
			if lines[pos[0]][pos[1]] == lines[pos2[0]][pos2[1]] {
				y_dis := pos2[0] - pos[0]
				x_dis := pos2[1] - pos[1]
				antinodes = append(antinodes, []int{pos[0] - (pos2[0] - pos[0]), pos[1] + (pos[1] - pos2[1])})
				antinodes = append(antinodes, []int{pos2[0] + (pos2[0] - pos[0]), pos2[1] - (pos[1] - pos2[1])})

				for i := 0; i <= len(lines)-1; i++ {
					distance_from_node := i - pos[0]
					if y_dis != 0 {
						if (distance_from_node*x_dis)%y_dis == 0 && pos[1]+((distance_from_node*x_dis)/y_dis) >= 0 && pos[1]+((distance_from_node*x_dis)/y_dis) <= len(lines[0])-1 {
							harmonics = append(harmonics, []int{pos[0] + distance_from_node, pos[1] + ((distance_from_node * x_dis) / y_dis)})
						}
					} else {
						for j := 0; j < len(lines[0]); j++ {
							harmonics = append(harmonics, []int{pos[0], j})
						}
					}
				}
			}
		}
	}
	var clean_antinodes [][]int
	for _, node := range antinodes {
		add := true
		if node[0] >= 0 && node[0] <= len(lines)-1 && node[1] >= 0 && node[1] <= len(lines[0])-1 {
			for _, c_node := range clean_antinodes {
				if c_node[0] == node[0] && c_node[1] == node[1] {
					add = false
					break
				}
			}
			if add {
				clean_antinodes = append(clean_antinodes, node)
			}
		}
	}

	for _, harmonic := range harmonics {
		add := true
		for _, c_node := range clean_antinodes {
			if c_node[0] == harmonic[0] && c_node[1] == harmonic[1] {
				add = false
				break
			}
		}
		if add {
			clean_antinodes = append(clean_antinodes, harmonic)
		}
	}

	fmt.Println("number of unique hotspots with harmonics: ", len(clean_antinodes))
}
