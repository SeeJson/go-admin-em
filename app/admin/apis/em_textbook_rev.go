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

type EmTextbookRev struct {
	api.Api
}

// GetPage 获取教材班级领取信息列表
// @Summary 获取教材班级领取信息列表
// @Description 获取教材班级领取信息列表
// @Tags 教材班级领取信息
// @Param revTime query time.Time false "领取日期"
// @Param classId query string false "班级名"
// @Param textbookId query string false "教材名"
// @Param classLeader query string false "班级负责人"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.EmTextbookRev}} "{"code": 200, "data": [...]}"
// @Router /api/v1/em-textbook-rev [get]
// @Security Bearer
func (e EmTextbookRev) GetPage(c *gin.Context) {
	req := dto.EmTextbookRevGetPageReq{}
	s := service.EmTextbookRev{}
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
	list := make([]models.EmTextbookRev, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取教材班级领取信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取教材班级领取信息
// @Summary 获取教材班级领取信息
// @Description 获取教材班级领取信息
// @Tags 教材班级领取信息
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.EmTextbookRev} "{"code": 200, "data": [...]}"
// @Router /api/v1/em-textbook-rev/{id} [get]
// @Security Bearer
func (e EmTextbookRev) Get(c *gin.Context) {
	req := dto.EmTextbookRevGetReq{}
	s := service.EmTextbookRev{}
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
	var object models.EmTextbookRev

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取教材班级领取信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建教材班级领取信息
// @Summary 创建教材班级领取信息
// @Description 创建教材班级领取信息
// @Tags 教材班级领取信息
// @Accept application/json
// @Product application/json
// @Param data body dto.EmTextbookRevInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/em-textbook-rev [post]
// @Security Bearer
func (e EmTextbookRev) Insert(c *gin.Context) {
	req := dto.EmTextbookRevInsertReq{}
	s := service.EmTextbookRev{}
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
		e.Error(500, err, fmt.Sprintf("创建教材班级领取信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改教材班级领取信息
// @Summary 修改教材班级领取信息
// @Description 修改教材班级领取信息
// @Tags 教材班级领取信息
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.EmTextbookRevUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/em-textbook-rev/{id} [put]
// @Security Bearer
func (e EmTextbookRev) Update(c *gin.Context) {
	req := dto.EmTextbookRevUpdateReq{}
	s := service.EmTextbookRev{}
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
		e.Error(500, err, fmt.Sprintf("修改教材班级领取信息失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除教材班级领取信息
// @Summary 删除教材班级领取信息
// @Description 删除教材班级领取信息
// @Tags 教材班级领取信息
// @Param data body dto.EmTextbookRevDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/em-textbook-rev [delete]
// @Security Bearer
func (e EmTextbookRev) Delete(c *gin.Context) {
	s := service.EmTextbookRev{}
	req := dto.EmTextbookRevDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除教材班级领取信息失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
