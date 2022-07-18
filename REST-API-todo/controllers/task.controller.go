package controller

import (
	"net/http"
	"strconv"

	"github.com/alexparco/rest-api-todo/model"
	"github.com/gin-gonic/gin"
)

type TaskController interface {
	GetTasks(c *gin.Context)
	GetTask(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type taskController struct {
	repo model.TaskRepo
}

func NewTaskController(taskRepo model.TaskRepo) TaskController {
	return &taskController{repo: taskRepo}
}

func (t *taskController) GetTasks(c *gin.Context) {
	tasks, err := t.repo.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (t *taskController) GetTask(c *gin.Context) {
	id := c.Param("id")

	task, err := t.repo.GetOne(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskController) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Status = false
	if err := t.repo.Create(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskController) UpdateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	task.Id, _ = strconv.Atoi(id)
	if err := t.repo.Update(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	task, err := t.repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
