# List Manager

This is a list manager built in Go, which performs the actions of adding, removing, checking whether a word exists in the list and listing out the list. The program runs in a loop. It uses tries to store the data.

## Functions and Tests

Add - Adds a string to the list, tested for generic adding, adding of a duplicate and adding and empty string.
Remove - Removes a string from the list, tested for generic removal, removal of a non-existant string and removal from an empty list
Check - Checks for the existance of a string, tested for existance, non-existance and empty string input
List - Lists out the strings in the list
Quit - An option provided in the main function in the loop to exit from the program

## Internal structure

The usage of tries to store data reduces the time complexity to perform a search action from O(n) (where n is the size of the list) to O(l) (where l is the length of the word).

## Usage

To run the main program
```bash
go run main.go
```

To run the tests
```bash
go test -v
```

## Working

The program runs in the form of a loop, through switch case statements within a for loop. The for loop ends when the quit option is chosen. The program throws an alert when a wrong command is typed out, and has alerts for other wrong entries as well.



