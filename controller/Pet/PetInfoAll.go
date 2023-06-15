package Pet

import (
	"fmt"
	"net/http"
	"strconv"
	"go-serve/controller/Table"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


func PetInfoAll(ctx *gin.Context){
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()
		ID := ctx.PostForm("ID")
		
		// db.AutoMigrate(&UserList{})s
			id, _ := strconv.Atoi(ID)
		
			var PI Table.PetInfo
			db.Debug().Table("pet_infos").Where("pet_id = ?",id).Find(&PI)
			var P Table.Pet
			db.Debug().Table("pets").Where("id = ?",id).Find(&P)
			var O Table.Order
			db.Debug().Table("orders").Where("id = ?",P.OrderId).Find(&O)
			// var u1 []GoodsType

			// db.Debug().Table("goods_types").Find(&u1)
	
			ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": PI,"order":O})
	
	}
}