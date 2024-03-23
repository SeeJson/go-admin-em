package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/film/models"
	"go-admin/app/film/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type EvFilmOnline struct {
	service.Service
}

// GetPage 获取EvFilmOnline列表
func (e *EvFilmOnline) GetPage(c *dto.EvFilmOnlineGetPageReq, p *actions.DataPermission, list *[]models.EvFilmOnline, count *int64) error {
	var err error
	var data models.EvFilmOnline

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EvFilmOnlineService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EvFilmOnline对象
func (e *EvFilmOnline) Get(d *dto.EvFilmOnlineGetReq, p *actions.DataPermission, model *models.EvFilmOnline) error {
	var data models.EvFilmOnline

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEvFilmOnline error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EvFilmOnline对象
func (e *EvFilmOnline) Insert(c *dto.EvFilmOnlineInsertReq) error {
	var err error
	var data models.EvFilmOnline
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EvFilmOnlineService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EvFilmOnline对象
func (e *EvFilmOnline) Update(c *dto.EvFilmOnlineUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.EvFilmOnline{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("EvFilmOnlineService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除EvFilmOnline
func (e *EvFilmOnline) Remove(d *dto.EvFilmOnlineDeleteReq, p *actions.DataPermission) error {
	var data models.EvFilmOnline

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveEvFilmOnline error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
