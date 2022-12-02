package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Total score first strategy : " + strconv.Itoa(firstStrategy()))
	fmt.Println("Total score second strategy : " + strconv.Itoa(secondStrategy()))
}

func firstStrategy() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalScore := 0

	for scanner.Scan() {
		str := scanner.Text()
		strSplit := strings.Fields(str)

		if strSplit[0] == "A" {
			if strSplit[1] == "X" {
				totalScore += 1 + 3
			} else if strSplit[1] == "Y" {
				totalScore += 2 + 6
			} else if strSplit[1] == "Z" {
				totalScore += 3 + 0
			}
		} else if strSplit[0] == "B" {
			if strSplit[1] == "X" {
				totalScore += 1 + 0
			} else if strSplit[1] == "Y" {
				totalScore += 2 + 3
			} else if strSplit[1] == "Z" {
				totalScore += 3 + 6
			}
		} else if strSplit[0] == "C" {
			if strSplit[1] == "X" {
				totalScore += 1 + 6
			} else if strSplit[1] == "Y" {
				totalScore += 2 + 0
			} else if strSplit[1] == "Z" {
				totalScore += 3 + 3
			}
		}
	}

	return totalScore
}

func secondStrategy() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalScore := 0

	for scanner.Scan() {
		str := scanner.Text()
		strSplit := strings.Fields(str)

		if strSplit[0] == "A" {
			if strSplit[1] == "X" {
				totalScore += 3 + 0
			} else if strSplit[1] == "Y" {
				totalScore += 1 + 3
			} else if strSplit[1] == "Z" {
				totalScore += 2 + 6
			}
		} else if strSplit[0] == "B" {
			if strSplit[1] == "X" {
				totalScore += 1 + 0
			} else if strSplit[1] == "Y" {
				totalScore += 2 + 3
			} else if strSplit[1] == "Z" {
				totalScore += 3 + 6
			}
		} else if strSplit[0] == "C" {
			if strSplit[1] == "X" {
				totalScore += 2 + 0
			} else if strSplit[1] == "Y" {
				totalScore += 3 + 3
			} else if strSplit[1] == "Z" {
				totalScore += 1 + 6
			}
		}
	}

	return totalScore
}
