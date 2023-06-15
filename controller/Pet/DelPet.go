package Pet

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DelPet(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		id := ctx.PostForm("ID")
		var u Table.Pet
		db.Debug().Table("pets").Where("id = ?", id).Find(&u)
		var o Table.Order
		db.Debug().Table("orders").Model(&o).Where("id = ?",u.OrderId).
		Updates(Table.Order{Status:2})
		var p []Table.Pet
		db.Debug().Table("pets").Where("master_id = ?", u.MasterID).Find(&p)
		var pi []Table.PetInfo
		db.Debug().Table("pet_infos").Where("pet_id = ?", id).Delete(&pi)
		fmt.Println("ssssssssssmmmmmms",len(p))
		if len(p) == 1{
			
			var m Table.MasterInfo
			db.Debug().Table("master_infos").Where("id = ?", u.MasterID).Delete(&m)
			fmt.Println("sssssssssssss",m)
		}
		db.Debug().Table("pets").Where("id = ?", id).Delete(&u)
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"删除成功","data":u})
	}
}