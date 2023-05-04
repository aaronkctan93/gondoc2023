package repository

import (
	"fmt"
	"gondoc2023/def"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:qwe123@tcp(127.0.0.1:30306)/dbtest?charset=utf8&parseTime=True&loc=Local"
var conn, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func GetUserData(c *gin.Context) {
	type GetUserIDStruct struct {
		UserID int64 `form:"user_id"`
	}
	var input GetUserIDStruct

	err := c.ShouldBindQuery(&input)
	if err != nil {
		fmt.Println("ShouldBindQuery error. ", err)
		return
	}

	var userData def.User
	err = conn.Where("user_id = ?", input.UserID).Debug().Find(&userData).Error
	if err != nil {
		fmt.Println("GetUserIDStruct error.", err)
		return
	}

	c.JSON(http.StatusOK, userData)
}

// POST
func CreateUser(c *gin.Context) {
	type CreateUser struct {
		Username string `json:"username"`
	}
	var createUser CreateUser

	// inputUsername := c.DefaultPostForm("username", "no data input")
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		fmt.Println("CreateUser error.", err)
		return
	}
	var u def.User

	u.Username = createUser.Username

	err = conn.Debug().Create(&u).Error
	if err != nil {
		fmt.Println("CreateUser error", err)
		return
	}

	c.String(http.StatusOK, "新增的 username 為: %s", createUser)

}

// Delete
func DeleteUser(c *gin.Context) {
	type DeleteUser struct {
		Username string `json:"username"`
	}
	var deleteUser DeleteUser

	err := c.ShouldBindJSON(&deleteUser)
	if err != nil {
		fmt.Println("DeteleUser error.", err)
		return
	}
	// var u User

	err = conn.Where("username=?", deleteUser.Username).Debug().Delete(&def.User{}).Error
	if err != nil {
		fmt.Println("Delete User error.", err)
		return
	}
	c.String(http.StatusOK, "資料已移除成功！")
}

// Update
func UpdateUser(c *gin.Context) {
	type userData struct {
		// UserID      int64  `json:"user_id,string"`
		Username    string `json:"username"`
		NewUsername string `json:"new_username"`
	}
	var UserData userData
	err := c.ShouldBindJSON(&UserData)
	if err != nil {
		fmt.Println("UpdateUser error.", err)
		return
	}

	var user def.User

	fmt.Println(UserData)

	err = conn.Model(&user).Where("user_id = ?", UserData.Username).Debug().Update("username", UserData.NewUsername).Error
	if err != nil {
		fmt.Println("Update User error", err)
		return
	}
	c.String(http.StatusOK, "新增的 username : %s, 已新增成功！", UserData.NewUsername)

}
