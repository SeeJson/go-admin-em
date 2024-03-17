package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EmTextbookSupGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:contains;column:name;table:em_textbook_sup" comment:"供应商名"`
	Status         string `form:"status"  search:"type:exact;column:status;table:em_textbook_sup" comment:"状态"`
	EmTextbookSupOrder
}

type EmTextbookSupOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:em_textbook_sup"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:em_textbook_sup"`
	Phone     string `form:"phoneOrder"  search:"type:order;column:phone;table:em_textbook_sup"`
	Addr      string `form:"addrOrder"  search:"type:order;column:addr;table:em_textbook_sup"`
	Time      string `form:"timeOrder"  search:"type:order;column:time;table:em_textbook_sup"`
	Status    string `form:"statusOrder"  search:"type:order;column:status;table:em_textbook_sup"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:em_textbook_sup"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:em_textbook_sup"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:em_textbook_sup"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:em_textbook_sup"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:em_textbook_sup"`
}

func (m *EmTextbookSupGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EmTextbookSupInsertReq struct {
	Id     int       `json:"-" comment:""` //
	Name   string    `json:"name" comment:"供应商名"`
	Phone  string    `json:"phone" comment:"供应商电话"`
	Addr   string    `json:"addr" comment:"供应商地址"`
	Time   time.Time `json:"time" comment:"供应商开始供应时间"`
	Status string    `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *EmTextbookSupInsertReq) Generate(model *models.EmTextbookSup) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Phone = s.Phone
	model.Addr = s.Addr
	model.Time = s.Time
	model.Status = s.Status
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EmTextbookSupInsertReq) GetId() interface{} {
	return s.Id
}

type EmTextbookSupUpdateReq struct {
	Id     int       `uri:"id" comment:""` //
	Name   string    `json:"name" comment:"供应商名"`
	Phone  string    `json:"phone" comment:"供应商电话"`
	Addr   string    `json:"addr" comment:"供应商地址"`
	Time   time.Time `json:"time" comment:"供应商开始供应时间"`
	Status string    `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *EmTextbookSupUpdateReq) Generate(model *models.EmTextbookSup) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Phone = s.Phone
	model.Addr = s.Addr
	model.Time = s.Time
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EmTextbookSupUpdateReq) GetId() interface{} {
	return s.Id
}

// EmTextbookSupGetReq 功能获取请求参数
type EmTextbookSupGetReq struct {
	Id int `uri:"id"`
}

func (s *EmTextbookSupGetReq) GetId() interface{} {
	return s.Id
}

// EmTextbookSupDeleteReq 功能删除请求参数
type EmTextbookSupDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EmTextbookSupDeleteReq) GetId() interface{} {
	return s.Ids
}
