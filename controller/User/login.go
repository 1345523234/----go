package User

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type User struct {
	ID uint
	UserName string
	PassWord string
	
}
 func Login(ctx *gin.Context){
	UserName := ctx.PostForm("username")
	PassWord := ctx.PostForm("password")
	fmt.Println("连接失败")
	fmt.Println(UserName)
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		// db.AutoMigrate(&Login{})
		// u1 := Login{1,"admin","123456"}
		// db.Create(&u1)
		db.AutoMigrate(&User{})
		//创建表 自动迁移 (把结构体和数据表进行对应)	
		// u1 := Login{1,"admin","123456"}
		// db.Create(&u1)
			//查询
		var u User
		// Struct
		db.Debug().Where(&User{UserName: UserName, PassWord: PassWord}).Find(&u)
		fmt.Println("u::",u)
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"token":"qweasdzxc","data":u })
		// Struct
		// db.Debug().First(&login)
		// fmt.Println(login)
		fmt.Println("连接成功！！！")
	}


}
