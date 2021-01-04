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
	fmt.Println("6. Exit")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose menu: ")
	inputBufio, _ := reader.ReadString('\n')
	// inputMenu = strconv.Atoi(inputBufio)

	fmt.Println(strconv.Atoi(inputBufio))
}
