package model

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"monstarlab/config"
)

// DBInstance is a singleton DB instance
type DBInstance struct {
	initializer func() interface{}
	instance    interface{}
	once        sync.Once
}

var (
	dbInstance *DBInstance
)

// Instance gets the singleton instance
func (i *DBInstance) Instance() interface{} {
	i.once.Do(func() {
		i.instance = i.initializer()
	})
	return i.instance
}

func dbInit() interface{} {
	lv := logger.Error
	if config.Server.Mode != gin.ReleaseMode {
		lv = logger.Info // output debug logs in dev mode
	}

	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(lv),
	}

	dsn := "host=0.0.0.0 user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to database")
	}

	return db
}

// DB returns the database instance
func DB() *gorm.DB {
	return dbInstance.Instance().(*gorm.DB)
}

func init() {
	dbInstance = &DBInstance{initializer: dbInit}
}
