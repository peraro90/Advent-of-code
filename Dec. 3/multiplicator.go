package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	println("multiplying numbers")

	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sum := 0
	r := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
	num := regexp.MustCompile("\\d{1,3}")

	var single_line string
	for _, line := range lines {
		single_line += line
	}

	factors := r.FindAllString(single_line, -1)
	for _, factor := range factors {
		numbers := num.FindAllString(factor, -1)
		fac0, _ := strconv.Atoi(numbers[0])
		fac1, _ := strconv.Atoi(numbers[1])
		sum += fac0 * fac1
	}

	println("first solution: ", sum)
	sum = 0

	do_regex := regexp.MustCompile("do\\(\\)")
	dont_regex := regexp.MustCompile("don't\\(\\)")
	dos := do_regex.FindAllStringIndex(single_line, -1)
	donts := dont_regex.FindAllStringIndex(single_line, -1)

	newstring := single_line[:min(donts[0][1], dos[0][1])]

	range_covered := min(donts[0][0], dos[0][0])
	for _, does := range dos {
		for _, do_not := range donts {
			if does[0] < range_covered {
				break
			}
			if do_not[0] > does[0] {
				range_covered = do_not[1]
				newstring += single_line[does[0]:do_not[1]]
				break
			}
			if does[0] > donts[len(donts)-1][0] {
				newstring += single_line[does[0]:]
				range_covered = len(single_line)
				break
			}
		}
	}

	factors = r.FindAllString(newstring, -1)
	for _, factor := range factors {
		numbers := num.FindAllString(factor, -1)
		fac0, _ := strconv.Atoi(numbers[0])
		fac1, _ := strconv.Atoi(numbers[1])
		sum += fac0 * fac1
	}

	println("final sum: ", sum)

}
