package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title     string // 标题
	Content   string // 内容
	Type      int    // 0: 无 1:工作 2:学习 3:生活
	StartTime string // 开始时间
	EndTime   string // 结束时间
	Finish    int    // 是否完成
}
