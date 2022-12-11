```go
package main

import (
"bytes"
"database/sql"
"encoding/json"
"log"
"os"
"time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"

)

// User 有一张 CreditCard，UserID 是外键
type User struct { // 拥有者
gorm.Model
Name string `gorm:"default:haotian"`
Email *string
Age uint8 `gorm:"default:30"`
Birthday *time.Time
MemberNumber sql.NullString
ActivatedAt sql.NullTime
Active sql.NullBool `gorm:"default:true"`
CreatedAt time.Time
UpdatedAt time.Time
CreditCard CreditCard // 一对一的关系

    FirstName string
    LastName  string

    FullName string `gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"`

    // ALTER TABLE `users` add  COLUMN `full_name` longtext  generated always as (concat(users.first_name,' ',users.last_name));

}

// 增加信用卡结构体
type CreditCard struct {
gorm.Model
Number string
UserID uint
}

type Result struct {
Name string
Age uint8
Date time.Time
Total int
}

func main() {
// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
dsn := "root:123456@tcp(localhost:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
    	panic("failed to connect database")
    }

    var users []Result
    // db.Distinct("name", "age").Order("name, age desc").Find(&users)
    // 2022/12/10 09:35:37 /Users/chenhailong/code/github/go/gorm-demo/main.go:59 Error 1146: Table 'gorm_demo.results' doesn't exist
    // [10.064ms] [rows:0] SELECT DISTINCT name,age FROM `results` ORDER BY name, age desc

    db.Model(&User{}).Distinct("name", "age").Order("name, age desc").Find(&users)

    b, _ := json.Marshal(users)

    var out bytes.Buffer

    err = json.Indent(&out, b, "", "\t")
    if err != nil {
    	log.Fatalln(err)
    }

    out.WriteTo(os.Stdout)

}

```
