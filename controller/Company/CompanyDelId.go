package Company

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CompanyDelId(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()
		id := ctx.PostForm("id")
		var u Table.Company
		
		db.Debug().Table("companies").Where("id = ?", id).Delete(&u)
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"删除成功" })
	}
}