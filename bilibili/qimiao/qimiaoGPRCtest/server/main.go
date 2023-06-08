package main

import (
	"context"
	"net"
	"qimiaoGPRCtest/pb/person"

	"google.golang.org/grpc"
)

type personServe struct {
	person.UnimplementedSearchServiceServer
}

func (*personServe) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{Name: "我收到了" + name + "的消息"}
	return res, nil
}
func (*personServe) SearchIn(person.SearchService_SearchInServer) error {
	return nil
}
func (*personServe) SearchOut(*person.PersonReq, person.SearchService_SearchOutServer) error {
	return nil
}
func (*personServe) SearchIO(person.SearchService_SearchIOServer) error {
	return nil
}
func main() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &personServe{})
	s.Serve(l)
}
