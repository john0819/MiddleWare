package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userv1 "github.com/john0819/John-Proto/gen/go/proto/user/v1"
)

func TestRunUserService(t *testing.T) {
	// 请求grpc服务
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := userv1.NewUserServiceClient(conn)

	// 调用GetUser方法
	response, err := client.GetUser(context.Background(), &userv1.GetUserRequest{
		Id: "222333",
	})
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}
	assert.Equal(t, response.User.Id, "222333")
	assert.Equal(t, response.User.Name, "John wong grpc")
}
