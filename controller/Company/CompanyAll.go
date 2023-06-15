package Company

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

func CompanyAll(ctx *gin.Context){
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()
		pno := ctx.PostForm("Pno")
		pSize := ctx.PostForm("Psize")
		CompanyContact := ctx.PostForm("companyContact")
		CompanyName := ctx.PostForm("companyName")
		// db.AutoMigrate(&UserList{})s
		if pno != "" && pSize != "" {
			currentPage, _ := strconv.Atoi(pno)
			pageSize, _ := strconv.Atoi(pSize)
			var u2 []Table.Company
			db.Debug().Limit(pageSize).Offset((currentPage - 1) * pageSize).Order("id desc").Find(&u2)
			//查询所有数据
			var u []Table.Company
			db.Debug().Order("id desc").Find(&u)
			var p Table.Page
			
			// var u1 []GoodsType
			// db.Debug().Table("goods_types").Find(&u1)
			if  CompanyContact != "" && CompanyName != ""{
				db.Debug().Where("company_contact = ? AND company_name =  ?",CompanyContact,CompanyName).Order("id desc").Find(&u)
				db.Debug().Where("company_contact = ? AND company_name =  ?",CompanyContact,CompanyName).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)				
			}else if CompanyContact != ""{
				db.Debug().Where("company_contact =  ?", CompanyContact).Order("id desc").Find(&u)
				db.Debug().Where("company_contact =  ?", CompanyContact).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)				
			}else if CompanyName != ""{
				db.Debug().Where("company_name =  ?", CompanyName).Order("id desc").Find(&u)
				db.Debug().Where("company_name =  ?", CompanyName).Order("id desc").Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&u2)				
			}
			p = Table.Page{
				PCount:        len(u) / pageSize,
				Pno:           currentPage,
				Psize:         pageSize,
				TotalElements: len(u)}
			page, _ := json.Marshal(p)
			ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": u2,"page": string(page)})
		}else{
			var u []Table.Company
			db.Debug().Order("id desc").Find(&u)
			ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": u,"page":""})
		}
		
	}
}