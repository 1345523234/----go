package Master

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UpdateMasterInfo(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()
		MasterName := ctx.PostForm("MasterName")
		IdCard := ctx.PostForm("IdCard")
		// PetType := ctx.PostForm("PetType")
		Age := ctx.PostForm("Age")
		Sex := ctx.PostForm("Sex")
		id := ctx.PostForm("ID")
		age, _ := strconv.Atoi(Age)

		var u Table.MasterInfo
		

		db.Debug().Table("master_infos").Model(&u).Where("id = ?", id).
			Updates(Table.MasterInfo{MasterName: MasterName, IdCard: IdCard,
			Sex:Sex,Age: age,
			})
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "修改", "data": u})
	}
}