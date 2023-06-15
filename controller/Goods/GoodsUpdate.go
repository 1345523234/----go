package Goods

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GoodsUpdate(ctx *gin.Context){
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		id := ctx.PostForm("ID")
		count1 := ctx.PostForm("Count")
	description := ctx.PostForm("Description")
	goodsTypeId1 := ctx.PostForm("GoodsTypeId")
	isOnSale1 := ctx.PostForm("IsOnSale")
	logo := ctx.PostForm("Logo")
	name := ctx.PostForm("Name")
	price1 := ctx.PostForm("Price")
	// pics := ctx.PostForm("pics")
	remark := ctx.PostForm("Remark")
	zheKou1 := ctx.PostForm("ZheKou")
	timeUnix:=time.Now().Unix()
		count, _ := strconv.Atoi(count1)
		goodsTypeId, _ := strconv.Atoi(goodsTypeId1)
		isOnSale, _ := strconv.Atoi(isOnSale1)
		zheKou, _ := strconv.Atoi(zheKou1)
		price, _ := strconv.Atoi(price1)
		var u Table.Goods
		db.Debug().Table("goods").Model(&u).Where("ID = ?",id).
		Updates(Table.Goods{
			Count:count,Description:description,GoodsTypeId:goodsTypeId,
			InsertTime: timeUnix,IsOnSale:isOnSale,Logo: logo,Name:name,
			Remark:remark,ZheKou:zheKou,Price:price ,
		})
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"修改成功" })
	}
}