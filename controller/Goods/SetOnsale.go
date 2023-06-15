package Goods

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetOnsale(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		id := ctx.PostForm("ID")
		isOnSale := ctx.PostForm("IsOnSale")
		onsale, _ := strconv.Atoi(isOnSale)
		// var u Goods
		fmt.Println(id,"id++++++++++++++")
		fmt.Println(isOnSale,"sds--------------")
		db.Debug().Table("goods").Where("ID = ?",id).Update("is_on_sale",onsale)
	
	}
}