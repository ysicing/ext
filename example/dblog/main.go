// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"github.com/ysicing/ext/exlog/dblog"
	"github.com/ysicing/ext/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func init() {
	logcfg := logger.Config{
		Simple:      true,
		HookFunc:    logger.Defaulthook(),
		JsonFormat:  false,
		CallerSkip:  false,
		ConsoleOnly: false,
		LogConfig: logger.LogConfig{
			LogPath: "./dblog",
		},
	}
	logger.InitLogger(&logcfg)
}

func main() {
	dblog1 := dblog.New(logger.Slog)
	db, err := gorm.Open(sqlite.Open("./db.sqlite"), &gorm.Config{Logger: dblog1})
	if err != nil {
		log.Println(err.Error())
	}
	db.AutoMigrate(&Product{})
	db.Debug().Create(&Product{Code: "D43", Price: 100})
	var product Product
	db.Debug().First(&product, 1)         // 根据整形主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	dblog2 := dblog.New(logger.Slog, true)
	tx := db.Session(&gorm.Session{Logger: dblog2})

	tx.Debug().Find(&Product{})
	tx.Find(&Product{})
}
