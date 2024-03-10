package dto

import (
    "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EmTextbookRevGetPageReq struct {
	dto.Pagination     `search:"-"`
    RevTime time.Time `form:"revTime"  search:"type:exact;column:rev_time;table:em_textbook_rev" comment:"领取日期"`
    ClassId string `form:"classId"  search:"type:contains;column:class_id;table:em_textbook_rev" comment:"班级名"`
    TextbookId string `form:"textbookId"  search:"type:exact;column:textbook_id;table:em_textbook_rev" comment:"教材名"`
    ClassLeader string `form:"classLeader"  search:"type:exact;column:class_leader;table:em_textbook_rev" comment:"班级负责人"`
    EmTextbookRevOrder
}

type EmTextbookRevOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:em_textbook_rev"`
    RevTime string `form:"revTimeOrder"  search:"type:order;column:rev_time;table:em_textbook_rev"`
    ClassId string `form:"classIdOrder"  search:"type:order;column:class_id;table:em_textbook_rev"`
    TextbookId string `form:"textbookIdOrder"  search:"type:order;column:textbook_id;table:em_textbook_rev"`
    Num string `form:"numOrder"  search:"type:order;column:num;table:em_textbook_rev"`
    ClassLeader string `form:"classLeaderOrder"  search:"type:order;column:class_leader;table:em_textbook_rev"`
    ClassLeaderPhone string `form:"classLeaderPhoneOrder"  search:"type:order;column:class_leader_phone;table:em_textbook_rev"`
    
}

func (m *EmTextbookRevGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EmTextbookRevInsertReq struct {
    Id int `json:"-" comment:""` // 
    RevTime time.Time `json:"revTime" comment:"领取日期"`
    ClassId string `json:"classId" comment:"班级名"`
    TextbookId string `json:"textbookId" comment:"教材名"`
    Num string `json:"num" comment:"数量"`
    ClassLeader string `json:"classLeader" comment:"班级负责人"`
    ClassLeaderPhone string `json:"classLeaderPhone" comment:"负责人电话"`
    common.ControlBy
}

func (s *EmTextbookRevInsertReq) Generate(model *models.EmTextbookRev)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.RevTime = s.RevTime
    model.ClassId = s.ClassId
    model.TextbookId = s.TextbookId
    model.Num = s.Num
    model.ClassLeader = s.ClassLeader
    model.ClassLeaderPhone = s.ClassLeaderPhone
}

func (s *EmTextbookRevInsertReq) GetId() interface{} {
	return s.Id
}

type EmTextbookRevUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    RevTime time.Time `json:"revTime" comment:"领取日期"`
    ClassId string `json:"classId" comment:"班级名"`
    TextbookId string `json:"textbookId" comment:"教材名"`
    Num string `json:"num" comment:"数量"`
    ClassLeader string `json:"classLeader" comment:"班级负责人"`
    ClassLeaderPhone string `json:"classLeaderPhone" comment:"负责人电话"`
    common.ControlBy
}

func (s *EmTextbookRevUpdateReq) Generate(model *models.EmTextbookRev)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.RevTime = s.RevTime
    model.ClassId = s.ClassId
    model.TextbookId = s.TextbookId
    model.Num = s.Num
    model.ClassLeader = s.ClassLeader
    model.ClassLeaderPhone = s.ClassLeaderPhone
}

func (s *EmTextbookRevUpdateReq) GetId() interface{} {
	return s.Id
}

// EmTextbookRevGetReq 功能获取请求参数
type EmTextbookRevGetReq struct {
     Id int `uri:"id"`
}
func (s *EmTextbookRevGetReq) GetId() interface{} {
	return s.Id
}

// EmTextbookRevDeleteReq 功能删除请求参数
type EmTextbookRevDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EmTextbookRevDeleteReq) GetId() interface{} {
	return s.Ids
}
