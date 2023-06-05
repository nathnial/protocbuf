package main

import (
	"context"
	"net"
	"qimiaoGPRCtest/pb/person"

	"google.golang.org/grpc"
)

type personServer struct {
	person.UnimplementedSearchServiceServer
}

func (*personServer) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	return nil, nil
}
func (*personServer) SearchIn(person.SearchService_SearchInServer) error {
	return nil
}
func (*personServer) SearchOut(*person.PersonReq, person.SearchService_SearchOutServer) error {
	return nil
}
func (*personServer) SearchIO(person.SearchService_SearchIOServer) error {
	return nil
}
func main() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &personServer{})
	s.Serve(l)
}
