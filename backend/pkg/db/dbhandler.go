package db

import (
	"gorm.io/gorm"
)

type IDbHandler struct {
	DB *gorm.DB
}

var instanceDbHandler *IDbHandler

func InitInstanceDbHandler(dbHandler *IDbHandler) {
	instanceDbHandler = dbHandler
}
func GetDbHandler() *IDbHandler {
	return instanceDbHandler
}
