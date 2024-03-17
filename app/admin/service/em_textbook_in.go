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

type EmTextbookIn struct {
	service.Service
}

// GetPage 获取EmTextbookIn列表
func (e *EmTextbookIn) GetPage(c *dto.EmTextbookInGetPageReq, p *actions.DataPermission, list *[]models.EmTextbookIn, count *int64) error {
	var err error
	var data models.EmTextbookIn

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EmTextbookInService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EmTextbookIn对象
func (e *EmTextbookIn) Get(d *dto.EmTextbookInGetReq, p *actions.DataPermission, model *models.EmTextbookIn) error {
	var data models.EmTextbookIn

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEmTextbookIn error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EmTextbookIn对象
func (e *EmTextbookIn) Insert(c *dto.EmTextbookInInsertReq) error {
	var err error
	var data models.EmTextbookIn
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EmTextbookInService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EmTextbookIn对象
func (e *EmTextbookIn) Update(c *dto.EmTextbookInUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.EmTextbookIn{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("EmTextbookInService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除EmTextbookIn
func (e *EmTextbookIn) Remove(d *dto.EmTextbookInDeleteReq, p *actions.DataPermission) error {
	var data models.EmTextbookIn

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveEmTextbookIn error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
