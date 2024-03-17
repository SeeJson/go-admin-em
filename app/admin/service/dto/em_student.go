package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EmStudentGetPageReq struct {
	dto.Pagination `search:"-"`
	Sex            string `form:"sex"  search:"type:exact;column:sex;table:em_student" comment:"性别"`
	Name           string `form:"name"  search:"type:contains;column:name;table:em_student" comment:"姓名"`
	Number         string `form:"number"  search:"type:contains;column:number;table:em_student" comment:"学号"`
	ClassName      string `form:"className"  search:"type:contains;column:class_name;table:em_student" comment:"班级名"`
	Major          string `form:"major"  search:"type:contains;column:major;table:em_student" comment:"专业名"`
	CollegeName    string `form:"collegeName"  search:"type:contains;column:college_name;table:em_student" comment:"学院名"`
	EmStudentOrder
}

type EmStudentOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:em_student"`
	Sex         string `form:"sexOrder"  search:"type:order;column:sex;table:em_student"`
	Name        string `form:"nameOrder"  search:"type:order;column:name;table:em_student"`
	Number      string `form:"numberOrder"  search:"type:order;column:number;table:em_student"`
	ClassName   string `form:"classNameOrder"  search:"type:order;column:class_name;table:em_student"`
	Major       string `form:"majorOrder"  search:"type:order;column:major;table:em_student"`
	CollegeName string `form:"collegeNameOrder"  search:"type:order;column:college_name;table:em_student"`
	Phone       string `form:"phoneOrder"  search:"type:order;column:phone;table:em_student"`
	Email       string `form:"emailOrder"  search:"type:order;column:email;table:em_student"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:em_student"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:em_student"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:em_student"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:em_student"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:em_student"`
}

func (m *EmStudentGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EmStudentInsertReq struct {
	Id          int    `json:"-" comment:""` //
	Sex         string `json:"sex" comment:"性别"`
	Name        string `json:"name" comment:"姓名"`
	Number      string `json:"number" comment:"学号"`
	ClassName   string `json:"className" comment:"班级名"`
	Major       string `json:"major" comment:"专业名"`
	CollegeName string `json:"collegeName" comment:"学院名"`
	Phone       string `json:"phone" comment:"手机号"`
	Email       string `json:"email" comment:"邮箱"`
	common.ControlBy
}

func (s *EmStudentInsertReq) Generate(model *models.EmStudent) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Sex = s.Sex
	model.Name = s.Name
	model.Number = s.Number
	model.ClassName = s.ClassName
	model.Major = s.Major
	model.CollegeName = s.CollegeName
	model.Phone = s.Phone
	model.Email = s.Email
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EmStudentInsertReq) GetId() interface{} {
	return s.Id
}

type EmStudentUpdateReq struct {
	Id          int    `uri:"id" comment:""` //
	Sex         string `json:"sex" comment:"性别"`
	Name        string `json:"name" comment:"姓名"`
	Number      string `json:"number" comment:"学号"`
	ClassName   string `json:"className" comment:"班级名"`
	Major       string `json:"major" comment:"专业名"`
	CollegeName string `json:"collegeName" comment:"学院名"`
	Phone       string `json:"phone" comment:"手机号"`
	Email       string `json:"email" comment:"邮箱"`
	common.ControlBy
}

func (s *EmStudentUpdateReq) Generate(model *models.EmStudent) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Sex = s.Sex
	model.Name = s.Name
	model.Number = s.Number
	model.ClassName = s.ClassName
	model.Major = s.Major
	model.CollegeName = s.CollegeName
	model.Phone = s.Phone
	model.Email = s.Email
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EmStudentUpdateReq) GetId() interface{} {
	return s.Id
}

// EmStudentGetReq 功能获取请求参数
type EmStudentGetReq struct {
	Id int `uri:"id"`
}

func (s *EmStudentGetReq) GetId() interface{} {
	return s.Id
}

// EmStudentDeleteReq 功能删除请求参数
type EmStudentDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EmStudentDeleteReq) GetId() interface{} {
	return s.Ids
}
