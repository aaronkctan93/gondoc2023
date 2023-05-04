package main

import (
	"fmt"
	"gondoc2023/def"
	"gondoc2023/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func main() {
// 	// mysql.Init()

// 	// 連線Redis成功
// 	Redis.NewClinet()

// 	// user := Mysql.User{}
// 	// fmt.Println(mysql.GetDB())
// 	/*
// 		http请求方法
// 		url路径
// 		控制器函数
// 	*/

// 	r := gin.Default()

// 	r.GET("/ping", GetPing)

// 	err := r.Run(":8082")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("test bottom")
// }

// func GetPing(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "pongs",
// 	})
// }

// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
var dsn = "root:qwe123@tcp(127.0.0.1:30306)/dbtest?charset=utf8&parseTime=True&loc=Local"

// 連線Mysql
var conn, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func main() {
	u := def.User{}

	r := gin.Default()

	r.GET("/user", repository.GetUserData)
	r.POST("/create", repository.CreateUser)
	r.DELETE("/delete", repository.DeleteUser)
	r.PUT("/update", repository.UpdateUser)

	err := r.Run(":8082")
	if err != nil {
		panic(err)
	}

	// err := conn.Where("user_id = ?", 10).
	// 	Debug().
	// 	Find(&u).Error
	// if err != nil {
	// 	fmt.Println("a s d")
	// 	return
	// }

	// update username: steve -> John
	// err = conn.Model(u).Where("user_id = ?", 10).Update("username", "John").Error
	// if err != nil {·
	// 	// panic(err)
	// 	e := fmt.Errorf("asd, err: %v", err)
	// 	fmt.Printf("e: %v\n", e)
	// 	return
	// }

	// fmt.Println("10號的使用者為：", u)

	// Delete username 是 John 的資料
	// err := conn.Where("username=?", "Chris").Debug().Delete(&User{}).Error
	// if err != nil {
	// 	fmt.Println("a s d")
	// 	return
	// }

	// err = conn.Find(&u).Error
	// if err != nil {
	// 	fmt.Errorf("讀取發生錯誤")
	// 	return
	// }

	// u1 := User{Username: "Steve"}
	// err = conn.Create(&u1).Error
	// if err != nil {
	// 	fmt.Errorf("資料新增發生錯字")
	// 	return
	// }

	fmt.Println("u:", u)

	fmt.Println()

}
