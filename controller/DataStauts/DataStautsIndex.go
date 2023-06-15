package DataStauts

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DataStautsIndex(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		var order [] Table.Order
		db.Debug().Table("orders").Find(&order)
		var pet [] Table.Pet
		db.Debug().Table("pets").Find(&pet)
		var uesr [] Table.UserList
		db.Debug().Table("user_lists").Find(&uesr)
		var WarehouseInfo [] Table.WarehouseInfo
		db.Debug().Table("warehouse_infos").Find(&WarehouseInfo)
		var companie [] Table.Company
		db.Debug().Table("companies").Find(&companie)
		var data Table.DataStauts
		data = Table.DataStauts{OrderLen:len(order),PetLen:len(pet),UesrLen:len(uesr),
			WarehouseInfoLen:len(WarehouseInfo),CompanieLen:len(companie)}
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"查询成功","order":order,"data":data})
	}
}