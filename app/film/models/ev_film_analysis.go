package models

import (
	"go-admin/common/models"
)

type EvFilmAnalysis struct {
	models.Model

	Name string `json:"name" gorm:"type:varchar(100);comment:name"`
	models.ModelTime
	models.ControlBy
}

func (EvFilmAnalysis) TableName() string {
	return "ev_film_analysis"
}

func (e *EvFilmAnalysis) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EvFilmAnalysis) GetId() interface{} {
	return e.Id
}
