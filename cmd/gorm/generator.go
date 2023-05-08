package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {


	g := gen.NewGenerator(gen.Config{
		OutPath: "./cmd/gorm/query",
		Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
	})

	db, _ := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/myblog?charset=utf8mb4&parseTime=True&loc=Local"))
	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(g.GenerateAllTable()...)
	g.GenerateAllTable()

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}