package models

import (
	"time"

	"go-admin/common/models"
)

type EmTextbookOn struct {
	models.Model

	TextbookId   string    `json:"textbookId" gorm:"type:int(10);comment:教材_id"`
	Num          string    `json:"num" gorm:"type:int(10);comment:数量"`
	TextbookUnit string    `json:"textbookUnit" gorm:"type:varchar(50);comment:经办单位"`
	OutTime      time.Time `json:"outTime" gorm:"type:timestamp;comment:出库时间"`
	models.ModelTime
	models.ControlBy
}

func (EmTextbookOn) TableName() string {
	return "em_textbook_on"
}

func (e *EmTextbookOn) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EmTextbookOn) GetId() interface{} {
	return e.Id
}
