package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day1Part2() {

	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)

	elfCalories := 0
	most1 := 0
	most2 := 0
	most3 := 0

	for fs.Scan() {
		item, err := strconv.Atoi(fs.Text())

		elfCalories += item

		if err != nil { //if error then we are on newline
			if elfCalories > most1 {
				most1 = elfCalories
			}

			if most1 > most2 {
				most1, most2 = most2, most1
			}

			if most2 > most3 {
				most3, most2 = most2, most3
			}

			elfCalories = 0
		}
	}

	fmt.Println(most1 + most2 + most3)

}
