package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EmClassInfoGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:contains;column:name;table:em_class_info" comment:"班级名"`
	Major          string `form:"major"  search:"type:contains;column:major;table:em_class_info" comment:"专业"`
	College        string `form:"college"  search:"type:contains;column:college;table:em_class_info" comment:"学院"`
	RePerson       string `form:"rePerson"  search:"type:contains;column:re_person;table:em_class_info" comment:"负责人"`
	EmClassInfoOrder
}

type EmClassInfoOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:em_class_info"`
	Name       string `form:"nameOrder"  search:"type:order;column:name;table:em_class_info"`
	Major      string `form:"majorOrder"  search:"type:order;column:major;table:em_class_info"`
	College    string `form:"collegeOrder"  search:"type:order;column:college;table:em_class_info"`
	RePerson   string `form:"rePersonOrder"  search:"type:order;column:re_person;table:em_class_info"`
	RePersonNo string `form:"rePersonNoOrder"  search:"type:order;column:re_person_no;table:em_class_info"`
	CreateBy   string `form:"createByOrder"  search:"type:order;column:create_by;table:em_class_info"`
	UpdateBy   string `form:"updateByOrder"  search:"type:order;column:update_by;table:em_class_info"`
	CreatedAt  string `form:"createdAtOrder"  search:"type:order;column:created_at;table:em_class_info"`
	UpdatedAt  string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:em_class_info"`
	DeletedAt  string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:em_class_info"`
}

func (m *EmClassInfoGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EmClassInfoInsertReq struct {
	Id         int    `json:"-" comment:""` //
	Name       string `json:"name" comment:"班级名"`
	Major      string `json:"major" comment:"专业"`
	College    string `json:"college" comment:"学院"`
	RePerson   string `json:"rePerson" comment:"负责人"`
	RePersonNo string `json:"rePersonNo" comment:"负责人学号"`
	common.ControlBy
}

func (s *EmClassInfoInsertReq) Generate(model *models.EmClassInfo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Major = s.Major
	model.College = s.College
	model.RePerson = s.RePerson
	model.RePersonNo = s.RePersonNo
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EmClassInfoInsertReq) GetId() interface{} {
	return s.Id
}

type EmClassInfoUpdateReq struct {
	Id         int    `uri:"id" comment:""` //
	Name       string `json:"name" comment:"班级名"`
	Major      string `json:"major" comment:"专业"`
	College    string `json:"college" comment:"学院"`
	RePerson   string `json:"rePerson" comment:"负责人"`
	RePersonNo string `json:"rePersonNo" comment:"负责人学号"`
	common.ControlBy
}

func (s *EmClassInfoUpdateReq) Generate(model *models.EmClassInfo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Major = s.Major
	model.College = s.College
	model.RePerson = s.RePerson
	model.RePersonNo = s.RePersonNo
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EmClassInfoUpdateReq) GetId() interface{} {
	return s.Id
}

// EmClassInfoGetReq 功能获取请求参数
type EmClassInfoGetReq struct {
	Id int `uri:"id"`
}

func (s *EmClassInfoGetReq) GetId() interface{} {
	return s.Id
}

// EmClassInfoDeleteReq 功能删除请求参数
type EmClassInfoDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EmClassInfoDeleteReq) GetId() interface{} {
	return s.Ids
}
