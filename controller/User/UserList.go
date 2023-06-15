package User

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Page struct {
	Pno int
	PCount int
	Psize int
	TotalElements int

}
func GetUserList(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		username := ctx.PostForm("Username")
		// var currentPage int
		pno := ctx.PostForm("Pno")
		pSize := ctx.PostForm("Psize")

		if pno != "" &&  pSize != ""{
			currentPage, _ := strconv.Atoi(pno)
			pageSize, _ := strconv.Atoi(pSize)
			var u2 []UserList
			db.Debug().Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)
			//查询所有数据
			var u []UserList
		db.Debug().Find(&u)
		var p Page
		p = Page{ 
			PCount: len(u)/pageSize,
			Pno: currentPage,
			Psize: pageSize,
			TotalElements: len(u)}
		page,_ := json.Marshal(p)
		if username != "" {
			var u1 []UserList
		db.Debug().Where("Username =  ?",username).Limit(pageSize).Order("id desc").Offset((currentPage - 1) * pageSize).Find(&u1)
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"data":u1 ,"page":string(page)})
		return
		}
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"data":u2, 
		"page":string(page)})	
		return
		}
		
		// db.AutoMigrate(&UserList{})s
		
			//查询
		var u []UserList
		db.Debug().Find(&u)
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"data":u })
		fmt.Println("连接成功！！！")
}
}