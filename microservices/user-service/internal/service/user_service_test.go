package service

import (
	"context"
	"testing"

	pb "taskmanager/microservices/user-service/pb"

	"github.com/stretchr/testify/assert"
)

func TestUserService_FullFlow(t *testing.T) {
	ctx := context.Background()
	repo := NewMockUserRepo()
	svc := NewUserService(repo)

	// 1️⃣ Create user
	createResp, err := svc.CreateUser(ctx, &pb.CreateUserRequest{
		Username: "john",
		Password: "123456",
	})
	assert.NoError(t, err)
	assert.Equal(t, "john", createResp.Username)
	assert.NotEmpty(t, createResp.Token)
	assert.NotZero(t, createResp.Id)

	// 2️⃣ Login thành công
	loginResp, err := svc.Login(ctx, &pb.LoginRequest{
		Username: "john",
		Password: "123456",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, loginResp.Token)

	// 3️⃣ Login sai password
	_, err = svc.Login(ctx, &pb.LoginRequest{
		Username: "john",
		Password: "wrongpass",
	})
	assert.Error(t, err)

	// 4️⃣ GetUser
	getResp, err := svc.GetUser(ctx, &pb.GetUserRequest{
		Id: createResp.Id,
	})
	assert.NoError(t, err)
	assert.Equal(t, "john", getResp.Username)

	// 5️⃣ UpdateUser (đổi username và password)
	updateResp, err := svc.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:       createResp.Id,
		Username: "johnny",
		Password: "newpass",
	})
	assert.NoError(t, err)
	assert.Equal(t, "johnny", updateResp.Username)

	// 6️⃣ Login lại với password mới
	loginResp2, err := svc.Login(ctx, &pb.LoginRequest{
		Username: "johnny",
		Password: "newpass",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, loginResp2.Token)

	// 7️⃣ ListUsers
	listResp, err := svc.ListUsers(ctx, &pb.ListUsersRequest{})
	assert.NoError(t, err)
	assert.Len(t, listResp.Users, 1)
	assert.Equal(t, "johnny", listResp.Users[0].Username)

	// 8️⃣ DeleteUser
	deleteResp, err := svc.DeleteUser(ctx, &pb.DeleteUserRequest{
		Id: createResp.Id,
	})
	assert.NoError(t, err)
	assert.True(t, deleteResp.Success)

	// 9️⃣ GetUser sau khi xoá (phải lỗi)
	_, err = svc.GetUser(ctx, &pb.GetUserRequest{Id: createResp.Id})
	assert.Error(t, err)

	// 🔟 ListUsers sau khi xoá (phải rỗng)
	listResp2, err := svc.ListUsers(ctx, &pb.ListUsersRequest{})
	assert.NoError(t, err)
	assert.Len(t, listResp2.Users, 0)
}
