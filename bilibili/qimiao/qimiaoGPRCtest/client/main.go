package main

import (
	"context"
	"fmt"
	"qimiaoGPRCtest/pb/person"

	"google.golang.org/grpc"
)

func main() {
	l, _ := grpc.Dial(":8888", grpc.WithInsecure())
	client := person.NewSearchServiceClient(l)
	res, err := client.Search(context.Background(), &person.PersonReq{Name: "qimiao 的 GRPC test class"})
	if err != nil {
		fmt.Println("Sorry client.Search has been failed!")
		panic(err)
	}
	fmt.Println(res)
}
