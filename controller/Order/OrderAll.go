package Order

import (
	"encoding/json"
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


func OrderAll(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()
		pno := ctx.PostForm("Pno")
		pSize := ctx.PostForm("Psize")
		Phone := ctx.PostForm("phone")
		Status := ctx.PostForm("status")
		// db.AutoMigrate(&UserList{})s
		if pno != "" && pSize != "" {
			currentPage, _ := strconv.Atoi(pno)
			pageSize, _ := strconv.Atoi(pSize)
			phone, _ := strconv.ParseInt(Phone,10,64)
			status, _ := strconv.Atoi(Status)
			fmt.Println(phone,"ssss")
			fmt.Println(status,"sssddsad")
			var u2 []Table.Order
			db.Debug().Limit(pageSize).Order("id desc").Offset((currentPage - 1) * pageSize).Find(&u2)
			//查询所有数据
			var u []Table.Order
			db.Debug().Find(&u)
			var p Table.Page
			
			var u1 []Table.GoodsType

			db.Debug().Table("goods_types").Order("id desc").Find(&u1)
			if Phone != "" && Status != ""{
				db.Debug().Where("phone = ? AND status =  ?",phone,status).Order("id desc").Find(&u)
				db.Debug().Where("phone = ? AND status =?",phone,status).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)
			}else if Phone != "" {
				db.Debug().Where("phone =  ?", phone).Order("id desc").Find(&u)
				db.Debug().Where("phone =  ?", phone).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)				
			}else if Status != ""{
				db.Debug().Where("status =  ?", status).Order("id desc").Find(&u)
				db.Debug().Where("status =  ?", status).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)				
			}
			p = Table.Page{
				PCount:        len(u) / pageSize,
				Pno:           currentPage,
				Psize:         pageSize,
				TotalElements: len(u)}
			page, _ := json.Marshal(p)
			ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": u2,"GoodsTypes":u1,"page": string(page)})
		}	
	}
}