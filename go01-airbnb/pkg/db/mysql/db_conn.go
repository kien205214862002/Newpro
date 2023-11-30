package mysql

import (
	"fmt"
	"go01-airbnb/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.MySQL.User,
		c.MySQL.Password,
		c.MySQL.Host,
		c.MySQL.Port,
		c.MySQL.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
