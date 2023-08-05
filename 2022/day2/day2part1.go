package day2

import (
	"bufio"
	"fmt"
	"os"
)

func Day2Part1() {

	file, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)

	decision := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	score := 0

	for fs.Scan() {
		oponent := fs.Text()[0:1]
		me := fs.Text()[2:3]

		if (me == "X" && oponent == "A") || (me == "Y" && oponent == "B") || (me == "Z" && oponent == "C") {
			//all draw outcomes
			score += 3 + decision[me]
		} else if (me == "X" && oponent == "C") || (me == "Y" && oponent == "A") || (me == "Z" && oponent == "B") {
			//all win for me outcomes
			score += 6 + decision[me]
		} else if (me == "X" && oponent == "B") || (me == "Y" && oponent == "C") || (me == "Z" && oponent == "A") {
			//all loss for me outcomes
			score += decision[me]
		}

	}
	fmt.Println(score)

}
