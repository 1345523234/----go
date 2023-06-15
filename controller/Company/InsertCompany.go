package Company

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)




func InsertCompany(ctx *gin.Context) {
	CompanyName := ctx.PostForm("companyName")
	CompanyAddress := ctx.PostForm("companyAddress")
	UnifiedSocialCreditCode := ctx.PostForm("unifiedSocialCreditCode")
	JuridicalPerson := ctx.PostForm("juridicalPerson")
	CompanyContact := ctx.PostForm("companyContact")
	Admissions := ctx.PostForm("admissions")
	// pics := ctx.PostForm("pics")
	remark := ctx.PostForm("remark")
	ContactPhone := ctx.PostForm("contactPhone")
	Email := ctx.PostForm("email")
	OpenAccount := ctx.PostForm("openAccount")
	Bank := ctx.PostForm("bank")
	BankAffiliated := ctx.PostForm("bankAffiliated")
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		// var gt Goods
		// db.Debug().Table("role_lists").Where("ID = ?",roleId).Find(&gt)
		// fmt.Println(re.Name)
		db.AutoMigrate(&Table.Company{})
		// ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"data":re })
		//创建表 自动迁移 (把结构体和数据表进行对应)	
		u1 := Table.Company{
			CompanyName:CompanyName,CompanyAddress:CompanyAddress,
			UnifiedSocialCreditCode:UnifiedSocialCreditCode,
			JuridicalPerson:JuridicalPerson,CompanyContact:CompanyContact,
			Admissions:Admissions,Remark:remark,ContactPhone:ContactPhone,
			OpenAccount:OpenAccount,Bank:Bank,BankAffiliated:BankAffiliated,
			Email:Email}
		db.Create(&u1)
	
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"添加成功" ,"data":u1})
		// Struct
		// db.Debug().First(&login)
		// fmt.Println(login)
		fmt.Println("连接成功！！！")
}
}