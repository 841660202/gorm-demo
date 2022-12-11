package model

import (
	"errors"
	"fmt"

	util "gorm-demo/util"

	"gorm.io/gorm"
)

type Del struct {
	gorm.Model
	Name  string
	IsDel bool
}

func Del_InsertData(db *gorm.DB) error {
	fmt.Println("Del_InsertData")
	db.Debug().Create(&Del{Name: "hello"})
	db.Debug().Create(&Del{Name: "world"})
	db.Debug().Create(&[]Del{{Name: "b1"}, {Name: "b2"}})
	return nil

}

// 删除一个
func Del_DeleteOne(db *gorm.DB) {
	fmt.Println("Del_DeleteOne")
}

// 批量删除
func Del_DeleteMany(db *gorm.DB) {
	fmt.Println("Del_DeleteMany")
}

// 永久删除
func Del_forever(db *gorm.DB) {
	fmt.Println("Del_forever")

	// 删除前
	//   mysql> select * from dels;
	// +----+-------------------------+-------------------------+-------------------------+--------+
	// | id | created_at              | updated_at              | deleted_at              | name   |
	// +----+-------------------------+-------------------------+-------------------------+--------+
	// |  1 | 2022-12-10 19:51:42.243 | 2022-12-10 20:11:02.841 | 2022-12-10 20:24:55.004 | name:1 |
	// |  2 | 2022-12-10 19:51:42.245 | 2022-12-10 20:11:02.841 | 2022-12-10 20:32:37.268 | name:2 |
	// |  3 | 2022-12-10 19:51:42.247 | 2022-12-10 20:11:02.841 | 2022-12-10 20:28:48.284 | name:3 |
	// |  4 | 2022-12-10 19:51:42.247 | 2022-12-10 20:11:02.841 | 2022-12-10 20:31:30.141 | name:4 |
	// |  5 | 2022-12-10 19:52:55.882 | 2022-12-10 19:52:55.882 | NULL                    | hello  |
	// |  6 | 2022-12-10 19:52:55.884 | 2022-12-10 19:52:55.884 | NULL                    | world  |
	// |  7 | 2022-12-10 19:52:55.889 | 2022-12-10 19:52:55.889 | NULL                    | b1     |
	// |  8 | 2022-12-10 19:52:55.889 | 2022-12-10 19:52:55.889 | NULL                    | b2     |
	// +----+-------------------------+-------------------------+-------------------------+--------+
	// 8 rows in set (0.00 sec)
	db.Debug().Delete(&Del{}, 7)
	// [11.530ms] [rows:1] UPDATE `dels` SET `deleted_at`='2022-12-11 09:25:31.532' WHERE `dels`.`id` = 7 AND `dels`.`deleted_at` IS NULL
	db.Debug().Unscoped().Delete(&Del{}, 8)
	// [6.727ms] [rows:1] DELETE FROM `dels` WHERE `dels`.`id` = 8

	// 删除后
	//   mysql> select * from dels;
	// +----+-------------------------+-------------------------+-------------------------+--------+
	// | id | created_at              | updated_at              | deleted_at              | name   |
	// +----+-------------------------+-------------------------+-------------------------+--------+
	// |  1 | 2022-12-10 19:51:42.243 | 2022-12-10 20:11:02.841 | 2022-12-10 20:24:55.004 | name:1 |
	// |  2 | 2022-12-10 19:51:42.245 | 2022-12-10 20:11:02.841 | 2022-12-10 20:32:37.268 | name:2 |
	// |  3 | 2022-12-10 19:51:42.247 | 2022-12-10 20:11:02.841 | 2022-12-10 20:28:48.284 | name:3 |
	// |  4 | 2022-12-10 19:51:42.247 | 2022-12-10 20:11:02.841 | 2022-12-10 20:31:30.141 | name:4 |
	// |  5 | 2022-12-10 19:52:55.882 | 2022-12-10 19:52:55.882 | NULL                    | hello  |
	// |  6 | 2022-12-10 19:52:55.884 | 2022-12-10 19:52:55.884 | NULL                    | world  |
	// |  7 | 2022-12-10 19:52:55.889 | 2022-12-10 19:52:55.889 | 2022-12-11 09:25:31.532 | b1     |
	// +----+-------------------------+-------------------------+-------------------------+--------+
	// 7 rows in set (0.00 sec)

}

// 查找所有内容，包括软删除
func Del_findUnscoped(db *gorm.DB) {

	//   mysql> select * from dels;
	// +----+-------------------------+-------------------------+-------------------------+--------+
	// | id | created_at              | updated_at              | deleted_at              | name   |
	// +----+-------------------------+-------------------------+-------------------------+--------+
	// |  1 | 2022-12-10 19:51:42.243 | 2022-12-10 20:11:02.841 | 2022-12-10 20:24:55.004 | name:1 |
	// |  2 | 2022-12-10 19:51:42.245 | 2022-12-10 20:11:02.841 | 2022-12-10 20:32:37.268 | name:2 |
	// |  3 | 2022-12-10 19:51:42.247 | 2022-12-10 20:11:02.841 | 2022-12-10 20:28:48.284 | name:3 |
	// |  4 | 2022-12-10 19:51:42.247 | 2022-12-10 20:11:02.841 | 2022-12-10 20:31:30.141 | name:4 |
	// |  5 | 2022-12-10 19:52:55.882 | 2022-12-10 19:52:55.882 | NULL                    | hello  |
	// |  6 | 2022-12-10 19:52:55.884 | 2022-12-10 19:52:55.884 | NULL                    | world  |
	// |  7 | 2022-12-10 19:52:55.889 | 2022-12-10 19:52:55.889 | 2022-12-11 09:25:31.532 | b1     |
	// +----+-------------------------+-------------------------+-------------------------+--------+
	// 7 rows in set (0.00 sec)
	var dels []Del
	db.Debug().Find(&dels)
	fmt.Println(len(dels))
	// [0.421ms] [rows:2] SELECT * FROM `dels` WHERE `dels`.`deleted_at` IS NULL
	// 2
	db.Debug().Unscoped().Find(&dels)
	fmt.Println(len(dels))
	// [0.671ms] [rows:7] SELECT * FROM `dels`
	// 7

}

// 原生语句
func Del_rowsql(db *gorm.DB) {
	type Result struct {
		ID   int
		Name string
		Age  int
	}
	var result Result
	db.Debug().Raw("SELECT id, name FROM dels WHERE id = ?", 3).Scan(&result)
	// [0.438ms] [rows:1] SELECT id, name FROM dels WHERE id = 3
	util.JsonPrint(result)
	//  {
	//     "ID": 3,
	//     "Name": "name:3",
	//     "Age": 0
	// }
	var name string
	db.Debug().Raw("select name from dels where id = ?", "1").Scan(&name)
	// [1.160ms] [rows:1] select name from dels where id = '1'

	fmt.Println(name)
	// name:1
}

// 事务自动提交
func Del_tx_autoCommit(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {
		err1 := tx.Create(&User{Name: "事务1"}).Error

		if err1 != nil {
			return err1
		}

		err2 := tx.Create(&User{Name: "事务2"}).Error
		if err2 != nil {
			return err2
		}

		return nil
	})
	//	mysql> select id, name from users;
	//
	// +----+---------+
	// | id | name    |
	// +----+---------+
	// |  1 | hali    |
	// |  2 | linzy   |
	// |  3 | linzy   |
	// |  4 | linzy   |
	// |  5 | 小明    |
	// |  6 | linzy   |
	// |  7 | linzy   |
	// |  8 | linzy   |
	// | 10 | linzy   |
	// | 11 | linzy   |
	// | 12 | linzy   |
	// | 13 | linzy   |
	// | 14 | linzy   |
	// | 15 | linzy   |
	// | 16 | linzy   |
	// | 17 | 事务1   |
	// | 18 | 事务2   |
	// +----+---------+
	// 17 rows in set (0.00 sec)
}

// 事务回滚
func Del_tx_rollback(db *gorm.DB) {
	db.Debug().Transaction(func(tx *gorm.DB) error {
		err1 := tx.Debug().Create(&User{Name: "事务3"}).Error

		if err1 != nil {
			return err1
		}

		err2 := tx.Debug().Create(&User{Name: "事务4"}).Error
		if err2 != nil {
			return err2
		}

		return errors.New("不要你成功！！！")
	})

	// User BeforeSave执行了
	// User AfterSave执行了

	// 2022/12/11 21:14:50 /Users/chenhailong/code/github/go/gorm-demo/model/del.go:151
	// [2.059ms] [rows:1] INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`active`,`first_name`,`last_name`) VALUES ('2022-12-11 21:14:50.385','2022-12-11 21:14:50.385',NULL,'事务3',NULL,30,NULL,NULL,NULL,true,'','')
	// User BeforeSave执行了
	// User AfterSave执行了

	// 2022/12/11 21:14:50 /Users/chenhailong/code/github/go/gorm-demo/model/del.go:157
	// [2.872ms] [rows:1] INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`active`,`first_name`,`last_name`) VALUES ('2022-12-11 21:14:50.387','2022-12-11 21:14:50.387',NULL,'事务4',NULL,30,NULL,NULL,NULL,true,”,”)

	// mysql> select id, name from users;
	// +----+---------+
	// | id | name    |
	// +----+---------+
	// |  1 | hali    |
	// |  2 | linzy   |
	// |  3 | linzy   |
	// |  4 | linzy   |
	// |  5 | 小明    |
	// |  6 | linzy   |
	// |  7 | linzy   |
	// |  8 | linzy   |
	// | 10 | linzy   |
	// | 11 | linzy   |
	// | 12 | linzy   |
	// | 13 | linzy   |
	// | 14 | linzy   |
	// | 15 | linzy   |
	// | 16 | linzy   |
	// | 17 | 事务1   |
	// | 18 | 事务2   |
	// +----+---------+
	// 17 rows in set (0.00 sec)

}

// 关闭事务可以提升性能，在只读的情况下关闭事务，这也是为什么读写分离
// 持续会话模式
// tx := db.Session(&Session{SkipDefaultTransaction: true})
