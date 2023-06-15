package Order

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


 
func OrderInsert(ctx *gin.Context) {
	Name := ctx.PostForm("name")
	Phone := ctx.PostForm("phone")
	SinglePrice := ctx.PostForm("singlePrice")
	Status := ctx.PostForm("status")
	GoodTypeID := ctx.PostForm("goodTypeID")
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()

		db.AutoMigrate(&Table.Order{})
		phone, _ := strconv.ParseInt(Phone,10,64)
		singlePrice, _ := strconv.Atoi(SinglePrice)
		status, _ := strconv.Atoi(Status)
		goodTypeID, _ := strconv.Atoi(GoodTypeID)
		
		timeUnix := time.Now().Unix()
		fmt.Println("cs"+strconv.FormatInt(timeUnix, 10) )
		u1 := Table.Order{
			Phone:phone,Name:Name,
			SinglePrice:singlePrice, 
			PostCode:"cs"+strconv.FormatInt(timeUnix, 10) ,
			InsertTime: timeUnix,GoodTypeID:goodTypeID,
			Status:status}
		db.Create(&u1)
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "添加成功", "data": u1})

		fmt.Println("连接成功！！！")
	}
}