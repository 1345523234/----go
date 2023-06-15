package PetdengInsert

import (
	"fmt"
	"go-serve/controller/Table"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


func PetInsert(ctx *gin.Context){
	
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		defer db.Close()
//创建订单表（以订单表为主表）准确的说是向表里追加
		Name := ctx.PostForm("name")
		Phone := ctx.PostForm("phone")
		SinglePrice := ctx.PostForm("singlePrice")
		GoodTypeID := ctx.PostForm("goodTypeID")
		PetName := ctx.PostForm("petName")
		PetTypeId := ctx.PostForm("petTypeId")
		PetDisease := ctx.PostForm("petDisease")
		PetInfoAge := ctx.PostForm("petInfoAge")
		PetInfoWeight := ctx.PostForm("petInfoWeight")
		Status := ctx.PostForm("status")
		PetVarieties := ctx.PostForm("petVarieties")
		IdCard := ctx.PostForm("idCard")
		goodsID:= ctx.PostForm("goodsID")
		OutpatientCare := ctx.PostForm("outpatientCare")
		DoctorID := ctx.PostForm("doctorId")
		RoleId := ctx.PostForm("roleId")
		Sex := ctx.PostForm("sex")
		Age := ctx.PostForm("age")
		// Doctor := ctx.PostForm("doctor")
		// RoleType := ctx.PostForm("roleType")
		
		// db.AutoMigrate(&Order{})
		phone, _ := strconv.ParseInt(Phone,10,64)
		singlePrice, _ := strconv.Atoi(SinglePrice)
		status, _ := strconv.Atoi(Status)
		goodTypeID, _ := strconv.Atoi(GoodTypeID)
		outpatientCare, _ := strconv.Atoi(OutpatientCare)
		roleId, _ := strconv.Atoi(RoleId)
		doctorID, _ := strconv.Atoi(DoctorID)
		age, _ := strconv.Atoi(Age)
		db.AutoMigrate(&Table.Order{})
		timeUnix := time.Now().Unix()
		var goods Table.Goods
		db.Debug().Table("goods").Where("id = ?",goodsID).Find(&goods)
		
		var wof Table.WarehouseInfo
		db.Debug().Table("warehouse_infos").Where("id = ?",goods.WarehouseInfoID).Find(&wof)

		var goodsType Table.GoodsType
		db.Debug().Table("goods_types").Where("id = ?",GoodTypeID).Find(&goodsType)

		var user Table.UserList
		db.Debug().Table("user_lists").Where("id = ?",doctorID).Find(&user)

		var role Table.RoleList
		db.Debug().Table("role_lists").Where("id = ?",roleId).Find(&role)
		db.AutoMigrate(&Table.MasterInfo{})

		var pt Table.PetType
		db.Debug().Table("pet_types").Where("id = ?",PetTypeId).Find(&pt)
		
		var msI Table.MasterInfo
		db.Debug().Table("master_infos").Where("id_card = ?",IdCard).Find(&msI)
		if goods.Count > 0{
			var wofs Table.WarehouseInfo
			db.Debug().Table("warehouse_infos").Model(&wofs).Where("id = ?", wof.ID).
			Updates(Table.WarehouseInfo{WarehouseTypeStockIng:wof.WarehouseTypeStockIng - 1})
			var g Table.Goods
			db.Debug().Table("goods").Model(&g).Where("id = ?", goods.ID).
			Updates(Table.Goods{Count:goods.Count - 1})
		}else{
			ctx.JSON(http.StatusOK, gin.H{"code": 505, "msg": "库存不足"})
			return
		}
		
		if  msI.ID == 0  {
		
			msI =Table.MasterInfo{MasterName:Name,IdCard:IdCard,
				PetName:PetName,PetType:pt.PetName,
				PetVarieties:PetVarieties,Price:singlePrice,
				Sex:Sex,Age:age}
				db.Create(&msI)
		}
		u1 := Table.Order{
			Phone:phone,Name:msI.MasterName,
			SinglePrice:singlePrice, 
			PostCode:"cs"+strconv.FormatInt(timeUnix, 10),
			InsertTime: timeUnix,GoodTypeID:goodTypeID,
			Status:status,DrugName:goods.Name,GoodType:goodsType.Name,
			OutpatientCare:outpatientCare,
			DoctorID:doctorID,RoleId:roleId,
			RoleType:role.Name,Doctor:user.Nickname,MasterID:msI.ID,
			GoodsID:goods.ID}
		db.Create(&u1)
		// u1 := &u
		// fmt.Println(u1,"ddddd",u.ID)
		
		//创建宠物表
		
		
		db.AutoMigrate(&Table.Pet{})
		p1 :=Table.Pet{PetName:PetName,MasterName:u1.Name,PetTypeId:PetTypeId,
			PetType:pt.PetName,PetDisease:PetDisease,OrderId:u1.ID,
			PetVarieties:PetVarieties,
			MasterID:msI.ID}
	db.Create(&p1)

	// //创建宠物详细信息表
	db.AutoMigrate(&Table.PetInfo{})
	 pI := Table.PetInfo{PetInfoAge:PetInfoAge,PetInfoType:pt.PetName,PetID:p1.ID,
		PetInfoDisease:PetDisease,PetInfoWeight:PetInfoWeight,PetName:PetName,
		Status:Status,PetVarieties:PetVarieties}
	db.Create(&pI)
	// //创建主人详细信息表
	// db.AutoMigrate(&MasterInfo{})
	// msI :=MasterInfo{MasterName:u1.Name,IdCard:IdCard,PetName:PetName,
	// 	PetType:pt.PetName,PetId:p1.ID,OrderId:u1.ID,PetVarieties:PetVarieties,
	// 	Price:singlePrice}
	// 	fmt.Println(msI,"ddddd")
	// 	db.Create(&msI)
		
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "添加成功"})
		fmt.Println("连接成功！！！")
	}
	

	//创建宠物信息表

	
}