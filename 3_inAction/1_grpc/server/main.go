package main

import (
	"fmt"
	"net"
	"strconv"
	"errors"
	"strings"
	"context"

	"Gogogo/3_inAction/1_grpc/myProto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Caculate(ctx context.Context, req *myProto.ReqMsg) (res *myProto.ResMsg, err error) {
	fmt.Printf("[客户端计算请求]: [%s]\n", req.MsgStr)

	var CaclType string
	var subs []string
	if strings.Contains(req.MsgStr, "+") {
		subs = strings.Split(req.MsgStr, "+")
		CaclType = "+"

	} else if strings.Contains(req.MsgStr, "-") {
		subs = strings.Split(req.MsgStr, "-")
		CaclType = "-"

	} else {
		return nil, errors.New("[ReqStr Error]")
	}

	if len(subs) != 2 {
		return nil, errors.New("[ReqStr Error]")
	}

	a, _ := strconv.Atoi(subs[0])
	b, _ := strconv.Atoi(subs[1])

	res = &myProto.ResMsg{}
	if CaclType == "+" {
		res.MsgInt = int32(a+b)
	} else {
		res.MsgInt = int32(a+b)
	}
	
	return
}

func main() {
	fmt.Println("gRpc Server Start...")

	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		fmt.Println("[Server Listen Failed]")
	}

	s := grpc.NewServer()

	myProto.RegisterMyGRpcServer(s, &server{})
	reflection.Register(s)
	err = s.Serve(listener)
	if err != nil {
		fmt.Printf("[Failed To Server]: [%s]\n", err)
	}
}