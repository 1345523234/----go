package Role

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RoleOrRoleId(ctx *gin.Context) {

	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		id := ctx.PostForm("ID")
		var u []Table.UserList
		db.Debug().Table("user_lists").Where("role_id = ?",id).Find(&u)
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"查询成功","data":u })
		fmt.Println("连接成功！！！")
}
}