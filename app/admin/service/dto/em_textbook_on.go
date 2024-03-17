package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EmTextbookOnGetPageReq struct {
	dto.Pagination `search:"-"`
	TextbookId     string `form:"textbookId"  search:"type:exact;column:textbook_id;table:em_textbook_on" comment:"教材_id"`
	TextbookUnit   string `form:"textbookUnit"  search:"type:contains;column:textbook_unit;table:em_textbook_on" comment:"经办单位"`
	EmTextbookOnOrder
}

type EmTextbookOnOrder struct {
	Id           string `form:"idOrder"  search:"type:order;column:id;table:em_textbook_on"`
	TextbookId   string `form:"textbookIdOrder"  search:"type:order;column:textbook_id;table:em_textbook_on"`
	Num          string `form:"numOrder"  search:"type:order;column:num;table:em_textbook_on"`
	TextbookUnit string `form:"textbookUnitOrder"  search:"type:order;column:textbook_unit;table:em_textbook_on"`
	OutTime      string `form:"outTimeOrder"  search:"type:order;column:out_time;table:em_textbook_on"`
	CreateBy     string `form:"createByOrder"  search:"type:order;column:create_by;table:em_textbook_on"`
	UpdateBy     string `form:"updateByOrder"  search:"type:order;column:update_by;table:em_textbook_on"`
	CreatedAt    string `form:"createdAtOrder"  search:"type:order;column:created_at;table:em_textbook_on"`
	UpdatedAt    string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:em_textbook_on"`
	DeletedAt    string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:em_textbook_on"`
}

func (m *EmTextbookOnGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EmTextbookOnInsertReq struct {
	Id           int       `json:"-" comment:""` //
	TextbookId   string    `json:"textbookId" comment:"教材_id"`
	Num          string    `json:"num" comment:"数量"`
	TextbookUnit string    `json:"textbookUnit" comment:"经办单位"`
	OutTime      time.Time `json:"outTime" comment:"出库时间"`
	common.ControlBy
}

func (s *EmTextbookOnInsertReq) Generate(model *models.EmTextbookOn) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.TextbookId = s.TextbookId
	model.Num = s.Num
	model.TextbookUnit = s.TextbookUnit
	model.OutTime = s.OutTime
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EmTextbookOnInsertReq) GetId() interface{} {
	return s.Id
}

type EmTextbookOnUpdateReq struct {
	Id           int       `uri:"id" comment:""` //
	TextbookId   string    `json:"textbookId" comment:"教材_id"`
	Num          string    `json:"num" comment:"数量"`
	TextbookUnit string    `json:"textbookUnit" comment:"经办单位"`
	OutTime      time.Time `json:"outTime" comment:"出库时间"`
	common.ControlBy
}

func (s *EmTextbookOnUpdateReq) Generate(model *models.EmTextbookOn) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.TextbookId = s.TextbookId
	model.Num = s.Num
	model.TextbookUnit = s.TextbookUnit
	model.OutTime = s.OutTime
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EmTextbookOnUpdateReq) GetId() interface{} {
	return s.Id
}

// EmTextbookOnGetReq 功能获取请求参数
type EmTextbookOnGetReq struct {
	Id int `uri:"id"`
}

func (s *EmTextbookOnGetReq) GetId() interface{} {
	return s.Id
}

// EmTextbookOnDeleteReq 功能删除请求参数
type EmTextbookOnDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EmTextbookOnDeleteReq) GetId() interface{} {
	return s.Ids
}
