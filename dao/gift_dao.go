package dao

import (
	"github.com/go-xorm/xorm"
	"learnLottery/models"
)

type GifDao struct {
	engine *xorm.Engine
}

func NewGiftDao(engine *xorm.Engine) *GifDao {
	return &GifDao{engine: engine}
}

func (d *GifDao) Get(id int) *models.LtGift {
	data := &models.LtGift{Id: id}

	ok, err := d.engine.Get(data)
	if err != nil {
		return nil
	}
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *GifDao) GetAll() []models.LtGift {
	datalist := make([]models.LtGift, 0)
	err := d.engine.Asc("sys_status").Asc("displayorder").Find(&datalist)
	if err != nil {
		return datalist
	}
	return datalist
}

func (d *GifDao) CountAll() int64 {
	count, err := d.engine.Count(&models.LtGift{})
	if err != nil {
		return 0
	} else {
		return count
	}
}

func (d *GifDao) Delete(id int) error {
	data := &models.LtGift{SysStatus: 1}
	_, err := d.engine.Id(id).Update(data)
	return err
}

func (d *GifDao) Update(data *models.LtGift, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}
