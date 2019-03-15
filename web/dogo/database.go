package dogo

import (
	"github.com/jinzhu/gorm"
	"fmt"
	// Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDatabase(cfg DatabaseConfig) (*gorm.DB,error) {
	if cfg.Type == "mysql" {
		args := fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.DatabaseName,
		)
		db,err := gorm.Open(cfg.Type,args)
		if err != nil {
			return db,err
		}
		// Max idle connections
		db.DB().SetMaxIdleConns(cfg.MaxIdleConns)

		// Max open connections
		db.DB().SetMaxOpenConns(cfg.MaxOpenConns)

		// Database logging
		db.LogMode(false)

		return db, nil
	}
	return nil, fmt.Errorf("Database type %s not supported", cfg.Type)
}

