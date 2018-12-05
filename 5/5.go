package main

import (
	"bufio"
	"fmt"
	"os"
)

func removeUnit(s []byte, r byte) []byte {
	var outSlice []byte
	for i := range s {
		// fmt.Println(string(s[i]))
		if s[i] != r && s[i] != r+32 {
			outSlice = append(outSlice, s[i])
		}
	}
	return outSlice
}

// reduce Recursively removes stuff
func reduce(s []byte) []byte {
	for i := range s[:len(s)-1] {
		if s[i]-s[i+1] == 32 || s[i+1]-s[i] == 32 {
			s = append(s[:i], s[i+2:]...)
			return reduce(s)
		}
	}
	return s
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()
	byteSlice := []byte(data)
	fmt.Println("The answer to part A:", len(reduce(byteSlice)))

	// Part B:

	tried := make(map[byte]int)
	for i := byte('A'); i <= byte('Z'); i++ {
		tried[i] = len(reduce(removeUnit([]byte(data), i)))
	}

	minimum := 100000
	var byteRemoved byte
	for i := range tried {
		if tried[i] < minimum {
			minimum = tried[i]
			byteRemoved = byte(i)
		}
	}
	fmt.Printf("The answer to Part B: %v (When removing %v)\n", minimum, string(byteRemoved))

}
