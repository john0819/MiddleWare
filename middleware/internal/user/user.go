package user

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	userv1 "github.com/john0819/John-Proto/gen/go/proto/user/v1"
	"github.com/zeromicro/go-zero/zrpc"
)

type UserService struct {
	userv1.UnimplementedUserServiceServer
}

func (s *UserService) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	log.Println("access to GetUser")
	return &userv1.GetUserResponse{
		User: &userv1.User{
			Id:   req.Id,
			Name: "John wong grpc",
		},
	}, nil
}

// 用zrpc启动一个grpc服务
func RunUserService(c zrpc.RpcServerConf) {
	server := zrpc.MustNewServer(c, func(grpcServer *grpc.Server) {
		userv1.RegisterUserServiceServer(grpcServer, &UserService{})
	})
	defer server.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	server.Start()
}

// todo: zrpc使用
