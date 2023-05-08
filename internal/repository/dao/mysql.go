package dao

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myblog/pkg/config"
	"myblog/pkg/zaplogger"
)

var DB *gorm.DB

func generateConnectionUrl() string {
	cfg := config.Get().MySqlConfig
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DatabaseName)
}

func New() error {
	url := generateConnectionUrl()

	var err error
	DB, err = gorm.Open(mysql.Open(url), &gorm.Config{
		Logger:          zaplogger.NewGormLogger(zap.L()),
		CreateBatchSize: 500, // 批量执行任务的时候, 每个批次执行任务的数量
	})

	if err != nil {
		return err
	}

	return nil

}
