package Order

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CompleteOrder(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		
		id := ctx.PostForm("id")
	
		var u Table.Order
		db.Debug().Table("order").Where("id = ?",id).Find(&u)
		db.Debug().Table("orders").Model(&u).Where("id = ?",id).
		Updates(Table.Order{Status: 3})
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"修改成功","data":u })
	}
}