package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cal := 0
	maxVal1 := 0
	maxVal2 := 0
	maxVal3 := 0
	for scanner.Scan() {
		if scanner.Text() != "" {
			str := scanner.Text()
			intVal, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			cal += intVal
		} else {
			if cal > maxVal1 {
				maxVal3 = maxVal2
				maxVal2 = maxVal1
				maxVal1 = cal
			} else if cal > maxVal2 {
				maxVal3 = maxVal2
				maxVal2 = cal
			} else if cal > maxVal3 {
				maxVal3 = cal
			}
			cal = 0
		}
	}

	fmt.Println("Max calories = " + strconv.Itoa(maxVal1))
	fmt.Println("3 elves max calories = " + strconv.Itoa(maxVal1+maxVal2+maxVal3))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
