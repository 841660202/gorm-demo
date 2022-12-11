```go

func main()  {
  dsn := "root:123456@tcp(localhost:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

  type APIUser struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	var users []User




  // 6、Order排序
	// db.Debug().Order("age desc, name").Find(&users)
  // SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY age desc, name


  // 多个 order
  // db.Debug().Order("age desc").Order("name").Find(&users)
  //  SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY age desc,name


  // todo: 语句

  // 7、Limit & Offset

  // db.Limit(3).Find(&users)
  // SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 3

	// db.Debug().Offset(1).Limit(1).Find(&users)
  //  SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 1 OFFSET 1

  // 这里不能使用(这应该是gorm的bug)
  // var apiUsers []APIUser
	// db.Debug().Offset(1).Limit(1).Find(&apiUsers)
  // SELECT * FROM `api_users` LIMIT 1 OFFSET 1

  // 增加 db.Model(&User{}) 就可以了
	// db.Model(&User{}).Debug().Offset(1).Limit(1).Find(&apiUsers)

	// var users2 []User
	// db.Debug().Limit(10).Find(&users).Limit(-1).Find(&users2)

	// fmt.Println(len(users))
	// fmt.Println(len(users2))

  // db.Debug().Limit(10).Find(&users).Limit(-1).Find(&apiUsers)
  // SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 10

  // SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL

	// 高级版 格式化输出json数据
	// b, _ := json.Marshal(users)

	// var out bytes.Buffer

	// err = json.Indent(&out, b, "", "\t")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// out.WriteTo(os.Stdout)



  type Result struct {
		Date  time.Time
		Total int
	}

// 8、Group & Having



// var result Result
// db.Debug().Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").Find(&result)
// SELECT name, sum(age) as total FROM `users` WHERE name LIKE 'group%' AND `users`.`deleted_at` IS NULL GROUP BY `name`



// db.Debug().Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&result)
// 大概是结果中没有主键id了，再按照这个排序，出错
// First、Last 方法会根据主键查找到第一个、最后一个记录， 它仅在通过结构体 struct 或提供 model 值进行查询时才起作用
//Error 1055: Expression #1 of ORDER BY clause is not in GROUP BY clause and contains nonaggregated column 'gorm_demo.users.id' which is not functionally dependent on columns in GROUP BY clause; this is incompatible with sql_mode=only_full_group_by
// SELECT name, sum(age) as total FROM `users` WHERE name LIKE 'group%' AND `users`.`deleted_at` IS NULL GROUP BY `name` ORDER BY `users`.`id` LIMIT 1

// 对于GROUP BY聚合操作，如果在SELECT中的列，没有在GROUP BY中出现，那么这个SQL是不合法的，因为列不在GROUP BY从句中
//  参考链接：https://blog.csdn.net/Naylor_5/article/details/124927555



// 也可以改成这样的，增加id分组

// db.Debug().Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").Group("id").First(&result)

// [3.378ms] [rows:0] SELECT name, sum(age) as total FROM `users` WHERE name LIKE 'group%' AND `users`.`deleted_at` IS NULL GROUP BY `name`,`id` ORDER BY `users`.`id` LIMIT 1


// 	rows, _ := db.Debug().Table("users").Select("date(created_at) as date, sum(age) as total").Group("date(created_at)").Rows()


//  for rows.Next() {
// 		var s []Result
// 		err = rows.Scan(rows, &s)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Printf("found row containing %q", s)
// 	}
// 	rows.Close()



// sql: Scan error on column index 0, name "date": unsupported Scan, storing driver.Value type <nil> into type *sql.Rows
// exit status 1


}

```
