package model

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gorm-demo/util"

	"gorm.io/gorm"
)

// User 有一张 CreditCard，UserID 是外键
type User struct { // 拥有者
	gorm.Model
	Name         string `gorm:"default:haotian"`
	Email        *string
	Age          uint8 `gorm:"default:30"`
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	Active       sql.NullBool `gorm:"default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CreditCard   CreditCard // 一对一的关系

	FirstName string
	LastName  string

	FullName string `gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"`

	// ALTER TABLE `results` add  COLUMN `full_name` longtext  generated always as (concat(results.first_name,' ',results.last_name));
}

type Result struct { // 拥有者
	Id   int8
	Name string `json:"name"`
	Age  int8

	// ALTER TABLE `results` add  COLUMN `full_name` longtext  generated always as (concat(results.first_name,' ',results.last_name));
}

// func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
// 	// JsonPrint(u)

// 	if u.ID == 1 {
// 		return errors.New("admin user not allowed to update")
// 	}
// 	fmt.Println("User BeforeUpdate执行了")
// 	return
// }

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("User11 BeforeUpdate执行了")
	util.JsonPrint(u)

	if u.Age < 100 || u.Name == "" {
		return errors.New("invalid Age or Name")
	}

	return nil
}

func (u *User) AfterUpdate(*gorm.DB) (err error) {
	fmt.Println("User AfterUpdate执行了")
	return
}
func (u *User) BeforeSave(*gorm.DB) (err error) {
	fmt.Println("User BeforeSave执行了")
	return
}
func (u *User) AfterSave(*gorm.DB) (err error) {
	fmt.Println("User AfterSave执行了")
	return
}

func User_NamedArg(db *gorm.DB) {

	/*=============================================== demo: 1========================================================*/

	// demo: 1
	// var user User

	// db.Where("first_name = @name OR last_name = @name", sql.Named("name", "haotian")).Find(&user)
	// db.Omit(clause.Associations).Where("first_name = @name OR last_name = @name", sql.Named("name", "haotian")).Find(&user)
	// db.Omit("MemberNumber").Where("first_name = @name OR last_name = @name", sql.Named("name", "haotian")).Find(&user)

	// 上面并不能去除掉查询时候的空关联
	/*=============================================== demo: 2========================================================*/

	// 改变查询后的数据结构，哈哈，如果不改变，一堆乱七八糟的没用的数据，消耗内存
	// var result3 Result
	// db.Debug().Model(&User{}).Where("first_name = @name OR last_name = @name", map[string]interface{}{"name": "haotian"}).First(&result3)
	// [3.756ms] [rows:1] SELECT `users`.`id`,`users`.`name`,`users`.`age` FROM `users` WHERE (first_name = 'haotian' OR last_name = 'haotian') AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	// {
	//         "Id": 15,
	//         "Name": "linzy",
	//         "Age": 23
	// }%

	// 给name增加tag
	//   [1.608ms] [rows:1] SELECT `users`.`id`,`users`.`name`,`users`.`age` FROM `users` WHERE (first_name = 'haotian' OR last_name = 'haotian') AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	// {
	//         "Id": 15,
	//         "name": "linzy",
	//         "Age": 23
	// }%
	// util.JsonPrint(result3)

	/*=============================================== demo: 3========================================================*/

	// 查询到结果的
	// var result4 Result
	// // 原生 SQL 及命名参数
	// db.Debug().Raw("SELECT * FROM users WHERE first_name = @name OR last_name = @name2",
	// 	sql.Named("name", "haotian"), sql.Named("name2", "haotian")).Find(&result4)

	// util.JsonPrint(result4)

	//   2022/12/11 10:23:04 /Users/chenhailong/code/github/go/gorm-demo/model/user.go:110
	// [2.968ms] [rows:1] SELECT * FROM users WHERE first_name = 'haotian' OR last_name = 'haotian'
	// {
	//         "Id": 15,
	//         "name": "linzy",
	//         "Age": 23
	// }%
	/*=============================================== demo: 4========================================================*/
	// demo: 4
	// 试试一个没查询到结果的

	// var result4 Result
	// // 原生 SQL 及命名参数
	// db.Debug().Raw("SELECT * FROM users WHERE first_name = @name OR last_name = @name2",
	// 	sql.Named("name", "haotian"), sql.Named("name2", "哈喽")).Find(&result4)

	// util.JsonPrint(result4)

	//   2022/12/11 10:24:12 /Users/chenhailong/code/github/go/gorm-demo/model/user.go:128
	// [3.451ms] [rows:0] SELECT * FROM users WHERE first_name = 'haotian' OR last_name = '哈喽'
	// {
	//         "Id": 0,
	//         "name": "",
	//         "Age": 0
	// }%

	// 看吧还是有数据的只不过都是空值
	/*=============================================== demo: 5========================================================*/
	// var result4 Result
	// // 原生 SQL 及命名参数
	// rows := db.Debug().Raw("SELECT * FROM users WHERE first_name = @name OR last_name = @name2",
	// 	sql.Named("name", "haotian"), sql.Named("name2", "哈喽")).Find(&result4)

	// if rows.RowsAffected == 0 {

	// 	fmt.Println("no found")
	// 	return
	// }

	// util.JsonPrint(result4)

	// 2022/12/11 10:29:40 /Users/chenhailong/code/github/go/gorm-demo/model/user.go:152
	// [3.050ms] [rows:0] SELECT * FROM users WHERE first_name = 'haotian' OR last_name = '哈喽'
	// no found

	/*=============================================== demo: 6========================================================*/

	// db.Exec("UPDATE users SET name1 = @name, name2 = @name2, name3 = @name",
	// 	sql.Named("name", "linzynew"), sql.Named("name2", "linzynew2"))
	// // UPDATE users SET name1 = "linzynew", name2 = "linzynew2", name3 = "linzynew"

	/*=============================================== demo: 7========================================================*/
	// map
	// var result7 Result
	// db.Debug().Raw("SELECT * FROM users WHERE (first_name = @name or last_name = @name2)",
	// 	map[string]interface{}{"name": "chen", "name2": "haotian"}).Find(&result7)

	// util.JsonPrint(result7)

	// [0.596ms] [rows:1] SELECT * FROM users WHERE (first_name = 'chen' or last_name = 'haotian')
	// {
	//         "Id": 15,
	//         "name": "linzy",
	//         "Age": 23
	// }%
	/*=============================================== demo: 7 关掉debug========================================================*/

	var result7 Result
	db.Raw("SELECT * FROM users WHERE (first_name = @name or last_name = @name2)",
		map[string]interface{}{"name": "chen", "name2": "haotian"}).Find(&result7)

	util.JsonPrint(result7)

	// [0.273ms] [rows:0] SELECT `id` FROM `dels` WHERE isNull(name) AND `dels`.`deleted_at` IS NULL
	// {
	//         "Id": 15,
	//         "name": "linzy",
	//         "Age": 23
	// }
	/*=============================================== demo: 8========================================================*/
	// 结构体
	// 	type NamedArgument struct {
	// 		FirstName string
	// 		LastName  string
	// 	}

	// 	var result8 Result

	// 	db.Debug().Raw("SELECT * FROM users WHERE (first_name = @FirstName AND last_name = @LastName)",
	// 		NamedArgument{FirstName: "chen", LastName: "haotian"}).Find(&result8)

	// 2022/12/11 10:32:34 /Users/chenhailong/code/github/go/gorm-demo/model/user.go:184
	// [0.672ms] [rows:1] SELECT * FROM users WHERE (first_name = 'chen' AND last_name = 'haotian')

}
