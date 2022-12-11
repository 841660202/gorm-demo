```go
package main

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
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

// 增加信用卡结构体
type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

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

	// db.Model(&User{}).Where("id = ?", 1).Update("active", nil) // null
	// db.Model(&User{}).Where("id = ?", 1).Update("active", 0)   // 0
	// db.Model(&User{}).Where("id = ?", 1).Update("active", 1)   //1
	db.Model(&User{}).Where("id = ?", 1).Update("active", 2) // 2
}

```
