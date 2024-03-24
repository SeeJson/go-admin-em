package models

import (
	"time"

	"go-admin/common/models"
)

type EvAreaFilmInfo struct {
	models.Model

	FilmId             string    `json:"filmId" gorm:"type:bigint;comment:电影id"`
	Level              string    `json:"level" gorm:"type:varchar(50);comment:城市级别"`
	NumPeople          string    `json:"numPeople" gorm:"type:bigint;comment:观看人次"`
	ReleaseTime        time.Time `json:"releaseTime" gorm:"type:datetime(3);comment:上映时间"`
	AccReleaseDays     string    `json:"accReleaseDays" gorm:"type:int;comment:累计上映天数"`
	ReleaseVenne       string    `json:"releaseVenne" gorm:"type:bigint;comment:上映场次"`
	ReleaseNumPeople   string    `json:"releaseNumPeople" gorm:"type:bigint;comment:上映人次"`
	OccupancyRate      string    `json:"occupancyRate" gorm:"type:bigint;comment:上座率"`
	BoxOffice          string    `json:"boxOffice" gorm:"type:bigint;comment:票房"`
	DianyBoxOffice     string    `json:"dianyBoxOffice" gorm:"type:bigint;comment:点映票房"`
	FirstDayeBoxOffice string    `json:"firstDayeBoxOffice" gorm:"type:bigint;comment:首日票房"`
	FirstWeekBoxOffice string    `json:"firstWeekBoxOffice" gorm:"type:bigint;comment:首周票房"`
	TwoWeekBoxOffice   string    `json:"twoWeekBoxOffice" gorm:"type:bigint;comment:第二周票房"`
	WordOfMouth        string    `json:"wordOfMouth" gorm:"type:bigint;comment:口碑"`
	Score              string    `json:"score" gorm:"type:float;comment:评分"`
	models.ModelTime
	models.ControlBy
}

func (EvAreaFilmInfo) TableName() string {
	return "ev_area_film_info"
}

func (e *EvAreaFilmInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EvAreaFilmInfo) GetId() interface{} {
	return e.Id
}
