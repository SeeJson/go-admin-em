package dto

import (
	"go-admin/app/film/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EvFilmOnlineGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:exact;column:name;table:ev_film_online" comment:"name"`
	EvFilmOnlineOrder
}

type EvFilmOnlineOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:ev_film_online"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:ev_film_online"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:ev_film_online"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:ev_film_online"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ev_film_online"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ev_film_online"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ev_film_online"`
}

func (m *EvFilmOnlineGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EvFilmOnlineInsertReq struct {
	Id   int    `json:"-" comment:""` //
	Name string `json:"name" comment:"name"`
	common.ControlBy
}

func (s *EvFilmOnlineInsertReq) Generate(model *models.EvFilmOnline) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EvFilmOnlineInsertReq) GetId() interface{} {
	return s.Id
}

type EvFilmOnlineUpdateReq struct {
	Id   int    `uri:"id" comment:""` //
	Name string `json:"name" comment:"name"`
	common.ControlBy
}

func (s *EvFilmOnlineUpdateReq) Generate(model *models.EvFilmOnline) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EvFilmOnlineUpdateReq) GetId() interface{} {
	return s.Id
}

// EvFilmOnlineGetReq 功能获取请求参数
type EvFilmOnlineGetReq struct {
	Id int `uri:"id"`
}

func (s *EvFilmOnlineGetReq) GetId() interface{} {
	return s.Id
}

// EvFilmOnlineDeleteReq 功能删除请求参数
type EvFilmOnlineDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EvFilmOnlineDeleteReq) GetId() interface{} {
	return s.Ids
}
