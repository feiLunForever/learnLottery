package services

import (
	"learnLottery/dao"
	"learnLottery/models"
)

type GiftService interface {
	CountAll() int64
	Get(id int) *models.LtGift
	Delete(id int) error
	Update(data *models.LtGift, columns []string) error
	Create(data *models.LtGift) error
}

type giftService struct {
	dao *dao.GiftDao
}

func (s *giftService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *giftService) Get(id int) *models.LtGift {
	return s.dao.Get(id)
}

func (s *giftService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *giftService) Update(data *models.LtGift, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *giftService) Create(data *models.LtGift) error {
	return s.dao.Create(data)
}

func NewGiftService() GiftService {
	return &giftService{
		dao: nil,
	}
}
