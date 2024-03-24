package apis

import (
	"fmt"
	"go-admin/app/film/models"
	"go-admin/app/film/service"
	"go-admin/app/film/service/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/common/actions"
)

type EvAreaFilmInfo struct {
	api.Api
}

// GetPage 获取地区影片票房信息列表
// @Summary 获取地区影片票房信息列表
// @Description 获取地区影片票房信息列表
// @Tags 地区影片票房信息
// @Param filmId query string false "电影id"
// @Param level query string false "城市级别"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.EvAreaFilmInfo}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ev-area-film-info [get]
// @Security Bearer
func (e EvAreaFilmInfo) GetPage(c *gin.Context) {
	req := dto.EvAreaFilmInfoGetPageReq{}
	s := service.EvAreaFilmInfo{}
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
	list := make([]models.EvAreaFilmInfo, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取地区影片票房信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取地区影片票房信息
// @Summary 获取地区影片票房信息
// @Description 获取地区影片票房信息
// @Tags 地区影片票房信息
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.EvAreaFilmInfo} "{"code": 200, "data": [...]}"
// @Router /api/v1/ev-area-film-info/{id} [get]
// @Security Bearer
func (e EvAreaFilmInfo) Get(c *gin.Context) {
	req := dto.EvAreaFilmInfoGetReq{}
	s := service.EvAreaFilmInfo{}
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
	var object models.EvAreaFilmInfo

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取地区影片票房信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建地区影片票房信息
// @Summary 创建地区影片票房信息
// @Description 创建地区影片票房信息
// @Tags 地区影片票房信息
// @Accept application/json
// @Product application/json
// @Param data body dto.EvAreaFilmInfoInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ev-area-film-info [post]
// @Security Bearer
func (e EvAreaFilmInfo) Insert(c *gin.Context) {
	req := dto.EvAreaFilmInfoInsertReq{}
	s := service.EvAreaFilmInfo{}
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
		e.Error(500, err, fmt.Sprintf("创建地区影片票房信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改地区影片票房信息
// @Summary 修改地区影片票房信息
// @Description 修改地区影片票房信息
// @Tags 地区影片票房信息
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.EvAreaFilmInfoUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ev-area-film-info/{id} [put]
// @Security Bearer
func (e EvAreaFilmInfo) Update(c *gin.Context) {
	req := dto.EvAreaFilmInfoUpdateReq{}
	s := service.EvAreaFilmInfo{}
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
		e.Error(500, err, fmt.Sprintf("修改地区影片票房信息失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除地区影片票房信息
// @Summary 删除地区影片票房信息
// @Description 删除地区影片票房信息
// @Tags 地区影片票房信息
// @Param data body dto.EvAreaFilmInfoDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ev-area-film-info [delete]
// @Security Bearer
func (e EvAreaFilmInfo) Delete(c *gin.Context) {
	s := service.EvAreaFilmInfo{}
	req := dto.EvAreaFilmInfoDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除地区影片票房信息失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
