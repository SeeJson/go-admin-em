package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EmCollegeInfoGetPageReq struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name"  search:"type:contains;column:name;table:em_college_info" comment:"学院名"`
    RePerson string `form:"rePerson"  search:"type:contains;column:re_person;table:em_college_info" comment:"负责人姓名"`
    RePhone string `form:"rePhone"  search:"type:contains;column:re_phone;table:em_college_info" comment:"负责人手机号"`
    ReNo string `form:"reNo"  search:"type:contains;column:re_no;table:em_college_info" comment:"负责人工号"`
    EmCollegeInfoOrder
}

type EmCollegeInfoOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:em_college_info"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:em_college_info"`
    Bk string `form:"bkOrder"  search:"type:order;column:bk;table:em_college_info"`
    RePerson string `form:"rePersonOrder"  search:"type:order;column:re_person;table:em_college_info"`
    RePhone string `form:"rePhoneOrder"  search:"type:order;column:re_phone;table:em_college_info"`
    ReNo string `form:"reNoOrder"  search:"type:order;column:re_no;table:em_college_info"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:em_college_info"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:em_college_info"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:em_college_info"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:em_college_info"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:em_college_info"`
    
}

func (m *EmCollegeInfoGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EmCollegeInfoInsertReq struct {
    Id int `json:"-" comment:""` // 
    Name string `json:"name" comment:"学院名"`
    Bk string `json:"bk" comment:"学院简介"`
    RePerson string `json:"rePerson" comment:"负责人姓名"`
    RePhone string `json:"rePhone" comment:"负责人手机号"`
    ReNo string `json:"reNo" comment:"负责人工号"`
    common.ControlBy
}

func (s *EmCollegeInfoInsertReq) Generate(model *models.EmCollegeInfo)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Bk = s.Bk
    model.RePerson = s.RePerson
    model.RePhone = s.RePhone
    model.ReNo = s.ReNo
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EmCollegeInfoInsertReq) GetId() interface{} {
	return s.Id
}

type EmCollegeInfoUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Name string `json:"name" comment:"学院名"`
    Bk string `json:"bk" comment:"学院简介"`
    RePerson string `json:"rePerson" comment:"负责人姓名"`
    RePhone string `json:"rePhone" comment:"负责人手机号"`
    ReNo string `json:"reNo" comment:"负责人工号"`
    common.ControlBy
}

func (s *EmCollegeInfoUpdateReq) Generate(model *models.EmCollegeInfo)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Bk = s.Bk
    model.RePerson = s.RePerson
    model.RePhone = s.RePhone
    model.ReNo = s.ReNo
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EmCollegeInfoUpdateReq) GetId() interface{} {
	return s.Id
}

// EmCollegeInfoGetReq 功能获取请求参数
type EmCollegeInfoGetReq struct {
     Id int `uri:"id"`
}
func (s *EmCollegeInfoGetReq) GetId() interface{} {
	return s.Id
}

// EmCollegeInfoDeleteReq 功能删除请求参数
type EmCollegeInfoDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EmCollegeInfoDeleteReq) GetId() interface{} {
	return s.Ids
}
