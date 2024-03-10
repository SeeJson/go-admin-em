package models

import (
    "time"

	"go-admin/common/models"

)

type EmTextbook struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(50);comment:教材名"` 
    Price string `json:"price" gorm:"type:int(10);comment:价格"` 
    Sup string `json:"sup" gorm:"type:int(10);comment:教材供应商"` 
    ProData time.Time `json:"proData" gorm:"type:timestamp;comment:出版日期"` 
    Author string `json:"author" gorm:"type:varchar(50);comment:出版作者"` 
    Press string `json:"press" gorm:"type:varchar(50);comment:出版社"` 
    models.ModelTime
    models.ControlBy
}

func (EmTextbook) TableName() string {
    return "em_textbook"
}

func (e *EmTextbook) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EmTextbook) GetId() interface{} {
	return e.Id
}