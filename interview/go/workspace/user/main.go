package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/go-redis/redis/v8"
	pb "github.com/sicp618/interview/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db  *gorm.DB
var redisClient *redis.Client

func initPostgres() {
    var err error
    dsn := "postgres://user:password@localhost:5432/nostr?TimeZone=Asia/Shanghai"
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to PostgreSQL: %v", err)
    }

    // 自动迁移模式
    err = db.AutoMigrate(&User{})
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }
}

func initRedis() {
    redisClient = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
}

type User struct {
    ID    int32  `gorm:"primaryKey"`
    Name  string `gorm:"column:name"`
    Email string `gorm:"column:email"`
}

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
		cacheKey := fmt.Sprintf("user:%d", req.Id)
		userCache, err := redisClient.Get(ctx, cacheKey).Result()
		if err == redis.Nil {
			var user User
			result := db.First(&user, req.Id)
			if result.Error != nil {
				if result.Error == gorm.ErrRecordNotFound {
					return nil, status.Errorf(codes.NotFound, "user not found")
				}
				return nil, status.Errorf(codes.Internal, "failed to query user: %v", result.Error)
			}
	
			err = redisClient.Set(ctx, cacheKey, fmt.Sprintf("%s,%s", user.Name, user.Email), time.Hour).Err()
			if err != nil {
				log.Printf("failed to set cache: %v", err)
			}
	
			return &pb.UserResponse{
				Data: &pb.UserInfo{
					Id:    user.ID,
					Name:  user.Name,
					Email: user.Email,
				},
			}, nil
		} else if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get cache: %v", err)
		}
	
		var name, email string
		fmt.Sscanf(userCache, "%s,%s", &name, &email)
	
		return &pb.UserResponse{
			Data: &pb.UserInfo{
				Id:    req.Id,
				Name:  name,
				Email: email,
			},
		}, nil
}

func main() {
	initPostgres()
    initRedis()

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