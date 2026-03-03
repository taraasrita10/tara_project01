package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	pb "statement6/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	if err == io.EOF {
		return "", err
	}

	if err != nil {
		fmt.Println("Error reading input:", err)
		return "", err
	}

	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		fmt.Println("Empty spaces are not valid input!")
		return "", fmt.Errorf("Empty input")
	}

	return trimmed, nil
}

func main() {
	conn, err := grpc.NewClient("trie-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTrieServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	quit := false

	for quit != true {
		fmt.Println("Commands : Add / List / Check / Remove / Quit")
		command, err := getInput(reader, "Enter the command : ")
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}

		command = strings.ToLower(command)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

		switch command {
		case "add":
			word, err := getInput(reader, "Enter the word : ")
			if err != nil {
				cancel()
				continue
			}
			res, err := c.Add(ctx, &pb.WordRequest{Word: word})
			if err != nil {
				fmt.Printf("Error adding word: %v\n", err)
			} else if res.GetSuccess() {
				fmt.Printf("%s has been added!\n", word)
			}

		case "list":
			res, err := c.List(ctx, &pb.Empty{})
			if err != nil {
				fmt.Printf("Error listing words: %v\n", err)
			} else {
				words := res.GetWords()
				if len(words) == 0 {
					fmt.Println("The list is empty!")
				} else {
					for _, w := range words {
						fmt.Println(w)
					}
				}
			}

		case "check":
			word, err := getInput(reader, "Enter the word : ")
			if err != nil {
				cancel()
				continue
			}
			res, err := c.Check(ctx, &pb.WordRequest{Word: word})
			if err != nil {
				fmt.Printf("Error checking word: %v\n", err)
			} else {
				if res.GetExists() {
					fmt.Printf("%s is present!\n", word)
				} else {
					fmt.Printf("%s is not present!\n", word)
				}
			}

		case "remove":
			word, err := getInput(reader, "Enter the word : ")
			if err != nil {
				cancel()
				continue
			}
			res, err := c.Remove(ctx, &pb.WordRequest{Word: word})
			if err != nil {
				fmt.Printf("Error removing word: %v\n", err)
			} else if res.GetSuccess() {
				fmt.Printf("%s has been removed!\n", word)
			} else {
				fmt.Printf("%s does not exist!\n", word)
			}

		case "quit":
			quit = true
		default:
			fmt.Println("The command entered is wrong, please check and enter again!")
		}
		cancel()
	}
}
