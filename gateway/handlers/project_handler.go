package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

// LIST: GET /projects?owner_id=&search=&page=&limit=
func (h *ProjectHandler) ListProjects(c *gin.Context) {
	// query params
	ownerIDStr := c.Query("owner_id")
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

	var ownerIDFilter *int
	if ownerIDStr != "" {
		if v, err := strconv.Atoi(ownerIDStr); err == nil {
			ownerIDFilter = &v
		}
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	res, err := h.ProjectClient.ListProjects(ctx, &pb.ListProjectsRequest{})
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	// filter in gateway
	items := make([]*pb.ProjectResponse, 0)
	for _, p := range res.GetProjects() {
		if ownerIDFilter != nil && int(p.GetOwnerId()) != *ownerIDFilter {
			continue
		}
		if search != "" && !containsFold(p.GetName(), search) {
			continue
		}
		items = append(items, p)
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

	c.JSON(http.StatusOK, gin.H{
		"projects": paged,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}

// LIST by user: GET /users/:id/projects?search=&page=&limit=
func (h *ProjectHandler) ListProjectsByUser(c *gin.Context) {
	userIDStr := c.Param("id")
	ownerID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	// call list and filter by owner
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	res, err := h.ProjectClient.ListProjects(ctx, &pb.ListProjectsRequest{})
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	items := make([]*pb.ProjectResponse, 0)
	for _, p := range res.GetProjects() {
		if int(p.GetOwnerId()) != ownerID {
			continue
		}
		if search != "" && !containsFold(p.GetName(), search) {
			continue
		}
		items = append(items, p)
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
	c.JSON(http.StatusOK, gin.H{"projects": paged, "total": total, "page": page, "limit": limit})
}

// helper: case-insensitive contains
func containsFold(s, sub string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(sub))
}

// helper: add or replace query values
func addOrReplaceQuery(raw string, kv map[string]string) string {
	q, _ := url.ParseQuery(raw)
	for k, v := range kv {
		q.Set(k, v)
	}
	return q.Encode()
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
