package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type EmTextbook struct {
	api.Api
}

// GetPage 获取教材信息列表
// @Summary 获取教材信息列表
// @Description 获取教材信息列表
// @Tags 教材信息
// @Param name query string false "教材名"
// @Param sup query string false "教材供应商"
// @Param proData query time.Time false "出版日期"
// @Param author query string false "出版作者"
// @Param press query string false "出版社"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.EmTextbook}} "{"code": 200, "data": [...]}"
// @Router /api/v1/em-textbook [get]
// @Security Bearer
func (e EmTextbook) GetPage(c *gin.Context) {
	req := dto.EmTextbookGetPageReq{}
	s := service.EmTextbook{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.EmTextbook, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取教材信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取教材信息
// @Summary 获取教材信息
// @Description 获取教材信息
// @Tags 教材信息
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.EmTextbook} "{"code": 200, "data": [...]}"
// @Router /api/v1/em-textbook/{id} [get]
// @Security Bearer
func (e EmTextbook) Get(c *gin.Context) {
	req := dto.EmTextbookGetReq{}
	s := service.EmTextbook{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.EmTextbook

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取教材信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建教材信息
// @Summary 创建教材信息
// @Description 创建教材信息
// @Tags 教材信息
// @Accept application/json
// @Product application/json
// @Param data body dto.EmTextbookInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/em-textbook [post]
// @Security Bearer
func (e EmTextbook) Insert(c *gin.Context) {
	req := dto.EmTextbookInsertReq{}
	s := service.EmTextbook{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建教材信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改教材信息
// @Summary 修改教材信息
// @Description 修改教材信息
// @Tags 教材信息
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.EmTextbookUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/em-textbook/{id} [put]
// @Security Bearer
func (e EmTextbook) Update(c *gin.Context) {
	req := dto.EmTextbookUpdateReq{}
	s := service.EmTextbook{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改教材信息失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除教材信息
// @Summary 删除教材信息
// @Description 删除教材信息
// @Tags 教材信息
// @Param data body dto.EmTextbookDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/em-textbook [delete]
// @Security Bearer
func (e EmTextbook) Delete(c *gin.Context) {
	s := service.EmTextbook{}
	req := dto.EmTextbookDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除教材信息失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
