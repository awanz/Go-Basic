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
	fmt.Println("6. Exit")
	fmt.Print("Choose menu: ")
	scanner := bufio.NewScanner(os.Stdin)                        // declar scanner
	scanner.Scan()                                               // wait input
	inputTerminal, _ := strconv.ParseInt(scanner.Text(), 10, 64) // when user input, program read text and convert to integer with strconv

	switch inputTerminal {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
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
	default:
		fmt.Println("Wrong input, try again!")
	}
}
