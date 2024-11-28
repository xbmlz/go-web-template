package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Args     string `json:"args"`
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %s", c.Host, c.Port, c.User, c.Password, c.Name, c.Args)
}

func Init(c *Config) (*gorm.DB, error) {
	ormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(c.DSN()), &gorm.Config{
		Logger: ormLogger,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MustInit(c *Config) *gorm.DB {
	db, err := Init(c)
	if err != nil {
		panic(err)
	}
	return db
}