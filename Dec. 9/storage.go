package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	println("Optimizing storage space")

	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var storage []string
	var alt_storage [][]int
	for i, letter := range lines[0] {
		temp, _ := strconv.Atoi(string(letter))
		var content []string
		for j := 0; j < temp; j++ {
			if i%2 != 0 {
				content = append(content, ".")
			} else {
				content = append(content, strconv.Itoa(i/2))
			}
		}

		if i%2 != 0 {
			if temp != 0 {
				alt_storage = append(alt_storage, []int{-1, temp})
			}
		} else {
			alt_storage = append(alt_storage, []int{i / 2, temp})
		}

		storage = append(storage, content...)
	}
	//var copy_storage = make([][]int, len(alt_storage))
	//copy(copy_storage, alt_storage)

	var copy_storage [][]int
	for _, s := range alt_storage {
		copy_storage = append(copy_storage, slices.Clone(s))
	}
	var unfragmented [][]int

	for _, cs := range slices.Backward(alt_storage) {
		if cs[0] == -1 {
			continue
		} else {
			for i, s := range copy_storage {
				if s[1] >= cs[1] && s[0] == -1 {
					copy_storage = slices.Insert(copy_storage, i, slices.Clone(cs))
					for a := i + 1; a <= len(copy_storage)-1; a++ {
						if copy_storage[a][0] == cs[0] {
							if a == len(copy_storage)-1 {
								copy_storage = slices.Delete(copy_storage, a, a+1)
								break
							} else {

								copy_storage[a][0] = -1
								if copy_storage[a+1][0] == -1 {
									copy_storage[a+1][1] += copy_storage[a][1]
									copy_storage = slices.Delete(copy_storage, a, a+1)
								}
								if copy_storage[a-1][0] == -1 {
									copy_storage[a-1][1] += copy_storage[a][1]
									copy_storage = slices.Delete(copy_storage, a, a+1)
								}
								break
							}
						}
					}
					s[1] -= cs[1]
					if s[1] == 0 {
						copy_storage = slices.Delete(copy_storage, i+1, i+2)
					}
					if copy_storage[len(copy_storage)-1][0] == -1 {
						copy_storage = copy_storage[:len(copy_storage)-1]
					}
					break
				}
				if cs[0] == s[0] {
					break
				}
			}

		}
	}

	unfragmented = copy_storage

	fmt.Println(unfragmented)
	sum := 0
	ticker := 0
	for _, member := range unfragmented {
		for i := 0; i < member[1]; i++ {
			if member[0] != -1 {
				sum += member[0] * ticker
			}
			ticker++
		}
	}
	fmt.Println(sum)

	var reorg [][]int
	for count, s := range alt_storage {
		if s[0] == -1 {
			for x := 0; x < s[1]; {
				if alt_storage[len(alt_storage)-1][1] <= s[1]-x {
					reorg = append(reorg, []int{alt_storage[len(alt_storage)-1][0], alt_storage[len(alt_storage)-1][1]})
					x += alt_storage[len(alt_storage)-1][1]
					alt_storage = alt_storage[:len(alt_storage)-1]
				} else {
					reorg = append(reorg, []int{alt_storage[len(alt_storage)-1][0], s[1] - x})
					alt_storage[len(alt_storage)-1][1] = alt_storage[len(alt_storage)-1][1] - (s[1] - x)
					x += s[1] - x
					continue
				}

				if alt_storage[len(alt_storage)-1][0] == -1 {
					alt_storage = alt_storage[:len(alt_storage)-1]
				}
			}
		} else {
			reorg = append(reorg, s)
		}

		if count == len(alt_storage)-1 {
			break
		}
	}

	//var unfragmented [][]int

	fmt.Println(reorg)
	sum = 0
	ticker = 0
	for _, member := range reorg {
		for i := 0; i < member[1]; i++ {
			sum += member[0] * ticker
			ticker++
		}
	}
	fmt.Println(sum)

}
