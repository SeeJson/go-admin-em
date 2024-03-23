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

type EvFilmScore struct {
	api.Api
}

// GetPage 获取电影评分表列表
// @Summary 获取电影评分表列表
// @Description 获取电影评分表列表
// @Tags 电影评分表
// @Param filmId query string false "电影id"
// @Param userId query string false "用户id"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.EvFilmScore}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ev-film-score [get]
// @Security Bearer
func (e EvFilmScore) GetPage(c *gin.Context) {
	req := dto.EvFilmScoreGetPageReq{}
	s := service.EvFilmScore{}
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
	list := make([]models.EvFilmScore, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取电影评分表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取电影评分表
// @Summary 获取电影评分表
// @Description 获取电影评分表
// @Tags 电影评分表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.EvFilmScore} "{"code": 200, "data": [...]}"
// @Router /api/v1/ev-film-score/{id} [get]
// @Security Bearer
func (e EvFilmScore) Get(c *gin.Context) {
	req := dto.EvFilmScoreGetReq{}
	s := service.EvFilmScore{}
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
	var object models.EvFilmScore

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取电影评分表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建电影评分表
// @Summary 创建电影评分表
// @Description 创建电影评分表
// @Tags 电影评分表
// @Accept application/json
// @Product application/json
// @Param data body dto.EvFilmScoreInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ev-film-score [post]
// @Security Bearer
func (e EvFilmScore) Insert(c *gin.Context) {
	req := dto.EvFilmScoreInsertReq{}
	s := service.EvFilmScore{}
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
		e.Error(500, err, fmt.Sprintf("创建电影评分表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改电影评分表
// @Summary 修改电影评分表
// @Description 修改电影评分表
// @Tags 电影评分表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.EvFilmScoreUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ev-film-score/{id} [put]
// @Security Bearer
func (e EvFilmScore) Update(c *gin.Context) {
	req := dto.EvFilmScoreUpdateReq{}
	s := service.EvFilmScore{}
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
		e.Error(500, err, fmt.Sprintf("修改电影评分表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除电影评分表
// @Summary 删除电影评分表
// @Description 删除电影评分表
// @Tags 电影评分表
// @Param data body dto.EvFilmScoreDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ev-film-score [delete]
// @Security Bearer
func (e EvFilmScore) Delete(c *gin.Context) {
	s := service.EvFilmScore{}
	req := dto.EvFilmScoreDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除电影评分表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
