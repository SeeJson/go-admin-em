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

type EvFilmInfo struct {
	service.Service
}

// GetPage 获取EvFilmInfo列表
func (e *EvFilmInfo) GetPage(c *dto.EvFilmInfoGetPageReq, p *actions.DataPermission, list *[]models.EvFilmInfo, count *int64) error {
	var err error
	var data models.EvFilmInfo

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EvFilmInfoService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetPage 获在在线电影列表
func (e *EvFilmInfo) GetOnline(c *dto.EvFilmInfoGetPageReq, p *actions.DataPermission, list *[]models.EvFilmInfo, count *int64) error {
	var err error
	var data models.EvFilmInfo

	e.Orm.Model(&data).Find(list).Limit(10)
	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EvFilmInfoService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// 获取总票房信息
func (e *EvFilmInfo) GetBasepf(c dto.EvFilmInfoGetPageReq, p *actions.DataPermission, mapFilmInfo map[string]interface{}, cloumnName string) (err error) {
	list := make([]models.EvFilmInfo, 0)
	c.BoxOffice = "desc"
	var data models.EvFilmInfo
	if err = e.Orm.Model(&data).Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		actions.Permission(data.TableName(), p),
	).Find(&list).Error; err != nil {
		e.Log.Errorf("GetBase %s GetPage fail, error:[%s] \r\n", cloumnName, err)
		return
	}
	baseFilmRate := make([]BaseFilmRate, 0)
	strName := make([]string, 0)
	strData := make([]string, 0)
	for _, info := range list {
		strName = append(strName, info.Name)
		strData = append(strData, info.BoxOffice)
		baseFilmRate = append(baseFilmRate, BaseFilmRate{
			Name:  info.Name,
			Value: info.BoxOffice,
		})
	}
	baseFilm := BaseFilm{
		Name: strName,
		Data: strData,
	}
	mapFilmInfo[cloumnName] = baseFilm
	mapFilmInfo[cloumnName+rata] = baseFilmRate
	return
}

// 获取人次
func (e *EvFilmInfo) GetBaserc(c dto.EvFilmInfoGetPageReq, p *actions.DataPermission, mapFilmInfo map[string]interface{}, cloumnName string) (err error) {
	list := make([]models.EvFilmInfo, 0)
	c.ReleaseNumPeople = "desc"
	var data models.EvFilmInfo
	if err = e.Orm.Model(&data).Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		actions.Permission(data.TableName(), p),
	).Find(&list).Error; err != nil {
		e.Log.Errorf("GetBase %s GetPage fail, error:[%s] \r\n", cloumnName, err)
		return
	}
	baseFilmRate := make([]BaseFilmRate, 0)
	strName := make([]string, 0)
	strData := make([]string, 0)
	for _, info := range list {
		strName = append(strName, info.Name)
		strData = append(strData, info.ReleaseNumPeople)
		baseFilmRate = append(baseFilmRate, BaseFilmRate{
			Name:  info.Name,
			Value: info.ReleaseNumPeople,
		})
	}
	baseFilm := BaseFilm{
		Name: strName,
		Data: strData,
	}
	mapFilmInfo[cloumnName] = baseFilm
	mapFilmInfo[cloumnName+rata] = baseFilmRate
	return
}

// 获取电影信息--场次
func (e *EvFilmInfo) GetBasecc(c dto.EvFilmInfoGetPageReq, p *actions.DataPermission, mapFilmInfo map[string]interface{}, cloumnName string) (err error) {
	list := make([]models.EvFilmInfo, 0)
	c.ReleaseVenne = "desc"
	var data models.EvFilmInfo
	if err = e.Orm.Model(&data).Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		actions.Permission(data.TableName(), p),
	).Find(&list).Error; err != nil {
		e.Log.Errorf("GetBase %s GetPage fail, error:[%s] \r\n", cloumnName, err)
		return
	}
	baseFilmRate := make([]BaseFilmRate, 0)
	strName := make([]string, 0)
	strData := make([]string, 0)
	for _, info := range list {
		strName = append(strName, info.Name)
		strData = append(strData, info.ReleaseVenne)
		baseFilmRate = append(baseFilmRate, BaseFilmRate{
			Name:  info.Name,
			Value: info.ReleaseVenne,
		})
	}
	baseFilm := BaseFilm{
		Name: strName,
		Data: strData,
	}
	mapFilmInfo[cloumnName] = baseFilm
	mapFilmInfo[cloumnName+rata] = baseFilmRate
	return
}

// 获取电影信息--上错率
func (e *EvFilmInfo) GetBaseszl(c dto.EvFilmInfoGetPageReq, p *actions.DataPermission, mapFilmInfo map[string]interface{}, cloumnName string) (err error) {
	list := make([]models.EvFilmInfo, 0)
	c.OccupancyRate = "desc"
	var data models.EvFilmInfo
	if err = e.Orm.Model(&data).Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		actions.Permission(data.TableName(), p),
	).Find(&list).Error; err != nil {
		e.Log.Errorf("GetBase %s GetPage fail, error:[%s] \r\n", cloumnName, err)
		return
	}
	baseFilmRate := make([]BaseFilmRate, 0)
	strName := make([]string, 0)
	strData := make([]string, 0)
	for _, info := range list {
		strName = append(strName, info.Name)
		strData = append(strData, info.OccupancyRate)
		baseFilmRate = append(baseFilmRate, BaseFilmRate{
			Name:  info.Name,
			Value: info.OccupancyRate,
		})
	}
	baseFilm := BaseFilm{
		Name: strName,
		Data: strData,
	}
	mapFilmInfo[cloumnName] = baseFilm
	mapFilmInfo[cloumnName+rata] = baseFilmRate
	return
}

// Get 获取EvFilmInfo对象
func (e *EvFilmInfo) Get(d *dto.EvFilmInfoGetReq, p *actions.DataPermission, model *models.EvFilmInfo) error {
	var data models.EvFilmInfo

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEvFilmInfo error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EvFilmInfo对象
func (e *EvFilmInfo) Insert(c *dto.EvFilmInfoInsertReq) error {
	var err error
	var data models.EvFilmInfo
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EvFilmInfoService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EvFilmInfo对象
func (e *EvFilmInfo) Update(c *dto.EvFilmInfoUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.EvFilmInfo{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("EvFilmInfoService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除EvFilmInfo
func (e *EvFilmInfo) Remove(d *dto.EvFilmInfoDeleteReq, p *actions.DataPermission) error {
	var data models.EvFilmInfo

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveEvFilmInfo error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
