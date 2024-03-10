package models

import (

	"go-admin/common/models"

)

type EmCollegeInfo struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(50);comment:学院名"` 
    Bk string `json:"bk" gorm:"type:varchar(200);comment:学院简介"` 
    RePerson string `json:"rePerson" gorm:"type:varchar(50);comment:负责人姓名"` 
    RePhone string `json:"rePhone" gorm:"type:varchar(50);comment:负责人手机号"` 
    ReNo string `json:"reNo" gorm:"type:int(20);comment:负责人工号"` 
    models.ModelTime
    models.ControlBy
}

func (EmCollegeInfo) TableName() string {
    return "em_college_info"
}

func (e *EmCollegeInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EmCollegeInfo) GetId() interface{} {
	return e.Id
}