package GoodsDiscount

import (
	"fmt"
	"go-serve/controller/Table"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GoodsDiscountDelId(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		id := ctx.PostForm("id")
		var u Table.Goods
		db.Debug().Table("goods").Where("id = ?", id).Delete(&u)
		
	}
}