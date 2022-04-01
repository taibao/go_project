package library

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type GormDB struct {
	*gorm.DB
}

func NewGormDB(conf *GormConfig) (*GormDB, error) {
	dsn := fmt.Sprintf(dataSourceNameFormat,
		conf.UserName,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
	)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		err = fmt.Errorf("gorm connection:[%s] Open error:%w", conf.ConnectionName, err)
		return nil, err
	}

	rawDB, err2 := gormDB.DB()
	if err2 != nil {
		err2 = fmt.Errorf("gorm connection:[%s] DB error:%w", conf.ConnectionName, err2)
		return nil, err2
	}

	rawDB.SetConnMaxIdleTime(time.Duration(conf.MaxLifeTime) * time.Second)
	rawDB.SetMaxOpenConns(conf.MaxOpenConn)
	rawDB.SetMaxIdleConns(conf.MaxIdleConn)

	return &GormDB{gormDB}, nil
}
