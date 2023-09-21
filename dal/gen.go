package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const dsn = "root:123456@(localhost:3306)/hr?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	db, _ := gorm.Open(mysql.Open(dsn))
	g.UseDB(db) // reuse your gorm db

	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
