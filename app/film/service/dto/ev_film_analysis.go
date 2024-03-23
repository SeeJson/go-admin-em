package dto

import (
	"go-admin/app/film/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EvFilmAnalysisGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:exact;column:name;table:ev_film_analysis" comment:"name"`
	EvFilmAnalysisOrder
}

type EvFilmAnalysisOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:ev_film_analysis"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:ev_film_analysis"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:ev_film_analysis"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:ev_film_analysis"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ev_film_analysis"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ev_film_analysis"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ev_film_analysis"`
}

func (m *EvFilmAnalysisGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EvFilmAnalysisInsertReq struct {
	Id   int    `json:"-" comment:""` //
	Name string `json:"name" comment:"name"`
	common.ControlBy
}

func (s *EvFilmAnalysisInsertReq) Generate(model *models.EvFilmAnalysis) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EvFilmAnalysisInsertReq) GetId() interface{} {
	return s.Id
}

type EvFilmAnalysisUpdateReq struct {
	Id   int    `uri:"id" comment:""` //
	Name string `json:"name" comment:"name"`
	common.ControlBy
}

func (s *EvFilmAnalysisUpdateReq) Generate(model *models.EvFilmAnalysis) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EvFilmAnalysisUpdateReq) GetId() interface{} {
	return s.Id
}

// EvFilmAnalysisGetReq 功能获取请求参数
type EvFilmAnalysisGetReq struct {
	Id int `uri:"id"`
}

func (s *EvFilmAnalysisGetReq) GetId() interface{} {
	return s.Id
}

// EvFilmAnalysisDeleteReq 功能删除请求参数
type EvFilmAnalysisDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EvFilmAnalysisDeleteReq) GetId() interface{} {
	return s.Ids
}
