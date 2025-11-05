package handlers

import (
	"net/http"
	"strconv"
	"strings"

	pb "taskmanager/gateway/pb"

	"github.com/gin-gonic/gin"
)

// UserHandler chứa gRPC client
type UserHandler struct {
	UserClient pb.UserServiceClient
}

// NewUserHandler khởi tạo handler
func NewUserHandler(client pb.UserServiceClient) *UserHandler {
	return &UserHandler{UserClient: client}
}

// Login
func (h *UserHandler) Login(c *gin.Context) {
	var req pb.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	res, err := h.UserClient.Login(c.Request.Context(), &req)
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// CREATE: POST /users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req pb.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	res, err := h.UserClient.CreateUser(c.Request.Context(), &req)
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// READ: GET /users/:id
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	res, err := h.UserClient.GetUser(c.Request.Context(), &pb.GetUserRequest{
		Id: int32(id64),
	})
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// UPDATE: PUT /users/:id
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var req pb.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	req.Id = int32(id64)

	res, err := h.UserClient.UpdateUser(c.Request.Context(), &req)
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// DELETE: DELETE /users/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	_, err = h.UserClient.DeleteUser(c.Request.Context(), &pb.DeleteUserRequest{
		Id: int32(id64),
	})
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

// LIST: GET /users?search=&page=&limit=
func (h *UserHandler) ListUsers(c *gin.Context) {
	search := c.Query("search")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, _ := strconv.Atoi(pageStr)
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(limitStr)
	if limit < 1 {
		limit = 10
	}

	res, err := h.UserClient.ListUsers(c.Request.Context(), &pb.ListUsersRequest{})
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	items := make([]*pb.UserResponse, 0)
	for _, u := range res.GetUsers() {
		if search != "" && !strings.Contains(strings.ToLower(u.GetUsername()), strings.ToLower(search)) {
			continue
		}
		items = append(items, u)
	}
	total := len(items)
	start := (page - 1) * limit
	if start > total {
		start = total
	}
	end := start + limit
	if end > total {
		end = total
	}
	paged := items[start:end]
	c.JSON(http.StatusOK, gin.H{"items": paged, "total": total, "page": page, "limit": limit})
}
