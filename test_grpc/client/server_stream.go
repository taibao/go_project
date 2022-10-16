package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	service "test_grpc/service"
)
// Address 连接地址
const Address string = ":8000"
var grpcClient service.StreamServerClient

func main(){
	// 连接服务器
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接服务端失败: %s", err)
		return
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = service.NewStreamServerClient(conn)
	listValue()
	fmt.Printf("调用成功")
}


// listValue 调用服务端的ListValue方法
func listValue() {
	// 创建发送结构体
	req := service.SimpleRequest{
		Data: "stream server grpc ",
	}
	// 调用我们的服务(ListValue方法)
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call ListStr err: %v", err)
	}
	for {
		//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
		res, err := stream.Recv()
		// 判断消息流是否已经结束
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.StreamValue)
	}
}