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

type EmCollegeInfo struct {
	api.Api
}

// GetPage 获取学院信息列表
// @Summary 获取学院信息列表
// @Description 获取学院信息列表
// @Tags 学院信息
// @Param name query string false "学院名"
// @Param rePerson query string false "负责人姓名"
// @Param rePhone query string false "负责人手机号"
// @Param reNo query string false "负责人工号"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.EmCollegeInfo}} "{"code": 200, "data": [...]}"
// @Router /api/v1/em-college-info [get]
// @Security Bearer
func (e EmCollegeInfo) GetPage(c *gin.Context) {
    req := dto.EmCollegeInfoGetPageReq{}
    s := service.EmCollegeInfo{}
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
	list := make([]models.EmCollegeInfo, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取学院信息失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取学院信息
// @Summary 获取学院信息
// @Description 获取学院信息
// @Tags 学院信息
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.EmCollegeInfo} "{"code": 200, "data": [...]}"
// @Router /api/v1/em-college-info/{id} [get]
// @Security Bearer
func (e EmCollegeInfo) Get(c *gin.Context) {
	req := dto.EmCollegeInfoGetReq{}
	s := service.EmCollegeInfo{}
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
	var object models.EmCollegeInfo

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取学院信息失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建学院信息
// @Summary 创建学院信息
// @Description 创建学院信息
// @Tags 学院信息
// @Accept application/json
// @Product application/json
// @Param data body dto.EmCollegeInfoInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/em-college-info [post]
// @Security Bearer
func (e EmCollegeInfo) Insert(c *gin.Context) {
    req := dto.EmCollegeInfoInsertReq{}
    s := service.EmCollegeInfo{}
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
		e.Error(500, err, fmt.Sprintf("创建学院信息失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改学院信息
// @Summary 修改学院信息
// @Description 修改学院信息
// @Tags 学院信息
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.EmCollegeInfoUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/em-college-info/{id} [put]
// @Security Bearer
func (e EmCollegeInfo) Update(c *gin.Context) {
    req := dto.EmCollegeInfoUpdateReq{}
    s := service.EmCollegeInfo{}
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
		e.Error(500, err, fmt.Sprintf("修改学院信息失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除学院信息
// @Summary 删除学院信息
// @Description 删除学院信息
// @Tags 学院信息
// @Param data body dto.EmCollegeInfoDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/em-college-info [delete]
// @Security Bearer
func (e EmCollegeInfo) Delete(c *gin.Context) {
    s := service.EmCollegeInfo{}
    req := dto.EmCollegeInfoDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除学院信息失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
