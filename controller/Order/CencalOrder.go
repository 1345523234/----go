package Order

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CencalOrder(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		
		id := ctx.PostForm("id")
	
		var u Table.Order
		db.Debug().Table("orders").Where("id = ?",id).Find(&u)
		var goods Table.Goods
		fmt.Println("连接失败",u.GoodsID)
		db.Debug().Table("goods").Where("id = ?",u.GoodsID).Find(&goods)
		var wof Table.WarehouseInfo
		db.Debug().Table("warehouse_infos").Where("id = ?",goods.WarehouseInfoID).Find(&wof)
		var wofs Table.WarehouseInfo
			db.Debug().Table("warehouse_infos").Model(&wofs).Where("id = ?", wof.ID).
			Updates(Table.WarehouseInfo{WarehouseTypeStockIng:wof.WarehouseTypeStockIng + 1})
			var g Table.Goods
			db.Debug().Table("goods").Model(&g).Where("id = ?", goods.ID).
			Updates(Table.Goods{Count:goods.Count + 1})
		db.Debug().Table("orders").Model(&u).Where("id = ?",id).
		Updates(Table.Order{Status: 2})
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"修改成功","data":u })
	}
}