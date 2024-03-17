package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type EmTextbookRev struct {
	service.Service
}

// GetPage 获取EmTextbookRev列表
func (e *EmTextbookRev) GetPage(c *dto.EmTextbookRevGetPageReq, p *actions.DataPermission, list *[]models.EmTextbookRev, count *int64) error {
	var err error
	var data models.EmTextbookRev

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EmTextbookRevService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EmTextbookRev对象
func (e *EmTextbookRev) Get(d *dto.EmTextbookRevGetReq, p *actions.DataPermission, model *models.EmTextbookRev) error {
	var data models.EmTextbookRev

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEmTextbookRev error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EmTextbookRev对象
func (e *EmTextbookRev) Insert(c *dto.EmTextbookRevInsertReq) error {
	var err error
	var data models.EmTextbookRev
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EmTextbookRevService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EmTextbookRev对象
func (e *EmTextbookRev) Update(c *dto.EmTextbookRevUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.EmTextbookRev{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("EmTextbookRevService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除EmTextbookRev
func (e *EmTextbookRev) Remove(d *dto.EmTextbookRevDeleteReq, p *actions.DataPermission) error {
	var data models.EmTextbookRev

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveEmTextbookRev error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
