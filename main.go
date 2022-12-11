package main

import (
	"gorm-demo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Result struct {
	ID   uint8
	Name string
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(localhost:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// db.AutoMigrate(&model.Del{})

	// 删除
	// model.Del_InsertData(db)
	// model.Del_DeleteOne(db)
	// model.Del_DeleteMany(db)
	// model.Del_forever(db)
	// model.Del_findUnscoped(db)
	// model.Del_rowsql(db)

	// 架构
	// conf
	// api
	// dao
	// service
	// model
	// utils

	// sql构建器
	// model.Del_rowsql(db)
	// 命名参数

	// model.User_NamedArg(db)

	// migrate

	// model.GetDataBase(db)
	/*=============================================== table ========================================================*/

	// model.ExistTableUser(db)
	// model.DropTableDemo(db)
	// model.RenameTableDemo(db)

	/*=============================================== column ========================================================*/

	// model.Del_AddColumn(db)
	// model.Del_dropColumn(db)
	// model.Del_renameColumn(db)
	// model.Del_existColumn(db)

	/*=============================================== 多态 ========================================================*/

	// model.Init_Polymorephic(db)
	// model.Add_Polymorephic(db)
	// model.Many_polymorphic(db)
	// model.Get_polymorphic(db)

	/*=============================================== many2many ========================================================*/

	// model.MM_migration(db)
	// model.MM_ADD(db)
	// model.MM_ADDInfoList(db)

	/*=============================================== 事务 ========================================================*/

	// model.Del_tx_autoCommit(db)
	model.Del_tx_rollback(db)
}
