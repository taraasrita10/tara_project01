package main

import (
	"context"
	"log"
	"net"
	trie "statement6/pkg"
	pb "statement6/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTrieServiceServer
	myTrie *trie.Trie
}

func (s *server) Add(ctx context.Context, req *pb.WordRequest) (*pb.StatusResponse, error) {
	s.myTrie.Add(req.Word)
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
	return &pb.StatusResponse{Success: true}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterTrieServiceServer(s, &server{myTrie: trie.NewTrie()})
	log.Println("Server running on port 50051...")
	s.Serve(lis)
}
