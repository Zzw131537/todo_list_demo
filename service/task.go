package service

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"todo_list_demo/model"

	"github.com/gin-gonic/gin"
)

type TaskService struct {
	Title     string `json:"title" form:"title"`
	Content   string `json:"content" form:"content"`
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
	Type      int    `json:"type" form:"type"`
	Finish    int    `json:"finish" form:"finish"`
	Sort      int    `json:"sort" form:"sort"`
	Num       int    `json:"num" form:"num"`
}

func Create(ctx *gin.Context) {
	var service TaskService
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.Create(ctx.Request.Context())
		ctx.JSON(200, res)
	} else {
		ctx.JSON(500, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
	}
}

// 创建任务
func (service *TaskService) Create(ctx context.Context) gin.H {
	if service.Title == "" {
		return gin.H{
			"code": 400,
			"msg":  "标题不能为空",
		}
	}

	model.DB.Model(&model.Task{}).Create(&model.Task{
		Type:      service.Type,
		Title:     service.Title,
		Content:   service.Content,
		StartTime: service.StartTime,
		EndTime:   service.EndTime,
		Finish:    service.Finish,
	})
	return gin.H{
		"code": 200,
		"msg":  "success",
	}
}

func (service *TaskService) Delete(ctx context.Context, id string) interface{} {
	newId, err := strconv.Atoi(id)
	if err != nil {
		return gin.H{
			"code": 400,
			"msg":  err.Error(),
		}
	}
	fmt.Println("newId is ", newId)
	model.DB.Where("id = ?", newId).Delete(&model.Task{})
	return gin.H{
		"code": 200,
		"msg":  "success",
	}

}

func (service *TaskService) UpdateFindish(ctx context.Context) gin.H {
	err := model.DB.Model(&model.Task{}).Update("finish", service.Finish).Error
	if err != nil {
		return gin.H{
			"code": 500,
			"msg":  err.Error(),
		}
	} else {
		return gin.H{
			"code": 200,
			"msg":  "success",
		}
	}
}

func (service *TaskService) ListTasks(ctx context.Context) gin.H {
	tasks := make([]model.Task, 0)
	if service.Sort == 0 { // 默认顺序
		model.DB.Model(&model.Task{}).Find(&tasks).Where("type = ?", service.Type)
	} else if service.Sort == 1 { // 按截止时间进行排序
		model.DB.Model(&model.Task{}).Find(&tasks).Where("type = ?", service.Type).Order("end_time asc")
	} else if service.Sort == 2 { // 按优先级进行排序
		model.DB.Model(&model.Task{}).Find(&tasks).Where("type = ?", service.Type).Order("num desc ")
	}

	return gin.H{
		"code": 200,
		"msg":  "success",
		"data": tasks,
	}

}

func DeleteById(c *gin.Context) {
	service := new(TaskService)
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
	}
}

func BookTask(c *gin.Context) {
	service := new(TaskService)
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateFindish(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
	}
}

func ListTasks(c *gin.Context) {
	service := new(TaskService)
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListTasks(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
	}
}
