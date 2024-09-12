package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sicp618/interview/nostr/redis"

	pb "github.com/sicp618/interview/proto/user"
	"google.golang.org/grpc"
)

func main() {
    fmt.Println("Nostr main")
    redis.Init()
	

    // 创建 gRPC 客户端连接
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewUserServiceClient(conn)

    // 创建请求上下文
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // 创建请求
    req := &pb.UserRequest{
        Id: 2,
    }

    // 发送请求
    resp, err := client.GetUser(ctx, req)
    if err != nil {
        log.Printf("could not get user: %v", err)
    }
    fmt.Printf("User: %v\n", resp)

	req.Id = 1
	resp, err = client.GetUser(ctx, req)
    if err != nil {
        log.Printf("could not get user: %v", err)
    }

    // 打印响应
    fmt.Printf("User: %v\n", resp)
}
