package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/film/models"
	"go-admin/app/film/service"
	"go-admin/app/film/service/dto"
	"go-admin/common/actions"
)

type EvFilmOnline struct {
	api.Api
}

// GetPage 获取上映电影列表
// @Summary 获取上映电影列表
// @Description 获取上映电影列表
// @Tags 上映电影
// @Param name query string false "name"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.EvFilmOnline}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ev-film-online [get]
// @Security Bearer
func (e EvFilmOnline) GetPage(c *gin.Context) {
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
	mapoxOffice := make(map[string]interface{})
	// 获取票房
	if err = s.GetBasepf(req, p, mapoxOffice, "box_office"); err != nil {
		e.Error(500, err, fmt.Sprintf("获取电影信息失败，\r\n失败信息 %s", err.Error()))
	}
	// 获取人次
	if err = s.GetBaserc(req, p, mapoxOffice, "release_num_people"); err != nil {
		e.Error(500, err, fmt.Sprintf("获取电影信息失败，\r\n失败信息 %s", err.Error()))
	}
	// 获取场次
	if err = s.GetBasecc(req, p, mapoxOffice, "release_venne"); err != nil {
		e.Error(500, err, fmt.Sprintf("获取电影信息失败，\r\n失败信息 %s", err.Error()))
	}
	// 获取上座率
	if err = s.GetBaseszl(req, p, mapoxOffice, "occupancy_rate"); err != nil {
		e.Error(500, err, fmt.Sprintf("获取电影信息失败，\r\n失败信息 %s", err.Error()))
	}

	e.OK(mapoxOffice, "查询成功")
}

// Get 获取上映电影
// @Summary 获取上映电影
// @Description 获取上映电影
// @Tags 上映电影
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.EvFilmOnline} "{"code": 200, "data": [...]}"
// @Router /api/v1/ev-film-online/{id} [get]
// @Security Bearer
func (e EvFilmOnline) Get(c *gin.Context) {
	req := dto.EvFilmOnlineGetReq{}
	s := service.EvFilmOnline{}
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
	var object models.EvFilmOnline

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取上映电影失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建上映电影
// @Summary 创建上映电影
// @Description 创建上映电影
// @Tags 上映电影
// @Accept application/json
// @Product application/json
// @Param data body dto.EvFilmOnlineInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ev-film-online [post]
// @Security Bearer
func (e EvFilmOnline) Insert(c *gin.Context) {
	req := dto.EvFilmOnlineInsertReq{}
	s := service.EvFilmOnline{}
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
		e.Error(500, err, fmt.Sprintf("创建上映电影失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改上映电影
// @Summary 修改上映电影
// @Description 修改上映电影
// @Tags 上映电影
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.EvFilmOnlineUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ev-film-online/{id} [put]
// @Security Bearer
func (e EvFilmOnline) Update(c *gin.Context) {
	req := dto.EvFilmOnlineUpdateReq{}
	s := service.EvFilmOnline{}
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
		e.Error(500, err, fmt.Sprintf("修改上映电影失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除上映电影
// @Summary 删除上映电影
// @Description 删除上映电影
// @Tags 上映电影
// @Param data body dto.EvFilmOnlineDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ev-film-online [delete]
// @Security Bearer
func (e EvFilmOnline) Delete(c *gin.Context) {
	s := service.EvFilmOnline{}
	req := dto.EvFilmOnlineDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除上映电影失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
