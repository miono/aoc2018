package main

import (
	"bufio"
	"fmt"
	"os"
)

type rect struct {
	ID     int
	startX int
	startY int
	width  int
	height int
}

func (r *rect) checkMySelfInGrid(g *grid) bool {
	ctrlVar := 0
	for i := r.startX; i < r.startX+r.width; i++ {
		for j := r.startY; j < r.startY+r.height; j++ {
			if g.matrix[i][j] != 1 {
				ctrlVar = 1
			}
		}
	}
	if ctrlVar == 0 {
		return true
	}
	return false
}

type grid struct {
	xySize int
	matrix map[int]map[int]int
}

func newGrid(xySize int) *grid {
	g := make(map[int]map[int]int)
	for i := 0; i < xySize; i++ {
		g[i] = make(map[int]int)
	}
	return &grid{xySize: xySize, matrix: g}
}

func (g *grid) addRectToGrid(r rect) {
	for i := r.startX; i < r.startX+r.width; i++ {
		for j := r.startY; j < r.startY+r.height; j++ {
			g.matrix[i][j]++
		}
	}
}
func (g *grid) print() {
	for i := 0; i < g.xySize; i++ {
		for j := 0; j < g.xySize; j++ {
			if g.matrix[i][j] < 1 {
				fmt.Print(0)
			} else {
				fmt.Print(g.matrix[j][i])
			}
		}
		fmt.Print("\n")
	}
}
func (g *grid) countOverlapping() int {
	overlapping := 0
	for i := 0; i < g.xySize; i++ {
		for j := 0; j < g.xySize; j++ {
			if g.matrix[i][j] > 1 {
				overlapping++
			}
		}
	}
	return overlapping

}

func main() {
	var rects []rect
	myGrid := newGrid(1050)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		var newRect rect
		_, err := fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &newRect.ID, &newRect.startX, &newRect.startY, &newRect.width, &newRect.height)
		if err != nil {
			panic(err)
		}
		myGrid.addRectToGrid(newRect)
		rects = append(rects, newRect)
	}
	fmt.Println("Answer to part A:", myGrid.countOverlapping())
	for _, r := range rects {
		if r.checkMySelfInGrid(myGrid) {
			fmt.Println("Answer to part B:", r.ID)
			break
		}
	}
}
