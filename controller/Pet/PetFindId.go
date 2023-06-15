package Pet

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func PetFindId(ctx *gin.Context){ 
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		id := ctx.PostForm("id")
		var u Table.Pet
		db.Debug().Table("pets").Where("id = ?",id).Find(&u)
		var PI Table.PetInfo
		db.Debug().Table("pet_infos").Where("pet_id = ?",id).Find(&PI)
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"查询成功","data":u,"petInfo":PI })
	}
}