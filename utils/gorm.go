package utils

import "gorm.io/gorm"

func ParseTime(model gorm.Model, format string) (createdTime string, updatedTime string) {
	createdTime = model.CreatedAt.Format(format)
	updatedTime = model.CreatedAt.Format(format)
	return createdTime, updatedTime
}
