package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func appendMessage(message string) string {
	str := fmt.Sprintf("%s", strings.ToUpper(message))
	return str

}

func copy_files_text() {
	now := time.Now()
	file, err := os.Create(strconv.Itoa(int(now.Month())) + "-" + strconv.Itoa(now.Day()) + "-" + strconv.Itoa(now.Year()) + ".txt")
	file.WriteString(read_typed())
	if err != nil {
		fmt.Println(err)
	}

}

func read_typed() string {
	pth := "Tasks/todos.txt"
	content, err := os.ReadFile(pth)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(content))

	return string(content)

}

func Get(el string) string {
	arg := fmt.Sprintf("%s", el)
	return arg
}

func main() {

	args := os.Args

	for {

		if Get(args[1]) == strings.ToUpper("y") {

			copy_files_text()
			break

		} else {

			copy_files_text()
			break
		}
	}

	fmt.Printf("%s", appendMessage("Completed"))

}
