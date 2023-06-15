package GoodsType

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


func GoodsTypeUpdate(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		name := ctx.PostForm("Name")
	    remark := ctx.PostForm("Remark")
		id := ctx.PostForm("ID")
		timeUnix:=time.Now().Unix()
		var u Table.GoodsType
		db.Debug().Table("goods_types").Model(&u).Where("id = ?",id).
		Updates(Table.GoodsType{InsertTime:timeUnix,Name: name,Remark: remark})
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"修改成功" })
	}
}