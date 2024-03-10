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

type EmTextbookSup struct {
	service.Service
}

// GetPage 获取EmTextbookSup列表
func (e *EmTextbookSup) GetPage(c *dto.EmTextbookSupGetPageReq, p *actions.DataPermission, list *[]models.EmTextbookSup, count *int64) error {
	var err error
	var data models.EmTextbookSup

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EmTextbookSupService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EmTextbookSup对象
func (e *EmTextbookSup) Get(d *dto.EmTextbookSupGetReq, p *actions.DataPermission, model *models.EmTextbookSup) error {
	var data models.EmTextbookSup

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEmTextbookSup error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EmTextbookSup对象
func (e *EmTextbookSup) Insert(c *dto.EmTextbookSupInsertReq) error {
    var err error
    var data models.EmTextbookSup
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EmTextbookSupService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EmTextbookSup对象
func (e *EmTextbookSup) Update(c *dto.EmTextbookSupUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.EmTextbookSup{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("EmTextbookSupService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除EmTextbookSup
func (e *EmTextbookSup) Remove(d *dto.EmTextbookSupDeleteReq, p *actions.DataPermission) error {
	var data models.EmTextbookSup

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveEmTextbookSup error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
