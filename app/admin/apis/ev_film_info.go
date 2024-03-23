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

type EvFilmInfo struct {
	api.Api
}

// GetPage 获取电影信息列表
// @Summary 获取电影信息列表
// @Description 获取电影信息列表
// @Tags 电影信息
// @Param name query string false "电影名"
// @Param releaseTime query time.Time false "上映时间"
// @Param country query string false "国家"
// @Param type query string false "电影类型"
// @Param status query string false "状态"
// @Param language query string false "语言"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.EvFilmInfo}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ev-film-info [get]
// @Security Bearer
func (e EvFilmInfo) GetPage(c *gin.Context) {
	req := dto.EvFilmInfoGetPageReq{}
	s := service.EvFilmInfo{}
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
	list := make([]models.EvFilmInfo, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取电影信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取电影信息
// @Summary 获取电影信息
// @Description 获取电影信息
// @Tags 电影信息
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.EvFilmInfo} "{"code": 200, "data": [...]}"
// @Router /api/v1/ev-film-info/{id} [get]
// @Security Bearer
func (e EvFilmInfo) Get(c *gin.Context) {
	req := dto.EvFilmInfoGetReq{}
	s := service.EvFilmInfo{}
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
	var object models.EvFilmInfo

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取电影信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建电影信息
// @Summary 创建电影信息
// @Description 创建电影信息
// @Tags 电影信息
// @Accept application/json
// @Product application/json
// @Param data body dto.EvFilmInfoInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ev-film-info [post]
// @Security Bearer
func (e EvFilmInfo) Insert(c *gin.Context) {
	req := dto.EvFilmInfoInsertReq{}
	s := service.EvFilmInfo{}
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
		e.Error(500, err, fmt.Sprintf("创建电影信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改电影信息
// @Summary 修改电影信息
// @Description 修改电影信息
// @Tags 电影信息
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.EvFilmInfoUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ev-film-info/{id} [put]
// @Security Bearer
func (e EvFilmInfo) Update(c *gin.Context) {
	req := dto.EvFilmInfoUpdateReq{}
	s := service.EvFilmInfo{}
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
		e.Error(500, err, fmt.Sprintf("修改电影信息失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除电影信息
// @Summary 删除电影信息
// @Description 删除电影信息
// @Tags 电影信息
// @Param data body dto.EvFilmInfoDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ev-film-info [delete]
// @Security Bearer
func (e EvFilmInfo) Delete(c *gin.Context) {
	s := service.EvFilmInfo{}
	req := dto.EvFilmInfoDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除电影信息失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
