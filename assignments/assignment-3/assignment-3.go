package main

import (
	"bufio"
	"fmt"
	helpers "go-academy/helpers"
	"log"
	"os"
	"strconv"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter a number")

	numStr, err := reader.ReadString(('\n'))

	if err != nil {
		log.Fatal(err)
	}

	numStr = helpers.RemoveNewlineFromStr(numStr)

	numInt, _ := strconv.Atoi(numStr)

	fmt.Println(numInt)

	if (int(numInt) < 1) || (int(numInt) > 10) {
		fmt.Printf("%d is not between 1 and 10\n", numInt)
	} else {
		fmt.Printf("%d is between 1 and 10\n", numInt)
	}

}
