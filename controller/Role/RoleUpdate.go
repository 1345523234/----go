package Role

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)



func UserUpdate(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		id := ctx.PostForm("ID")
		name := ctx.PostForm("Name")
		var u Table.RoleList
		timeUnix:=time.Now().Unix()
		db.Debug().Table("role_lists").Model(&u).Where("ID = ?",id).
		Updates(Table.RoleList{Name:name,InsertTime: timeUnix})
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"修改成功" })
	}
}