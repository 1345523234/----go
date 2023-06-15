package User

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserList struct {
	ID uint
	UserID uint
	InsertTime int64
	Face string
	Nickname string
	Password string
	RoleName string
	RoleId string
	Username string
	Deptname string
}

// type UserList struct {
// 	ID uint
// 	Name string
// 	Gender string
// 	Hobby string
// }
func UserListAdd(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	nickname := ctx.PostForm("nickname")
	roleId := ctx.PostForm("roleId")

	// fmt.Println(nickname)
	// fmt.Println(rolename)
	// fmt.Println(roleId)
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		var re RoleList
		db.Debug().Table("role_lists").Where("ID = ?",roleId).Find(&re)
		fmt.Println(re.Name)
		db.AutoMigrate(&UserList{})
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"data":re })
		u:=&User{UserName: username, PassWord: password} 
		// Struct
		db.Create(&u)
		timeUnix:=time.Now().Unix()
		//创建表 自动迁移 (把结构体和数据表进行对应)	
		u1 := UserList{Username:username,Password:password,InsertTime:timeUnix,
			Nickname:nickname,RoleName:re.Name,RoleId:roleId,
			UserID:u.ID}
		db.Create(&u1)
		
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"添加成功" })
		// Struct
		// db.Debug().First(&login)
		// fmt.Println(login)
		fmt.Println("连接成功！！！")
}
}