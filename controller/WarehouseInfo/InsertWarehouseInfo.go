package WarehouseInfo

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InsertWarehouseInfo(ctx *gin.Context) {
	WarehouseAddress := ctx.PostForm("WarehouseAddress")
	CompanyID := ctx.PostForm("CompanyID")
	CompanyName := ctx.PostForm("CompanyName")
	WarehouseName := ctx.PostForm("WarehouseName")
	WarehouseType := ctx.PostForm("WarehouseType")
	WarehouseTypeStock := ctx.PostForm("WarehouseTypeStock")
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()
		companyID, _ := strconv.Atoi(CompanyID)
		warehouseTypeStock, _ := strconv.Atoi(WarehouseTypeStock)
		db.AutoMigrate(&Table.WarehouseInfo{})
		u1 := Table.WarehouseInfo{WarehouseAddress:WarehouseAddress,CompanyID:companyID,
			CompanyName:CompanyName,WarehouseName:WarehouseName,WarehouseType:WarehouseType,
			WarehouseTypeStock:warehouseTypeStock,WarehouseTypeStockIng:0}
		db.Create(&u1)
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "添加成功", "data": u1})

		fmt.Println("连接成功！！！")
	}
}