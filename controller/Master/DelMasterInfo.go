package Master

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DelMasterInfo(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()
		ID := ctx.PostForm("ID")
		// db.AutoMigrate(&UserList{})s
			id, _ := strconv.Atoi(ID)
			var M Table.MasterInfo
			var p []Table.Pet
			var PI []Table.PetInfo
			var o []Table.Order
			db.Debug().Table("master_infos").Where("id = ?",id).Find(&M)
			db.Debug().Table("pets").Where("master_id = ?",M.ID).Find(&p)
			for i := 0; i < len(p); i++ { 
				fmt.Println("sdasdasdasPI.ID",p[i].ID,"12154512")
				db.Debug().Table("pet_infos").Where("pet_id = ?",p[i].ID).Delete(&PI)
				db.Debug().Table("orders").Model(&o).Where("id = ?",p[i].OrderId).
				Updates(Table.Order{Status: 2})
			}
			fmt.Println("sdasdasdasP",p)
			fmt.Println("sdasdasdasLen(P)",len(p))
			fmt.Println("sdasdasdasM",M)
			db.Debug().Table("pets").Where("master_id = ?",M.ID).Delete(&p)
			db.Debug().Table("master_infos").Where("id = ?",id).Delete(&M)
			// var u1 []GoodsType

			// db.Debug().Table("goods_types").Find(&u1)
	
			ctx.JSON(http.StatusOK, gin.H{"code": 200,"msg":"删除成功"})
	
	}
}