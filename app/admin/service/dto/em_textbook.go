package dto

import (
    "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EmTextbookGetPageReq struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name"  search:"type:contains;column:name;table:em_textbook" comment:"教材名"`
    Sup string `form:"sup"  search:"type:exact;column:sup;table:em_textbook" comment:"教材供应商"`
    ProData time.Time `form:"proData"  search:"type:exact;column:pro_data;table:em_textbook" comment:"出版日期"`
    Author string `form:"author"  search:"type:contains;column:author;table:em_textbook" comment:"出版作者"`
    Press string `form:"press"  search:"type:contains;column:press;table:em_textbook" comment:"出版社"`
    EmTextbookOrder
}

type EmTextbookOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:em_textbook"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:em_textbook"`
    Price string `form:"priceOrder"  search:"type:order;column:price;table:em_textbook"`
    Sup string `form:"supOrder"  search:"type:order;column:sup;table:em_textbook"`
    ProData string `form:"proDataOrder"  search:"type:order;column:pro_data;table:em_textbook"`
    Author string `form:"authorOrder"  search:"type:order;column:author;table:em_textbook"`
    Press string `form:"pressOrder"  search:"type:order;column:press;table:em_textbook"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:em_textbook"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:em_textbook"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:em_textbook"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:em_textbook"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:em_textbook"`
    
}

func (m *EmTextbookGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EmTextbookInsertReq struct {
    Id int `json:"-" comment:""` // 
    Name string `json:"name" comment:"教材名"`
    Price string `json:"price" comment:"价格"`
    Sup string `json:"sup" comment:"教材供应商"`
    ProData time.Time `json:"proData" comment:"出版日期"`
    Author string `json:"author" comment:"出版作者"`
    Press string `json:"press" comment:"出版社"`
    common.ControlBy
}

func (s *EmTextbookInsertReq) Generate(model *models.EmTextbook)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Price = s.Price
    model.Sup = s.Sup
    model.ProData = s.ProData
    model.Author = s.Author
    model.Press = s.Press
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EmTextbookInsertReq) GetId() interface{} {
	return s.Id
}

type EmTextbookUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Name string `json:"name" comment:"教材名"`
    Price string `json:"price" comment:"价格"`
    Sup string `json:"sup" comment:"教材供应商"`
    ProData time.Time `json:"proData" comment:"出版日期"`
    Author string `json:"author" comment:"出版作者"`
    Press string `json:"press" comment:"出版社"`
    common.ControlBy
}

func (s *EmTextbookUpdateReq) Generate(model *models.EmTextbook)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Price = s.Price
    model.Sup = s.Sup
    model.ProData = s.ProData
    model.Author = s.Author
    model.Press = s.Press
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EmTextbookUpdateReq) GetId() interface{} {
	return s.Id
}

// EmTextbookGetReq 功能获取请求参数
type EmTextbookGetReq struct {
     Id int `uri:"id"`
}
func (s *EmTextbookGetReq) GetId() interface{} {
	return s.Id
}

// EmTextbookDeleteReq 功能删除请求参数
type EmTextbookDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EmTextbookDeleteReq) GetId() interface{} {
	return s.Ids
}
