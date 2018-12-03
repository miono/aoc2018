package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var boxes []string

func checkForN(boxID string, n int) int {
	letters := map[rune]int{}
	for _, char := range boxID {
		letters[char]++
	}
	for _, v := range letters {
		if v == n {
			return 1
		}
	}
	return 0
}

func findMatchingBox(box1 string) string {
	for _, box2 := range boxes {
		numDiff := 0
		for i := range box2 {
			if box2[i] != box1[i] {
				numDiff++
			}
		}
		if numDiff == 1 {
			return box2
		}
	}

	return ""
}

func normalizeBoxes(boxes []string) string {
	var outString string
	for i := range boxes[0] {
		if boxes[0][i] == boxes[1][i] {
			outString = strings.Join([]string{outString, string(boxes[0][i])}, "")
		}

	}
	return outString

}

func main() {
	var numTwos int
	var numThrees int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		boxes = append(boxes, scanner.Text())
	}
	matchFound := 0
	var matchingBoxes []string
	for _, box := range boxes {
		numTwos += checkForN(box, 2)
		numThrees += checkForN(box, 3)
		if matchFound == 0 {
			matchingBox := findMatchingBox(box)
			if len(matchingBox) > 0 {
				matchingBoxes = append(matchingBoxes, box, matchingBox)
				matchFound = 1
			}
		}
	}
	fmt.Printf("The answer to A is: %d\n", numTwos*numThrees)
	fmt.Println("The matching boxes are:", matchingBoxes)
	fmt.Println("They are normalized to:", normalizeBoxes(matchingBoxes))
}
