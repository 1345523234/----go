package Master

import (
	"fmt"
	"net/http"
	"strconv"

	"go-serve/controller/Table"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func MasterInfo(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()
		ID := ctx.PostForm("ID")
		// db.AutoMigrate(&UserList{})s
			id, _ := strconv.Atoi(ID)
			var P Table.MasterInfo
			db.Debug().Table("master_infos").Where("id = ?",id).Find(&P)
			var O []Table.Order
			db.Debug().Table("orders").Where("master_id = ?",id).Find(&O)
			// var u1 []GoodsType

			// db.Debug().Table("goods_types").Find(&u1)
	
			ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": P,"order":O})
	
	}
}