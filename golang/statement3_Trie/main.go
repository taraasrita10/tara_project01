package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type Node struct {
	children map[rune]*Node
	isEnd    bool
}

func newNode() *Node {
	return &Node{
		children: make(map[rune]*Node),
		isEnd:    false,
	}
}

type Trie struct {
	root *Node
}

func newTrie() *Trie {
	return &Trie{
		root: newNode(),
	}
}

func (t *Trie) Add(word string) {
	currNode := t.root
	for _, char := range word {
		_, exists := currNode.children[char]
		if !exists {
			currNode.children[char] = newNode()
		}
		currNode = currNode.children[char]
	}
	currNode.isEnd = true
}

func (t *Trie) Check(word string) bool {
	currNode := t.root
	for _, char := range word {
		_, exists := currNode.children[char]
		if !exists {
			return false
		}
		currNode = currNode.children[char]
	}
	return currNode.isEnd
}

func (t *Trie) RemoveHelper(currNode *Node, word string, index int) bool {
	if index == len(word) {
		if !currNode.isEnd {
			return false
		}
		currNode.isEnd = false
		return len(currNode.children) == 0
	}
	char := rune(word[index])
	nextNode, exists := currNode.children[char]
	if !exists {
		return false
	}
	deleteChild := t.RemoveHelper(nextNode, word, index+1)
	if deleteChild {
		delete(currNode.children, char)
		return len(currNode.children) == 0 && currNode.isEnd
	}

	return false
}

func (t *Trie) Remove(word string) {
	t.RemoveHelper(t.root, word, 0)
}

func (t *Trie) ListHelper(currNode *Node, prefix []rune, results *[]string) {
	if currNode.isEnd {
		*results = append(*results, string(prefix))
	}
	for char, nextNode := range currNode.children {
		newPrefix := append(prefix, char)
		t.ListHelper(nextNode, newPrefix, results)
	}
}

func (t *Trie) List() []string {
	results := []string{}
	t.ListHelper(t.root, []rune{}, &results)
	slices.Sort(results)
	return results
}

func getInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return strings.TrimSpace(input), err
	}
	return strings.TrimSpace(input), nil
}

func main() {
	myTrie := newTrie()
	reader := bufio.NewReader(os.Stdin)
	quit := false
	for quit != true {
		fmt.Println("Commands : Add / List / Check / Remove / Quit")
		command, err := (getInput(reader, "Enter the command : "))
		command = strings.ToLower(command)
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
			if word == "" {
				fmt.Println("Empty strings are not valid input!")
				continue
			}
			if myTrie.Check(word) {
				fmt.Println("String does not exist!")
				continue
			}
			myTrie.Add(word)
		case "list":
			results := myTrie.List()
			for _, w := range results {
				fmt.Println(w)
			}
		case "check":
			word, err := getInput(reader, "Enter the word : ")
			if err != nil && err != io.EOF {
				fmt.Println("Error reading input!")
				break
			}
			if word == "" {
				fmt.Println("Empty strings are not valid input!")
			} else {
				if myTrie.Check(word) {
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
			if word == "" {
				fmt.Println("Empty strings are not valid input!")
				continue
			}
			if !myTrie.Check(word) {
				fmt.Println("String does not exist!")
				continue
			}
			myTrie.Remove(word)
		case "quit":
			quit = true
		default:
			fmt.Println("The command entered is wrong, please check and enter again!")
		}
	}
}
