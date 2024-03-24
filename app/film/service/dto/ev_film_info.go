package dto

import (
	"go-admin/app/film/models"
	"time"

	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EvFilmInfoGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string    `form:"name"  search:"type:contains;column:name;table:ev_film_info" comment:"电影名"`
	ReleaseTime    time.Time `form:"releaseTime"  search:"type:exact;column:release_time;table:ev_film_info" comment:"上映时间"`
	Country        string    `form:"country"  search:"type:exact;column:country;table:ev_film_info" comment:"国家"`
	Type           string    `form:"type"  search:"type:exact;column:type;table:ev_film_info" comment:"电影类型"`
	Status         string    `form:"status"  search:"type:exact;column:status;table:ev_film_info" comment:"状态"`
	Language       string    `form:"language"  search:"type:exact;column:language;table:ev_film_info" comment:"语言"`
	EvFilmInfoOrder
}

type EvFilmInfoOrder struct {
	Id                 string `form:"idOrder"  search:"type:order;column:id;table:ev_film_info"`
	Name               string `form:"nameOrder"  search:"type:order;column:name;table:ev_film_info"`
	Alias              string `form:"aliasOrder"  search:"type:order;column:alias;table:ev_film_info"`
	LeadActor          string `form:"leadActorOrder"  search:"type:order;column:lead_actor;table:ev_film_info"`
	Director           string `form:"directorOrder"  search:"type:order;column:director;table:ev_film_info"`
	BriefIntroduction  string `form:"briefIntroductionOrder"  search:"type:order;column:brief_introduction;table:ev_film_info"`
	Cover              string `form:"coverOrder"  search:"type:order;column:cover;table:ev_film_info"`
	ReleaseTime        string `form:"releaseTimeOrder"  search:"type:order;column:release_time;table:ev_film_info"`
	ReleaseRelease     string `form:"releaseReleaseOrder"  search:"type:order;column:release_release;table:ev_film_info"`
	Country            string `form:"countryOrder"  search:"type:order;column:country;table:ev_film_info"`
	Type               string `form:"typeOrder"  search:"type:order;column:type;table:ev_film_info"`
	Status             string `form:"statusOrder"  search:"type:order;column:status;table:ev_film_info"`
	Language           string `form:"languageOrder"  search:"type:order;column:language;table:ev_film_info"`
	FilmLength         string `form:"filmLengthOrder"  search:"type:order;column:film_length;table:ev_film_info"`
	TppNum             string `form:"tppNumOrder"  search:"type:order;column:tpp_num;table:ev_film_info"`
	MyNum              string `form:"myNumOrder"  search:"type:order;column:my_num;table:ev_film_info"`
	DbNum              string `form:"dbNumOrder"  search:"type:order;column:db_num;table:ev_film_info"`
	MtNum              string `form:"mtNumOrder"  search:"type:order;column:mt_num;table:ev_film_info"`
	AccReleaseDays     string `form:"accReleaseDaysOrder"  search:"type:order;column:acc_release_days;table:ev_film_info"`
	ReleaseVenne       string `form:"releaseVenneOrder"  search:"type:order;column:release_venne;table:ev_film_info"`
	ReleaseNumPeople   string `form:"releaseNumPeopleOrder"  search:"type:order;column:release_num_people;table:ev_film_info"`
	OccupancyRate      string `form:"occupancyRateOrder"  search:"type:order;column:occupancy_rate;table:ev_film_info"`
	TotalBoxOffice     string `form:"totalBoxOfficeOrder"  search:"type:order;column:total_box_office;table:ev_film_info"`
	BoxOffice          string `form:"boxOfficeOrder"  search:"type:order;column:box_office;table:ev_film_info"`
	DianyBoxOffice     string `form:"dianyBoxOfficeOrder"  search:"type:order;column:diany_box_office;table:ev_film_info"`
	FirstDayeBoxOffice string `form:"firstDayeBoxOfficeOrder"  search:"type:order;column:first_daye_box_office;table:ev_film_info"`
	FirstWeekBoxOffice string `form:"firstWeekBoxOfficeOrder"  search:"type:order;column:first_week_box_office;table:ev_film_info"`
	TwoWeekBoxOffice   string `form:"twoWeekBoxOfficeOrder"  search:"type:order;column:two_week_box_office;table:ev_film_info"`
	WordOfMouth        string `form:"wordOfMouthOrder"  search:"type:order;column:word_of_mouth;table:ev_film_info"`
	Score              string `form:"scoreOrder"  search:"type:order;column:score;table:ev_film_info"`
	CreateBy           string `form:"createByOrder"  search:"type:order;column:create_by;table:ev_film_info"`
	UpdateBy           string `form:"updateByOrder"  search:"type:order;column:update_by;table:ev_film_info"`
	CreatedAt          string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ev_film_info"`
	UpdatedAt          string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ev_film_info"`
	DeletedAt          string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ev_film_info"`
}

func (m *EvFilmInfoGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EvFilmInfoInsertReq struct {
	Id                 int       `json:"-" comment:""` //
	Name               string    `json:"name" comment:"电影名"`
	Alias              string    `json:"alias" comment:"别名"`
	LeadActor          string    `json:"leadActor" comment:"主演"`
	Director           string    `json:"director" comment:"导演"`
	BriefIntroduction  string    `json:"briefIntroduction" comment:"简介"`
	Cover              string    `json:"cover" comment:"封面"`
	ReleaseTime        time.Time `json:"releaseTime" comment:"上映时间"`
	ReleaseRelease     string    `json:"releaseRelease" comment:"上映地区"`
	Country            string    `json:"country" comment:"国家"`
	Type               string    `json:"type" comment:"电影类型"`
	Status             string    `json:"status" comment:"状态"`
	Language           string    `json:"language" comment:"语言"`
	FilmLength         string    `json:"filmLength" comment:"片长"`
	TppNum             string    `json:"tppNum" comment:"淘票票想看数"`
	MyNum              string    `json:"myNum" comment:"猫眼想看数"`
	DbNum              string    `json:"dbNum" comment:"豆瓣想看数"`
	MtNum              string    `json:"mtNum" comment:"美团想看数"`
	AccReleaseDays     string    `json:"accReleaseDays" comment:"累计上映天数"`
	ReleaseVenne       string    `json:"releaseVenne" comment:"上映场次"`
	ReleaseNumPeople   string    `json:"releaseNumPeople" comment:"上映人次"`
	OccupancyRate      string    `json:"occupancyRate" comment:"上座率"`
	TotalBoxOffice     string    `json:"totalBoxOffice" comment:"总票房"`
	BoxOffice          string    `json:"boxOffice" comment:"当前票房"`
	DianyBoxOffice     string    `json:"dianyBoxOffice" comment:"点映票房"`
	FirstDayeBoxOffice string    `json:"firstDayeBoxOffice" comment:"首日票房"`
	FirstWeekBoxOffice string    `json:"firstWeekBoxOffice" comment:"首周票房"`
	TwoWeekBoxOffice   string    `json:"twoWeekBoxOffice" comment:"第二周票房"`
	WordOfMouth        string    `json:"wordOfMouth" comment:"口碑"`
	Score              string    `json:"score" comment:"评分"`
	common.ControlBy
}

func (s *EvFilmInfoInsertReq) Generate(model *models.EvFilmInfo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Alias = s.Alias
	model.LeadActor = s.LeadActor
	model.Director = s.Director
	model.BriefIntroduction = s.BriefIntroduction
	model.Cover = s.Cover
	model.ReleaseTime = s.ReleaseTime
	model.ReleaseRelease = s.ReleaseRelease
	model.Country = s.Country
	model.Type = s.Type
	model.Status = s.Status
	model.Language = s.Language
	model.FilmLength = s.FilmLength
	model.TppNum = s.TppNum
	model.MyNum = s.MyNum
	model.DbNum = s.DbNum
	model.MtNum = s.MtNum
	model.AccReleaseDays = s.AccReleaseDays
	model.ReleaseVenne = s.ReleaseVenne
	model.ReleaseNumPeople = s.ReleaseNumPeople
	model.OccupancyRate = s.OccupancyRate
	model.TotalBoxOffice = s.TotalBoxOffice
	model.BoxOffice = s.BoxOffice
	model.DianyBoxOffice = s.DianyBoxOffice
	model.FirstDayeBoxOffice = s.FirstDayeBoxOffice
	model.FirstWeekBoxOffice = s.FirstWeekBoxOffice
	model.TwoWeekBoxOffice = s.TwoWeekBoxOffice
	model.WordOfMouth = s.WordOfMouth
	model.Score = s.Score
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EvFilmInfoInsertReq) GetId() interface{} {
	return s.Id
}

type EvFilmInfoUpdateReq struct {
	Id                 int       `uri:"id" comment:""` //
	Name               string    `json:"name" comment:"电影名"`
	Alias              string    `json:"alias" comment:"别名"`
	LeadActor          string    `json:"leadActor" comment:"主演"`
	Director           string    `json:"director" comment:"导演"`
	BriefIntroduction  string    `json:"briefIntroduction" comment:"简介"`
	Cover              string    `json:"cover" comment:"封面"`
	ReleaseTime        time.Time `json:"releaseTime" comment:"上映时间"`
	ReleaseRelease     string    `json:"releaseRelease" comment:"上映地区"`
	Country            string    `json:"country" comment:"国家"`
	Type               string    `json:"type" comment:"电影类型"`
	Status             string    `json:"status" comment:"状态"`
	Language           string    `json:"language" comment:"语言"`
	FilmLength         string    `json:"filmLength" comment:"片长"`
	TppNum             string    `json:"tppNum" comment:"淘票票想看数"`
	MyNum              string    `json:"myNum" comment:"猫眼想看数"`
	DbNum              string    `json:"dbNum" comment:"豆瓣想看数"`
	MtNum              string    `json:"mtNum" comment:"美团想看数"`
	AccReleaseDays     string    `json:"accReleaseDays" comment:"累计上映天数"`
	ReleaseVenne       string    `json:"releaseVenne" comment:"上映场次"`
	ReleaseNumPeople   string    `json:"releaseNumPeople" comment:"上映人次"`
	OccupancyRate      string    `json:"occupancyRate" comment:"上座率"`
	TotalBoxOffice     string    `json:"totalBoxOffice" comment:"总票房"`
	BoxOffice          string    `json:"boxOffice" comment:"当前票房"`
	DianyBoxOffice     string    `json:"dianyBoxOffice" comment:"点映票房"`
	FirstDayeBoxOffice string    `json:"firstDayeBoxOffice" comment:"首日票房"`
	FirstWeekBoxOffice string    `json:"firstWeekBoxOffice" comment:"首周票房"`
	TwoWeekBoxOffice   string    `json:"twoWeekBoxOffice" comment:"第二周票房"`
	WordOfMouth        string    `json:"wordOfMouth" comment:"口碑"`
	Score              string    `json:"score" comment:"评分"`
	common.ControlBy
}

func (s *EvFilmInfoUpdateReq) Generate(model *models.EvFilmInfo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Alias = s.Alias
	model.LeadActor = s.LeadActor
	model.Director = s.Director
	model.BriefIntroduction = s.BriefIntroduction
	model.Cover = s.Cover
	model.ReleaseTime = s.ReleaseTime
	model.ReleaseRelease = s.ReleaseRelease
	model.Country = s.Country
	model.Type = s.Type
	model.Status = s.Status
	model.Language = s.Language
	model.FilmLength = s.FilmLength
	model.TppNum = s.TppNum
	model.MyNum = s.MyNum
	model.DbNum = s.DbNum
	model.MtNum = s.MtNum
	model.AccReleaseDays = s.AccReleaseDays
	model.ReleaseVenne = s.ReleaseVenne
	model.ReleaseNumPeople = s.ReleaseNumPeople
	model.OccupancyRate = s.OccupancyRate
	model.TotalBoxOffice = s.TotalBoxOffice
	model.BoxOffice = s.BoxOffice
	model.DianyBoxOffice = s.DianyBoxOffice
	model.FirstDayeBoxOffice = s.FirstDayeBoxOffice
	model.FirstWeekBoxOffice = s.FirstWeekBoxOffice
	model.TwoWeekBoxOffice = s.TwoWeekBoxOffice
	model.WordOfMouth = s.WordOfMouth
	model.Score = s.Score
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EvFilmInfoUpdateReq) GetId() interface{} {
	return s.Id
}

// EvFilmInfoGetReq 功能获取请求参数
type EvFilmInfoGetReq struct {
	Id int `uri:"id"`
}

func (s *EvFilmInfoGetReq) GetId() interface{} {
	return s.Id
}

// EvFilmInfoDeleteReq 功能删除请求参数
type EvFilmInfoDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EvFilmInfoDeleteReq) GetId() interface{} {
	return s.Ids
}
