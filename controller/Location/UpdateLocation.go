package Location

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UpdateLocation(ctx *gin.Context) {
	BeginTime := ctx.PostForm("BeginTime")
	EndTime := ctx.PostForm("EndTime")
	City := ctx.PostForm("City")
	District := ctx.PostForm("District")
	Lat := ctx.PostForm("Lat")
	Lng := ctx.PostForm("Lng")
	Nation := ctx.PostForm("Nation")
	Province := ctx.PostForm("Province")
	Street := ctx.PostForm("Street")
	StreetNumber := ctx.PostForm("StreetNumber")
	R := ctx.PostForm("R")
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		db.AutoMigrate(&Table.Location{})
	
		//创建表 自动迁移 (把结构体和数据表进行对应)
		
		lat,_ := strconv.ParseFloat(Lat,64)

		lng,_ := strconv.ParseFloat(Lng,64)


		r,_ := strconv.Atoi(R)
		var u Table.Location
			db.Debug().Table("locations").Model(&u).Where("id = ?",1).
		Updates(Table.Location{BeginTime:BeginTime,EndTime:EndTime,City:City,District:District,
			Lat:lat,Lng:lng,Nation:Nation,Province:Province,Street:Street,StreetNumber:StreetNumber,
			R:r})
	

		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"修改成功" })
		// Struct
		// db.Debug().First(&login)
		// fmt.Println(login)
		fmt.Println("连接成功！！！")
}
}