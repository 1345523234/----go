package Master

import (
	"encoding/json"
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// type MasterInfo struct {
// 	ID int
// 	MasterName string
// 	IdCard string
// 	PetName string
// 	PetType string
// 	PetId int
// 	OrderId int
// 	PetVarieties string
// 	Price int
// 	Sex string
// 	Age int
//  }

func MasterAll(ctx *gin.Context){
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()
		pno := ctx.PostForm("Pno")
		pSize := ctx.PostForm("Psize")
		MasterName := ctx.PostForm("masterName")
		// db.AutoMigrate(&UserList{})s
		if pno != "" && pSize != "" {
			currentPage, _ := strconv.Atoi(pno)
			pageSize, _ := strconv.Atoi(pSize)
			var u2 []Table.MasterInfo
			db.Debug().Limit(pageSize).Offset((currentPage - 1) * pageSize).Order("id desc").Find(&u2)
			//查询所有数据
			var u []Table.MasterInfo
			db.Debug().Order("id desc").Find(&u)
			var p Table.Page
			
			// var u1 []GoodsType
			// db.Debug().Table("goods_types").Find(&u1)
			if MasterName != ""{
				db.Debug().Where("master_name =  ?", MasterName).Order("id desc").Find(&u)
				db.Debug().Where("master_name =  ?", MasterName).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)				
			}
			p = Table.Page{
				PCount:        len(u) / pageSize,
				Pno:           currentPage,
				Psize:         pageSize,
				TotalElements: len(u)}
			page, _ := json.Marshal(p)
			ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": u2,"page": string(page)})
		}	
	}
}