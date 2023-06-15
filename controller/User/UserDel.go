package User

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UserDel(ctx *gin.Context){
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		id := ctx.PostForm("ID")
		var i User
		var u UserList
		db.Debug().Table("user_lists").Where("ID = ?", id).Find(&u)
		db.Debug().Table("users").Where("ID = ?", u.UserID).Delete(&i)
		db.Debug().Table("user_lists").Where("ID = ?", id).Delete(&u)
		fmt.Println("u--------",u)
	}
}