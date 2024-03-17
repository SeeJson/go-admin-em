package models

import (
	"go-admin/common/models"
)

type EmMajorInfo struct {
	models.Model

	MajorName   string `json:"majorName" gorm:"type:varchar(10);comment:所属专业"`
	MajorBk     string `json:"majorBk" gorm:"type:varchar(200);comment:简介"`
	CollegeName string `json:"collegeName" gorm:"type:varchar(12);comment:所属学院"`
	models.ModelTime
	models.ControlBy
}

func (EmMajorInfo) TableName() string {
	return "em_major_info"
}

func (e *EmMajorInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EmMajorInfo) GetId() interface{} {
	return e.Id
}
