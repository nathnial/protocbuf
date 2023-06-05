package main

import (
	"context"
	"fmt"
	hello_proto "helloProto/pb"

	"google.golang.org/grpc"
)

func main() {
	// 如果不设置第二个参数会有ERROR: no transport security set (use grpc.WithTransportCredentials(insecure.NewCredentials()) explicitly or set credentials)
	// grpc.WithInsecure() 已经被弃用
	// grpc.WithTransportCredentials(credentials.NewTLS(nil))
	// insecure.NewCredentials()
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc.Dial has been failed!")
		panic(err)
	}
	defer conn.Close()
	cli := hello_proto.NewHelloGRPCClient(conn)
	r, _ := cli.SayHi(context.Background(), &hello_proto.Req{Message: "~~~~~~ 我是 GRPC 客户端 ~~~~~~"})
	fmt.Println(r.GetMessage())
}
