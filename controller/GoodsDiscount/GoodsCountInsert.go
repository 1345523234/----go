package GoodsDiscount

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GoodsCountInsert(ctx *gin.Context) {
	
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
	
		description := ctx.PostForm("Description")
	
	name := ctx.PostForm("Name")
	price1 := ctx.PostForm("Price")
	// pics := ctx.PostForm("pics")
	remark := ctx.PostForm("Remark")
	zheKou1 := ctx.PostForm("ZheKou")
	timeUnix:= ctx.PostForm("InsertTime")
		
		zheKou, _ := strconv.Atoi(zheKou1)
		price, _ := strconv.Atoi(price1)
		layuot := "2006-01-02 15:04:05"
			t1,_ := time.Parse(layuot,timeUnix)
		
		db.AutoMigrate(&Table.Goods{})
		// ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"data":re })

		//创建表 自动迁移 (把结构体和数据表进行对应)	
		u1 := Table.Goods{
			Description:description,
			InsertTime: t1.Unix(),Name:name,
			Remark:remark,ZheKou:zheKou,Price:price ,
		}
		db.Create(&u1)
	
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"添加成功" ,"data":u1})
		// Struct
		// db.Debug().First(&login)
		// fmt.Println(login)
		fmt.Println("连接成功！！！")
}
}