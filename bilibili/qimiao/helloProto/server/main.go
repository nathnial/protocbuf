package main

import (
	"context"
	"fmt"
	hello_proto "helloProto/pb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	hello_proto.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, req *hello_proto.Req) (res *hello_proto.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_proto.Res{Message: "~~~~~~我是从服务端返回的 Grpc 的内容~~~~~~"}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("net.Listen has been failed")
		panic(err)
	}
	s := grpc.NewServer()
	hello_proto.RegisterHelloGRPCServer(s, &server{})
	s.Serve(l)
}
