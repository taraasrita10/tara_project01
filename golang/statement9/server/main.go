package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"
	trie "statement6/pkg"
	pb "statement6/proto"

	"google.golang.org/grpc"
)

const dataPath = "/data/trie_db.txt"

func loadTrie(trie *trie.Trie) {
	file, err := os.Open(dataPath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatalf("failed to open storage: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		trie.Add(word)
	}
}

type server struct {
	pb.UnimplementedTrieServiceServer
	myTrie *trie.Trie
}

func (s *server) Add(ctx context.Context, req *pb.WordRequest) (*pb.StatusResponse, error) {
	s.myTrie.Add(req.Word)

	f, err := os.OpenFile(dataPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return &pb.StatusResponse{Success: false}, err
	}
	defer f.Close()

	if _, err := f.WriteString(req.Word + "\n"); err != nil {
		return &pb.StatusResponse{Success: false}, err
	}

	return &pb.StatusResponse{Success: true}, nil
}

func (s *server) Check(ctx context.Context, req *pb.WordRequest) (*pb.CheckResponse, error) {
	exists := s.myTrie.Check(req.Word)
	return &pb.CheckResponse{Exists: exists}, nil
}

func (s *server) List(ctx context.Context, req *pb.Empty) (*pb.ListResponse, error) {
	words := s.myTrie.List()
	return &pb.ListResponse{Words: words}, nil
}

func (s *server) Remove(ctx context.Context, req *pb.WordRequest) (*pb.StatusResponse, error) {
	s.myTrie.Remove(req.Word)

	f, err := os.OpenFile(dataPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return &pb.StatusResponse{Success: false}, err
	}
	defer f.Close()

	words := s.myTrie.List()
	for _, w := range words {
		f.WriteString(w + "\n")
	}

	return &pb.StatusResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	myTrieInstance := trie.NewTrie()
	loadTrie(myTrieInstance)

	s := grpc.NewServer()

	pb.RegisterTrieServiceServer(s, &server{myTrie: myTrieInstance})

	log.Println("Server running on port 50051 and storage loaded...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
