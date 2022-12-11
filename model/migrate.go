package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Demo struct {
	gorm.Model
	Name string
}

func GetDataBase(db *gorm.DB) {
	str := db.Migrator().CurrentDatabase()
	fmt.Println("数据库名称：", str)
}

func ExistTableUser(db *gorm.DB) {
	isExist := db.Migrator().HasTable(&User{})
	// isExist := db.Migrator().HasTable("users")
	if !isExist {
		fmt.Printf("users 表不存在\n")
		return
	}
	fmt.Printf("users 表存在\n")
}

func DropTableDemo(db *gorm.DB) {
	err := db.Debug().Migrator().DropTable(&Demo{})
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return
	}
	// err := db.Migrator().DropTable("users")
	fmt.Printf("demos 表删除成功\n")

	// 2022/12/11 11:21:53 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:30
	// [0.158ms] [rows:0] SET FOREIGN_KEY_CHECKS = 0;

	// 2022/12/11 11:21:53 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:30
	// [1.426ms] [rows:0] DROP TABLE IF EXISTS `demos` CASCADE

	// 2022/12/11 11:21:53 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:30
	// [0.600ms] [rows:0] SET FOREIGN_KEY_CHECKS = 1;

	// 表不存在的情况，没有执行删除，表没了，相当于删除了
}

type DemoInfo struct {
	gorm.Model
	Name string `gorm:"default: salaheiyou"`
}

func RenameTableDemo(db *gorm.DB) {
	// db.Migrator().RenameTable("users", "user_infos")
	//若是users表存在则改名为user_infos表，反之亦然

	// mysql> show tables;
	// +---------------------+
	// | Tables_in_gorm_demo |
	// +---------------------+
	// | credit_cards        |
	// | dels                |
	// | mydbops_lab_test    |
	// | users               |
	// +---------------------+
	// 4 rows in set (0.00 sec)

	// if b := db.Migrator().HasTable(&Demo{}); b {
	//   db.Migrator().RenameTable(&Demo{}, &DemoInfo{})
	//   fmt.Printf("demos 表名修改成功\n")
	// } else {

	//   db.Migrator().RenameTable(&DemoInfo{}, &Demo{})
	//   fmt.Printf("demo_infos 表名修改成功\n")
	// }

	// 特意写一个没有的表，哈～
	// 2022/12/11 11:26:19 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:60 Error 1146: Table 'gorm_demo.demo_infos' doesn't exist
	// [1.079ms] [rows:0] ALTER TABLE `demo_infos` RENAME TO `demos`
	// demo_infos 表名修改成功

	c := db.Debug().Migrator().HasTable(&DemoInfo{})
	b := db.Debug().Migrator().HasTable(&Demo{})
	if b || c {

		if b := db.Migrator().HasTable(&Demo{}); b {
			db.Migrator().RenameTable(&Demo{}, &DemoInfo{})
			fmt.Printf("demos 表名修改成功\n")
		} else {

			db.Migrator().RenameTable(&DemoInfo{}, &Demo{})
			fmt.Printf("demo_infos 表名修改成功\n")
		}
	} else {
		db.AutoMigrate(&Demo{})
		fmt.Println("都没有，则初始化～")

	}

	//   2022/12/11 11:21:53 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:30
	// [0.158ms] [rows:0] SET FOREIGN_KEY_CHECKS = 0;

	// 2022/12/11 11:21:53 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:30
	// [1.426ms] [rows:0] DROP TABLE IF EXISTS `demos` CASCADE
	// ♠ /Users/chenhailong/code/github/go/gorm-demo $ go run main.go

	// 2022/12/11 11:30:19 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:67
	// [0.079ms] [rows:-] SELECT DATABASE()

	// 2022/12/11 11:30:19 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:67
	// [1.123ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'gorm_demo%' ORDER BY SCHEMA_NAME='gorm_demo' DESC,SCHEMA_NAME limit 1

	// 2022/12/11 11:30:19 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:67
	// [0.376ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'gorm_demo' AND table_name = 'demo_infos' AND table_type = 'BASE TABLE'

	// 2022/12/11 11:30:19 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:68
	// [0.401ms] [rows:-] SELECT DATABASE()

	// 2022/12/11 11:30:19 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:68
	// [1.618ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'gorm_demo%' ORDER BY SCHEMA_NAME='gorm_demo' DESC,SCHEMA_NAME limit 1

	// 2022/12/11 11:30:19 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:68
	// [0.703ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'gorm_demo' AND table_name = 'demos' AND table_type = 'BASE TABLE'

	// mysql> show tables;
	// +---------------------+
	// | Tables_in_gorm_demo |
	// +---------------------+
	// | credit_cards        |
	// | dels                |
	// | demos               |
	// | mydbops_lab_test    |
	// | users               |
	// +---------------------+
	// 5 rows in set (0.01 sec)

	// mysql> desc demos;
	// +------------+---------------------+------+-----+---------+----------------+
	// | Field      | Type                | Null | Key | Default | Extra          |
	// +------------+---------------------+------+-----+---------+----------------+
	// | id         | bigint(20) unsigned | NO   | PRI | NULL    | auto_increment |
	// | created_at | datetime(3)         | YES  |     | NULL    |                |
	// | updated_at | datetime(3)         | YES  |     | NULL    |                |
	// | deleted_at | datetime(3)         | YES  | MUL | NULL    |                |
	// | name       | longtext            | YES  |     | NULL    |                |
	// +------------+---------------------+------+-----+---------+----------------+
	// 5 rows in set (0.00 sec)

	/*=============================================== demo: 再次执行========================================================*/

	//   ♠ /Users/chenhailong/code/github/go/gorm-demo $ go run main.go

	// 2022/12/11 11:49:16 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:82
	// [0.197ms] [rows:-] SELECT DATABASE()

	// 2022/12/11 11:49:16 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:82
	// [1.975ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'gorm_demo%' ORDER BY SCHEMA_NAME='gorm_demo' DESC,SCHEMA_NAME limit 1

	// 2022/12/11 11:49:16 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:82
	// [2.490ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'gorm_demo' AND table_name = 'demo_infos' AND table_type = 'BASE TABLE'

	// 2022/12/11 11:49:16 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:83
	// [0.605ms] [rows:-] SELECT DATABASE()

	// 2022/12/11 11:49:16 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:83
	// [10.005ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'gorm_demo%' ORDER BY SCHEMA_NAME='gorm_demo' DESC,SCHEMA_NAME limit 1

	// 2022/12/11 11:49:16 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:83
	// [2.578ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'gorm_demo' AND table_name = 'demos' AND table_type = 'BASE TABLE'
	// demos 表名修改成功
	//   mysql> show tables;
	// +---------------------+
	// | Tables_in_gorm_demo |
	// +---------------------+
	// | credit_cards        |
	// | dels                |
	// | demo_infos          |
	// | mydbops_lab_test    |
	// | users               |
	// +---------------------+
	// 5 rows in set (0.00 sec)

	// mysql> desc demo_infos;
	// +------------+---------------------+------+-----+---------+----------------+
	// | Field      | Type                | Null | Key | Default | Extra          |
	// +------------+---------------------+------+-----+---------+----------------+
	// | id         | bigint(20) unsigned | NO   | PRI | NULL    | auto_increment |
	// | created_at | datetime(3)         | YES  |     | NULL    |                |
	// | updated_at | datetime(3)         | YES  |     | NULL    |                |
	// | deleted_at | datetime(3)         | YES  | MUL | NULL    |                |
	// | name       | longtext            | YES  |     | NULL    |                |
	// +------------+---------------------+------+-----+---------+----------------+
	// 5 rows in set (0.00 sec)

}

func Del_AddColumn(db *gorm.DB) {
	err := db.Debug().Migrator().AddColumn(&Del{}, "IsDel")
	// [61.220ms] [rows:0] ALTER TABLE `dels` ADD `is_del` boolean
	if err != nil {
		fmt.Printf("添加字段错误,err:%s\n", err)
		return
	}

	//	2022/12/11 12:09:05 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:201 Error 1060: Duplicate column name 'is_del'
	//
	// [0.256ms] [rows:0] ALTER TABLE `dels` ADD `is_del` boolean
	// 添加字段错误,err:Error 1060: Duplicate column name 'is_del'
}
func Del_dropColumn(db *gorm.DB) {
	err := db.Debug().Migrator().DropColumn(&Del{}, "IsDel")
	// [56.084ms] [rows:0] ALTER TABLE `dels` DROP COLUMN `is_del`
	if err != nil {
		fmt.Printf("删除字段错误,err:%s\n", err)
		return
	}

	//	2022/12/11 12:10:49 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:213 Error 1091: Can't DROP 'is_del'; check that column/key exists
	//
	// [0.311ms] [rows:0] ALTER TABLE `dels` DROP COLUMN `is_del`
	// 删除字段错误,err:Error 1091: Can't DROP 'is_del'; check that column/key exists
}

// 1. 先改结构体, 2.RenameColumn更改表, 这两步没有关系，但是会影响数据库及之后的操作
func Del_renameColumn(db *gorm.DB) {
	err := db.Migrator().RenameColumn(&Del{}, "user_name", "name")
	if err != nil {
		fmt.Printf("修改字段名错误,err:%s\n", err)
		return
	}
}
func Del_existColumn(db *gorm.DB) {
	isExistField := db.Debug().Migrator().HasColumn(&User{}, "name")
	fmt.Printf("name字段是否存在:%t\n", isExistField)
	isExistField = db.Debug().Migrator().HasColumn(&User{}, "user_name")
	fmt.Printf("user_name:%t\n", isExistField)

	// 2022/12/11 12:23:42 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:236
	// [0.148ms] [rows:-] SELECT DATABASE()

	// 2022/12/11 12:23:42 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:236
	// [5.694ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'gorm_demo%' ORDER BY SCHEMA_NAME='gorm_demo' DESC,SCHEMA_NAME limit 1

	// 2022/12/11 12:23:42 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:236
	// [1.492ms] [rows:-] SELECT count(*) FROM INFORMATION_SCHEMA.columns WHERE table_schema = 'gorm_demo' AND table_name = 'users' AND column_name = 'name'
	// name字段是否存在:true

	// 2022/12/11 12:23:42 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:238
	// [0.633ms] [rows:-] SELECT DATABASE()

	// 2022/12/11 12:23:42 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:238
	// [1.333ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'gorm_demo%' ORDER BY SCHEMA_NAME='gorm_demo' DESC,SCHEMA_NAME limit 1

	// 2022/12/11 12:23:42 /Users/chenhailong/code/github/go/gorm-demo/model/migrate.go:238
	// [2.149ms] [rows:-] SELECT count(*) FROM INFORMATION_SCHEMA.columns WHERE table_schema = 'gorm_demo' AND table_name = 'users' AND column_name = 'user_name'
	// user_name:false
}
