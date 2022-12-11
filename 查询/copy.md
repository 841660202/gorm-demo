package main

import (
"bytes"
"database/sql"
"encoding/json"
"fmt"
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

func main() {
// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
dsn := "root:123456@tcp(localhost:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
    	panic("failed to connect database")
    }

    // db.Debug().AutoMigrate(&User{})
    // db.Debug().AutoMigrate(&CreditCard{})

    // 插入一条数据
    // birtime := time.Now()
    // user := User{Name: "linzy", Age: 23, Birthday: &birtime}

    // result := db.Create(&user) // 通过数据的指针来创建

    // fmt.Print(result)

    // 批量插入数据
    // var users = []User{{Name: "linzy1"}, {Name: "linzy2"}, {Name: "linzy3"}}
    // db.Create(&users)

    // for _, user := range users {
    // 	// 1,2,3
    // 	fmt.Println(user.ID)
    // }

    // map 单条数据
    // db.Model(&User{}).Create(map[string]interface{}{
    // 	"Name": "map_linzy", "Age": 23,
    // })

    //map 创建多条记录
    // db.Model(&User{}).Create([]map[string]interface{}{
    // 	{"Name": "map_linzy_1", "Age": 23},
    // 	{"Name": "map_linzy_2", "Age": 66},
    // 	{"Name": "map_linzy_3", "Age": 88},
    // })

    // 创建关联数据
    // db.Create(&User{
    // 	Name:       "linzy",
    // 	CreditCard: CreditCard{Number: "123456789"},
    // })

    // 跳过CreditCard关联
    // user := User{
    // 	Name:       "omit_linzy",
    // 	CreditCard: CreditCard{Number: "omit_123456789"},
    // }

    // db.Omit("CreditCard").Create(&user)

    // 跳过所有关联
    // user := User{
    // 	Name:       "omit_linzy_1",
    // 	CreditCard: CreditCard{Number: "omit_123456789"},
    // }
    // db.Omit(clause.Associations).Create(&user)

    // 对于1对1的表，即使多次插入数据，形成一堆多，多对多的情况，在表中也是不存在的，有意思吧
    // mysql> select * from credit_cards;
    // +----+-------------------------+-------------------------+------------+-----------+---------+
    // | id | created_at              | updated_at              | deleted_at | number    | user_id |
    // +----+-------------------------+-------------------------+------------+-----------+---------+
    // |  1 | 2022-12-08 22:18:41.953 | 2022-12-08 22:18:41.953 | NULL       | 123456789 |      10 |
    // |  2 | 2022-12-08 22:20:49.231 | 2022-12-08 22:20:49.231 | NULL       | 123456789 |      11 |
    // +----+-------------------------+-------------------------+------------+-----------+---------+
    // 2 rows in set (0.00 sec)

    // mysql> select * from users;
    // +----+-------------------------+-------------------------+------------+-------------+-------+------+-------------------------+---------------+--------------+
    // | id | created_at              | updated_at              | deleted_at | name        | email | age  | birthday                | member_number | activated_at |
    // +----+-------------------------+-------------------------+------------+-------------+-------+------+-------------------------+---------------+--------------+
    // |  1 | 2022-12-08 22:11:55.228 | 2022-12-08 22:11:55.228 | NULL       | linzy       | NULL  |   23 | 2022-12-08 22:11:55.216 | NULL          | NULL         |
    // |  2 | 2022-12-08 22:14:04.208 | 2022-12-08 22:14:04.208 | NULL       | linzy1      | NULL  |    0 | NULL                    | NULL          | NULL         |
    // |  3 | 2022-12-08 22:14:04.208 | 2022-12-08 22:14:04.208 | NULL       | linzy2      | NULL  |    0 | NULL                    | NULL          | NULL         |
    // |  4 | 2022-12-08 22:14:04.208 | 2022-12-08 22:14:04.208 | NULL       | linzy3      | NULL  |    0 | NULL                    | NULL          | NULL         |
    // |  5 | NULL                    | NULL                    | NULL       | map_linzy   | NULL  |   23 | NULL                    | NULL          | NULL         |
    // |  6 | NULL                    | NULL                    | NULL       | map_linzy_1 | NULL  |   23 | NULL                    | NULL          | NULL         |
    // |  7 | NULL                    | NULL                    | NULL       | map_linzy_2 | NULL  |   66 | NULL                    | NULL          | NULL         |
    // |  8 | NULL                    | NULL                    | NULL       | map_linzy_3 | NULL  |   88 | NULL                    | NULL          | NULL         |
    // | 10 | 2022-12-08 22:18:41.952 | 2022-12-08 22:18:41.952 | NULL       | linzy       | NULL  |    0 | NULL                    | NULL          | NULL         |
    // | 11 | 2022-12-08 22:20:49.230 | 2022-12-08 22:20:49.230 | NULL       | linzy       | NULL  |    0 | NULL                    | NULL          | NULL         |
    // | 12 | 2022-12-08 22:20:49.244 | 2022-12-08 22:20:49.244 | NULL       | omit_linzy  | NULL  |    0 | NULL                    | NULL          | NULL         |
    // | 13 | 2022-12-08 22:21:26.668 | 2022-12-08 22:21:26.668 | NULL       | omit_linzy  | NULL  |    0 | NULL                    | NULL          | NULL         |
    // +----+-------------------------+-------------------------+------------+-------------+-------+------+-------------------------+---------------+--------------+
    // 12 rows in set (0.00 sec)

    // mysql> select * from users;
    // +----+-------------------------+-------------------------+------------+--------------+-------+------+-------------------------+---------------+--------------+
    // | id | created_at              | updated_at              | deleted_at | name         | email | age  | birthday                | member_number | activated_at |
    // +----+-------------------------+-------------------------+------------+--------------+-------+------+-------------------------+---------------+--------------+
    // |  1 | 2022-12-08 22:11:55.228 | 2022-12-08 22:11:55.228 | NULL       | linzy        | NULL  |   23 | 2022-12-08 22:11:55.216 | NULL          | NULL         |
    // |  2 | 2022-12-08 22:14:04.208 | 2022-12-08 22:14:04.208 | NULL       | linzy1       | NULL  |    0 | NULL                    | NULL          | NULL         |
    // |  3 | 2022-12-08 22:14:04.208 | 2022-12-08 22:14:04.208 | NULL       | linzy2       | NULL  |    0 | NULL                    | NULL          | NULL         |
    // |  4 | 2022-12-08 22:14:04.208 | 2022-12-08 22:14:04.208 | NULL       | linzy3       | NULL  |    0 | NULL                    | NULL          | NULL         |
    // |  5 | NULL                    | NULL                    | NULL       | map_linzy    | NULL  |   23 | NULL                    | NULL          | NULL         |
    // |  6 | NULL                    | NULL                    | NULL       | map_linzy_1  | NULL  |   23 | NULL                    | NULL          | NULL         |
    // |  7 | NULL                    | NULL                    | NULL       | map_linzy_2  | NULL  |   66 | NULL                    | NULL          | NULL         |
    // |  8 | NULL                    | NULL                    | NULL       | map_linzy_3  | NULL  |   88 | NULL                    | NULL          | NULL         |
    // | 10 | 2022-12-08 22:18:41.952 | 2022-12-08 22:18:41.952 | NULL       | linzy        | NULL  |    0 | NULL                    | NULL          | NULL         |
    // | 11 | 2022-12-08 22:20:49.230 | 2022-12-08 22:20:49.230 | NULL       | linzy        | NULL  |    0 | NULL                    | NULL          | NULL         |
    // | 12 | 2022-12-08 22:20:49.244 | 2022-12-08 22:20:49.244 | NULL       | omit_linzy   | NULL  |    0 | NULL                    | NULL          | NULL         |
    // | 13 | 2022-12-08 22:21:26.668 | 2022-12-08 22:21:26.668 | NULL       | omit_linzy   | NULL  |    0 | NULL                    | NULL          | NULL         |
    // | 14 | 2022-12-08 22:24:16.415 | 2022-12-08 22:24:16.415 | NULL       | omit_linzy_1 | NULL  |    0 | NULL                    | NULL          | NULL         |
    // +----+-------------------------+-------------------------+------------+--------------+-------+------+-------------------------+---------------+--------------+
    // 13 rows in set (0.00 sec)

    // mysql> select * from credit_cards;
    // +----+-------------------------+-------------------------+------------+-----------+---------+
    // | id | created_at              | updated_at              | deleted_at | number    | user_id |
    // +----+-------------------------+-------------------------+------------+-----------+---------+
    // |  1 | 2022-12-08 22:18:41.953 | 2022-12-08 22:18:41.953 | NULL       | 123456789 |      10 |
    // |  2 | 2022-12-08 22:20:49.231 | 2022-12-08 22:20:49.231 | NULL       | 123456789 |      11 |
    // +----+-------------------------+-------------------------+------------+-----------+---------+
    // 2 rows in set (0.00 sec)

    // fullname测试
    // user := User{Name: "linzy", Age: 23, FirstName: "chen", LastName: "haotian"}

    // db.Create(&user) // 通过数据的指针来创建
    // var user User
    // db.Last(&user)
    // a, _ := json.Marshal(user)
    // fmt.Print(string(a))

    var users []User

    // db.Where("name <> ?", "linzy").Find(&users)
    // row := db.Debug().Where("name LIKE ?", "%in%").Find(&users)
    // a, _ := json.Marshal(users)
    // fmt.Println(string(a))

    // row := db.Debug().Where("updated_at < ?", time.Now().AddDate(0, 0, -1)).Find(&users)
    // fmt.Println(len(users))
    // fmt.Println(row.RowsAffected)

    // last := time.Now().AddDate(0, 0, -1)
    // today := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
    // row := db.Debug().Where("created_at BETWEEN ? AND ?", last, today).Find(&users)
    // a, _ := json.Marshal(users)
    // fmt.Println(string(a))
    // fmt.Println(row.RowsAffected)

    // db.Where([]int64{1}).Find(&users)

    // a, _ := json.Marshal(users)
    // fmt.Println(string(a))

    // 结构体零值 不会作为查询条件
    // db.Where(&User{Name: "linzy", Age: 0}).Find(&users)
    // a, _ := json.Marshal(users)
    // fmt.Println(string(a))

    // db.Where(map[string]interface{}{"Name": "linzy", "Age": 0}).Find(&users)

    // a, _ := json.Marshal(users)
    // fmt.Println(string(a))

    // fmt.Println(len(users))

    // db.Not([]int64{1, 2, 3}).First(&users)
    // a, _ := json.Marshal(users)
    // fmt.Println(string(a))

    // fmt.Println(len(users))

    type APIUser struct {
    	ID   uint   `json:"id"`
    	Name string `json:"name"`
    }
    // var apiUsers []APIUser
    // 查询时会自动选择 `id`, `name` 字段
    // db.Model(&User{}).Limit(10).Find(&apiUsers)
    // find用数组接，返回数组；用对象接，返回对象
    // a, _ := json.Marshal(apiUsers)
    // fmt.Println(string(a))
    fmt.Println(users)
    // 低级版 手动格式化输出
    // fmt.Println("[")
    // for _, v := range apiUsers {
    // 	a, _ := json.Marshal(v)
    // 	fmt.Println("  " + string(a))
    // }
    // fmt.Println("]")

    // db.Model(&User{}).Debug().Offset(1).Limit(1).Find(&APIUser{})
    // var users2 []APIUser

    // db.Debug().Limit(10).Find(&users).Limit(-1).Find(&apiUsers)

    // fmt.Println(len(users))
    // fmt.Println(len(apiUsers))

    // db.Debug().Order("age desc").Order("name").Find(&users)

    type Result struct {
    	Date  time.Time
    	Total int
    }

    // var result Result
    // db.Debug().Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name like ?", "linzy%").Find(&result)
    // rows, _ := db.Debug().Table("users").Select("date(created_at) as date, sum(age) as total").Group("date(created_at)").Rows()

    // 高级版 格式化输出json数据
    // b, _ := json.Marshal(rows)

    // var out bytes.Buffer

    // err = json.Indent(&out, b, "", "\t")
    // if err != nil {
    // 	log.Fatalln(err)
    // }

    // out.WriteTo(os.Stdout)
    // var results []Result
    // for rows.Next() {
    // 	// var user User
    // 	var user Result
    // 	// err = rows.Scan(rows, &s)
    // 	db.ScanRows(rows, &user)
    // 	if err != nil {
    // 		log.Fatal(err)
    // 	}
    // 	results = append(results, user)

    // 	// log.Printf("found row containing %v", user)

    // }
    // rows.Close()

    b, _ := json.Marshal(results)

    var out bytes.Buffer

    err = json.Indent(&out, b, "", "\t")
    if err != nil {
    	log.Fatalln(err)
    }

    out.WriteTo(os.Stdout)

}
