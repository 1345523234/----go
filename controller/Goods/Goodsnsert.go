package Goods

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)




func GoodsInsert(ctx *gin.Context) {
	count1 := ctx.PostForm("count")
	description := ctx.PostForm("description")
	goodsTypeId1 := ctx.PostForm("goodsTypeId")
	isOnSale1 := ctx.PostForm("isOnSale")
	logo := ctx.PostForm("logo")
	name := ctx.PostForm("name")
	// pics := ctx.PostForm("pics")
	remark := ctx.PostForm("remark")
	zheKou1 := ctx.PostForm("zheKou")
	price1 := ctx.PostForm("price")
	WarehouseInfo := ctx.PostForm("warehouseInfo")
	WarehouseInfoID := ctx.PostForm("warehouseInfoID")
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		// var gt Goods
		// db.Debug().Table("role_lists").Where("ID = ?",roleId).Find(&gt)
		// fmt.Println(re.Name)
		db.AutoMigrate(&Table.Goods{})
		// ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"data":re })

		timeUnix:=time.Now().Unix()
		count, _ := strconv.Atoi(count1)
		goodsTypeId, _ := strconv.Atoi(goodsTypeId1)
		isOnSale, _ := strconv.Atoi(isOnSale1)
		zheKou, _ := strconv.Atoi(zheKou1)
		price, _ := strconv.Atoi(price1)
		warehouseInfoID, _ := strconv.Atoi(WarehouseInfoID)
		
		//创建表 自动迁移 (把结构体和数据表进行对应)	
		u1 := Table.Goods{
			Count:count,Description:description,GoodsTypeId:goodsTypeId,
			InsertTime: timeUnix,IsOnSale:isOnSale,Logo: logo,Name:name,
			Remark:remark,ZheKou:zheKou,Price: price,WarehouseInfoID:warehouseInfoID,
			WarehouseInfo:WarehouseInfo}
		db.Create(&u1)
		var u Table.WarehouseInfo
		db.Debug().Table("warehouse_infos").Where("id = ?",warehouseInfoID).Find(&u)
		if u.WarehouseTypeStockIng + count > u.WarehouseTypeStock{
			ctx.JSON(http.StatusOK,gin.H{"code": 400 ,"msg":"所选仓库库存量不足" })
			return
		}
		var u2 Table.WarehouseInfo
		db.Debug().Table("warehouse_infos").Model(&u2).Where("id = ?",warehouseInfoID).
		Updates(Table.WarehouseInfo{WarehouseTypeStockIng:u.WarehouseTypeStockIng + count})
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"添加成功" ,"data":u1})
		
		// Struct
		// db.Debug().First(&login)
		// fmt.Println(login)
		fmt.Println("连接成功！！！")
}
}