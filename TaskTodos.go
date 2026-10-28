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

func emptyFile(args []string) string {

	if args[1] == "-e" {
		now := time.Now()
		file, _ := os.Create(strconv.Itoa(int(now.Month())) + "-" + strconv.Itoa(now.Day()) + "-" + strconv.Itoa(now.Year()) + ".txt")
		file.WriteString(read_typed())

	}

	return args[1]
}

func copy_files_text() {
	now := time.Now()
	file, err := os.Create(strconv.Itoa(int(now.Month())) + "-" + strconv.Itoa(now.Day()) + "-" + strconv.Itoa(now.Year()) + ".txt")
	file.WriteString(read_typed())

	err_handle := ErrorHandle{
		e: err,
	}

	if err_handle.e != nil {
		fmt.Println(err_handle.e.Error())
	}

}

func read_typed() string {
	pth := "Tasks/todos.txt"
	content, err := os.ReadFile(pth)

	err_handle := ErrorHandle{
		e: err,
	}

	if err_handle.e != nil {
		fmt.Println(err_handle.e.Error())
	}

	fmt.Println(string(content))

	return string(content)

}

func Get(el string) string {
	arg := fmt.Sprintf("%s", el)
	return arg
}

type ErrorHandle struct {
	e error
}

func main() {

	args := os.Args

	for {

		if len(Get(args[0])) < 1 {

			fmt.Println(strings.ToUpper("You Didn't Pass Arguments"))
			emptyFile(args)
		} else if Get(args[1]) == "-e" {
			emptyFile(args)

		}

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
