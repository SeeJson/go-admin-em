package models

import (
	"go-admin/common/models"
)

type EmClassInfo struct {
	models.Model

	Name       string `json:"name" gorm:"type:varchar(50);comment:班级名"`
	Major      string `json:"major" gorm:"type:varchar(30);comment:专业"`
	College    string `json:"college" gorm:"type:varchar(30);comment:学院"`
	RePerson   string `json:"rePerson" gorm:"type:varchar(30);comment:负责人"`
	RePersonNo string `json:"rePersonNo" gorm:"type:int(30);comment:负责人学号"`
	models.ModelTime
	models.ControlBy
}

func (EmClassInfo) TableName() string {
	return "em_class_info"
}

func (e *EmClassInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EmClassInfo) GetId() interface{} {
	return e.Id
}
