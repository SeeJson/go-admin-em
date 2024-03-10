package models

import (
    "time"

	"go-admin/common/models"

)

type EmTextbookIn struct {
    models.Model
    
    TextbookId string `json:"textbookId" gorm:"type:int(10);comment:教材_id"` 
    InTime time.Time `json:"inTime" gorm:"type:timestamp;comment:入库时间"` 
    Price string `json:"price" gorm:"type:int(10);comment:定价"` 
    Num string `json:"num" gorm:"type:int(10);comment:数量"` 
    TextbookPro string `json:"textbookPro" gorm:"type:int(10);comment:供应商"` 
    TextbookUnit string `json:"textbookUnit" gorm:"type:varchar(20);comment:经办单位"` 
    models.ModelTime
    models.ControlBy
}

func (EmTextbookIn) TableName() string {
    return "em_textbook_in"
}

func (e *EmTextbookIn) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EmTextbookIn) GetId() interface{} {
	return e.Id
}