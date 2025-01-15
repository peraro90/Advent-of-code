package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Button struct {
	x, y, cost int
}

type ClawMachine struct {
	a, b       Button
	prize      []int
	total_cost int
}

func main() {
	println("Playing claw machines")

	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	regex_button := regexp.MustCompile("\\d+")

	var clawmachine []ClawMachine
	for i := 0; i < len(lines); i = i + 4 {
		var cm ClawMachine
		var a_button, b_button Button

		x := regex_button.FindString(strings.Split(strings.Split(lines[i], ":")[1], ",")[0])
		y := regex_button.FindString(strings.Split(strings.Split(lines[i], ":")[1], ",")[1])

		a_button.x, _ = strconv.Atoi(x)
		a_button.y, _ = strconv.Atoi(y)
		a_button.cost = 3

		x = regex_button.FindString(strings.Split(strings.Split(lines[i+1], ":")[1], ",")[0])
		y = regex_button.FindString(strings.Split(strings.Split(lines[i+1], ":")[1], ",")[1])

		b_button.x, _ = strconv.Atoi(x)
		b_button.y, _ = strconv.Atoi(y)
		b_button.cost = 1

		cm.a = a_button
		cm.b = b_button

		px, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(lines[i+2], ":")[1], ",")[0], "=")[1])
		py, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(lines[i+2], ":")[1], ",")[1], "=")[1])

		cm.prize = []int{px, py}
		cm.total_cost = 401 //max possible cost+1
		clawmachine = append(clawmachine, cm)
	}

	total := 0
	for n, cm := range clawmachine {
		for i := 0; i <= 100; i++ {
			for j := 0; j <= 100; j++ {
				if i*cm.a.x+j*cm.b.x == cm.prize[0] && i*cm.a.y+j*cm.b.y == cm.prize[1] {
					cm.total_cost = min(i*cm.a.cost+j*cm.b.cost, cm.total_cost)
					fmt.Println("ClawMachine ", n, " wins by a: ", i, " b: ", j)
					break
				}
				if i*cm.a.x+j*cm.b.x > cm.prize[0] || i*cm.a.y+j*cm.b.y > cm.prize[1] {
					break
				}
			}
		}
		if cm.total_cost < 401 {
			total += cm.total_cost
		}
	}

	fmt.Println("Total cost: ", total)

	fmt.Println("Unit conversion error. \n Fixing prize location")
	for i, cm := range clawmachine {
		clawmachine[i].prize[0] = cm.prize[0] + 10000000000000
		clawmachine[i].prize[1] = cm.prize[1] + 10000000000000
		clawmachine[i].total_cost = -1
	}
	fmt.Println("Done")
	total = 0

	//Solved as if a matrix
	for n, cm := range clawmachine {

		determinant := cm.a.x*cm.b.y - cm.a.y*cm.b.x
		a := cm.b.y*cm.prize[0] - cm.b.x*cm.prize[1]
		b := -cm.a.y*cm.prize[0] + cm.a.x*cm.prize[1]

		if a%determinant == 0 && b%determinant == 0 {
			cm.total_cost = (a/determinant)*cm.a.cost + (b/determinant)*cm.b.cost
		}

		if cm.total_cost != -1 {
			total += cm.total_cost
			fmt.Println("ClawMachine ", n, " gives prize with min ", cm.total_cost, " tokens")
		}
	}
	fmt.Println(total)
}
