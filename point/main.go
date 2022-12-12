package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	str := "x = " + string(rune(points.x/10%10+48)) + string(rune(points.x%10+48)) + ", y = " + string(rune(points.y/10%10+48)) + string(rune(points.y%10+48))

	printStr(str)
}
