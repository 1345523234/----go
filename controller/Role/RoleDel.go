package Role

import (
	"fmt"
	"go-serve/controller/Table"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RoleDel(ctx *gin.Context){
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		id := ctx.PostForm("id")
		var u Table.RoleList
		db.Debug().Table("role_lists").Where("ID = ?", id).Delete(&u)
		fmt.Println("u--------",u)
	}
}