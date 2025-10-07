package controller

import (
	"DailyChecklist/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 url     --> controller  --> logic   -->    model
请求来了  -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	// 2. 存入数据库
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg": "success",
		//	"data": todo,
		//})
	}
}

func GetTodoList(c *gin.Context) {
	// 查询todo这个表里的所有未存档数据
	todoList, err := models.GetActiveTodos()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}

// ArchiveATodo 存档待办事项
func ArchiveATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.ArchiveATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "存档成功"})
	}
}

// GetArchivedTodos 获取存档的待办事项
func GetArchivedTodos(c *gin.Context) {
	todoList, err := models.GetArchivedTodos()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

// GetTodoStats 获取统计数据
func GetTodoStats(c *gin.Context) {
	stats, err := models.GetTodoStats()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, stats)
	}
}

// GetArchivedTodoStats 获取存档统计数据
func GetArchivedTodoStats(c *gin.Context) {
	stats, err := models.GetArchivedTodoStats()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, stats)
	}
}

// ArchivePageHandler 存档页面处理器
func ArchivePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "archive.html", nil)
}

// GetTodoHistoryStats 获取历史统计数据
func GetTodoHistoryStats(c *gin.Context) {
	period := c.DefaultQuery("period", "week")
	stats, err := models.GetTodoHistoryStats(period)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, stats)
	}
}
