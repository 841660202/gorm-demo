```Go
package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	spew "github.com/davecgh/go-spew/spew"
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


func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("User11 BeforeUpdate执行了")
	JsonPrint(u)

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

func JsonPrint(val any) {

	b, _ := json.Marshal(val)

	var out bytes.Buffer

	err := json.Indent(&out, b, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	out.WriteTo(os.Stdout)
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(localhost:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var user User

	db.Debug().Select("id", "name").Model(&user).First(&user, "name = ?", "hali")
	// JsonPrint(user)

	spew.Println(user.ID)

	// user.Name = "hali"
	spew.Println("更新前")

  // 神坑
	// db.Debug().Model(&User{}).Where("id = ?", user.ID).Updates(user) // hook 获取属性全部是零值

  // ok 1
	// db.Debug().Where("id = ?", user.ID).Updates(user)


  // ok 2
  // 把有数据的user地址塞进去
  db.Debug().Model(&user).Where("id = ?", user.ID).Updates(user)

	spew.Println("更新后")

}

```
