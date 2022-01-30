package libs

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (c *DbConfig) InitMysqlDB() *gorm.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database, c.Charset)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	return db
}
