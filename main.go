package main

import (

	// "gin-vue/gin-serve/controller"

	"go-serve/controller/Company"
	"go-serve/controller/DataStauts"
	"go-serve/controller/Goods"
	"go-serve/controller/GoodsDiscount"
	"go-serve/controller/GoodsType"
	"go-serve/controller/Location"
	"go-serve/controller/Master"
	"go-serve/controller/Order"
	"go-serve/controller/Pet"
	"go-serve/controller/PetType"
	"go-serve/controller/PetdengInsert"
	"go-serve/controller/Role"
	"go-serve/controller/User"
	"go-serve/controller/WarehouseInfo"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// func (v Login) TableName() string {
//     return "login"
// }
func main(){
	r := gin.Default()
	//登录接口
	r.POST("/gin-serve/login",User.Login)
	//新增userlist接口 username password deptname roleId
	r.POST("/gin-serve/user/insert",User.UserListAdd)
	//查询userlist接口 参数可带username
	r.POST("/gin-serve/user/list",User.GetUserList)
	//根据id查询用户信息
	r.POST("/gin-serve/user/find/id",User.UserFindId)
	//修改用户信息 ID
	r.POST("/gin-serve/user/update",User.UserUpdate)
	//删除用户
	r.POST("/gin-serve/user/del",User.UserDel)
	//角色信息
	r.POST("/gin-serve/role/list/all",Role.RoleListAll)
	//查询教程信息 ID
	r.POST("/gin-serve/role/find/id",Role.RoleFindId)
	//修改角色信息
	r.POST("/gin-serve/role/update",Role.UserUpdate)
	//删除角色 ID
	r.POST("/gin-serve/role/delete",Role.RoleDel)
	//新增角色
	r.POST("/gin-serve/role/insert",Role.RoleInsert)
	//根据role_id实现二级联动
	r.POST("/gin-serve/roleType/role/id",Role.RoleOrRoleId)
	//药品管理 获取所有药品信息 可分页 
	r.POST("/gin-serve/goods/list/page",Goods.GoodsAll)
	//新增药品信息
	r.POST("/gin-serve/goods/list/insert",Goods.GoodsInsert)
	r.POST("/gin-serve/goods-type/list/page",GoodsType.GoodsTypeAll)
	//新增药品
	r.POST("/gin-serve/goods-type/list/insert",GoodsType.GoodsTypeInsert)
	//上架下架药品
	r.POST("/gin-serve/goods/set/onsale",Goods.SetOnsale)
	//删除商品
	r.POST("/gin-serve/goods/list/del",Goods.DelGoods)
	//更新商品信息
	r.POST("/gin-serve/goods/list/update",Goods.GoodsUpdate)
	//根据id获取商品信息
	r.POST("/gin-serve/goods/find/id",Goods.GoodsFindId)
	//根据goodstypeID来查询商品一个二级联动
	r.POST("/gin-serve/goods/goods-type/find",Goods.GoodTypeQueryGoods)
	//根据id获取商品类型信息
	
	r.POST("/gin-serve/goods-type/find/id",GoodsType.GoodsTypeFindId)
	//修改商品信息
	r.POST("/gin-serve/goods-type/update",GoodsType.GoodsTypeUpdate)
	//根据商品信息id删除商品信息
	r.POST("/gin-serve/goods-type/delete",GoodsType.GoodsTypeDel)
	
	//获取折扣商品
	r.POST("/gin-serve/goods/goodsDiscount/all",GoodsDiscount.GoodsDiscountAll)
	//根据id查看折扣商品信息
	r.POST("/gin-serve/goods/goodsDiscount/find/id",GoodsDiscount.GoodsDiscountFindId)
	//修改商品信息
	r.POST("/gin-serve/goods/goodsDiscount/updata",GoodsDiscount.GoodsDiscountUpdata)
	//删除商品信息
	r.POST("/gin-serve/goods/goodsDiscount/del/id",GoodsDiscount.GoodsDiscountDelId)
	//新增折扣商品信息
	r.POST("/gin-serve/goods/goodsDiscount/insert",GoodsDiscount.GoodsCountInsert)
	//新增订单信息
	r.POST("/gin-serve/order/insert",Order.OrderInsert)
	//获取订单信息
	r.POST("/gin-serve/order/all",Order.OrderAll)
	//取消订单 根据id
	r.POST("/gin-serve/order/cencal/id",Order.CencalOrder)
	//完成订单 根据id
	r.POST("/gin-serve/order/Complete/id",Order.CompleteOrder)
	//删除订单
	r.POST("/gin-serve/order/del/id",Order.DelOrder)
	//宠物类型的增删改查
	r.POST("/gin-serve/petType/all",PetType.PetTypeAll)
	r.POST("/gin-serve/petType/insert",PetType.PetTypeInsert)
	r.POST("/gin-serve/petType/update",PetType.PetTypeUpdate)
	r.POST("/gin-serve/petType/del",PetType.PetTypeDel)
	r.POST("/gin-serve/petType/find/id",PetType.PetTypeFindId)
	//新增宠物订单宠物主人宠物三表关系表
	r.POST("/gin-serve/petOrmaster/Petdeng/Insert",PetdengInsert.PetInsert)
	//查询宠物表
	r.POST("/gin-serve/pet/all",Pet.PetFindAll)
	//根据宠物ID获取宠物详细信息
	r.POST("/gin-serve/pet/petInfo/all",Pet.PetInfoAll)
	//根据ID获取宠物信息
	r.POST("/gin-serve/pet/find/id",Pet.PetFindId)
	//修改宠物信息 
	r.POST("/gin-serve/pet/pet/update",Pet.EdiltPetId)
	//删除宠物信息
	r.POST("/gin-serve/pet/pet/del",Pet.DelPet)
	//获取主人详细信息表
	r.POST("/gin-serve/master/all",Master.MasterAll)
	//获取主人详细信息
	r.POST("/gin-serve/master/detail/all",Master.MasterInfo)
	//根据id获取主人信息
	r.POST("/gin-serve/master/update/id",Master.UpdateMasterInfo)
	//删除主人信息
	r.POST("/gin-serve/master/del/id",Master.DelMasterInfo)
	//新增供货商信息
	r.POST("/gin-serve/cpmpany/insert",Company.InsertCompany)
	r.POST("/gin-serve/cpmpany/all",Company.CompanyAll)
	//修改接口
	r.POST("/gin-serve/cpmpay/update/id",Company.CompanyUpdate)
	//根据id查询
	r.POST("/gin-serve/cpmpay/find/id",Company.CompanyFindid)
	//根据id删除
	r.POST("/gin-serve/cpmpay/del/id",Company.CompanyDelId)
	//新增仓库信息
	r.POST("/gin-serve/WarehouseInfo/insert",WarehouseInfo.InsertWarehouseInfo)
	//查看
	r.POST("/gin-serve/WarehouseInfo/all",WarehouseInfo.WarehouseInfoAll)
	//根据id查找
	r.POST("/gin-serve/WarehouseInfo/find/id",WarehouseInfo.WarehouseInfoFindID)
	//修改
	r.POST("/gin-serve/WarehouseInfo/update/id",WarehouseInfo.WarehouseInfoUpdate)
	//删除
	r.POST("/gin-serve/WarehouseInfo/del/id",WarehouseInfo.WarehouseInfoDel)
	//新增位置信息
	r.POST("/gin-serve/loaction/insert",Location.InsertLocation)
	r.POST("/gin-serve/loaction/All",Location.LocationAll)
	r.POST("/gin-serve/loaction/update",Location.UpdateLocation)
	//数据
	r.POST("/gin-serve/DataStauts/index",DataStauts.DataStautsIndex)
	r.Run(":8090")
}