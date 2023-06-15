package PetType

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func PetTypeInsert(ctx *gin.Context) {
	PetName := ctx.PostForm("petName")
	remark := ctx.PostForm("remark")
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()

		db.AutoMigrate(&Table.PetType{})

		timeUnix := time.Now().Unix()
		u1 := Table.PetType{PetName:PetName,InsertTime:timeUnix,Remark:remark }
		db.Create(&u1)
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "添加成功", "data": u1})

		fmt.Println("连接成功！！！")
	}
}