package main

import "fmt"

func drawQuad(
	maxLine int,
	maxColumn int,
	cornerTopRight rune,
	cornerTopLeft rune,
	cornerBottomRight rune,
	cornerBottomLeft rune,
	// Change the name to horizontal edge
	edgeLine rune,
	// Change the name to vertical edge
	edgeColumn rune,
) {
	whiteSpace := rune(' ')
	for line := 1; line <= maxLine; line++ {
		if line == 1 {
			drawRow(cornerTopLeft, cornerTopRight, edgeLine, maxColumn)
		} else if line == maxLine {
			drawRow(cornerBottomLeft, cornerBottomRight, edgeLine, maxColumn)
		} else {
			drawRow(edgeColumn, edgeColumn, whiteSpace, maxColumn)
		}
	}
}

func drawRow(
	start rune,
	end rune,
	middle rune,
	maxColumn int,
) {
	for column := 1; column <= maxColumn; column++ {
		if column == 1 {
			fmt.Print(string(start))
		} else if column == maxColumn {
			fmt.Print(string(end))
		} else {
			fmt.Print(string(middle))
		}
	}
	fmt.Println()
}

func quadA(
	maxLine int,
	maxColumn int,
) {
	cornerTopRight := 'o'
	cornerTopLeft := cornerTopRight
	cornerBottomRight := cornerTopRight
	cornerBottomLeft := cornerTopRight
	edgeLine := '-'
	edgeColumn := '|'
	drawQuad(maxLine, maxColumn, cornerTopRight, cornerTopLeft, cornerBottomRight,cornerBottomLeft, edgeLine, edgeColumn)
}

func quadB(
	maxLine int,
	maxColumn int,
) {
	cornerTopRight := '\\'
	cornerTopLeft := '/'
	cornerBottomRight := cornerTopLeft
	cornerBottomLeft := cornerTopRight
	edgeLine := '*'
	edgeColumn := edgeLine
	drawQuad(maxLine, maxColumn, cornerTopRight, cornerTopLeft, cornerBottomRight,cornerBottomLeft, edgeLine, edgeColumn)
}

func quadC(
	maxLine int,
	maxColumn int,
) {
	cornerTopRight := 'A'
	cornerTopLeft := cornerTopRight
	cornerBottomRight := 'C'
	cornerBottomLeft := cornerBottomRight
	edgeLine := 'B'
	edgeColumn := edgeLine
	drawQuad(maxLine, maxColumn, cornerTopRight, cornerTopLeft, cornerBottomRight,cornerBottomLeft, edgeLine, edgeColumn)
}

func quadD(
	maxLine int,
	maxColumn int,
) {
	cornerTopRight := 'C'
	cornerTopLeft := 'A'
	cornerBottomRight := cornerTopRight
	cornerBottomLeft := cornerTopLeft
	edgeLine := 'B'
	edgeColumn := edgeLine
	drawQuad(maxLine, maxColumn, cornerTopRight, cornerTopLeft, cornerBottomRight,cornerBottomLeft, edgeLine, edgeColumn)
}

func quadE(
	maxLine int,
	maxColumn int,
) {
	cornerTopRight := 'C'
	cornerTopLeft := 'A'
	cornerBottomRight := cornerTopLeft
	cornerBottomLeft := cornerTopRight
	edgeLine := 'B'
	edgeColumn := edgeLine
	drawQuad(maxLine, maxColumn, cornerTopRight, cornerTopLeft, cornerBottomRight,cornerBottomLeft, edgeLine, edgeColumn)
}

func main() {
	quadA(3, 3)
	fmt.Println()
	quadB(1, 5)
	fmt.Println()
	quadC(2, 4)
	fmt.Println()
	quadD(7, 5)
	fmt.Println()
	quadE(6, 2)
	fmt.Println()
	quadE(4, 1)
}
