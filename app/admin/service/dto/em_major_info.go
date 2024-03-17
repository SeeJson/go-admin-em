package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EmMajorInfoGetPageReq struct {
	dto.Pagination `search:"-"`
	MajorName      string `form:"majorName"  search:"type:exact;column:major_name;table:em_major_info" comment:"所属专业"`
	CollegeName    string `form:"collegeName"  search:"type:exact;column:college_name;table:em_major_info" comment:"所属学院"`
	EmMajorInfoOrder
}

type EmMajorInfoOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:em_major_info"`
	MajorName   string `form:"majorNameOrder"  search:"type:order;column:major_name;table:em_major_info"`
	MajorBk     string `form:"majorBkOrder"  search:"type:order;column:major_bk;table:em_major_info"`
	CollegeName string `form:"collegeNameOrder"  search:"type:order;column:college_name;table:em_major_info"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:em_major_info"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:em_major_info"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:em_major_info"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:em_major_info"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:em_major_info"`
}

func (m *EmMajorInfoGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EmMajorInfoInsertReq struct {
	Id          int    `json:"-" comment:""` //
	MajorName   string `json:"majorName" comment:"所属专业"`
	MajorBk     string `json:"majorBk" comment:"简介"`
	CollegeName string `json:"collegeName" comment:"所属学院"`
	common.ControlBy
}

func (s *EmMajorInfoInsertReq) Generate(model *models.EmMajorInfo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.MajorName = s.MajorName
	model.MajorBk = s.MajorBk
	model.CollegeName = s.CollegeName
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EmMajorInfoInsertReq) GetId() interface{} {
	return s.Id
}

type EmMajorInfoUpdateReq struct {
	Id          int    `uri:"id" comment:""` //
	MajorName   string `json:"majorName" comment:"所属专业"`
	MajorBk     string `json:"majorBk" comment:"简介"`
	CollegeName string `json:"collegeName" comment:"所属学院"`
	common.ControlBy
}

func (s *EmMajorInfoUpdateReq) Generate(model *models.EmMajorInfo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.MajorName = s.MajorName
	model.MajorBk = s.MajorBk
	model.CollegeName = s.CollegeName
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EmMajorInfoUpdateReq) GetId() interface{} {
	return s.Id
}

// EmMajorInfoGetReq 功能获取请求参数
type EmMajorInfoGetReq struct {
	Id int `uri:"id"`
}

func (s *EmMajorInfoGetReq) GetId() interface{} {
	return s.Id
}

// EmMajorInfoDeleteReq 功能删除请求参数
type EmMajorInfoDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EmMajorInfoDeleteReq) GetId() interface{} {
	return s.Ids
}
