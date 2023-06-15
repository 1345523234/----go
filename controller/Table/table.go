package Table

// 订单表
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
	Doctor string
	DoctorID int 
	RoleId int
  	RoleType string
	MasterID int
	PetID int
	GoodsID uint
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
	MasterID int

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
	PetVarieties string
	Price int
	Sex string
	Age int
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
	WarehouseInfo string
	WarehouseInfoID int
}
type GoodsType struct {
	ID            uint
	InsertTime    int64
	Name          string
	Remark        string
}
type RoleList struct {
	ID         uint
	InsertTime int64
	Name       string
}

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
type Page struct {
	Pno           int
	PCount        int
	Psize         int
	TotalElements int
}

type Company struct {
	ID int
	CompanyName string
	CompanyAddress string
	UnifiedSocialCreditCode string
	JuridicalPerson string
	CompanyContact string
	Admissions string
	ContactPhone string
	Email string
	OpenAccount string
	Bank string
	BankAffiliated string
	Remark string
}

type WarehouseInfo struct {
	ID int
	WarehouseAddress string
	CompanyID int
	CompanyName string
	WarehouseName string
	WarehouseType string
	WarehouseTypeStock int
	WarehouseTypeStockIng int 
}
type Location struct {
	ID int
	BeginTime string
	EndTime string
	City string
	District string
	Lat float64
	Lng float64
	Nation string
	Province string
	Street string
	StreetNumber string
	R int
}

type DataStauts struct {
	OrderLen int
	PetLen int
	UesrLen int
	WarehouseInfoLen int
	CompanieLen int
}

