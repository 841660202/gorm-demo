package model

import (
	"gorm-demo/util"

	"gorm.io/gorm"
)

type Customer struct {
	ID   int
	Name string
	//polymorphic指定多态类型，比如模型名
	Address []Address `gorm:"polymorphic:Address;"`
	// polymorphic 自动去
}

type Order struct {
	ID      int
	Name    string
	Address Address `gorm:"polymorphic:Address;polymorphicValue:or"`
}

type Address struct {
	ID          int
	Name        string
	AddressID   int
	AddressType string
}

// 如果改了 gorm:"polymorphic:Address， Address中还是ownerId,则会报错

// 2022/12/11 14:17:27 /Users/chenhailong/code/github/go/gorm-demo/model/polymorphic.go:118
// [error] failed to parse value &model.Customer{ID:0, Name:"", Address:[]model.Address(nil)}, got error invalid polymorphic type gorm-demo/model.Address for gorm-demo/model.Customer on field Address, missing field userID

// 2022/12/11 14:17:27 /Users/chenhailong/code/github/go/gorm-demo/model/polymorphic.go:118
// [error] invalid polymorphic type gorm-demo/model.Address for gorm-demo/model.Customer on field Address, missing field userID

// 2022/12/11 14:17:27 /Users/chenhailong/code/github/go/gorm-demo/model/polymorphic.go:119
// [error] invalid polymorphic type gorm-demo/model.Address for gorm-demo/model.Customer on field Address, missing field userID

// 保持多态表一致

//  查看更改后表数据

// mysql> select * from addresses;
// +----+-----------------+----------+------------+------------+--------------+
// | id | name            | owner_id | owner_type | address_id | address_type |
// +----+-----------------+----------+------------+------------+--------------+
// |  1 | 翻斗花园        |        1 | customers  |       NULL | NULL         |
// |  2 | 火星幼儿园      |        1 | orders     |       NULL | NULL         |
// |  3 | 翻斗花园        |        2 | customers  |       NULL | NULL         |
// |  4 | 火星幼儿园      |        2 | orders     |       NULL | NULL         |
// |  5 | 西溪北苑        |        3 | customers  |       NULL | NULL         |
// |  6 | 西溪新区        |        3 | customers  |       NULL | NULL         |
// |  7 | 西溪北苑        |        4 | customers  |       NULL | NULL         |
// |  8 | 西溪新区        |        4 | customers  |       NULL | NULL         |
// |  9 | -- 新区         |     NULL | NULL       |          5 | customers    |
// | 10 | -- 北苑         |     NULL | NULL       |          5 | customers    |
// +----+-----------------+----------+------------+------------+--------------+
// 10 rows in set (0.00 sec)

func Init_Polymorephic(db *gorm.DB) {
	db.AutoMigrate(&Customer{}, &Order{}, &Address{})
}
func Add_Polymorephic(db *gorm.DB) {
	// db.Debug().Create(&Customer{
	// 	Name: "linzy",
	// 	Address: Address{
	// 		Name: "翻斗花园",
	// 	},
	// })

	// db.Debug().Create(&Order{
	// 	Name: "忘崽牛奶",
	// 	Address: Address{
	// 		Name: "火星幼儿园",
	// 	},
	// })
	// 原文链接：https://blog.csdn.net/weixin_46618592/article/details/127225395

	// 下面四句sql运行的比较奇怪，不应该先插入 customers 与 orders 表，返回主键，再插入addresses表吗？

	// 2022/12/11 13:55:13 /Users/chenhailong/code/github/go/gorm-demo/model/polymorphic.go:29
	// [0.441ms] [rows:1] INSERT INTO `addresses` (`name`,`owner_id`,`owner_type`) VALUES ('翻斗花园',1,'customers') ON DUPLICATE KEY UPDATE `owner_type`=VALUES(`owner_type`),`owner_id`=VALUES(`owner_id`)

	// 2022/12/11 13:55:13 /Users/chenhailong/code/github/go/gorm-demo/model/polymorphic.go:29
	// [4.763ms] [rows:1] INSERT INTO `customers` (`name`) VALUES ('linzy')

	// 2022/12/11 13:55:13 /Users/chenhailong/code/github/go/gorm-demo/model/polymorphic.go:36
	// [8.377ms] [rows:1] INSERT INTO `addresses` (`name`,`owner_id`,`owner_type`) VALUES ('火星幼儿园',1,'orders') ON DUPLICATE KEY UPDATE `owner_type`=VALUES(`owner_type`),`owner_id`=VALUES(`owner_id`)

	// 2022/12/11 13:55:13 /Users/chenhailong/code/github/go/gorm-demo/model/polymorphic.go:36
	// [17.463ms] [rows:1] INSERT INTO `orders` (`name`) VALUES ('忘崽牛奶')

	// 再执行上面的方法，会继续插入数据

	// 2022/12/11 13:55:13 /Users/chenhailong/code/github/go/gorm-demo/model/polymorphic.go:29
	// [0.441ms] [rows:1] INSERT INTO `addresses` (`name`,`owner_id`,`owner_type`) VALUES ('翻斗花园',1,'customers') ON DUPLICATE KEY UPDATE `owner_type`=VALUES(`owner_type`),`owner_id`=VALUES(`owner_id`)
	// ♠ /Users/chenhailong/code/github/go/gorm-demo $ go run main.go

	// 2022/12/11 14:01:37 /Users/chenhailong/code/github/go/gorm-demo/model/polymorphic.go:29
	// [1.056ms] [rows:1] INSERT INTO `addresses` (`name`,`owner_id`,`owner_type`) VALUES ('翻斗花园',2,'customers') ON DUPLICATE KEY UPDATE `owner_type`=VALUES(`owner_type`),`owner_id`=VALUES(`owner_id`)

	// 2022/12/11 14:01:37 /Users/chenhailong/code/github/go/gorm-demo/model/polymorphic.go:29
	// [3.093ms] [rows:1] INSERT INTO `customers` (`name`) VALUES ('linzy')

	// 2022/12/11 14:01:37 /Users/chenhailong/code/github/go/gorm-demo/model/polymorphic.go:36
	// [6.399ms] [rows:1] INSERT INTO `addresses` (`name`,`owner_id`,`owner_type`) VALUES ('火星幼儿园',2,'orders') ON DUPLICATE KEY UPDATE `owner_type`=VALUES(`owner_type`),`owner_id`=VALUES(`owner_id`)

	// 2022/12/11 14:01:37 /Users/chenhailong/code/github/go/gorm-demo/model/polymorphic.go:36
	// [11.769ms] [rows:1] INSERT INTO `orders` (`name`) VALUES ('忘崽牛奶')

	//   mysql> select * from addresses;
	// +----+-----------------+----------+------------+
	// | id | name            | owner_id | owner_type |
	// +----+-----------------+----------+------------+
	// |  1 | 翻斗花园        |        1 | customers  |
	// |  2 | 火星幼儿园      |        1 | orders     |
	// +----+-----------------+----------+------------+
	// 2 rows in set (0.00 sec)

	// mysql> select * from addresses;
	// +----+-----------------+----------+------------+
	// | id | name            | owner_id | owner_type |
	// +----+-----------------+----------+------------+
	// |  1 | 翻斗花园        |        1 | customers  |
	// |  2 | 火星幼儿园      |        1 | orders     |
	// |  3 | 翻斗花园        |        2 | customers  |
	// |  4 | 火星幼儿园      |        2 | orders     |
	// +----+-----------------+----------+------------+
	// 4 rows in set (0.00 sec)

	// mysql> select * from orders;
	// +----+--------------+
	// | id | name         |
	// +----+--------------+
	// |  1 | 忘崽牛奶     |
	// |  2 | 忘崽牛奶     |
	// +----+--------------+
	// 2 rows in set (0.00 sec)

	// mysql> select * from customers;
	// +----+-------+
	// | id | name  |
	// +----+-------+
	// |  1 | linzy |
	// |  2 | linzy |
	// +----+-------+
	// 2 rows in set (0.00 sec)
}

// 多态并不保证唯一，只会重复的执行指定内容，满不满足 has_many并不关心
func Many_polymorphic(db *gorm.DB) {

	db.AutoMigrate(&Customer{}, &Order{}, &Address{})
	db.Create(&Customer{
		Name: "linzy",
		Address: []Address{
			{Name: "-- 新区"},
			{Name: "-- 北苑"},
		},
	})

}
func Get_polymorphic(db *gorm.DB) {

	var c []Customer

	db.Debug().Preload("Address").Find(&c)

	util.JsonPrint(c)

	// 改了多态结构，数据没有迁移

	//   2022/12/11 14:24:25 /Users/chenhailong/code/github/go/gorm-demo/model/polymorphic.go:167
	// [1.138ms] [rows:5] SELECT * FROM `customers`
	// [
	//         {
	//                 "ID": 1,
	//                 "Name": "linzy",
	//                 "Address": []
	//         },
	//         {
	//                 "ID": 2,
	//                 "Name": "linzy",
	//                 "Address": []
	//         },
	//         {
	//                 "ID": 3,
	//                 "Name": "linzy",
	//                 "Address": []
	//         },
	//         {
	//                 "ID": 4,
	//                 "Name": "linzy",
	//                 "Address": []
	//         },
	//         {
	//                 "ID": 5,
	//                 "Name": "linzy",
	//                 "Address": [
	//                         {
	//                                 "ID": 9,
	//                                 "Name": "-- 新区",
	//                                 "AddressID": 5,
	//                                 "AddressType": "customers"
	//                         },
	//                         {
	//                                 "ID": 10,
	//                                 "Name": "-- 北苑",
	//                                 "AddressID": 5,
	//                                 "AddressType": "customers"
	//                         }
	//                 ]
	//         }
	// ]%

	// 数据没有迁移
	// mysql> select * from addresses;
	// +----+-----------------+----------+------------+------------+--------------+
	// | id | name            | owner_id | owner_type | address_id | address_type |
	// +----+-----------------+----------+------------+------------+--------------+
	// |  1 | 翻斗花园        |        1 | customers  |       NULL | NULL         |
	// |  2 | 火星幼儿园      |        1 | orders     |       NULL | NULL         |
	// |  3 | 翻斗花园        |        2 | customers  |       NULL | NULL         |
	// |  4 | 火星幼儿园      |        2 | orders     |       NULL | NULL         |
	// |  5 | 西溪北苑        |        3 | customers  |       NULL | NULL         |
	// |  6 | 西溪新区        |        3 | customers  |       NULL | NULL         |
	// |  7 | 西溪北苑        |        4 | customers  |       NULL | NULL         |
	// |  8 | 西溪新区        |        4 | customers  |       NULL | NULL         |
	// |  9 | -- 新区         |     NULL | NULL       |          5 | customers    |
	// | 10 | -- 北苑         |     NULL | NULL       |          5 | customers    |
	// +----+-----------------+----------+------------+------------+--------------+

	// mysql> replace into addresses (id, name,address_id,address_type)  select id, name,owner_id,owner_type from addresses where owner_id is not null;
	// mysql> select * from addresses
	//   -> ;
	// +----+-----------------+----------+------------+------------+--------------+
	// | id | name            | owner_id | owner_type | address_id | address_type |
	// +----+-----------------+----------+------------+------------+--------------+
	// |  1 | 翻斗花园        |     NULL | NULL       |          1 | customers    |
	// |  2 | 火星幼儿园      |     NULL | NULL       |          1 | orders       |
	// |  3 | 翻斗花园        |     NULL | NULL       |          2 | customers    |
	// |  4 | 火星幼儿园      |     NULL | NULL       |          2 | orders       |
	// |  5 | 西溪北苑        |     NULL | NULL       |          3 | customers    |
	// |  6 | 西溪新区        |     NULL | NULL       |          3 | customers    |
	// |  7 | 西溪北苑        |     NULL | NULL       |          4 | customers    |
	// |  8 | 西溪新区        |     NULL | NULL       |          4 | customers    |
	// |  9 | -- 新区         |     NULL | NULL       |          5 | customers    |
	// | 10 | -- 北苑         |     NULL | NULL       |          5 | customers    |
	// +----+-----------------+----------+------------+------------+--------------+
	// 10 rows in set (0.00 sec)

	// 删除表中没有使用的列
	// mysql> alter table addresses drop column owner_id;
	// Query OK, 0 rows affected (0.06 sec)
	// Records: 0  Duplicates: 0  Warnings: 0

	// mysql> select * from addresses
	//     -> ;
	// +----+-----------------+------------+------------+--------------+
	// | id | name            | owner_type | address_id | address_type |
	// +----+-----------------+------------+------------+--------------+
	// |  1 | 翻斗花园        | NULL       |          1 | customers    |
	// |  2 | 火星幼儿园      | NULL       |          1 | orders       |
	// |  3 | 翻斗花园        | NULL       |          2 | customers    |
	// |  4 | 火星幼儿园      | NULL       |          2 | orders       |
	// |  5 | 西溪北苑        | NULL       |          3 | customers    |
	// |  6 | 西溪新区        | NULL       |          3 | customers    |
	// |  7 | 西溪北苑        | NULL       |          4 | customers    |
	// |  8 | 西溪新区        | NULL       |          4 | customers    |
	// |  9 | -- 新区         | NULL       |          5 | customers    |
	// | 10 | -- 北苑         | NULL       |          5 | customers    |
	// +----+-----------------+------------+------------+--------------+
	// 10 rows in set (0.00 sec)

}
