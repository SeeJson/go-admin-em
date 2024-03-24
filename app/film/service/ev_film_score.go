package service

import (
	"errors"
	"go-admin/app/film/models"
	"go-admin/app/film/service/dto"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type EvFilmScore struct {
	service.Service
}

// GetPage 获取EvFilmScore列表
func (e *EvFilmScore) GetPage(c *dto.EvFilmScoreGetPageReq, p *actions.DataPermission, list *[]models.EvFilmScore, count *int64) error {
	var err error
	var data models.EvFilmScore

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EvFilmScoreService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EvFilmScore对象
func (e *EvFilmScore) Get(d *dto.EvFilmScoreGetReq, p *actions.DataPermission, model *models.EvFilmScore) error {
	var data models.EvFilmScore

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEvFilmScore error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EvFilmScore对象
func (e *EvFilmScore) Insert(c *dto.EvFilmScoreInsertReq) error {
	var err error
	var data models.EvFilmScore
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EvFilmScoreService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EvFilmScore对象
func (e *EvFilmScore) Update(c *dto.EvFilmScoreUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.EvFilmScore{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("EvFilmScoreService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除EvFilmScore
func (e *EvFilmScore) Remove(d *dto.EvFilmScoreDeleteReq, p *actions.DataPermission) error {
	var data models.EvFilmScore

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveEvFilmScore error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
