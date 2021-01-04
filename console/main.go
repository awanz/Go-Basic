package main

import (
	"bufio" // read input user
	"fmt"
	"os"      // os input
	"strconv" // for convert string to integer
)

func main() {
	// var inputMenu int
	fmt.Println("===============================")
	fmt.Println("             Employe")
	fmt.Println("===============================")
	fmt.Println("1. Add data")
	fmt.Println("2. List Data")
	fmt.Println("3. Update Data")
	fmt.Println("4. Delete Data")
	fmt.Println("5. Report")
	fmt.Println("6. Example Code Loop")
	fmt.Println("7. Example Code Array")
	fmt.Println("8. Example Code Slice")
	fmt.Println("9. Example Code Map")
	fmt.Println("10. The Biggest Number with Function")
	fmt.Print("Choose menu: ")
	scanner := bufio.NewScanner(os.Stdin)                        // declar scanner
	scanner.Scan()                                               // wait input
	inputTerminal, _ := strconv.ParseInt(scanner.Text(), 10, 64) // when user input, program read text and convert to integer with strconv

	switch inputTerminal {
	case 1:
		fmt.Println("1")
	case 2:
		display()
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	case 6:
		fmt.Println("   1. For")
		fmt.Println("   2. While")
		fmt.Println("   3. Range")
		fmt.Print("   Choose menu: ")
		scanner.Scan()
		inputTerminal, _ = strconv.ParseInt(scanner.Text(), 10, 64)

		switch inputTerminal {
		case 1:
			for index := 0; index < 10; index++ {
				fmt.Println("   Loop:", index)
			}
		case 2:
			var i int = 0
			for i < 3 {
				fmt.Println("   Loop:", i)
				i++
			}
		case 3:
			color := []string{"red", "yellow", "blue"}

			for index, dt := range color {
				fmt.Println("   Loop:", index, dt)
			}
		}
	case 7:
		var word [2]string
		word[0] = "Hello"
		word[1] = "World!"

		fmt.Println(word[0] + " " + word[1])

		var number [2][2]int

		number[0][0] = 1
		number[0][1] = 2
		number[1][0] = 3
		number[1][1] = 4

		fmt.Print(number[0][0])
		fmt.Print(number[0][1])
		fmt.Print(number[1][0])
		fmt.Println(number[1][1])
	case 8:
		s := make([]string, 3)
		s[0] = "a"
		s[1] = "b"
		s[2] = "c"
		fmt.Println("set:", s)
		fmt.Println("get:", s[2])
		fmt.Println("len:", len(s)) // total array
	case 9:
		var mahasiswa = map[string]string{}

		mahasiswa["Nama"] = "Didik Prabowo"
		mahasiswa["alamat"] = "Wonogiri"

		fmt.Println(mahasiswa)
	case 10:
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Number one: ")
		scanner.Scan()
		inputTerminalA, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		fmt.Print("Number two: ")
		scanner.Scan()
		inputTerminalB, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		fmt.Print("Result: ")
		fmt.Println(biggestNumber(int(inputTerminalA), int(inputTerminalB)))
	default:
		fmt.Println("Wrong input, try again!")
	}
}

func display() {
	fmt.Println("Display!")
}

func biggestNumber(a int, b int) int {
	var result int
	if a > b {
		result = a
	} else {
		result = b
	}
	return result
}
