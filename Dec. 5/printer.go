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

func main() {
	println("solving crossword puzzle")

	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	rules := true
	var ordered_rules [][]int
	var printer_queue []string
	for _, line := range lines {
		if string(line) == "" {
			rules = false
		}

		if rules {
			rule_string := strings.Split(line, "|")
			x, _ := strconv.Atoi(rule_string[0])
			y, _ := strconv.Atoi(rule_string[1])
			ordered_rules = append(ordered_rules, []int{x, y})

		} else {
			printer_queue = append(printer_queue, line)
		}
	}

	sum := 0
	error_sum := 0
	for _, queue := range printer_queue {
		elements := strings.Split(queue, ",")
		printable := true
		for i, element := range elements {
			if i == 0 {
				continue
			}
			for j := range i {
				for _, rule := range ordered_rules {
					current_value, _ := strconv.Atoi(element)
					early_value, _ := strconv.Atoi(elements[j])
					if current_value == rule[0] && early_value == rule[1] {
						printable = false
						break
					}
				}
				if !printable {
					break
				}
			}
			if !printable {
				break
			}
		}
		if printable {
			mid_value, _ := strconv.Atoi(elements[len(elements)/2])
			sum += mid_value
		} else {
			for {
				printable := true
				sorted_printer_queue := elements
				for i, element := range sorted_printer_queue {
					if i == 0 {
						continue
					}
					for j := range i {
						for _, rule := range ordered_rules {
							current_value, _ := strconv.Atoi(element)
							early_value, _ := strconv.Atoi(elements[j])
							if current_value == rule[0] && early_value == rule[1] {
								printable = false
								sorted_printer_queue[j] = strconv.Itoa(current_value)
								sorted_printer_queue[i] = strconv.Itoa(early_value)
								break
							}
						}
						if !printable {
							break
						}
					}
					if !printable {
						break
					}
				}
				if printable {
					mid_value, _ := strconv.Atoi(sorted_printer_queue[len(sorted_printer_queue)/2])
					error_sum += mid_value
					break
				}
			}
		}
	}
	fmt.Println("Total= ", sum)
	fmt.Println("ERROR CORRECTION_TOTAL= ", error_sum)

}
