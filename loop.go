package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
)

func convertToBin(n int) string {
	result := ""
	if n == 0 {
		return fmt.Sprint("unsupported operation:", n)
	} else {
		for ; n > 0; n /= 2 {
			lsb := n % 2
			result = strconv.Itoa(lsb) + result
		}
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

//死循环
func forever(){
	for{
		printFile("test.txt")
	}
}

func main() {
	fmt.Println(convertToBin(0))
	fmt.Println(convertToBin(520))
	printFile("test.txt")
	//forever()
}
