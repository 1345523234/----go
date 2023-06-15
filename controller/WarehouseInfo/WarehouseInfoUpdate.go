package WarehouseInfo

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func WarehouseInfoUpdate(ctx *gin.Context){
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		id := ctx.PostForm("ID")
		WarehouseAddress := ctx.PostForm("WarehouseAddress")
	CompanyID := ctx.PostForm("CompanyID")
	CompanyName := ctx.PostForm("CompanyName")
	WarehouseName := ctx.PostForm("WarehouseName")
	WarehouseType := ctx.PostForm("WarehouseType")
	WarehouseTypeStock := ctx.PostForm("WarehouseTypeStock")

	companyID, _ := strconv.Atoi(CompanyID)
	warehouseTypeStock, _ := strconv.Atoi(WarehouseTypeStock)

	var u1 Table.WarehouseInfo
		db.Debug().Table("warehouse_infos").Where("id = ?",id).Find(&u1)
		if(u1.WarehouseTypeStockIng > warehouseTypeStock){
			ctx.JSON(http.StatusOK,gin.H{"code": 400 ,"msg":"库存量不能小于当前存量" })
			return
		}
		var u Table.WarehouseInfo
		db.Debug().Table("warehouse_infos").Model(&u).Where("id = ?",id).
		Updates(Table.WarehouseInfo{WarehouseAddress:WarehouseAddress,CompanyID:companyID,
			CompanyName:CompanyName,WarehouseName:WarehouseName,WarehouseType:WarehouseType,
			WarehouseTypeStock:warehouseTypeStock})
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"修改成功" })
	}
}