package routers

import (
	"DailyChecklist/controller"
	"DailyChecklist/setting"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	r.GET("/archive", controller.ArchivePageHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
		// 存档某一个待办事项
		v1Group.POST("/todo/:id/archive", controller.ArchiveATodo)
		// 获取存档的待办事项
		v1Group.GET("/todo/archived", controller.GetArchivedTodos)
		// 获取统计数据
		v1Group.GET("/todo/stats", controller.GetTodoStats)
		// 获取历史统计数据
		v1Group.GET("/todo/history", controller.GetTodoHistoryStats)
	}
	return r
}
