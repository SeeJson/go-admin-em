package dto

import (
    "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EmTextbookInGetPageReq struct {
	dto.Pagination     `search:"-"`
    TextbookId string `form:"textbookId"  search:"type:exact;column:textbook_id;table:em_textbook_in" comment:"教材_id"`
    TextbookPro string `form:"textbookPro"  search:"type:exact;column:textbook_pro;table:em_textbook_in" comment:"供应商"`
    TextbookUnit string `form:"textbookUnit"  search:"type:contains;column:textbook_unit;table:em_textbook_in" comment:"经办单位"`
    EmTextbookInOrder
}

type EmTextbookInOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:em_textbook_in"`
    TextbookId string `form:"textbookIdOrder"  search:"type:order;column:textbook_id;table:em_textbook_in"`
    InTime string `form:"inTimeOrder"  search:"type:order;column:in_time;table:em_textbook_in"`
    Price string `form:"priceOrder"  search:"type:order;column:price;table:em_textbook_in"`
    Num string `form:"numOrder"  search:"type:order;column:num;table:em_textbook_in"`
    TextbookPro string `form:"textbookProOrder"  search:"type:order;column:textbook_pro;table:em_textbook_in"`
    TextbookUnit string `form:"textbookUnitOrder"  search:"type:order;column:textbook_unit;table:em_textbook_in"`
    
}

func (m *EmTextbookInGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EmTextbookInInsertReq struct {
    Id int `json:"-" comment:""` // 
    TextbookId string `json:"textbookId" comment:"教材_id"`
    InTime time.Time `json:"inTime" comment:"入库时间"`
    Price string `json:"price" comment:"定价"`
    Num string `json:"num" comment:"数量"`
    TextbookPro string `json:"textbookPro" comment:"供应商"`
    TextbookUnit string `json:"textbookUnit" comment:"经办单位"`
    common.ControlBy
}

func (s *EmTextbookInInsertReq) Generate(model *models.EmTextbookIn)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.TextbookId = s.TextbookId
    model.InTime = s.InTime
    model.Price = s.Price
    model.Num = s.Num
    model.TextbookPro = s.TextbookPro
    model.TextbookUnit = s.TextbookUnit
}

func (s *EmTextbookInInsertReq) GetId() interface{} {
	return s.Id
}

type EmTextbookInUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    TextbookId string `json:"textbookId" comment:"教材_id"`
    InTime time.Time `json:"inTime" comment:"入库时间"`
    Price string `json:"price" comment:"定价"`
    Num string `json:"num" comment:"数量"`
    TextbookPro string `json:"textbookPro" comment:"供应商"`
    TextbookUnit string `json:"textbookUnit" comment:"经办单位"`
    common.ControlBy
}

func (s *EmTextbookInUpdateReq) Generate(model *models.EmTextbookIn)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.TextbookId = s.TextbookId
    model.InTime = s.InTime
    model.Price = s.Price
    model.Num = s.Num
    model.TextbookPro = s.TextbookPro
    model.TextbookUnit = s.TextbookUnit
}

func (s *EmTextbookInUpdateReq) GetId() interface{} {
	return s.Id
}

// EmTextbookInGetReq 功能获取请求参数
type EmTextbookInGetReq struct {
     Id int `uri:"id"`
}
func (s *EmTextbookInGetReq) GetId() interface{} {
	return s.Id
}

// EmTextbookInDeleteReq 功能删除请求参数
type EmTextbookInDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EmTextbookInDeleteReq) GetId() interface{} {
	return s.Ids
}
