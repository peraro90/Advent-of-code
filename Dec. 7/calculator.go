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

func find_value(result int, current_val int, values []string, use_conc bool) (int, bool) {
	val, _ := strconv.Atoi(values[0])
	temp_string := strconv.Itoa(current_val) + strconv.Itoa(val)
	concat, _ := strconv.Atoi(temp_string)
	if len(values) == 1 {
		if current_val+val == result {
			return 0, true
		} else if current_val*val == result {
			return 0, true
		} else if concat == result && use_conc {
			return 0, true
		} else {
			return -1, false
		}
	} else {
		addition, add_res := find_value(result, current_val+val, values[1:], use_conc)
		if add_res {
			return addition, add_res
		}
		mult, mult_res := find_value(result, current_val*val, values[1:], use_conc)
		if mult_res {
			return mult, mult_res
		}
		if use_conc {
			conc, conc_res := find_value(result, concat, values[1:], use_conc)
			if conc_res {
				return conc, conc_res
			}
		}
	}

	return -1, false
}

func main() {
	println("Playing with elephants")

	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sum := 0
	conc_sum := 0
	for _, line := range lines {
		result, _ := strconv.Atoi(strings.Split(line, ":")[0])
		variables := strings.Split(strings.Split(line, ": ")[1], " ")

		_, is_true := find_value(result, 0, variables, false)
		if is_true {
			sum += result
		}
		_, is_true = find_value(result, 0, variables, true)
		if is_true {
			conc_sum += result
		}
	}

	fmt.Println(sum)
	fmt.Println(conc_sum)

}
