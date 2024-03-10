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

type EmCollegeInfo struct {
	service.Service
}

// GetPage 获取EmCollegeInfo列表
func (e *EmCollegeInfo) GetPage(c *dto.EmCollegeInfoGetPageReq, p *actions.DataPermission, list *[]models.EmCollegeInfo, count *int64) error {
	var err error
	var data models.EmCollegeInfo

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EmCollegeInfoService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EmCollegeInfo对象
func (e *EmCollegeInfo) Get(d *dto.EmCollegeInfoGetReq, p *actions.DataPermission, model *models.EmCollegeInfo) error {
	var data models.EmCollegeInfo

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEmCollegeInfo error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EmCollegeInfo对象
func (e *EmCollegeInfo) Insert(c *dto.EmCollegeInfoInsertReq) error {
    var err error
    var data models.EmCollegeInfo
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EmCollegeInfoService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EmCollegeInfo对象
func (e *EmCollegeInfo) Update(c *dto.EmCollegeInfoUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.EmCollegeInfo{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("EmCollegeInfoService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除EmCollegeInfo
func (e *EmCollegeInfo) Remove(d *dto.EmCollegeInfoDeleteReq, p *actions.DataPermission) error {
	var data models.EmCollegeInfo

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveEmCollegeInfo error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
