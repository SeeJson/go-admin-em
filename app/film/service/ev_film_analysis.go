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

type EvFilmAnalysis struct {
	service.Service
}

// GetPage 获取EvFilmAnalysis列表
func (e *EvFilmAnalysis) GetPage(c *dto.EvFilmAnalysisGetPageReq, p *actions.DataPermission, list *[]models.EvFilmAnalysis, count *int64) error {
	var err error
	var data models.EvFilmAnalysis

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EvFilmAnalysisService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EvFilmAnalysis对象
func (e *EvFilmAnalysis) Get(d *dto.EvFilmAnalysisGetReq, p *actions.DataPermission, model *models.EvFilmAnalysis) error {
	var data models.EvFilmAnalysis

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEvFilmAnalysis error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EvFilmAnalysis对象
func (e *EvFilmAnalysis) Insert(c *dto.EvFilmAnalysisInsertReq) error {
	var err error
	var data models.EvFilmAnalysis
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EvFilmAnalysisService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EvFilmAnalysis对象
func (e *EvFilmAnalysis) Update(c *dto.EvFilmAnalysisUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.EvFilmAnalysis{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("EvFilmAnalysisService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除EvFilmAnalysis
func (e *EvFilmAnalysis) Remove(d *dto.EvFilmAnalysisDeleteReq, p *actions.DataPermission) error {
	var data models.EvFilmAnalysis

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveEvFilmAnalysis error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
