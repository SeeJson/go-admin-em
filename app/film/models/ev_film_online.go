package models

import (
	"go-admin/common/models"
)

type EvFilmOnline struct {
	models.Model

	Name string `json:"name" gorm:"type:varchar(100);comment:name"`
	models.ModelTime
	models.ControlBy
}

func (EvFilmOnline) TableName() string {
	return "ev_film_online"
}

func (e *EvFilmOnline) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EvFilmOnline) GetId() interface{} {
	return e.Id
}
