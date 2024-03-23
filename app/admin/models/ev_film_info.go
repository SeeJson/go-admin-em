package models

import (
	"time"

	"go-admin/common/models"
)

type EvFilmInfo struct {
	models.Model

	Name               string    `json:"name" gorm:"type:varchar(50);comment:电影名"`
	Alias              string    `json:"alias" gorm:"type:varchar(50);comment:别名"`
	LeadActor          string    `json:"leadActor" gorm:"type:varchar(100);comment:主演"`
	Director           string    `json:"director" gorm:"type:varchar(30);comment:导演"`
	BriefIntroduction  string    `json:"briefIntroduction" gorm:"type:text;comment:简介"`
	Cover              string    `json:"cover" gorm:"type:varchar(100);comment:封面"`
	ReleaseTime        time.Time `json:"releaseTime" gorm:"type:datetime(3);comment:上映时间"`
	ReleaseRelease     string    `json:"releaseRelease" gorm:"type:varchar(50);comment:上映地区"`
	Country            string    `json:"country" gorm:"type:int;comment:国家"`
	Type               string    `json:"type" gorm:"type:int;comment:电影类型"`
	Status             string    `json:"status" gorm:"type:int;comment:状态"`
	Language           string    `json:"language" gorm:"type:int;comment:语言"`
	FilmLength         string    `json:"filmLength" gorm:"type:int;comment:片长"`
	TppNum             string    `json:"tppNum" gorm:"type:bigint;comment:淘票票想看数"`
	MyNum              string    `json:"myNum" gorm:"type:bigint;comment:猫眼想看数"`
	DbNum              string    `json:"dbNum" gorm:"type:bigint;comment:豆瓣想看数"`
	MtNum              string    `json:"mtNum" gorm:"type:bigint;comment:美团想看数"`
	AccReleaseDays     string    `json:"accReleaseDays" gorm:"type:int;comment:累计上映天数"`
	ReleaseVenne       string    `json:"releaseVenne" gorm:"type:bigint;comment:上映场次"`
	ReleaseNumPeople   string    `json:"releaseNumPeople" gorm:"type:bigint;comment:上映人次"`
	OccupancyRate      string    `json:"occupancyRate" gorm:"type:bigint;comment:上座率"`
	TotalBoxOffice     string    `json:"totalBoxOffice" gorm:"type:bigint;comment:总票房"`
	BoxOffice          string    `json:"boxOffice" gorm:"type:bigint;comment:当前票房"`
	DianyBoxOffice     string    `json:"dianyBoxOffice" gorm:"type:bigint;comment:点映票房"`
	FirstDayeBoxOffice string    `json:"firstDayeBoxOffice" gorm:"type:bigint;comment:首日票房"`
	FirstWeekBoxOffice string    `json:"firstWeekBoxOffice" gorm:"type:bigint;comment:首周票房"`
	TwoWeekBoxOffice   string    `json:"twoWeekBoxOffice" gorm:"type:bigint;comment:第二周票房"`
	WordOfMouth        string    `json:"wordOfMouth" gorm:"type:bigint;comment:口碑"`
	Score              string    `json:"score" gorm:"type:float;comment:评分"`
	models.ModelTime
	models.ControlBy
}

func (EvFilmInfo) TableName() string {
	return "ev_film_info"
}

func (e *EvFilmInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EvFilmInfo) GetId() interface{} {
	return e.Id
}
