package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	const filename = "test.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	if2()

	fmt.Println(
		learnswitch(100),
		learnswitch(80),
		learnswitch(60),
	)
}

func if2() {
	const filename = "test.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n%s\n\n", contents)
	}
}

func learnswitch(score int) string {
	result := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		return "F"
	case score < 80:
		return "C"
	case score < 90:
		return "B"
	case score <= 100:
		return "A"
	}
	return result
}
