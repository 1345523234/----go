package dao

//订单表
type Order struct {
	ID int
	Phone int64
	Name string
	SinglePrice int //订单金额
	PostCode string //订单编号
	InsertTime int64
	GoodType string
	GoodTypeID int
	Status int
	DrugName string
	OutpatientCare  int
 }
 //宠物表
 type Pet struct {
	ID int
	PetName string
	MasterName string
	PetType string
	PetDisease string
	OrderId int
	masterID int
	PetVarieties string
	PetTypeId string
 }
 //具体宠物信息表
 type PetInfo struct {
	ID int
	PetInfoAge string
	PetInfoType string
	PetID int
	PetInfoDisease string
	PetInfoWeight string
	PetName string
	Status string
	PetVarieties string
	
 }
 type MasterInfo struct {
	ID int
	MasterName string
	IdCard string
	PetName string
	PetType string
	PetId int
	OrderId int
	PetVarieties string
	Price int
 }
 type PetType struct {
	ID            uint
	InsertTime    int64
	PetName       string
	Remark        string
}
type Goods struct {
	ID uint
	Count int
	Description string
	GoodsTypeId int
	GoodsTypeName string
	InsertTime int64
	IsOnSale int
	Logo string
	Name string
	Price int
	Remark string
	Pics string
	ZheKou int
}
type GoodsType struct {
	ID            uint
	InsertTime    int64
	Name          string
	Remark        string
}