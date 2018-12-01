package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var frequencies []int

func main() {
	var instructions []string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		instructions = append(instructions, scanner.Text())
	}
	fmt.Printf("Solution to Part A: %d\n", calcPartA(instructions))
	fmt.Printf("Solution to Part B: %d\n", calcPartB(instructions))
}

func calcPartA(instructions []string) int {
	var sum int
	for _, instruction := range instructions {
		operator, amount := extractInstruction(instruction)
		if operator == "+" {
			sum = sum + amount
		} else if operator == "-" {
			sum = sum - amount
		}

	}
	return sum
}

func calcPartB(instructions []string) int {
	var i int
	var sum int
	for {
		operator, amount := extractInstruction(instructions[i%len(instructions)])
		i++
		if operator == "+" {
			sum = sum + amount
		} else if operator == "-" {
			sum = sum - amount
		}

		if addToFrequencyList(sum) {
			return sum
		}
	}
}

// extractInstruction takes an instruction and returns the operator and the amount
func extractInstruction(instr string) (string, int) {
	operator := instr[:1]
	amount, err := strconv.Atoi(instr[1:])
	if err != nil {
		panic(err)
	}
	return operator, amount
}

// addToFrequencyList takes a frequency and returns true or false depending on if we have a duplicate entry
func addToFrequencyList(f int) bool {
	for _, i := range frequencies {
		if i == f {
			return true
		}
	}
	frequencies = append(frequencies, f)
	return false

}
