package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EvFilmScoreGetPageReq struct {
	dto.Pagination `search:"-"`
	FilmId         string `form:"filmId"  search:"type:exact;column:film_id;table:ev_film_score" comment:"电影id"`
	UserId         string `form:"userId"  search:"type:exact;column:user_id;table:ev_film_score" comment:"用户id"`
	EvFilmScoreOrder
}

type EvFilmScoreOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:ev_film_score"`
	FilmId    string `form:"filmIdOrder"  search:"type:order;column:film_id;table:ev_film_score"`
	UserId    string `form:"userIdOrder"  search:"type:order;column:user_id;table:ev_film_score"`
	Score     string `form:"scoreOrder"  search:"type:order;column:score;table:ev_film_score"`
	ReValue   string `form:"reValueOrder"  search:"type:order;column:re_value;table:ev_film_score"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:ev_film_score"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:ev_film_score"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ev_film_score"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ev_film_score"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ev_film_score"`
}

func (m *EvFilmScoreGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EvFilmScoreInsertReq struct {
	Id      int    `json:"-" comment:""` //
	FilmId  string `json:"filmId" comment:"电影id"`
	UserId  string `json:"userId" comment:"用户id"`
	Score   string `json:"score" comment:"评分"`
	ReValue string `json:"reValue" comment:"推荐值"`
	common.ControlBy
}

func (s *EvFilmScoreInsertReq) Generate(model *models.EvFilmScore) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.FilmId = s.FilmId
	model.UserId = s.UserId
	model.Score = s.Score
	model.ReValue = s.ReValue
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EvFilmScoreInsertReq) GetId() interface{} {
	return s.Id
}

type EvFilmScoreUpdateReq struct {
	Id      int    `uri:"id" comment:""` //
	FilmId  string `json:"filmId" comment:"电影id"`
	UserId  string `json:"userId" comment:"用户id"`
	Score   string `json:"score" comment:"评分"`
	ReValue string `json:"reValue" comment:"推荐值"`
	common.ControlBy
}

func (s *EvFilmScoreUpdateReq) Generate(model *models.EvFilmScore) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.FilmId = s.FilmId
	model.UserId = s.UserId
	model.Score = s.Score
	model.ReValue = s.ReValue
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EvFilmScoreUpdateReq) GetId() interface{} {
	return s.Id
}

// EvFilmScoreGetReq 功能获取请求参数
type EvFilmScoreGetReq struct {
	Id int `uri:"id"`
}

func (s *EvFilmScoreGetReq) GetId() interface{} {
	return s.Id
}

// EvFilmScoreDeleteReq 功能删除请求参数
type EvFilmScoreDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EvFilmScoreDeleteReq) GetId() interface{} {
	return s.Ids
}
