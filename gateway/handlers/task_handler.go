package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	pb "taskmanager/gateway/pb"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	TaskClient pb.TaskServiceClient
}

func NewTaskHandler(client pb.TaskServiceClient) *TaskHandler {
	return &TaskHandler{TaskClient: client}
}

// CREATE: POST /tasks
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req pb.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	res, err := h.TaskClient.CreateTask(ctx, &req)
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// READ: GET /tasks/:id
func (h *TaskHandler) GetTask(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	res, err := h.TaskClient.GetTask(ctx, &pb.GetTaskRequest{Id: int32(id64)})
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// UPDATE: PUT /tasks/:id
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	var req pb.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	req.Id = int32(id64)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	res, err := h.TaskClient.UpdateTask(ctx, &req)
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// DELETE: DELETE /tasks/:id
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	_, err = h.TaskClient.DeleteTask(ctx, &pb.DeleteTaskRequest{Id: int32(id64)})
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}
