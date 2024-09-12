package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/sicp618/interview/proto/user"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	if in.Id == 1 {
		return &pb.UserResponse {
			Data: &pb.UserInfo {
				Id: 1,
				Name: "jack",
			},
		}, nil
	}

	return nil, status.Errorf(codes.NotFound, "no find")
}



func main() {
    fmt.Println("Starting gRPC server...")

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, &Server{})

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}