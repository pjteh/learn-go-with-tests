package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Hello(name string) string {
	if name != "" {
		return "Hello, " + name + "!"
	}
	return "Hello, World!"
}

func World() string {
	return "Awesome!"
}

type Book struct {
	title  string
	author string
	copies int
}

func main() {

	// var choice int
	// var title, newtitle, newauthor, author string
	// var copies, newcopies int

	// for {
	// 	newBook := Book{}
	// 	fmt.Println("1. Add book")
	// 	fmt.Println("2. Update book")
	// 	fmt.Println("Enter your choice: ")
	// 	fmt.Scanln(&choice)
	// 	switch choice {
	// 	case 1:
	// 		fmt.Println("Enter Title: ")
	// 		fmt.Scanln(&title)
	// 		fmt.Print("Enter Author: ")
	// 		fmt.Scanln(&author)
	// 		fmt.Print("Enter Copies: ")
	// 		fmt.Scanln(&copies)
	// 		newBook.title = title
	// 		newBook.author = author
	// 		newBook.copies = copies
	// 		fmt.Println(newBook)
	// 	case 2:
	// 		fmt.Println("Enter Title: ")
	// 		fmt.Scanln(&title)
	// 		fmt.Println("Update Title: ")
	// 		fmt.Scanln(&newtitle)
	// 		fmt.Println("Update Author: ")
	// 		fmt.Scanln(&newauthor)
	// 		fmt.Println("Update Copies: ")
	// 		fmt.Scanln(&newcopies)

	// 		if newBook.title == title {
	// 			newBook.title = newtitle
	// 			newBook.author = newauthor
	// 			newBook.copies = newcopies
	// 		}
	// 		fmt.Println(newBook)

	// 	}
	// }

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Who is your fave author?")
	author, _ := reader.ReadString('\n')
	author = author[:len(author)-1] // Trim the newline character

	fmt.Println("What book is your fave?")
	title, _ := reader.ReadString('\n')
	title = title[:len(title)-1]

	fmt.Println("How many copies you have?")
	input, _ := reader.ReadString('\n')
	copies, err := strconv.Atoi(input[:len(input)-1])
	if err != nil {
		fmt.Println("Invalid number of copies")
		return
	}

	fmt.Printf("You have %v copies of %v written by %v\n", copies, title, author)
}
