package handlers

import (
	"context"
	"net/http"
	"strconv"
	"strings"
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

// LIST: GET /tasks?project_id=&done=&priority=&search=&page=&limit=
func (h *TaskHandler) ListTasks(c *gin.Context) {
	projStr := c.Query("project_id")
	doneStr := c.Query("done")
	priStr := c.Query("priority")
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

	var projID *int
	if projStr != "" {
		if v, err := strconv.Atoi(projStr); err == nil {
			projID = &v
		}
	}
	var done *bool
	if doneStr != "" {
		if v, err := strconv.ParseBool(doneStr); err == nil {
			done = &v
		}
	}
	var pri *int
	if priStr != "" {
		if v, err := strconv.Atoi(priStr); err == nil {
			pri = &v
		}
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	res, err := h.TaskClient.ListTasks(ctx, &pb.ListTasksRequest{})
	if err != nil {
		handleGrpcError(c, err)
		return
	}

	items := make([]*pb.TaskResponse, 0)
	for _, t := range res.GetTasks() {
		if projID != nil && int(t.GetProjectId()) != *projID {
			continue
		}
		if done != nil && t.GetDone() != *done {
			continue
		}
		if pri != nil && int(t.GetPriority()) != *pri {
			continue
		}
		if search != "" && !strings.Contains(strings.ToLower(t.GetTitle()), strings.ToLower(search)) {
			continue
		}
		items = append(items, t)
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
	c.JSON(http.StatusOK, gin.H{"tasks": paged, "total": total, "page": page, "limit": limit})
}

// LIST by user: GET /users/:id/tasks?search=&done=&priority=&page=&limit=
// Note: We infer tasks of a user by finding their projects (owner_id = user) and selecting tasks with project_id in that set.
func (h *TaskHandler) ListTasksByUser(c *gin.Context, projectClient pb.ProjectServiceClient) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	search := c.Query("search")
	doneStr := c.Query("done")
	priStr := c.Query("priority")
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
	var done *bool
	if doneStr != "" {
		if v, err := strconv.ParseBool(doneStr); err == nil {
			done = &v
		}
	}
	var pri *int
	if priStr != "" {
		if v, err := strconv.Atoi(priStr); err == nil {
			pri = &v
		}
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 6*time.Second)
	defer cancel()
	// 1) list projects and collect IDs owned by user
	pres, err := projectClient.ListProjects(ctx, &pb.ListProjectsRequest{})
	if err != nil {
		handleGrpcError(c, err)
		return
	}
	projSet := map[int]struct{}{}
	for _, p := range pres.GetProjects() {
		if int(p.GetOwnerId()) == userID {
			projSet[int(p.GetId())] = struct{}{}
		}
	}
	// 2) list tasks and filter by project_id in user's projects
	tres, err := h.TaskClient.ListTasks(ctx, &pb.ListTasksRequest{})
	if err != nil {
		handleGrpcError(c, err)
		return
	}
	items := make([]*pb.TaskResponse, 0)
	for _, t := range tres.GetTasks() {
		if _, ok := projSet[int(t.GetProjectId())]; !ok {
			continue
		}
		if done != nil && t.GetDone() != *done {
			continue
		}
		if pri != nil && int(t.GetPriority()) != *pri {
			continue
		}
		if search != "" && !strings.Contains(strings.ToLower(t.GetTitle()), strings.ToLower(search)) {
			continue
		}
		items = append(items, t)
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
	c.JSON(http.StatusOK, gin.H{"tasks": paged, "total": total, "page": page, "limit": limit})
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
