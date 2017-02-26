package dao

import (
	"github.com/jinzhu/gorm"
	olioDAO "github.com/rachoac/service-skeleton-go/olio/dao"
)

type BaseDAO struct {
	connectionManager olioDAO.ConnectionProvider
}

func NewBaseDao(connectionManager olioDAO.ConnectionProvider) *BaseDAO {
	dao := BaseDAO{
		connectionManager,
	}
	return &dao
}

func (d *BaseDAO) Db() *gorm.DB {
	return d.connectionManager.GetDb()
}
