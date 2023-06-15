package Role

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func RoleInsert(ctx *gin.Context) {
	name := ctx.PostForm("Name")
	fmt.Println(name)
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		db.AutoMigrate(&Table.RoleList{})
		timeUnix:=time.Now().Unix()
		//创建表 自动迁移 (把结构体和数据表进行对应)
		u1 := Table.RoleList{InsertTime:timeUnix,Name: name }
		db.Create(&u1)

		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"添加成功" })
		// Struct
		// db.Debug().First(&login)
		// fmt.Println(login)
		fmt.Println("连接成功！！！")
}
}