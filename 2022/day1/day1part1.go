package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day1Part1() {

	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)

	elfCalories := 0
	mostCallories := 0

	for fs.Scan() {
		item, err := strconv.Atoi(fs.Text())

		elfCalories += item

		if err != nil { //if error then we are on newline
			if elfCalories > mostCallories {
				mostCallories = elfCalories
			}
			elfCalories = 0
		}
	}

	fmt.Println(mostCallories)

}
