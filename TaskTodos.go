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

func main() {

	copy_files_text()

	fmt.Printf("%s", appendMessage("Completed"))

}
