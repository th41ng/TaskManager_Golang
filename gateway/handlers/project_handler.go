package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	pb "taskmanager/gateway/pb"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	ProjectClient pb.ProjectServiceClient
}

func NewProjectHandler(client pb.ProjectServiceClient) *ProjectHandler {
	return &ProjectHandler{ProjectClient: client}
}

// CREATE: POST /projects
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var req pb.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(fmt.Errorf("invalid request body: %v", err)) // thêm dòng này
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	res, err := h.ProjectClient.CreateProject(ctx, &req)
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// READ: GET /projects/:id
func (h *ProjectHandler) GetProject(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	res, err := h.ProjectClient.GetProject(ctx, &pb.GetProjectRequest{Id: int32(id64)})
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// UPDATE: PUT /projects/:id
func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	var req pb.UpdateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	req.Id = int32(id64)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	res, err := h.ProjectClient.UpdateProject(ctx, &req)
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// DELETE: DELETE /projects/:id
func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	_, err = h.ProjectClient.DeleteProject(ctx, &pb.DeleteProjectRequest{Id: int32(id64)})
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "project deleted successfully"})
}
