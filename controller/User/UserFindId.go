package User

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UserFindId(ctx *gin.Context){
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		id := ctx.PostForm("id")
		var u UserList
		db.Table("user_lists").Where("ID = ?",id).Find(&u)
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"查询成功","data":u })
	}
}