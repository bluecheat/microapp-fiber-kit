package domains

import (
	"gorm.io/gorm"
)

type Board struct {
	gorm.Model

	Title   string `json:"title"`
	Content string `json:"content"`
	Writer  string `json:"writer"`

	User User `gorm:"foreignKey:Email;references:writer"`
}
