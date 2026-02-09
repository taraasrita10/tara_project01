package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Takes input from STDIN till the new line character occurs
// Trims whitespaces
func getInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Checks for the existance of a string in the list
// Returns a boolean value indicating its existance
func check(word string, arr *[]string) bool {
	return slices.Contains(*arr, word)
}

// Lists out the contents of the list
func list(arr *[]string) {
	if len(*arr) == 0 {
		fmt.Println("The list is empty!")
		return
	}
	for _, word := range *arr {
		fmt.Println(word)
	}
}

// Adds the given word to the list
func add(word string, arr *[]string) {
	if word == "" {
		fmt.Println("Empty spaces are not valid input!")
		return
	}
	if check(word, arr) == true {
		fmt.Printf("%s is already present in the string, duplicates cannot be added again!\n", word)
		return
	}
	*arr = append(*arr, word)
	fmt.Printf("%s has been added!\n", word)
}

// Removes the given word from the list
func remove(word string, arr *[]string) {
	if word == "" {
		fmt.Println("Empty spaces are not valid input!")
		return
	}
	if len(*arr) == 0 {
		fmt.Println("The list is empty!")
		return
	}
	for i, w := range *arr {
		if w == word {
			*arr = append((*arr)[:i], (*arr)[i+1:]...)
			fmt.Printf("%s has been removed!\n", word)
			return
		}
	}
	fmt.Printf("%s does not exist!\n", word)
}

func main() {
	var arr []string
	reader := bufio.NewReader(os.Stdin)
	quit := false
	var command string
	for quit != true {
		fmt.Println("Commands : Add / List / Check / Remove / Quit")
		command = strings.ToLower(getInput(reader, "Enter the command : "))
		switch command {
		case "add":
			word := getInput(reader, "Enter the word : ")
			add(word, &arr)
		case "list":
			list(&arr)
		case "check":
			var word string
			word = getInput(reader, "Enter the word : ")
			if word == "" {
				fmt.Println("Empty spaces are not valid input!")
			} else {
				if check(word, &arr) {
					fmt.Printf("%s is present!\n", word)
				} else {
					fmt.Printf("%s is not present!\n", word)
				}
			}
		case "remove":
			word := getInput(reader, "Enter the word : ")
			remove(word, &arr)
		case "quit":
			quit = true
		default:
			fmt.Println("The command entered is wrong, please check and enter again!")
		}
	}
}
