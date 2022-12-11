package model

import "gorm.io/gorm"

type CCreditCard struct {
	ID     uint
	Number string `gorm:"index:unique;size:255"`
	/*=============================================== 教程顺序 ========================================================*/

	Infos []Info `gorm:"many2many:card_infos;foreignKey:Number;joinForeignKey:card_number;references:Name;joinReferences:namem"`
	//   mysql> show create table c_credit_cards\G
	// *************************** 1. row ***************************
	//        Table: c_credit_cards
	// Create Table: CREATE TABLE `c_credit_cards` (
	//   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	//   `number` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
	//   PRIMARY KEY (`id`),
	//   KEY `unique` (`number`)
	// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci
	// 1 row in set (0.00 sec)

	// mysql> show create table infos\G
	// *************************** 1. row ***************************
	//        Table: infos
	// Create Table: CREATE TABLE `infos` (
	//   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	//   `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
	//   `age` bigint(20) DEFAULT NULL,
	//   PRIMARY KEY (`id`),
	//   KEY `unique` (`name`)
	// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci
	// 1 row in set (0.00 sec)

	// mysql> show create table card_infos\G
	// *************************** 1. row ***************************
	//
	//	Table: card_infos
	//
	// Create Table: CREATE TABLE `card_infos` (
	//
	//	`card_number` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
	//	`name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
	//	PRIMARY KEY (`card_number`,`name`),
	//	KEY `fk_card_infos_info` (`name`),
	//	CONSTRAINT `fk_card_infos_c_credit_card` FOREIGN KEY (`card_number`) REFERENCES `c_credit_cards` (`number`),
	//	CONSTRAINT `fk_card_infos_info` FOREIGN KEY (`name`) REFERENCES `infos` (`name`)
	//
	// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci
	// 1 row in set (0.00 sec)

	/*=============================================== 顺序无关 ========================================================*/

	// 顺序无关
	// Infos []Info `gorm:"many2many:card_infos;joinForeignKey:card_number;joinReferences:namem;foreignKey:Number;references: Name"`

	// mysql> show create table card_infos\G
	// *************************** 1. row ***************************
	//
	//	Table: card_infos
	//
	// Create Table: CREATE TABLE `card_infos` (
	//
	//	`card_number` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
	//	`namem` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
	//	PRIMARY KEY (`card_number`,`namem`),
	//	KEY `fk_card_infos_info` (`namem`),
	//	CONSTRAINT `fk_card_infos_c_credit_card` FOREIGN KEY (`card_number`) REFERENCES `c_credit_cards` (`number`),
	//	CONSTRAINT `fk_card_infos_info` FOREIGN KEY (`namem`) REFERENCES `infos` (`name`)
	//
	// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci
	// 1 row in set (0.00 sec)

	/*=============================================== foreignKey:Number -> references:Number ========================================================*/

	// foreignKey:Number -> references:Number
	// Infos []Info `gorm:"many2many:card_infos;joinForeignKey:card_number;joinReferences:namem;references:Number;references: Name"`

	// mysql> show create table card_infos\G
	// *************************** 1. row ***************************
	//        Table: card_infos
	// Create Table: CREATE TABLE `card_infos` (
	//   `card_number` bigint(20) unsigned NOT NULL,
	//   `namem` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
	//   PRIMARY KEY (`card_number`,`namem`),
	//   KEY `fk_card_infos_info` (`namem`),
	//   CONSTRAINT `fk_card_infos_c_credit_card` FOREIGN KEY (`card_number`) REFERENCES `c_credit_cards` (`id`),
	//   CONSTRAINT `fk_card_infos_info` FOREIGN KEY (`namem`) REFERENCES `infos` (`name`)
	// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci
	// 1 row in set (0.00 sec)

	// Infos []Info `gorm:"many2many:card_infos;joinForeignKey:card_number;joinReferences:namem;foreignKey:Number;foreignKey: Name"`

	//   2022/12/11 15:42:22 /Users/chenhailong/code/github/go/gorm-demo/model/many2many.go:96
	// [error] invalid foreign key: Name

	// 2022/12/11 15:42:22 /Users/chenhailong/code/github/go/gorm-demo/model/many2many.go:96
	// [error] failed to parse value &model.CCreditCard{ID:0x0, Number:"", Infos:[]model.Info(nil)}, got error invalid foreign key: Name

	// 2022/12/11 15:42:22 /Users/chenhailong/code/github/go/gorm-demo/model/many2many.go:96
	// [error] invalid foreign key: Name
}

type Info struct {
	ID   uint
	Name string `gorm:"index:unique;size:255"`
	Age  int
}

func MM_migration(db *gorm.DB) {
	db.AutoMigrate(&CCreditCard{}, &Info{})

}

func MM_ADD(db *gorm.DB) {
	db.Debug().Create(&CCreditCard{
		Number: "123456",
		Infos: []Info{
			{
				ID:   1,
				Name: "linzy",
				Age:  18,
			},
		},
	})
}
func MM_ADDInfoList(db *gorm.DB) {
	db.Debug().Create(&CCreditCard{
		Number: "456789",
		Infos: []Info{
			{
				ID:   2,
				Name: "slyyy",
				Age:  66,
			},
			{
				ID:   3,
				Name: "qhgwueiq",
				Age:  1,
			},
		},
	})
}
