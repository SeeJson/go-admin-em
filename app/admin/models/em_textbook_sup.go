package models

import (
    "time"

	"go-admin/common/models"

)

type EmTextbookSup struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(50);comment:供应商名"` 
    Phone string `json:"phone" gorm:"type:int(12);comment:供应商电话"` 
    Addr string `json:"addr" gorm:"type:varchar(50);comment:供应商地址"` 
    Time time.Time `json:"time" gorm:"type:timestamp;comment:供应商开始供应时间"` 
    Status string `json:"status" gorm:"type:int(1);comment:状态"` 
    models.ModelTime
    models.ControlBy
}

func (EmTextbookSup) TableName() string {
    return "em_textbook_sup"
}

func (e *EmTextbookSup) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EmTextbookSup) GetId() interface{} {
	return e.Id
}