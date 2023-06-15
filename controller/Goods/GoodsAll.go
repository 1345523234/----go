package Goods

import (
	"encoding/json"
	"fmt"
	"go-serve/controller/Table"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Page struct {
	Pno           int
	PCount        int
	Psize         int
	TotalElements int
}

func GoodsAll(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()
		pno := ctx.PostForm("Pno")
		pSize := ctx.PostForm("Psize")
		GoodsTypeId := ctx.PostForm("GoodsTypeId")
		IsOnSale := ctx.PostForm("IsOnSale")
		Name := ctx.PostForm("Name")
		// db.AutoMigrate(&UserList{})
		if pno != "" && pSize != "" {
			currentPage, _ := strconv.Atoi(pno)
			pageSize, _ := strconv.Atoi(pSize)

			var u []Table.Goods
			db.Debug().Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u)
			//查询所有数据
			var u1 []Table.Goods
			db.Debug().Find(&u1)
			var p Table.Page
			//分页
			var u2 []Table.Goods
			//所有
			var u3 []Table.Goods
			if GoodsTypeId!="" && Name != ""{
				db.Debug().Where("goods_type_id = ? AND name =  ?",GoodsTypeId,Name).Order("id desc").Find(&u3)
				db.Debug().Where("goods_type_id = ? AND name =?",GoodsTypeId,Name).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)
			}else if IsOnSale!="" && Name != "" {
				db.Debug().Where("is_on_sale = ? AND is_on_sale =  ?",Name,IsOnSale).Order("id desc").Find(&u3)
				db.Debug().Where("is_on_sale = ? AND is_on_sale =  ?",Name,IsOnSale).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)
			}else if GoodsTypeId != "" && IsOnSale!="" && Name != ""{
				db.Debug().Where("goods_type_id = ? AND is_on_sale AND name =  ?",GoodsTypeId,IsOnSale,Name).Order("id desc").Find(&u3)
				db.Debug().Where("goods_type_id = ? AND is_on_sale = ? AND name =?",GoodsTypeId,IsOnSale,Name).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)
			}else if GoodsTypeId != "" && IsOnSale!="" {
				db.Debug().Where("goods_type_id = ? AND is_on_sale =  ?",GoodsTypeId,IsOnSale).Order("id desc").Find(&u3)
				db.Debug().Where("goods_type_id = ? AND is_on_sale =  ?",GoodsTypeId,IsOnSale).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)
			}else if GoodsTypeId != "" {
				db.Debug().Where("goods_type_id = ?",GoodsTypeId).Order("id desc").Find(&u3)
				db.Debug().Where("goods_type_id = ?",GoodsTypeId).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)
			}else if IsOnSale != "" {
				db.Debug().Where("is_on_sale = ?",IsOnSale).Order("id desc").Find(&u3)
				db.Debug().Where("is_on_sale = ?",IsOnSale).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)
			}else if Name != "" {
				db.Debug().Where("name = ?",Name).Order("id desc").Find(&u3)
				db.Debug().Where("name = ?",Name).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)
			}
			if u2 != nil {
				u = u2
				u1 = u3
			}
			p = Table.Page{PCount:len(u1) / pageSize,Pno:currentPage,
				Psize:pageSize,TotalElements: len(u1)}
			page, _ := json.Marshal(p)
			ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": u,
				"page": string(page)})
			return
		}
		var u1 []Table.Goods
		db.Debug().Order("id desc").Find(&u1)
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": u1})
	}
}

