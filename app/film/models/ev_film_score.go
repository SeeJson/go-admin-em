package models

import (
	"go-admin/common/models"
)

type EvFilmScore struct {
	models.Model

	FilmId  string `json:"filmId" gorm:"type:bigint;comment:电影id"`
	UserId  string `json:"userId" gorm:"type:bigint;comment:用户id"`
	Score   string `json:"score" gorm:"type:int;comment:评分"`
	ReValue string `json:"reValue" gorm:"type:float;comment:推荐值"`
	models.ModelTime
	models.ControlBy
}

func (EvFilmScore) TableName() string {
	return "ev_film_score"
}

func (e *EvFilmScore) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EvFilmScore) GetId() interface{} {
	return e.Id
}
