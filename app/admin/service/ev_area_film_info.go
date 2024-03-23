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

type EvAreaFilmInfo struct {
	service.Service
}

// GetPage 获取EvAreaFilmInfo列表
func (e *EvAreaFilmInfo) GetPage(c *dto.EvAreaFilmInfoGetPageReq, p *actions.DataPermission, list *[]models.EvAreaFilmInfo, count *int64) error {
	var err error
	var data models.EvAreaFilmInfo

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EvAreaFilmInfoService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EvAreaFilmInfo对象
func (e *EvAreaFilmInfo) Get(d *dto.EvAreaFilmInfoGetReq, p *actions.DataPermission, model *models.EvAreaFilmInfo) error {
	var data models.EvAreaFilmInfo

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEvAreaFilmInfo error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EvAreaFilmInfo对象
func (e *EvAreaFilmInfo) Insert(c *dto.EvAreaFilmInfoInsertReq) error {
	var err error
	var data models.EvAreaFilmInfo
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EvAreaFilmInfoService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EvAreaFilmInfo对象
func (e *EvAreaFilmInfo) Update(c *dto.EvAreaFilmInfoUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.EvAreaFilmInfo{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("EvAreaFilmInfoService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除EvAreaFilmInfo
func (e *EvAreaFilmInfo) Remove(d *dto.EvAreaFilmInfoDeleteReq, p *actions.DataPermission) error {
	var data models.EvAreaFilmInfo

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveEvAreaFilmInfo error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
