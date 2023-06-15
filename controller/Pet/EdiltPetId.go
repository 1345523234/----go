package Pet

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func EdiltPetId(ctx *gin.Context){ 
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}else{
		defer db.Close()
		petName := ctx.PostForm("PetName")
		petTypeID := ctx.PostForm("PetTypeId")
		// PetType := ctx.PostForm("PetType")
		PetInfoWeight := ctx.PostForm("PetInfoWeight")
		PetInfoAge := ctx.PostForm("PetInfoAge")
		PetVarieties := ctx.PostForm("PetVarieties")
		Status := ctx.PostForm("Status")
		id := ctx.PostForm("ID")
		PetTypeID, _ := strconv.Atoi(petTypeID)
		

		var u Table.Pet
		var py Table.PetType
		var PI Table.PetInfo

		db.Debug().Table("pet_types").Where("id = ?",PetTypeID).Find(&py)
		db.Debug().Table("pet_infos").Model(&u).Where("pet_id = ?",id).
		Updates(Table.PetInfo{PetName:petName,PetInfoAge:PetInfoAge,
			PetInfoType:py.PetName,PetInfoWeight:PetInfoWeight,
			PetVarieties:PetVarieties,
		Status: Status,})

		db.Debug().Table("pets").Model(&u).Where("id = ?",id).
		Updates(Table.Pet{PetName:petName,PetTypeId:petTypeID,
			PetVarieties:PetVarieties,
			PetType:py.PetName})
		ctx.JSON(http.StatusOK,gin.H{"code": 200 ,"msg":"修改","data":u,"petInfo":PI})
	}
}