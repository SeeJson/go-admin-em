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

type EmMajorInfo struct {
	service.Service
}

// GetPage 获取EmMajorInfo列表
func (e *EmMajorInfo) GetPage(c *dto.EmMajorInfoGetPageReq, p *actions.DataPermission, list *[]models.EmMajorInfo, count *int64) error {
	var err error
	var data models.EmMajorInfo

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EmMajorInfoService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EmMajorInfo对象
func (e *EmMajorInfo) Get(d *dto.EmMajorInfoGetReq, p *actions.DataPermission, model *models.EmMajorInfo) error {
	var data models.EmMajorInfo

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEmMajorInfo error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EmMajorInfo对象
func (e *EmMajorInfo) Insert(c *dto.EmMajorInfoInsertReq) error {
	var err error
	var data models.EmMajorInfo
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EmMajorInfoService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EmMajorInfo对象
func (e *EmMajorInfo) Update(c *dto.EmMajorInfoUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.EmMajorInfo{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("EmMajorInfoService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除EmMajorInfo
func (e *EmMajorInfo) Remove(d *dto.EmMajorInfoDeleteReq, p *actions.DataPermission) error {
	var data models.EmMajorInfo

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveEmMajorInfo error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
