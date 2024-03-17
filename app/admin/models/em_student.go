package models

import (
	"go-admin/common/models"
)

type EmStudent struct {
	models.Model

	Sex         string `json:"sex" gorm:"type:int(1);comment:性别"`
	Name        string `json:"name" gorm:"type:varchar(20);comment:姓名"`
	Number      string `json:"number" gorm:"type:int(20);comment:学号"`
	ClassName   string `json:"className" gorm:"type:varchar(50);comment:班级名"`
	Major       string `json:"major" gorm:"type:varchar(50);comment:专业名"`
	CollegeName string `json:"collegeName" gorm:"type:varchar(50);comment:学院名"`
	Phone       string `json:"phone" gorm:"type:varchar(12);comment:手机号"`
	Email       string `json:"email" gorm:"type:varchar(50);comment:邮箱"`
	models.ModelTime
	models.ControlBy
}

func (EmStudent) TableName() string {
	return "em_student"
}

func (e *EmStudent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EmStudent) GetId() interface{} {
	return e.Id
}
