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

type EmStudent struct {
	api.Api
}

// GetPage 获取学生信息列表
// @Summary 获取学生信息列表
// @Description 获取学生信息列表
// @Tags 学生信息
// @Param sex query string false "性别"
// @Param name query string false "姓名"
// @Param number query string false "学号"
// @Param className query string false "班级名"
// @Param major query string false "专业名"
// @Param collegeName query string false "学院名"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.EmStudent}} "{"code": 200, "data": [...]}"
// @Router /api/v1/em-student [get]
// @Security Bearer
func (e EmStudent) GetPage(c *gin.Context) {
    req := dto.EmStudentGetPageReq{}
    s := service.EmStudent{}
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
	list := make([]models.EmStudent, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取学生信息失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取学生信息
// @Summary 获取学生信息
// @Description 获取学生信息
// @Tags 学生信息
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.EmStudent} "{"code": 200, "data": [...]}"
// @Router /api/v1/em-student/{id} [get]
// @Security Bearer
func (e EmStudent) Get(c *gin.Context) {
	req := dto.EmStudentGetReq{}
	s := service.EmStudent{}
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
	var object models.EmStudent

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取学生信息失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建学生信息
// @Summary 创建学生信息
// @Description 创建学生信息
// @Tags 学生信息
// @Accept application/json
// @Product application/json
// @Param data body dto.EmStudentInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/em-student [post]
// @Security Bearer
func (e EmStudent) Insert(c *gin.Context) {
    req := dto.EmStudentInsertReq{}
    s := service.EmStudent{}
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
		e.Error(500, err, fmt.Sprintf("创建学生信息失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改学生信息
// @Summary 修改学生信息
// @Description 修改学生信息
// @Tags 学生信息
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.EmStudentUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/em-student/{id} [put]
// @Security Bearer
func (e EmStudent) Update(c *gin.Context) {
    req := dto.EmStudentUpdateReq{}
    s := service.EmStudent{}
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
		e.Error(500, err, fmt.Sprintf("修改学生信息失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除学生信息
// @Summary 删除学生信息
// @Description 删除学生信息
// @Tags 学生信息
// @Param data body dto.EmStudentDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/em-student [delete]
// @Security Bearer
func (e EmStudent) Delete(c *gin.Context) {
    s := service.EmStudent{}
    req := dto.EmStudentDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除学生信息失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
