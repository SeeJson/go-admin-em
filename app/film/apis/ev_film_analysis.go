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

type EvFilmAnalysis struct {
	api.Api
}

// GetPage 获取电影分析列表
// @Summary 获取电影分析列表
// @Description 获取电影分析列表
// @Tags 电影分析
// @Param name query string false "name"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.EvFilmAnalysis}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ev-film-analysis [get]
// @Security Bearer
func (e EvFilmAnalysis) GetPage(c *gin.Context) {
	req := dto.EvFilmAnalysisGetPageReq{}
	s := service.EvFilmAnalysis{}
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
	list := make([]models.EvFilmAnalysis, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取电影分析失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取电影分析
// @Summary 获取电影分析
// @Description 获取电影分析
// @Tags 电影分析
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.EvFilmAnalysis} "{"code": 200, "data": [...]}"
// @Router /api/v1/ev-film-analysis/{id} [get]
// @Security Bearer
func (e EvFilmAnalysis) Get(c *gin.Context) {
	req := dto.EvFilmAnalysisGetReq{}
	s := service.EvFilmAnalysis{}
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
	var object models.EvFilmAnalysis

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取电影分析失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建电影分析
// @Summary 创建电影分析
// @Description 创建电影分析
// @Tags 电影分析
// @Accept application/json
// @Product application/json
// @Param data body dto.EvFilmAnalysisInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ev-film-analysis [post]
// @Security Bearer
func (e EvFilmAnalysis) Insert(c *gin.Context) {
	req := dto.EvFilmAnalysisInsertReq{}
	s := service.EvFilmAnalysis{}
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
		e.Error(500, err, fmt.Sprintf("创建电影分析失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改电影分析
// @Summary 修改电影分析
// @Description 修改电影分析
// @Tags 电影分析
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.EvFilmAnalysisUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ev-film-analysis/{id} [put]
// @Security Bearer
func (e EvFilmAnalysis) Update(c *gin.Context) {
	req := dto.EvFilmAnalysisUpdateReq{}
	s := service.EvFilmAnalysis{}
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
		e.Error(500, err, fmt.Sprintf("修改电影分析失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除电影分析
// @Summary 删除电影分析
// @Description 删除电影分析
// @Tags 电影分析
// @Param data body dto.EvFilmAnalysisDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ev-film-analysis [delete]
// @Security Bearer
func (e EvFilmAnalysis) Delete(c *gin.Context) {
	s := service.EvFilmAnalysis{}
	req := dto.EvFilmAnalysisDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除电影分析失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
