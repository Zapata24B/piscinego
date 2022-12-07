package main

import (
	"fmt"
	// "github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	numberArgs := len(args)
	firstArg := args[0]
	if numberArgs == 0 || firstArg == "-h" || firstArg == "--help" {
		fmt.Print("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
	} else if firstArg == "-o" || firstArg == "--order" {
		result := orderTask(args[1])
		fmt.Println(result)
	} else if firstArg[:3] == "-i=" || firstArg[:9] == "--insert=" {
		startInsert := 9
		if firstArg[:3] == "-i=" {
			startInsert = 3
		}
		insertValue := firstArg[startInsert:]
		lastArg := args[len(args)-1]
		nextArg := args[1]
		result := insertTask(lastArg, insertValue)
		if nextArg == "-o" || nextArg == "--order" {
			result = orderTask(result)
		}
		fmt.Println(result)
	} else {
		fmt.Println(firstArg)
	}
}

func insertTask(str1 string, str2 string) string {
	return str1 + str2
}

func orderTask(str string) string {
	runes := []rune(str)
	for i := range runes {
		for j := range runes {
			if runes[i] < runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}
