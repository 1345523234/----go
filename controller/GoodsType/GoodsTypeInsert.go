package GoodsType

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func GoodsTypeInsert(ctx *gin.Context) {
	name := ctx.PostForm("name")
	remark := ctx.PostForm("remark")
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()

		db.AutoMigrate(&Table.GoodsType{})

		timeUnix := time.Now().Unix()
		u1 := Table.GoodsType{Name:name,InsertTime:timeUnix,Remark:remark }
		db.Create(&u1)
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "添加成功", "data": u1})

		fmt.Println("连接成功！！！")
	}
}