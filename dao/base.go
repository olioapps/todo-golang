package dao

import (
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	olioDAO "github.com/olioapps/service-skeleton-go/olio/dao"
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

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func (d *BaseDAO) toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
