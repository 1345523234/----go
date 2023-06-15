package Location

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"


	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func LocationAll(ctx *gin.Context){
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()
		
		// db.AutoMigrate(&UserList{})s
		
			
		var u1 []Table.Location
		db.Debug().Table("locations").Find(&u1)
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": u1})
		}
		
	}
