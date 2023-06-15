package GoodsDiscount

import (
	"encoding/json"
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// db.Not("name", "jinzhu").First(&user)
	


func GoodsDiscountAll(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()
		pno := ctx.PostForm("Pno")
		pSize := ctx.PostForm("Psize")
		username := ctx.PostForm("Username")
		// db.AutoMigrate(&UserList{})s
		if pno != "" && pSize != "" {
			currentPage, _ := strconv.Atoi(pno)
			pageSize, _ := strconv.Atoi(pSize)
			var t = "2023-03-06 12:00:00"
			 layuot := "2006-01-02 15:04:05"
			t2,_ := time.Parse(layuot,t)
			fmt.Println(t2.Unix())
		
			var u2 []Table.Goods
			db.Debug().Not("zhe_kou", 0).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)
			//查询所有数据
			var u []Table.Goods
			db.Debug().Not("zhe_kou", 0).Order("id desc").Find(&u)
			var p Table.Page
			p = Table.Page{
				PCount:        len(u) / pageSize,
				Pno:           currentPage,
				Psize:         pageSize,
				TotalElements: len(u)}
			page, _ := json.Marshal(p)
			if username != "" {
				var u1 []Table.Goods
				db.Debug().Where("Name =  ?", username).Order("id desc").Find(&u1)
				ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": u1, "page": string(page)})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": u2,
				"page": string(page)})
			return
		}
		var u1 []Table.Goods
		// db.Debug().Find(&u1)
		db.Debug().Not("zhe_kou", 0).Order("id desc").Find(&u1)
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": u1})
	}
}
