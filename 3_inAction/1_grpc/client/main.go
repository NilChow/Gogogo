package main

import (
	"fmt"
	"context"

	"Gogogo/3_inAction/1_grpc/myProto"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("gRpc Client Start...")

	conn, err := grpc.Dial("127.0.0.1:5555", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("[Dial To gRpc Server Failed]: [%s]\n", err)
	}
	defer conn.Close()

	c := myProto.NewMyGRpcClient(conn)

	req := "3+5"

	res, err := c.Caculate(context.Background(), &myProto.ReqMsg{MsgStr : req})
	if err != nil {
		fmt.Printf("[gRpc Error]: [%s]\n", err)
		panic("[gRpc Error]")
	}
	fmt.Printf("[gRpc Success] [%s=%d]\n", req, res.MsgInt)
}