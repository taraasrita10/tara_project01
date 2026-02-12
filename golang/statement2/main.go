package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// Takes input from STDIN till the new line character occurs
// Trims whitespaces
func getInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return strings.TrimSpace(input), nil
	}
	return strings.TrimSpace(input), nil
}

// Checks for the existance of a string in the list
// Returns a boolean value indicating its existance
func check(word string, arr []string) bool {
	return slices.Contains(arr, word)
}

// Lists out the contents of the list
func list(arr []string, w io.Writer) {
	if len(arr) == 0 {
		fmt.Fprintln(w, "The list is empty!")
		return
	}
	for _, word := range arr {
		fmt.Fprintln(w, word)
	}
}

// Adds the given word to the list
func add(word string, arr []string) []string {
	if word == "" {
		fmt.Println("Empty spaces are not valid input!")
		return nil
	}
	if check(word, arr) == true {
		fmt.Printf("%s is already present in the string, duplicates cannot be added again!\n", word)
		return nil
	}
	arr = append(arr, word)
	fmt.Printf("%s has been added!\n", word)
	return arr
}

// Removes the given word from the list
func remove(word string, arr []string) []string {
	if word == "" {
		fmt.Println("Empty spaces are not valid input!")
		return nil
	}
	if len(arr) == 0 {
		fmt.Println("The list is empty!")
		return nil
	}
	for i, w := range arr {
		if w == word {
			arr = append(arr[:i], arr[i+1:]...)
			fmt.Printf("%s has been removed!\n", word)
			return arr
		}
	}
	fmt.Printf("%s does not exist!\n", word)
	return nil
}

func main() {
	var arr []string
	reader := bufio.NewReader(os.Stdin)
	quit := false
	for quit != true {
		fmt.Println("Commands : Add / List / Check / Remove / Quit")
		command, err := getInput(reader, "Enter the command : ")
		if err != nil && err != io.EOF {
			fmt.Println("Error reading input!")
			break
		}
		switch command {
		case "add":
			word, err := getInput(reader, "Enter the word : ")
			if err != nil && err != io.EOF {
				fmt.Println("Error reading input!")
				break
			}
			arr = add(word, arr)
		case "list":
			list(arr, os.Stdout)
		case "check":
			var word string
			word, err := getInput(reader, "Enter the word : ")
			if err != nil && err != io.EOF {
				fmt.Println("Error reading input!")
				break
			}
			if word == "" {
				fmt.Println("Empty spaces are not valid input!")
			} else {
				if check(word, arr) {
					fmt.Printf("%s is present!\n", word)
				} else {
					fmt.Printf("%s is not present!\n", word)
				}
			}
		case "remove":
			word, err := getInput(reader, "Enter the word : ")
			if err != nil && err != io.EOF {
				fmt.Println("Error reading input!")
				break
			}
			arr = remove(word, arr)
		case "quit":
			quit = true
		default:
			fmt.Println("The command entered is wrong, please check and enter again!")
		}
	}
}
