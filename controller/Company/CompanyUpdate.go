package Company

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


func CompanyUpdate(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		CompanyName := ctx.PostForm("CompanyName")
	CompanyAddress := ctx.PostForm("CompanyAddress")
	UnifiedSocialCreditCode := ctx.PostForm("UnifiedSocialCreditCode")
	JuridicalPerson := ctx.PostForm("JuridicalPerson")
	CompanyContact := ctx.PostForm("CompanyContact")
	Admissions := ctx.PostForm("Admissions")
	// pics := ctx.PostForm("pics")
	remark := ctx.PostForm("Remark")
	ContactPhone := ctx.PostForm("ContactPhone")
	Email := ctx.PostForm("Email")
	OpenAccount := ctx.PostForm("OpenAccount")
	Bank := ctx.PostForm("Bank")
	BankAffiliated := ctx.PostForm("BankAffiliated")
	
	id := ctx.PostForm("ID")
		var u Table.Company
		// db.Debug().Table("role_lists").Model(&u).Where("ID = ?",id).
		// Updates(Table.RoleList{Name:name,InsertTime: timeUnix})
		db.Debug().Table("companies").Model(&u).Where("id = ?",id).
		Updates(Table.Company{	CompanyName:CompanyName,CompanyAddress:CompanyAddress,
			UnifiedSocialCreditCode:UnifiedSocialCreditCode,
			JuridicalPerson:JuridicalPerson,CompanyContact:CompanyContact,
			Admissions:Admissions,Remark:remark,ContactPhone:ContactPhone,
			OpenAccount:OpenAccount,Bank:Bank,BankAffiliated:BankAffiliated,
			Email:Email})
		db.Debug().Where("id = ?",id).Find(&u)
			fmt.Println("连接失败-+------------------------",CompanyName)
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"修改成功" })
	}
}