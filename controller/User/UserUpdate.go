package User

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type RoleList struct {
	ID         uint
	InsertTime int64
	Name       string
}

func UserUpdate(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		id := ctx.PostForm("ID")
		roleId := ctx.PostForm("RoleId")
		nickname := ctx.PostForm("Nickname")
		password := ctx.PostForm("Password")
		username := ctx.PostForm("Username")
		var re RoleList
		db.Debug().Table("role_lists").Where("ID = ?",roleId).Find(&re)
		var u UserList
		timeUnix:=time.Now().Unix()
		db.Table("user_lists").Model(&u).Where("ID = ?",id).
		Updates(UserList{Nickname:nickname,RoleName:re.Name,InsertTime:timeUnix,Username:username,Password:password})
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"修改成功" })
	}
}