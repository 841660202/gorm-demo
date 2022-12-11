```Go
    // var dels []model.Del
    // db.Debug().Select("id").Where("isNull(name)").Find(&dels) //
    // // [0.283ms] [rows:4] SELECT `id` FROM `dels` WHERE isNull(name) AND `dels`.`deleted_at` IS NULL

    // for i, v := range dels {
    // 	// v.Name = fmt.Sprintf("name:%d", v.ID) // 直接改不能成功
    // 	dels[i].Name = fmt.Sprintf("name:%d", v.ID)

    // 	// fmt.Println(fmt.Sprintf("name:%d", v.ID))
    // }

```
