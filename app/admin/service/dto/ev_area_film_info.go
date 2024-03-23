package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EvAreaFilmInfoGetPageReq struct {
	dto.Pagination `search:"-"`
	FilmId         string `form:"filmId"  search:"type:exact;column:film_id;table:ev_area_film_info" comment:"电影id"`
	Level          string `form:"level"  search:"type:exact;column:level;table:ev_area_film_info" comment:"城市级别"`
	EvAreaFilmInfoOrder
}

type EvAreaFilmInfoOrder struct {
	Id                 string `form:"idOrder"  search:"type:order;column:id;table:ev_area_film_info"`
	FilmId             string `form:"filmIdOrder"  search:"type:order;column:film_id;table:ev_area_film_info"`
	Level              string `form:"levelOrder"  search:"type:order;column:level;table:ev_area_film_info"`
	NumPeople          string `form:"numPeopleOrder"  search:"type:order;column:num_people;table:ev_area_film_info"`
	ReleaseTime        string `form:"releaseTimeOrder"  search:"type:order;column:release_time;table:ev_area_film_info"`
	AccReleaseDays     string `form:"accReleaseDaysOrder"  search:"type:order;column:acc_release_days;table:ev_area_film_info"`
	ReleaseVenne       string `form:"releaseVenneOrder"  search:"type:order;column:release_venne;table:ev_area_film_info"`
	ReleaseNumPeople   string `form:"releaseNumPeopleOrder"  search:"type:order;column:release_num_people;table:ev_area_film_info"`
	OccupancyRate      string `form:"occupancyRateOrder"  search:"type:order;column:occupancy_rate;table:ev_area_film_info"`
	BoxOffice          string `form:"boxOfficeOrder"  search:"type:order;column:box_office;table:ev_area_film_info"`
	DianyBoxOffice     string `form:"dianyBoxOfficeOrder"  search:"type:order;column:diany_box_office;table:ev_area_film_info"`
	FirstDayeBoxOffice string `form:"firstDayeBoxOfficeOrder"  search:"type:order;column:first_daye_box_office;table:ev_area_film_info"`
	FirstWeekBoxOffice string `form:"firstWeekBoxOfficeOrder"  search:"type:order;column:first_week_box_office;table:ev_area_film_info"`
	TwoWeekBoxOffice   string `form:"twoWeekBoxOfficeOrder"  search:"type:order;column:two_week_box_office;table:ev_area_film_info"`
	WordOfMouth        string `form:"wordOfMouthOrder"  search:"type:order;column:word_of_mouth;table:ev_area_film_info"`
	Score              string `form:"scoreOrder"  search:"type:order;column:score;table:ev_area_film_info"`
	CreateBy           string `form:"createByOrder"  search:"type:order;column:create_by;table:ev_area_film_info"`
	UpdateBy           string `form:"updateByOrder"  search:"type:order;column:update_by;table:ev_area_film_info"`
	CreatedAt          string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ev_area_film_info"`
	UpdatedAt          string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ev_area_film_info"`
	DeletedAt          string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ev_area_film_info"`
}

func (m *EvAreaFilmInfoGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EvAreaFilmInfoInsertReq struct {
	Id                 int       `json:"-" comment:""` //
	FilmId             string    `json:"filmId" comment:"电影id"`
	Level              string    `json:"level" comment:"城市级别"`
	NumPeople          string    `json:"numPeople" comment:"观看人次"`
	ReleaseTime        time.Time `json:"releaseTime" comment:"上映时间"`
	AccReleaseDays     string    `json:"accReleaseDays" comment:"累计上映天数"`
	ReleaseVenne       string    `json:"releaseVenne" comment:"上映场次"`
	ReleaseNumPeople   string    `json:"releaseNumPeople" comment:"上映人次"`
	OccupancyRate      string    `json:"occupancyRate" comment:"上座率"`
	BoxOffice          string    `json:"boxOffice" comment:"票房"`
	DianyBoxOffice     string    `json:"dianyBoxOffice" comment:"点映票房"`
	FirstDayeBoxOffice string    `json:"firstDayeBoxOffice" comment:"首日票房"`
	FirstWeekBoxOffice string    `json:"firstWeekBoxOffice" comment:"首周票房"`
	TwoWeekBoxOffice   string    `json:"twoWeekBoxOffice" comment:"第二周票房"`
	WordOfMouth        string    `json:"wordOfMouth" comment:"口碑"`
	Score              string    `json:"score" comment:"评分"`
	common.ControlBy
}

func (s *EvAreaFilmInfoInsertReq) Generate(model *models.EvAreaFilmInfo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.FilmId = s.FilmId
	model.Level = s.Level
	model.NumPeople = s.NumPeople
	model.ReleaseTime = s.ReleaseTime
	model.AccReleaseDays = s.AccReleaseDays
	model.ReleaseVenne = s.ReleaseVenne
	model.ReleaseNumPeople = s.ReleaseNumPeople
	model.OccupancyRate = s.OccupancyRate
	model.BoxOffice = s.BoxOffice
	model.DianyBoxOffice = s.DianyBoxOffice
	model.FirstDayeBoxOffice = s.FirstDayeBoxOffice
	model.FirstWeekBoxOffice = s.FirstWeekBoxOffice
	model.TwoWeekBoxOffice = s.TwoWeekBoxOffice
	model.WordOfMouth = s.WordOfMouth
	model.Score = s.Score
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EvAreaFilmInfoInsertReq) GetId() interface{} {
	return s.Id
}

type EvAreaFilmInfoUpdateReq struct {
	Id                 int       `uri:"id" comment:""` //
	FilmId             string    `json:"filmId" comment:"电影id"`
	Level              string    `json:"level" comment:"城市级别"`
	NumPeople          string    `json:"numPeople" comment:"观看人次"`
	ReleaseTime        time.Time `json:"releaseTime" comment:"上映时间"`
	AccReleaseDays     string    `json:"accReleaseDays" comment:"累计上映天数"`
	ReleaseVenne       string    `json:"releaseVenne" comment:"上映场次"`
	ReleaseNumPeople   string    `json:"releaseNumPeople" comment:"上映人次"`
	OccupancyRate      string    `json:"occupancyRate" comment:"上座率"`
	BoxOffice          string    `json:"boxOffice" comment:"票房"`
	DianyBoxOffice     string    `json:"dianyBoxOffice" comment:"点映票房"`
	FirstDayeBoxOffice string    `json:"firstDayeBoxOffice" comment:"首日票房"`
	FirstWeekBoxOffice string    `json:"firstWeekBoxOffice" comment:"首周票房"`
	TwoWeekBoxOffice   string    `json:"twoWeekBoxOffice" comment:"第二周票房"`
	WordOfMouth        string    `json:"wordOfMouth" comment:"口碑"`
	Score              string    `json:"score" comment:"评分"`
	common.ControlBy
}

func (s *EvAreaFilmInfoUpdateReq) Generate(model *models.EvAreaFilmInfo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.FilmId = s.FilmId
	model.Level = s.Level
	model.NumPeople = s.NumPeople
	model.ReleaseTime = s.ReleaseTime
	model.AccReleaseDays = s.AccReleaseDays
	model.ReleaseVenne = s.ReleaseVenne
	model.ReleaseNumPeople = s.ReleaseNumPeople
	model.OccupancyRate = s.OccupancyRate
	model.BoxOffice = s.BoxOffice
	model.DianyBoxOffice = s.DianyBoxOffice
	model.FirstDayeBoxOffice = s.FirstDayeBoxOffice
	model.FirstWeekBoxOffice = s.FirstWeekBoxOffice
	model.TwoWeekBoxOffice = s.TwoWeekBoxOffice
	model.WordOfMouth = s.WordOfMouth
	model.Score = s.Score
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EvAreaFilmInfoUpdateReq) GetId() interface{} {
	return s.Id
}

// EvAreaFilmInfoGetReq 功能获取请求参数
type EvAreaFilmInfoGetReq struct {
	Id int `uri:"id"`
}

func (s *EvAreaFilmInfoGetReq) GetId() interface{} {
	return s.Id
}

// EvAreaFilmInfoDeleteReq 功能删除请求参数
type EvAreaFilmInfoDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EvAreaFilmInfoDeleteReq) GetId() interface{} {
	return s.Ids
}
