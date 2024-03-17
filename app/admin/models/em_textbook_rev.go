package models

import (
	"time"

	"go-admin/common/models"
)

type EmTextbookRev struct {
	models.Model

	RevTime          time.Time `json:"revTime" gorm:"type:timestamp;comment:领取日期"`
	ClassId          string    `json:"classId" gorm:"type:int(10);comment:班级名"`
	TextbookId       string    `json:"textbookId" gorm:"type:int(10);comment:教材名"`
	Num              string    `json:"num" gorm:"type:int(11);comment:数量"`
	ClassLeader      string    `json:"classLeader" gorm:"type:varchar(50);comment:班级负责人"`
	ClassLeaderPhone string    `json:"classLeaderPhone" gorm:"type:int(12);comment:负责人电话"`
	models.ModelTime
	models.ControlBy
}

func (EmTextbookRev) TableName() string {
	return "em_textbook_rev"
}

func (e *EmTextbookRev) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EmTextbookRev) GetId() interface{} {
	return e.Id
}
